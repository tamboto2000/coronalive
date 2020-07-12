package main

import (
	"coronalive/routes"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gorilla/handlers"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//get Heroku port
	port := os.Getenv("PORT")

	//CORS
	headersCORS := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Connection", "Accept", "Upgrade-Insecure-Requests"})
	originsCORS := handlers.AllowedOrigins([]string{"*"})
	methodsCORS := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(originsCORS, headersCORS, methodsCORS)(routes.Router)))
}
