package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	_ = godotenv.Load("config/.env")
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	router := httprouter.New()
	router.POST("/calculate", PositiveNumbersMiddleware(CalculateHandler))
	log.Println("Server started on port: ", port)
	log.Fatal(http.ListenAndServe(host+":"+port, router))

}
