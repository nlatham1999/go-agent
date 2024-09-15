package api

import (
	"encoding/json"
	"net/http"
)

func (a *Api) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (a *Api) setUpHandler(w http.ResponseWriter, r *http.Request) {
	a.SetupFunc()

	model := convertModelToApiModel(a.ModelFunc())

	//return the model as json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}

func (a *Api) goHandler(w http.ResponseWriter, r *http.Request) {
	a.GoFunc()
	w.WriteHeader(http.StatusOK)
}

func (a *Api) modelHandler(w http.ResponseWriter, r *http.Request) {
	model := convertModelToApiModel(a.ModelFunc())

	//return the model as json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}
