package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// each user has its own photo table

type PhotoItem struct {
	Id          int    `json:"id"`
	Hash        string `json:"hash"` //  sha1 hash. If original file exist. Use the hash of the original file. If not, use the hash of the Jpeg file.
	HasOriginal bool   `json:"has_original"`
	OriginalExt string `json:"original_ext"`
	Deleted     bool   `json:"deleted"`
	Tags        string `json:"tags"`
	Category    string `json:"category"`
}
type PhotoDbConfig struct {
	Address       string `json:"address"`
	User          string `json:"user"`
	Password      string `json:"password"`
	PhotoDatabase string `json:"photo_database"`
}

func InitPhotoDb(config PhotoDbConfig) (db_photo *sql.DB, err error) {
	sql_endpoint := fmt.Sprintf("%s:%s@%s/", config.User, config.Password, config.Address)
	db, err := sql.Open("mysql", sql_endpoint)
	if err != nil {
		panic(err)
	}
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", config.PhotoDatabase)
	_, err = db.Exec(query)
	err = db.Close()
	if err != nil {
		return nil, err
	}

	sql_endpoint = fmt.Sprintf("%s:%s@%s/%s", config.User, config.Password, config.Address, config.PhotoDatabase)
	db_photo, err = sql.Open("mysql", sql_endpoint)
	if err != nil {
		panic(err)
	}
	return db_photo, nil
}

func initPhotoTable(photo_db *sql.DB, user User) error {
	if user.Id == 0 {
		return errors.New("user id is 0")
	}
	table_name := fmt.Sprintf("photo_%d", user.Id)

	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
    		id           INT UNSIGNED UNIQUE AUTO_INCREMENT,
    		hash         VARCHAR(255) UNIQUE NOT NULL,
    		has_original BOOLEAN NOT NULL DEFAULT FALSE,
			original_ext VARCHAR(255) NOT NULL DEFAULT "",
    		deleted      BOOLEAN NOT NULL DEFAULT FALSE,
    		tags         VARCHAR(2048),
    		category     VARCHAR(2048),
    		created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    		PRIMARY KEY (id))`, table_name)
	_, err := photo_db.Exec(query)
	return err
}

func addPhoto(photo_db *sql.DB, user User, photo PhotoItem) error {
	table_name := fmt.Sprintf("photo_%d", user.Id)
	query := fmt.Sprintf(`INSERT INTO %s (hash, has_original, original_ext) VALUES (?, ?, ?)`, table_name)
	_, err := photo_db.Exec(query, photo.Hash, photo.HasOriginal, photo.OriginalExt)
	return err
}

func getPhoto(photo_db *sql.DB, user User, id int) (PhotoItem, error) {
	table_name := fmt.Sprintf("photo_%d", user.Id)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = ?`, table_name)
	row := photo_db.QueryRow(query, id)
	var photo PhotoItem
	err := row.Scan(&photo.Id, &photo.Hash, &photo.HasOriginal, &photo.OriginalExt, &photo.Deleted, &photo.Tags, &photo.Category)
	return photo, err
}

func deletePhoto(photo_db *sql.DB, user User, id string) error {
	table_name := fmt.Sprintf("photo_%d", user.Id)
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, table_name)
	_, err := photo_db.Exec(query, id)
	return err
}
func InsertPhotoUser(photo_db *sql.DB, user User, photo PhotoItem) error {
	err := initPhotoTable(photo_db, user)
	if err != nil {
		return err
	}
	if user.Id == 0 || user.Name == "" {
		return errors.New("invalid user")
	}
	return addPhoto(photo_db, user, photo)
}

func GetPhotoUser(photo_db *sql.DB, user User, id int) (PhotoItem, error) {
	if user.Id == 0 || user.Name == "" {
		return PhotoItem{}, errors.New("invalid user")
	}
	return getPhoto(photo_db, user, id)
}
func DeletePhotoUser(photo_db *sql.DB, user User, id string) error {
	if user.Id == 0 || user.Name == "" {
		return errors.New("invalid user")
	}
	return deletePhoto(photo_db, user, id)
}
