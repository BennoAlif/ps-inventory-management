package main

import (
	"log"

	"github.com/BennoAlif/ps-cats-social/src/drivers/db"
	"github.com/BennoAlif/ps-cats-social/src/http"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConnection := db.CreateConnection()

	h := http.New(&http.Http{
		DB: dbConnection,
	})

	h.Launch()
}
