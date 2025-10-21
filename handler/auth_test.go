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
		UserTable:  "test",
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
	token := database.V3GenerateToken("test@example.com")
	req.Header.Set("Authorization", "Bearer "+token)

	// Mock the database query
	rows := sqlmock.NewRows([]string{"id", "email", "name"}).
		AddRow(1, "test@example.com", "Test User")
	mock.ExpectQuery("SELECT id, email, name FROM test WHERE email=?").
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
	token = database.V3GenerateToken("not-in-db@example.com")
	req.Header.Set("Authorization", "Bearer "+token)

	// Mock the database query
	mock.ExpectQuery("SELECT id, email, name FROM test WHERE email=?").
		WithArgs("not-in-db@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "name"}))

	r.ServeHTTP(w, req)

	// Check that the request was unauthorized
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, w.Code)
	}

}
