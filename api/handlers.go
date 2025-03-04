package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func (a *Api) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (a *Api) setUpHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Setting up model")

	if a.goRepeatRunning {
		a.stopRepeating <- struct{}{}
		a.goRepeatRunning = false
	}

	a.funcMutext.Lock()
	defer a.funcMutext.Unlock()

	if a.concurrentCall {
		http.Error(w, "concurrent call", http.StatusInternalServerError)
		return
	}
	a.concurrentCall = true

	a.tickValue = -1
	a.stepData = map[int]*Model{}
	a.oldestValue = 0

	err := a.Model.SetUp()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	a.concurrentCall = false

}

func (a *Api) goHandler(w http.ResponseWriter, r *http.Request) {
	a.funcMutext.Lock()
	defer a.funcMutext.Unlock()

	if a.concurrentCall {
		http.Error(w, "concurrent call", http.StatusInternalServerError)
		return
	}
	a.concurrentCall = true

	a.tickValue = -1

	a.Model.Go()
	a.storeStepData()
	w.WriteHeader(http.StatusOK)

	a.concurrentCall = false
}

func (a *Api) goRepeatHandler(w http.ResponseWriter, r *http.Request) {
	a.goRepeatMutex.Lock()
	defer a.goRepeatMutex.Unlock()

	if a.concurrentCall {
		http.Error(w, "concurrent call", http.StatusInternalServerError)
		return
	}
	a.concurrentCall = true

	if a.goRepeatRunning {
		// Stop the loop
		a.stopRepeating <- struct{}{}
		a.goRepeatRunning = false
	} else {
		// Start the loop
		a.stopRepeating = make(chan struct{})
		a.goRepeatRunning = true
		a.tickValue = -1
		go a.loop()
	}

	w.WriteHeader(http.StatusOK)

	a.concurrentCall = false
}

func (a *Api) storeStepData() {
	if a.settings.StoreSteps {
		// Store the model
		model := convertModelToApiModel(a.Model.Model())
		a.stepData[model.Ticks] = model

		// Remove the oldest step if we have reached the maximum number of steps
		if len(a.stepData) > a.settings.MaxSteps {
			found := false
			for !found {
				if _, ok := a.stepData[a.oldestValue]; ok {
					delete(a.stepData, a.oldestValue)
					a.oldestValue++
					found = true
				} else {
					a.oldestValue++
				}
			}
		}
	}
}

func (a *Api) loop() {
	for {
		a.funcMutext.Lock()
		select {
		case <-a.stopRepeating:
			a.funcMutext.Unlock()
			return
		default:
			if a.Model.Stop() {
				a.goRepeatRunning = false
				a.funcMutext.Unlock()
				return
			}
			a.Model.Go()
			a.storeStepData()
			time.Sleep(a.simulationSpeed) // Simulate some work
		}
		a.funcMutext.Unlock()
	}
}

func (a *Api) HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("content").Parse(indexHTML)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":   "Go Agent",
		"Widgets": a.buildWidgets(),
	}

	tmpl.Execute(w, data)
}

