package main

import (
	"log"
	"net/http"

	"github.com/ardith/golang_101/handlers"
)

func main() {
	router := chi.NewRouter()
	router.Get("/api/jobs", handlers.GetJobs)
	//run it on port 8080
	err := http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
