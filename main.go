package main

import (
	"coronalive/routes"
	"log"
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Fatal(http.ListenAndServe(":8000", routes.Router))
}
