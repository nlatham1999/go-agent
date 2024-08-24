package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nlatham1999/go-agent/internal/api"
)

// antpath "github.com/nlatham1999/go-agent/examples/ant-path"

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", api.HomeHandler)
	r.HandleFunc("/health", api.HealthCheckHandler)
	r.HandleFunc("/setup", api.SetUpHandler)
	r.HandleFunc("/go", api.GoHandler)

	// antpath.SetUp()

	// antpath.Go()

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080", // Address and port for the server to listen on
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server on :8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
