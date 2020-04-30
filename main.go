package main

import (
	"log"
	"net/http"

	"github.com/dvbnrg/JokeGenerator/routes"
)

func main() {
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":80", router))
}
