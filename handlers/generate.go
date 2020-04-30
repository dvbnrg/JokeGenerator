package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/dvbnrg/JokeGenerator/services"
)

func Generate(w http.ResponseWriter, r *http.Request) {

	j, err := services.GetJoke()

	if err != nil {
		log.Printf("joke generate error: %v", err)
	}

	if j.Type == "success" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		buff := bytes.Buffer{}
		json.NewEncoder(&buff).Encode(j)
		w.Write(buff.Bytes())
	}
}
