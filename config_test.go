package main

import (
	"log"
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {
	// read json config file
	// print current environment path
	log.Println("Current environment path: ", os.Getenv("GOPATH"))
	config := ReadConfig()
	log.Println(config.BlogDatabase)
}
