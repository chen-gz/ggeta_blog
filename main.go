package main

import (
	"embed"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go_blog/database"
	hd "go_blog/handler"
	"log"
	"net/http"
)

//go:embed front_end/dist/*
var frontend embed.FS

func ginServer() {

	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Authorization", "Content-Length", "Content-Type", "Origin", "Access-Control-Allow-Headers",
			"Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials"},
		AllowCredentials: true,
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://localhost:2009",
			"https://blog.ggeta.com",
			"https://ggeta.com",
			"*",
		},
	}))

	config := ReadConfig()
	log.Println(config.BlogDatabase)
	db_blog := database.InitV4(config.BlogDatabase)
	db_user, _ := database.UserDbInit(config.UserDatabase)
	db_photo, _ := database.InitPhotoDb(config.PhotoDatabase)
	// dbVideo := database.InitVideoDb(config.VideoDb)
	database.InitPhotoTableV2(db_photo, database.User{Id: 2})
	log.Println(config.Minio)
	minio_client := hd.InitMinioClient(config.Minio)
	//photo_minio_client := hd.InitPhotoMinioClient(config.PhotoMinio)
	// videoMinioClient := hd.InitVideoMinioClient(config.VideoMinio)
	r.POST("/api/blog_file/v1/get_presigned_url", func(c *gin.Context) {
		hd.GetPresignedUrl(c, db_user, db_blog, minio_client)
	})
	r.GET("/api/blog_file/v1/get_file_lists/:id", func(c *gin.Context) {
		hd.GetFileList(c, db_user, db_blog)
	})

	r.GET("/api/photo/v2/get_photo_list", func(c *gin.Context) {
		hd.GetPhotoIdsV2(c, db_user, db_photo)
	})
	r.POST("/api/v4/login", func(c *gin.Context) {
		hd.V4Login(c, db_user)
	})
	r.POST("/api/v4/verify_token", func(c *gin.Context) {
		hd.V4VerifyToken(c, db_user)
	})
	r.POST("/api/v4/get_post", func(c *gin.Context) {
		hd.V4GetPost(c, db_user, db_blog)
	})
	r.POST("/api/v4/search_posts", func(c *gin.Context) {
		hd.V4SearchPosts(c, db_user, db_blog)
	})
	r.POST("/api/v4/update_post", func(c *gin.Context) {
		hd.V4UpdatePost(c, db_user, db_blog)
	})
	r.POST("/api/v4/new_post", func(c *gin.Context) {
		hd.V4NewPost(c, db_user, db_blog)
	})
	r.POST("/api/v4/get_distinct", func(c *gin.Context) {
		hd.V4GetDistinct(c, db_user, db_blog)
	})

	r.POST("/api/post/v5/render", func(c *gin.Context) {
		hd.V5Render(c, db_user)
	})
	///////////////////////////////////////////////////////////////////////////////////// v4 api
	r.GET("/assets/*filepath", func(c *gin.Context) {
		//c.FileFromFS("/assets/", frontendBox)
		if data, err := frontend.ReadFile("front_end/dist/assets" + c.Param("filepath")); err == nil {
			if c.Param("filepath")[len(c.Param("filepath"))-3:] == ".js" {
				c.Data(200, "application/javascript", data)
			} else if c.Param("filepath")[len(c.Param("filepath"))-4:] == ".css" {
				c.Data(200, "text/css", data)
			} else if c.Param("filepath")[len(c.Param("filepath"))-4:] == ".svg" {
				c.Data(200, "image/svg+xml", data)
			} else {
				c.Data(200, "application/octet-stream", data)
			}
		} else {
			c.String(404, "File not found")
		}
	})
	///////////////////////////////////////////////////////////////////////////////////// shows api
	r.POST("/api/shows/v1/get_list", func(c *gin.Context) {
		hd.GetList(c, db_user, minio_client)
	})
	r.POST("/api/shows/v1/get_presigned_url", func(c *gin.Context) {
		hd.GetShowPresignedUrl(c, minio_client, db_user)
	})
	r.POST("/api/v1/get_presigned_url", func(c *gin.Context) {
		hd.GetPresignedUrlNew(c, minio_client, db_user)
	})

	r.NoRoute(func(c *gin.Context) {
		c.FileFromFS("front_end/dist/", http.FS(frontend))
	})
	r.Run(":2009") // listen and serve on
}

func main() {
	ginServer()
}
