// server/server.go
package server

import (
	"log"
	"net/http"

	"test8989/internal/router"
)

func Start(port string, host string) {
	router.Init()

	log.Println("Server started on port: ", port)
	log.Fatal(http.ListenAndServe(host+":"+port, router.Router))
}
