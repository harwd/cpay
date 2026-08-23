package main

import (
	"log"
	"net/http"

	"github.com/mlloc/cflow/internal/api"
)

func main() {
	router := api.NewRouter()

	log.Println("starting api on :8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}