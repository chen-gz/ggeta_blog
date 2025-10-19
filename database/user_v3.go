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
	Address      string `json:"address"`
	User         string `json:"user"`
	Password     string `json:"password"`
	UserDatabase string `json:"user_database"`
	UserTable    string `json:"user_table"`
	SecreteKey   []byte `json:"secrete_key"`
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
	sql_endpoint := fmt.Sprintf("%s:%s@%s/", userDbConfig.User, userDbConfig.Password, userDbConfig.Address)
	db, err := sql.Open("mysql", sql_endpoint)
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", userDbConfig.UserDatabase)
	_, err = db.Exec(query)
	if err != nil {
		db.Close()
		return nil, err
	}
	err = db.Close()
	if err != nil {
		return nil, err
	}

	sql_endpoint = fmt.Sprintf("%s:%s@%s/%s", userDbConfig.User, userDbConfig.Password, userDbConfig.Address, userDbConfig.UserDatabase)
	db_user, err = sql.Open("mysql", sql_endpoint)
	if err != nil {
		return nil, err
	}
	query = fmt.Sprintf(` CREATE TABLE IF NOT EXISTS %s (
    	id         INT UNSIGNED AUTO_INCREMENT,
    	email      VARCHAR(255) UNIQUE NOT NULL,
    	name       VARCHAR(255),
    	password   VARCHAR(255),
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    	PRIMARY KEY (id)
		);`, userDbConfig.UserTable)

	_, err = db_user.Exec(query)

	if err != nil {
		db_user.Close()
		return nil, err
	}
	return db_user, nil
}

func UserAdd(db_user *sql.DB, user User, password string) error {
	query := fmt.Sprintf("INSERT INTO %s (email, name, password) VALUES (?, ?, ?)", userDbConfig.UserTable)
	_, err := db_user.Exec(query, user.Email, user.Name, password)
	return err
}

// GetUserByEmail get user by email
// If cannot find user, return empty user
func GetUserByEmail(dbUser *sql.DB, email string) User {
	var user User
	query := fmt.Sprintf("SELECT id, email, name FROM %s WHERE email=?", userDbConfig.UserTable)
	err := dbUser.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.Name)
	if err != nil {
		log.Println(err)
		return User{}
	}
	return user
}

func V1VerifyToken(token string) (bool, string) {
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

// V3GetUserByAuthHeader get user by auth header
// if auth type is Bearer, get token and verify it
// if auth type is Basic, return empty user
// if auth header is invalid, return empty user
func V3GetUserByAuthHeader(db_user *sql.DB, auth string) User {
	// if auth type is Bearer get token
	if len(auth) < 7 {
		return User{}
	}
	if auth[0:7] == "Bearer " {
		token := auth[7:]
		valid, email := V1VerifyToken(token)
		if !valid {
			return User{}
		} else {
			return GetUserByEmail(db_user, email)
		}
	}
	return User{}
}

func V3Login(db_user *sql.DB, email string, password string) bool {
	// select rwo from users where email = email and password = password
	query := fmt.Sprintf("SELECT email FROM %s WHERE email=? AND password=?", userDbConfig.UserTable)
	err := db_user.QueryRow(query, email, password).Scan(&email)
	if err != nil {
		log.Println("V3Login: ", err)
		return false
	}
	return true
}

func V3GenerateToken(email string) string {
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
