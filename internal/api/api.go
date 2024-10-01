package api

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Api struct {
	Sim ModelInterface

	running    bool
	stop       chan struct{}
	mu         sync.Mutex
	funcMutext sync.Mutex
}

func NewApi(sim ModelInterface) *Api {

	return &Api{
		Sim: sim,
	}
}

func (a *Api) Serve() {

	if a.Sim.Model() == nil {
		fmt.Println("Model is nil")
	}

	a.Sim.Init()

	if a.Sim.Model() == nil {
		fmt.Println("Model is nil")
	}

	r := mux.NewRouter()

	r.HandleFunc("/", a.HomeHandler)
	r.HandleFunc("/health", a.healthCheckHandler)
	r.HandleFunc("/setup", a.setUpHandler).Methods("POST")
	r.HandleFunc("/go", a.goHandler).Methods("POST")
	r.HandleFunc("/gorepeat", a.goRepeatHandler).Methods("POST")
	r.HandleFunc("/model", a.modelHandler)

	//frontend handlers
	r.HandleFunc("/load", a.loadHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080", // Address and port for the server to listen on
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server on http://127.0.0.1:8080/")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
