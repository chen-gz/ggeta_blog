package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type PostFile struct {
	Id       int    `json:"id"`
	PostId   int    `json:"post_id"`
	FileName string `json:"file_name"`
	FileUrl  string `json:"file_url"`
}

func initializeFileTable(db *sql.DB) error {

	log.Println(blogDbConfig.BlogFileTable)
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id         INTEGER PRIMARY KEY AUTOINCREMENT,
    		post_id    INT UNSIGNED NOT NULL,
    		file_name  VARCHAR(2048) NOT NULL,
		file_url   VARCHAR(2048) NOT NULL
		)`, blogDbConfig.BlogFileTable)
	_, err := db.Exec(query)
	return err
}

func insertFile(db *sql.DB, post_id int, file_name string, file_url string) error {
	query := fmt.Sprintf(`INSERT INTO %s (post_id, file_name, file_url) VALUES (?, ?, ?)`, blogDbConfig.BlogFileTable)
	_, err := db.Exec(query, post_id, file_name, file_url)
	return err
}
func seachFile(db *sql.DB, post_id int) ([]PostFile, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE post_id = ?`, blogDbConfig.BlogFileTable)
	rows, err := db.Query(query, post_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var files []PostFile
	for rows.Next() {
		var file PostFile
		err := rows.Scan(&file.Id, &file.PostId, &file.FileName, &file.FileUrl)
		if err != nil {
			return files, err
		}
		files = append(files, file)
	}
	return files, nil
}

func deleteFile(db *sql.DB, post_id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE post_id = ?`, blogDbConfig.BlogFileTable)
	_, err := db.Exec(query, post_id)
	return err
}

func InsertFileUser(db_blog *sql.DB, user User, post_id int, file_name string, file_url string) error {
	// check permission, should be same as updatePost
	if !UpdatePostPermissionCheck(db_blog, user, post_id) {
		return errors.New("permission denied")
	}
	return insertFile(db_blog, post_id, file_name, file_url)
}

func SearchFile(db *sql.DB, user User, post_id int) ([]PostFile, error) {
	// check permission, should be same as updatePost
	if !UpdatePostPermissionCheck(db, user, post_id) {
		return nil, errors.New("permission denied")
	}
	return seachFile(db, post_id)
}
