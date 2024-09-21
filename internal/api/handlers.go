package api

import (
	"encoding/json"
	"net/http"
)

func (a *Api) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (a *Api) setUpHandler(w http.ResponseWriter, r *http.Request) {
	a.Sim.SetUp()

	w.WriteHeader(http.StatusOK)
}

func (a *Api) goHandler(w http.ResponseWriter, r *http.Request) {
	a.Sim.Go()
	w.WriteHeader(http.StatusOK)
}

func (a *Api) modelHandler(w http.ResponseWriter, r *http.Request) {
	model := convertModelToApiModel(a.Sim.Model())

	//return the model as json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}
