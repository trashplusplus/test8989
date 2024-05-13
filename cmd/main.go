package main

import (
	"log"
	"os"
	"test8989/internal/server"

	"github.com/joho/godotenv"
)

func readConf() (string, string) {

	err := godotenv.Load("../configs/.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	return port, host

}

func main() {
	server.Start(readConf())
}
