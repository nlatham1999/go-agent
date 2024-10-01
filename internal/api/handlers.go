package api

import (
	"encoding/json"
	"net/http"
	"time"
)

func (a *Api) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (a *Api) setUpHandler(w http.ResponseWriter, r *http.Request) {
	a.funcMutext.Lock()
	defer a.funcMutext.Unlock()

	a.Sim.SetUp()

	w.WriteHeader(http.StatusOK)
}

func (a *Api) goHandler(w http.ResponseWriter, r *http.Request) {
	a.Sim.Go()
	w.WriteHeader(http.StatusOK)
}

func (a *Api) goRepeatHandler(w http.ResponseWriter, r *http.Request) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.running {
		// Stop the loop
		a.stop <- struct{}{}
		a.running = false
		w.Write([]byte("Simulation stopped"))
	} else {
		// Start the loop
		a.stop = make(chan struct{})
		a.running = true
		go a.loop()
		w.Write([]byte("Simulation started"))
	}

	w.WriteHeader(http.StatusOK)
}

func (a *Api) loop() {
	for {
		a.funcMutext.Lock()
		select {
		case <-a.stop:
			a.funcMutext.Unlock()
			return
		default:
			if a.Sim.Stop() {
				a.running = false
				a.funcMutext.Unlock()
				return
			}
			a.Sim.Go()
			time.Sleep(10 * time.Millisecond) // Simulate some work
		}
		a.funcMutext.Unlock()
	}
}

func (a *Api) modelHandler(w http.ResponseWriter, r *http.Request) {
	model := convertModelToApiModel(a.Sim.Model())

	//return the model as json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}
