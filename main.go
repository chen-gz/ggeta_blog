package main

import (
	"encoding/json"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go_blog/database"
	hd "go_blog/handler"
	"go_blog/interfaces"
	"log"
	"os"
)

type Config struct {
	BlogDatabase  database.BlogDbConfig  `json:"blog_database"`
	UserDatabase  database.UserDbConfig  `json:"user_database"`
	PhotoDatabase database.PhotoDbConfig `json:"photo_database"`
	VideoDb       interfaces.DbConfig    `json:"video_db"`
}

// read config.json and return Config struct
func ReadConfig() Config {
	var config Config
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Println("Error opening config file:", err)
		return Config{}
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&config); err != nil {
		log.Println("Error decoding JSON:", err)
		return Config{}
	}

	config.UserDatabase.SecreteKey = []byte(os.Getenv("SECRETE_KEY"))

	log.Println(config)
	return config
}

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
	r.POST("/api/v4/login", func(c *gin.Context) {
		hd.V4Login(c, db_user)
	})
	r.POST("/api/v4/verify_token", func(c *gin.Context) {
		hd.V4VerifyToken(c, db_user)
	})

	apiV1 := r.Group("/api")
	apiV1.Use(hd.AuthMiddleware(db_user))
	{
		apiV1.GET("/photo/v2/get_photo_list", func(c *gin.Context) {
			hd.GetPhotoIdsV2(c, db_user, db_photo)
		})
		apiV1.POST("/v4/get_post", func(c *gin.Context) {
			hd.V4GetPost(c, db_user, db_blog)
		})
		apiV1.POST("/v4/update_post", func(c *gin.Context) {
			hd.V4UpdatePost(c, db_user, db_blog)
		})
		apiV1.POST("/v4/new_post", func(c *gin.Context) {
			hd.V4NewPost(c, db_user, db_blog)
		})
		apiV1.POST("/v4/get_distinct", func(c *gin.Context) {
			hd.V4GetDistinct(c, db_user, db_blog)
		})
		apiV1.POST("/post/v5/render", func(c *gin.Context) {
			hd.V5Render(c, db_user)
		})
	}
	r.Run(":2009") // listen and serve on
}

func main() {
	ginServer()
}