func (a *Api) loadHandler(w http.ResponseWriter, r *http.Request) {

	a.funcMutext.Lock()
	defer a.funcMutext.Unlock()

	// queryParams := r.URL.Query()

	// Get the 'width' and 'height' parameters from the query string
	// widthStr := queryParams.Get("width")
	// heightStr := queryParams.Get("height")
	// width, err := strconv.Atoi(widthStr)
	// if err != nil {
	// 	http.Error(w, "Invalid width parameter", http.StatusBadRequest)
	// 	return
	// }

	// height, err := strconv.Atoi(heightStr)
	// if err != nil {
	// 	http.Error(w, "Invalid height parameter", http.StatusBadRequest)
	// 	return
	// }

	// either load the model at the current tick or at the tick stored
	var model *Model
	var ok bool
	if a.tickValue != -1 {
		model, ok = a.stepData[a.tickValue]
		if !ok {
			model = convertModelToApiModel(a.Model.Model())
		}
	} else {
		model = convertModelToApiModel(a.Model.Model())
	}

	//return the model as json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)

	// // Get the HTML template for rendering
	// tmpl := a.getFrontend(width, height, model)

	// // Execute the template
	// _, err = w.Write([]byte(tmpl))
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

func (a *Api) loadStatsHandler(w http.ResponseWriter, r *http.Request) {

	a.funcMutext.Lock()
	defer a.funcMutext.Unlock()

	tmpl := a.renderStats()

	// Execute the template
	_, err := w.Write([]byte(tmpl))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *Api) updateSpeedHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	// Get the 'speed' parameter from the query string
	speedStr := queryParams.Get("speed")
	speed, err := strconv.Atoi(speedStr)
	if err != nil {
		http.Error(w, "Invalid speed parameter", http.StatusBadRequest)
		return
	}

	// Update the speed
	a.simulationSpeed = time.Duration(speed) * time.Millisecond

	w.WriteHeader(http.StatusOK)
}

func (a *Api) modelHandler(w http.ResponseWriter, r *http.Request) {
	a.funcMutext.Lock()
	defer a.funcMutext.Unlock()

	model := convertModelToApiModel(a.Model.Model())

	//return the model as json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}

func (a *Api) modelAtHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	step := queryParams.Get("step")
	stepInt, err := strconv.Atoi(step)
	if err != nil {
		http.Error(w, "Invalid step parameter", http.StatusBadRequest)
		return
	}

	model, ok := a.stepData[stepInt]
	if !ok {
		http.Error(w, "Step not found", http.StatusNotFound)
		return
	}

	//return the model as json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}

func (a *Api) updateDynamicVariableHandler(w http.ResponseWriter, r *http.Request) {

	a.funcMutext.Lock()
	defer a.funcMutext.Unlock()

	// Parse the query parameters
	queryParams := r.URL.Query()

	// Loop through all query parameters
	for name, values := range queryParams {
		// Assuming there's only one value per query parameter (HTMX serializes like this)
		var value string
		if len(values) > 0 {
			value = values[0]
		}

		fmt.Println("Updating dynamic variable", name, "with value", value)

		// go through widgets and update the dynamic variable
		for _, widget := range a.Model.Widgets() {
			if widget.TargetVariable == name {

				// If the widget is a button, just call the target function
				if widget.WidgetType == "button" {
					widget.Target()
					continue
				}

				// Update the dynamic variable
				if widget.WidgetValueType == "int" {
					intValue, err := strconv.Atoi(value)
					if err != nil {
						http.Error(w, "Invalid value for dynamic variable", http.StatusBadRequest)
					}
					if widget.ValuePointerInt == nil {
						http.Error(w, "Invalid value pointer for dynamic variable", http.StatusBadRequest)
						continue
					}
					*widget.ValuePointerInt = intValue
				} else if widget.WidgetValueType == "float" {
					floatValue, err := strconv.ParseFloat(value, 64)
					if err != nil {
						http.Error(w, "Invalid value for dynamic variable", http.StatusBadRequest)
						continue
					}
					if widget.ValuePointerFloat == nil {
						http.Error(w, "Invalid value pointer for dynamic variable", http.StatusBadRequest)
						continue
					}
					*widget.ValuePointerFloat = floatValue
				} else {
					if widget.ValuePointerString == nil {
						http.Error(w, "Invalid value pointer for dynamic variable", http.StatusBadRequest)
						continue
					}
					*widget.ValuePointerString = value
				}
			}
		}

		// Here you can process the variable name and value dynamically, e.g., store them, respond, etc.
	}

	// Respond to the client (for demonstration purposes)
	w.WriteHeader(http.StatusOK)
}

func (a *Api) setTickValueHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	tickValue := queryParams.Get("tick")
	tick, err := strconv.Atoi(tickValue)
	if err != nil {
		http.Error(w, "invalid tick value", http.StatusBadRequest)
	}

	a.tickValue = tick
	w.WriteHeader(http.StatusOK)
}
