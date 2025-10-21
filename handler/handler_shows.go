package handler

import (
	"context"
	"database/sql"
	"fmt"
	"go_blog/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func GetList(c *gin.Context, db_user *sql.DB, minioClient *minio.Client) {
	type request struct {
		ShowName string `json:"show_name"`
	}
	type response struct {
		Shows []string `json:"shows"`
	}
	user := database.GetUserByAuthHeader(db_user, c.Request.Header.Get("Authorization"))
	if user.Id == 0 {
		c.JSON(http.StatusForbidden, response{Shows: []string{}})
		return
	}
	// get show name
	var req request
	if c.BindJSON(&req) != nil {
		c.JSON(http.StatusBadRequest, response{Shows: []string{}})
		return
	}

	info := minioClient.ListObjects(c, "shows", minio.ListObjectsOptions{
		Prefix:    req.ShowName,
		Recursive: true,
	})
	// print info
	var shows []string
	for obj := range info {
		shows = append(shows, obj.Key)
	}
	c.JSON(http.StatusOK, response{Shows: shows})
}

func GetShowPresignedUrl(c *gin.Context, minioClient *minio.Client, db_user *sql.DB) {
	// verify user

	user := database.GetUserByAuthHeader(db_user, c.Request.Header.Get("Authorization"))
	if user.Id == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
		return
	}
	// get path from json
	type request struct {
		Path string `json:"path"`
	}
	var req request
	if c.BindJSON(&req) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid path"})
		return
	}
	//objectName := fmt.Sprintf("%s", req.Path)
	objectName := req.Path
	fmt.Println("objectName: ", objectName)
	presignedUrl, err := minioClient.PresignedGetObject(context.Background(),
		"shows", objectName, time.Hour*2, nil)
	if err != nil {
		//return "", err
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"show_url": presignedUrl.String()})
	//return presignedUrl.String(), nil

}

func GetPresignedUrlNew(c *gin.Context, minioClient *minio.Client, db_user *sql.DB) {
	// verify user
	user := database.GetUserByAuthHeader(db_user, c.Request.Header.Get("Authorization"))
	if user.Id == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
		return
	}
	// get path from json
	type request struct {
		Bucket string `json:"bucket"`
		Path   string `json:"path"`
	}
	var req request
	if c.BindJSON(&req) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid path"})
		return
	}
	//objectName := fmt.Sprintf("%s", req.Path)
	objectName := req.Path
	fmt.Println("objectName: ", objectName)
	presignedUrl, err := minioClient.PresignedGetObject(context.Background(),
		req.Bucket, objectName, time.Hour*2, nil)
	if err != nil {
		//return "", err
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"presigned_url": presignedUrl.String()})

}
