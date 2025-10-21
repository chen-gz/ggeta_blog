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
	SqlitePath string `json:"sqlite_path"`
}

func InitPhotoDb(config PhotoDbConfig) (db_photo *sql.DB, err error) {
	db_photo, err = sql.Open("sqlite3", config.SqlitePath)
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
		id           INTEGER PRIMARY KEY AUTOINCREMENT,
    		hash         VARCHAR(255) UNIQUE NOT NULL,
    		has_original BOOLEAN NOT NULL DEFAULT FALSE,
			original_ext VARCHAR(255) NOT NULL DEFAULT "",
    		deleted      BOOLEAN NOT NULL DEFAULT FALSE,
    		tags         VARCHAR(2048) NOT NULL DEFAULT "",
    		category     VARCHAR(2048) NOT NULL DEFAULT "",
    		created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`, table_name)
	_, err := photo_db.Exec(query)
	return err
}

func addPhoto(photo_db *sql.DB, user User, photo PhotoItem) error {
	table_name := fmt.Sprintf("photo_%d", user.Id)
	query := fmt.Sprintf(`INSERT INTO %s (hash, has_original, original_ext) VALUES (?, ?, ?)`, table_name)
	_, err := photo_db.Exec(query, photo.Hash, photo.HasOriginal, photo.OriginalExt)
	return err
}
func updatePhoto(photo_db *sql.DB, user User, photo PhotoItem) error {
	// only some field can be updated - tags, category, deleted
	table_name := fmt.Sprintf("photo_%d", user.Id)
	query := fmt.Sprintf(`UPDATE %s SET tags = ?, category = ?, deleted = ? WHERE id = ?`, table_name)
	_, err := photo_db.Exec(query, photo.Tags, photo.Category, photo.Deleted, photo.Id)
	return err
}

func getPhoto(photo_db *sql.DB, user User, id int) (PhotoItem, error) {
	table_name := fmt.Sprintf("photo_%d", user.Id)
	query := fmt.Sprintf(`SELECT id, hash, has_original, original_ext, deleted, tags, category FROM %s WHERE id = ?`, table_name)
	//log.Println("getPhoto: ", query, "id: ", id)
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

func GetPhotoUser(photoDb *sql.DB, user User, id int) (PhotoItem, error) {
	if user.Id == 0 || user.Name == "" {
		return PhotoItem{}, errors.New("invalid user")
	}
	return getPhoto(photoDb, user, id)
}
func DeletePhotoUser(photoDb *sql.DB, user User, id string) error {
	if user.Id == 0 || user.Name == "" {
		return errors.New("invalid user")
	}
	return deletePhoto(photoDb, user, id)
}

func GetAllPhotoList(photoDb *sql.DB, user User) (ids []int, err error) {
	tableName := fmt.Sprintf("photo_%d", user.Id)
	query := fmt.Sprintf(`SELECT id FROM %s WHERE deleted = FALSE`, tableName)
	rows, err := photoDb.Query(query)
	if err != nil {
		return []int{}, err
	}
	defer rows.Close()
	var id int
	//var ids []int
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return []int{}, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func UpdatePhotoUser(photoDb *sql.DB, user User, photo PhotoItem) error {
	if user.Id == 0 || user.Name == "" {
		return errors.New("invalid user")
	}
	return updatePhoto(photoDb, user, photo)
}

func GetDeletedPhotoList(photoDb *sql.DB, user User) (ids []int, err error) {
	tableName := fmt.Sprintf("photo_%d", user.Id)
	query := fmt.Sprintf(`SELECT id FROM %s WHERE deleted = TRUE`, tableName)
	rows, err := photoDb.Query(query)
	if err != nil {
		return []int{}, err
	}
	defer rows.Close()
	var id int
	//var ids []int
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return []int{}, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}
