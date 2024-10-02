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
			time.Sleep(a.speed) // Simulate some work
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

	queryParams := r.URL.Query()

	// Get the 'width' and 'height' parameters from the query string
	widthStr := queryParams.Get("width")
	heightStr := queryParams.Get("height")
	width, err := strconv.Atoi(widthStr)
	if err != nil {
		http.Error(w, "Invalid width parameter", http.StatusBadRequest)
		return
	}

	height, err := strconv.Atoi(heightStr)
	if err != nil {
		http.Error(w, "Invalid height parameter", http.StatusBadRequest)
		return
	}

	// Get the HTML template for rendering
	tmpl := a.getFrontend(width, height)

	// Execute the template
	_, err = w.Write([]byte(tmpl))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *Api) updateSpeedHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	fmt.Println("Update speed")

	// Get the 'speed' parameter from the query string
	speedStr := queryParams.Get("speed")
	speed, err := strconv.Atoi(speedStr)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid speed parameter", http.StatusBadRequest)
		return
	}

	// Update the speed
	a.speed = time.Duration(speed) * time.Millisecond

	w.WriteHeader(http.StatusOK)
}

func (a *Api) modelHandler(w http.ResponseWriter, r *http.Request) {
	model := convertModelToApiModel(a.Sim.Model())

	//return the model as json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}

func (a *Api) updateDynamicVariableHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters
	queryParams := r.URL.Query()

	// Loop through all query parameters
	for name, values := range queryParams {
		// Assuming there's only one value per query parameter (HTMX serializes like this)
		value := values[0]
		// go through widgets and update the dynamic variable
		for _, widget := range a.Widgets {
			if widget.TargetVariable == name {
				// Update the dynamic variable
				if widget.WidgetValueType == "int" {
					intValue, err := strconv.Atoi(value)
					if err != nil {
						http.Error(w, "Invalid value for dynamic variable", http.StatusBadRequest)
					}
					a.Sim.Model().SetGlobal(name, intValue)
				} else {
					a.Sim.Model().SetGlobal(name, value)
				}
			}
		}

		// Here you can process the variable name and value dynamically, e.g., store them, respond, etc.
	}

	// Respond to the client (for demonstration purposes)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Dynamic variables processed"))
}
