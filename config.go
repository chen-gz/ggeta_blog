package main

import (
	"encoding/json"
	"go_blog/database"
	"go_blog/handler"
	"go_blog/interfaces"
	"log"
	"os"
)

type Config struct {
	BlogDatabase  database.BlogDbConfig  `json:"blog_database"`
	UserDatabase  database.UserDbConfig  `json:"user_database"`
	PhotoDatabase database.PhotoDbConfig `json:"photo_database"`
	Minio         handler.MinioConfig    `json:"minio"`
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

	log.Println(config)
	return config
}
