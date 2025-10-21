package handler

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go_blog/database"
	"go_blog/interfaces"
	"net/http"
	"time"
)

func AddVideo(c *gin.Context, dbUser *sql.DB, dbVideo *sql.DB,
	minioClient *minio.Client, md5 string, sha256 string, title string, ext string) {
	type response struct {
		Video        interfaces.VideoItem `json:"video"`
		Message      string               `json:"message"`
		PresignedUrl string               `json:"presigned_url"`
	}
	// check parameter, md5 and sha256 is required, title is optional
	if len(md5) != 32 || len(sha256) != 64 || len(ext) == 0 {
		c.JSON(http.StatusBadRequest, response{
			Message: "invalid md5 or sha256",
		})
		return
	}
	user := database.GetUserByAuthHeader(dbUser, c.Request.Header.Get("Authorization"))
	if user.Id == 0 {
		c.JSON(http.StatusForbidden, response{Message: "permission denied"})
		return
	}
	video := database.GetVideoByMd5Sha256(dbVideo, user.Id, md5, sha256)
	if video.Id != 0 {
		c.JSON(http.StatusBadRequest, response{
			Message: "video already exist",
		})
		return
	}
	// get presigned url
	presignedUrl, err := putPresignedUrl(minioClient, user, md5, sha256, video.Ext)
	if err != nil {
		c.JSON(http.StatusBadRequest, response{
			Message: "get presigned url error",
		})
		return
	}
	// insert video
	video = interfaces.VideoItem{
		UserId: user.Id,
		Md5:    md5,
		Sha256: sha256,
		Ext:    ext,
		Title:  title,
	}
	err = database.InsertVideoUser(dbVideo, video)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response{
			Message: "insert video error",
		})
		return
	}
	c.JSON(http.StatusOK, response{
		Video:        video,
		Message:      "success",
		PresignedUrl: presignedUrl,
	})
}
func putPresignedUrl(minioClient *minio.Client, user database.User, md5 string, sha256 string, ext string) (string, error) {
	objectName := fmt.Sprintf("%d/%s/%s.%s", user.Id, md5[0:5], sha256[0:5], ext)
	fmt.Println("objectName: ", objectName)
	presignedUrl, err := minioClient.PresignedPutObject(context.Background(), VideoMinioConfig.BucketName, objectName, time.Hour*2)
	if err != nil {
		return "", err
	}
	return presignedUrl.String(), nil
}
func getPresignedUrl(minioClient *minio.Client, user database.User, md5 string, sha256 string, ext string) (string, error) {
	objectName := fmt.Sprintf("%d/%s/%s.%s", user.Id, md5[0:5], sha256[0:5], ext)
	fmt.Println("objectName: ", objectName)
	presignedUrl, err := minioClient.PresignedGetObject(context.Background(), VideoMinioConfig.BucketName, objectName, time.Hour*2, nil)
	if err != nil {
		return "", err
	}
	return presignedUrl.String(), nil
}
func getPresignedCoverImg(client *minio.Client, user database.User, md5 string, sha256 string) (string, error) {
	objectName := fmt.Sprintf("%d/%s/%s.jpg", user.Id, md5[0:5], sha256[0:5])
	fmt.Println("objectName: ", objectName)
	presignedUrl, err := client.PresignedGetObject(context.Background(), VideoMinioConfig.BucketName, objectName, time.Hour*2, nil)
	if err != nil {
		return "", err
	}
	return presignedUrl.String(), nil
}

var VideoMinioConfig MinioConfig

func InitVideoMinioClient(_config MinioConfig) (client *minio.Client) {
	VideoMinioConfig = _config
	minioClient, err := minio.New(_config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(_config.AccessKeyID, _config.SecreteAccessKey, ""),
		Secure: true,
	})
	if err != nil {
		panic(err)
	}
	return minioClient
}

