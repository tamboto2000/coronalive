package main

import (
	"coronalive/routes"
	"log"
	"net/http"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, routes.Router))
}
