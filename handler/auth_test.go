package handler

import (
	"github.com/gin-gonic/gin"
	"go_blog/database"
	"net/http"
	"net/http/httptest"
	"testing"
)

import (
	"github.com/DATA-DOG/go-sqlmock"
)

import (
	"bytes"
	"encoding/json"
	"errors"
)

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a mock user database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	database.SetUserDbConfig(database.UserDbConfig{
		SqlitePath: ":memory:",
		SecreteKey: []byte("test"),
	})

	// Create a new router
	r := gin.New()
	r.Use(AuthMiddleware(db))
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Create a test request with a valid token
	req, _ := http.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	token := database.GenerateToken("test@example.com")
	req.Header.Set("Authorization", "Bearer "+token)

	// Mock the database query
	rows := sqlmock.NewRows([]string{"id", "email", "name"}).
		AddRow(1, "test@example.com", "Test User")
	mock.ExpectQuery("SELECT id, email, name FROM user WHERE email=?").
		WithArgs("test@example.com").
		WillReturnRows(rows)

	r.ServeHTTP(w, req)

	// Check that the request was successful
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Create a test request with an invalid token
	req, _ = http.NewRequest("GET", "/test", nil)
	w = httptest.NewRecorder()
	req.Header.Set("Authorization", "Bearer invalid-token")
	r.ServeHTTP(w, req)

	// Check that the request was unauthorized
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, w.Code)
	}

	// Create a test request with a valid token but user not in database
	req, _ = http.NewRequest("GET", "/test", nil)
	w = httptest.NewRecorder()
	token = database.GenerateToken("not-in-db@example.com")
	req.Header.Set("Authorization", "Bearer "+token)

	// Mock the database query
	mock.ExpectQuery("SELECT id, email, name FROM user WHERE email=?").
		WithArgs("not-in-db@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "name"}))

	r.ServeHTTP(w, req)

	// Check that the request was unauthorized
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, w.Code)
	}

}

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	database.SetUserDbConfig(database.UserDbConfig{
		SqlitePath: ":memory:",
		SecreteKey: []byte("test"),
	})

	r := gin.New()
	r.POST("/login", func(c *gin.Context) {
		Login(c, db)
	})

	// Test successful login
	loginDetails := map[string]string{
		"email":    "test@example.com",
		"password": "password",
	}
	jsonValue, _ := json.Marshal(loginDetails)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	rows := sqlmock.NewRows([]string{"email"}).AddRow("test@example.com")
	mock.ExpectQuery("SELECT email FROM user WHERE email=\\? AND password=\\?").
		WithArgs("test@example.com", "password").
		WillReturnRows(rows)

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Test failed login
	loginDetails = map[string]string{
		"email":    "test@example.com",
		"password": "wrong_password",
	}
	jsonValue, _ = json.Marshal(loginDetails)
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	mock.ExpectQuery("SELECT email FROM user WHERE email=\\? AND password=\\?").
		WithArgs("test@example.com", "wrong_password").
		WillReturnError(errors.New("not found"))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, w.Code)
	}

	// Test invalid request
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
}
