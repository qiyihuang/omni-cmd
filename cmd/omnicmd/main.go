package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/qiyihuang/omni-cmd/config"
	"github.com/qiyihuang/omni-cmd/server"
)

const version = "0.2.3"

func main() {
	if os.Getenv("ENV") != "production" && os.Getenv("ENV") != "test" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	config.Load()

	log.Println("Omni-cmd server up.      Version: " + version)
	log.Fatal(server.Start())
}
