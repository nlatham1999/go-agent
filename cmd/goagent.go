package goagent

import "net/http"

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Send an HTTP 200 response for any request to the health check endpoint
	w.WriteHeader(http.StatusOK)
}

func main() {
	// Register the health check handler with the path "/health"
	http.HandleFunc("/health", healthCheckHandler)

	// Start the HTTP server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
