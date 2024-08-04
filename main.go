package main

import (
	"net/http"

	antpath "github.com/nlatham1999/go-agent/examples/ant-path"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Send an HTTP 200 response for any request to the health check endpoint
	w.WriteHeader(http.StatusOK)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {

	antpath.SetUp()

	antpath.Go()
}
