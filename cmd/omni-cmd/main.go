package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/qiyihuang/omni-cmd/internal/config"
	"github.com/qiyihuang/omni-cmd/server"
)

const version = "0.1.3"

func main() {
	if os.Getenv("ENV") != "production" && os.Getenv("ENV") != "test" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	config.Load()

	log.Println("Omnibox Cmd server up.      Version: " + version)
	log.Fatal(server.Run())
}
