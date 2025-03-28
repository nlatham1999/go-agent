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
	models map[string]ModelInterface

	currentModel        ModelInterface
	currentModelWidgets map[string]Widget

	funcMutext      sync.Mutex // Mutex for when we are running a model function
	simulationSpeed time.Duration

	goRepeatRunning bool
	stopRepeating   chan struct{}
	goRepeatMutex   sync.Mutex // Mutex for the goRepeatHandler
	settings        ApiSettings
	concurrentCall  bool

	stepData    map[int]*Model
	oldestValue int

	tickValue int //used for loading the model frontend
}

type ApiSettings struct {
	ButtonTitles       map[string]string
	ButtonDescriptions map[string]string
	StoreSteps         bool // Whether to store steps
	MaxSteps           int  // Maximum number of steps to store. Default is 1000
}

func NewApi(models map[string]ModelInterface, settings ApiSettings) (*Api, error) {

	// make sure there isn't a model with empty string key
	if _, ok := models[""]; ok {
		return nil, fmt.Errorf("Model with empty string key")
	}

	if settings.StoreSteps && settings.MaxSteps == 0 {
		settings.MaxSteps = 1000 // Default value
	}

	return &Api{
		models:          models,
		simulationSpeed: 100 * time.Millisecond,
		settings:        settings,
		stepData:        map[int]*Model{},
	}, nil
}

func (a *Api) Serve() {

	// init all the models
	for _, model := range a.models {
		model.Init()
	}

	r := mux.NewRouter()

	r.HandleFunc("/", a.HomeHandler)

	//handler by name in the url
	r.HandleFunc("/run/{model}", a.ModelPageHandler)
	r.HandleFunc("/health", a.healthCheckHandler)
	r.HandleFunc("/setup", a.setUpHandler).Methods("POST")
	r.HandleFunc("/go", a.goHandler).Methods("POST")
	r.HandleFunc("/gorepeat", a.goRepeatHandler).Methods("POST")
	r.HandleFunc("/model", a.modelHandler)
	r.HandleFunc("/modelat", a.modelAtHandler)

	//frontend handlers
	r.HandleFunc("/loadstats", a.loadStatsHandler)
	r.HandleFunc("/updatespeed", a.updateSpeedHandler)
	r.HandleFunc("/updatedynamic", a.updateDynamicVariableHandler)
	r.HandleFunc("/settick", a.setTickValueHandler)

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