func process(minioClient *minio.Client, user database.User, md5 string, sha256 string) (err error) {
	// get movie info and get m3u8
	//
	//objectName := fmt.Sprintf("%d/%s/%s", user.Id, md5[0:5], sha256[0:5])
	// download movie from minio
	//object := minio.GetObjectOptions{}

	//object, err := minioClient.GetObject(context.Background(), VideoMinioConfig.BucketName, objectName, minio.GetObjectOptions{})
	//if err != nil {
	//	panic(err)
	//}
	//return err

	// convert to m3u8 use ffmpeg

	// upload m3u8 to minio
	return nil
}

func GetVideoList(c *gin.Context, dbUser *sql.DB, dbVideo *sql.DB) {
	type response struct {
		Videos  []interfaces.VideoItem `json:"videos"`
		Message string                 `json:"message"`
	}
	user := database.GetUserByAuthHeader(dbUser, c.Request.Header.Get("Authorization"))
	if user.Id == 0 {
		c.JSON(http.StatusForbidden, response{Message: "permission denied"})
		return
	}
	videoList := database.GetVideoUser(dbVideo, user)
	c.JSON(http.StatusOK, response{
		Videos:  videoList,
		Message: "success",
	})

}

func GetVideo(c *gin.Context, dbUser *sql.DB, dbVideo *sql.DB, minioClient *minio.Client, md5 string, sha256 string, id int) {
	type response struct {
		Video        interfaces.VideoItem `json:"video"`
		Message      string               `json:"message"`
		PresignedUrl string               `json:"presigned_url"`
	}
	if (len(md5) != 32 || len(sha256) != 64) && id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}
	user := database.GetUserByAuthHeader(dbUser, c.Request.Header.Get("Authorization"))
	if user.Id == 0 {
		c.JSON(http.StatusForbidden, response{Message: "permission denied"})
		return
	}
	// id is the first priority, then sha256, then md5
	var video interfaces.VideoItem
	if id != 0 {
		video = database.GetVideoById(dbVideo, user.Id, id)
	} else {
		video = database.GetVideoByMd5Sha256(dbVideo, user.Id, md5, sha256)
	}
	if video.Id == 0 {
		c.JSON(http.StatusNotFound, response{
			Message: "video not found",
		})
		return
	}
	url, err := getPresignedUrl(minioClient, user, video.Md5, video.Sha256, video.Ext)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response{
			Message: "get presigned url error",
		})
		return
	}
	c.JSON(http.StatusOK, response{
		Video:        video,
		Message:      "success",
		PresignedUrl: url,
	})
}

func GetVideoMeta(c *gin.Context, dbUser *sql.DB, dbVideo *sql.DB, minioClient *minio.Client, md5 string, sha256 string, id int) {
	type response struct {
		Video             interfaces.VideoItem `json:"video"`
		Message           string               `json:"message"`
		PresignedCoverImg string               `json:"presigned_cover_img"`
	}
	if (len(md5) != 32 || len(sha256) != 64) && id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}
	user := database.GetUserByAuthHeader(dbUser, c.Request.Header.Get("Authorization"))
	if user.Id == 0 {
		c.JSON(http.StatusForbidden, response{Message: "permission denied"})
		return
	}
	// id is the first priority, then sha256, then md5
	var video interfaces.VideoItem
	if id != 0 {
		video = database.GetVideoById(dbVideo, user.Id, id)
	} else {
		video = database.GetVideoByMd5Sha256(dbVideo, user.Id, md5, sha256)
	}
	if video.Id == 0 {
		c.JSON(http.StatusNotFound, response{
			Message: "video not found",
		})
		return
	}
	url, err := getPresignedCoverImg(minioClient, user, video.Md5, video.Sha256)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response{
			Message: "get presigned url error",
		})
		return
	}
	c.JSON(http.StatusOK, response{
		Video:             video,
		PresignedCoverImg: url,
		Message:           "success",
	})
}
