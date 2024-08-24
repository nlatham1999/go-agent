package api

import (
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Health check endpoint hit")
	w.WriteHeader(http.StatusOK)
}

func SetUpHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Set up endpoint hit")
	w.WriteHeader(http.StatusOK)
}

func GoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Go endpoint hit")
	w.WriteHeader(http.StatusOK)
}
