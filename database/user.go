package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

type UserDbConfig struct {
	SqlitePath string `json:"sqlite_path"`
	SecreteKey []byte `json:"secrete_key"`
}

//var secreteKey = []byte("bcb967bec859b86e96564992792636bb442548af35a2e3374cee7a0f92542c18")

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

var userDbConfig UserDbConfig

func SetUserDbConfig(config UserDbConfig) {
	userDbConfig = config
}

func UserDbInit(config UserDbConfig) (db_user *sql.DB, err error) {
	SetUserDbConfig(config)
	log.Println("Initializing user database ...", userDbConfig)
	db_user, err = sql.Open("sqlite3", userDbConfig.SqlitePath)
	if err != nil {
		return nil, err
	}
	query := ` CREATE TABLE IF NOT EXISTS user (
	id         INTEGER PRIMARY KEY AUTOINCREMENT,
    	email      VARCHAR(255) UNIQUE NOT NULL,
    	name       VARCHAR(255),
    	password   VARCHAR(255),
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`
	_, err = db_user.Exec(query)
	if err != nil {
		db_user.Close()
		return nil, err
	}
	return db_user, nil
}

func UserAdd(db_user *sql.DB, user User, password string) error {
	query := "INSERT INTO user (email, name, password) VALUES (?, ?, ?)"
	_, err := db_user.Exec(query, user.Email, user.Name, password)
	return err
}

// GetUserByEmail get user by email
// If cannot find user, return empty user
func GetUserByEmail(dbUser *sql.DB, email string) User {
	var user User
	query := "SELECT id, email, name FROM user WHERE email=?"
	err := dbUser.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.Name)
	if err != nil {
		log.Println(err)
		return User{}
	}
	return user
}

func VerifyToken(token string) (bool, string) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return userDbConfig.SecreteKey, nil
	})
	if err != nil {
		log.Println("verify token failed: ", err)
		return false, ""
	}
	valid := parsedToken.Valid
	email := parsedToken.Claims.(jwt.MapClaims)["email"].(string)
	return valid, email
}

// GetUserByAuthHeader get user by auth header
// if auth type is Bearer, get token and verify it
// if auth type is Basic, return empty user
// if auth header is invalid, return empty user
func GetUserByAuthHeader(db_user *sql.DB, auth string) User {
	// if auth type is Bearer get token
	if len(auth) < 7 {
		return User{}
	}
	if auth[0:7] == "Bearer " {
		token := auth[7:]
		valid, email := VerifyToken(token)
		if !valid {
			return User{}
		} else {
			return GetUserByEmail(db_user, email)
		}
	}
	return User{}
}

func Login(db_user *sql.DB, email string, password string) bool {
	// select rwo from users where email = email and password = password
	query := "SELECT email FROM user WHERE email=? AND password=?"
	err := db_user.QueryRow(query, email, password).Scan(&email)
	if err != nil {
		log.Println("Login: ", err)
		return false
	}
	return true
}

func GenerateToken(email string) string {
	log.Println("Generating token for user: ", email, " ...")
	signingMethod := jwt.SigningMethodHS256 // HS256 is an instance of HMAC
	claims := jwt.MapClaims{
		"email": email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	token := jwt.NewWithClaims(signingMethod, claims)
	signedToken, err := token.SignedString(userDbConfig.SecreteKey)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Generated token success")
	return signedToken
}
