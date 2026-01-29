package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

func (a *Api) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (a *Api) setUpHandler(w http.ResponseWriter, r *http.Request) {

	// if we are currently running a goRepeat, stop it
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

	err := a.currentModel.SetUp()
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

	a.currentModel.Go()
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
		model := convertModelToApiModel(a.currentModel.Model())
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
			if a.currentModel.Stop() {
				a.goRepeatRunning = false
				a.funcMutext.Unlock()
				return
			}
			a.currentModel.Go()
			a.storeStepData()
			time.Sleep(a.simulationSpeed) // Simulate some work
		}
		a.funcMutext.Unlock()
	}
}

func (a *Api) HomeHandler(w http.ResponseWriter, r *http.Request) {
	a.currentModel = nil

	htmlTmpl, err := template.New("content").Parse(homePageHtml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"ModelList": a.buildModelList(),
	}
	htmlTmpl.Execute(w, data)

	styleTmpl, err := template.New("content").Parse(homePageStyle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	styleTmpl.Execute(w, nil)
}

func (a *Api) ModelPageHandler(w http.ResponseWriter, r *http.Request) {

	// get the model name from the url
	vars := mux.Vars(r)
	modelName := vars["model"]

	a.currentModel = a.models[modelName]
	if a.currentModel == nil {
		http.Error(w, "Model not found", http.StatusNotFound)
		return
	}

	a.currentModelWidgets = map[string]Widget{}
	for _, widget := range a.currentModel.Widgets() {
		a.currentModelWidgets[widget.Id] = widget
	}

	tmpl, err := template.New("content").Parse(modelPageHtml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)

	// load the threejs html as a string
	jsTml, err := template.New("content").Parse(modelPageThreeJS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsTml.Execute(w, nil)

	// load the scripts
	scriptsTmpl, err := template.New("content").Parse(modelPageScripts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	scriptsTmpl.Execute(w, nil)

	// load the style
	styleTmpl, err := template.New("content").Parse(modelPageStyle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	styleTmpl.Execute(w, nil)
}

func (a *Api) loadStatsHandler(w http.ResponseWriter, r *http.Request) {

	a.funcMutext.Lock()
	defer a.funcMutext.Unlock()

	if a.currentModel == nil {
		http.Error(w, "Model not instantiated", http.StatusNotFound)
		return
	}

	// get the stats
	stats := a.currentModel.Stats()

	if stats == nil {
		stats = map[string]interface{}{}
	}

	// add in the tick
	stats["ticks"] = a.currentModel.Model().Ticks

	//return the stats as json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stats)
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
	speed = 100 - speed

	// Update the speed
	a.simulationSpeed = time.Duration(speed) * time.Millisecond

	w.WriteHeader(http.StatusOK)
}

func (a *Api) modelHandler(w http.ResponseWriter, r *http.Request) {
	a.funcMutext.Lock()
	defer a.funcMutext.Unlock()

	if a.currentModel == nil {
		http.Error(w, "Model not instantiated", http.StatusNotFound)
		return
	}

	model := convertModelToApiModel(a.currentModel.Model())

	//return the model as json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}

func (a *Api) modelAtHandler(w http.ResponseWriter, r *http.Request) {

	if a.currentModel == nil {
		http.Error(w, "Model not instantiated", http.StatusNotFound)
		return
	}

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

		widget, found := a.currentModelWidgets[name]
		if !found {
			fmt.Println("Widget not found", name)
			continue
		}

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
		} else if widget.WidgetValueType == "bool" {
			boolValue, err := strconv.ParseBool(value)
			if err != nil {
				http.Error(w, "Invalid value for dynamic variable", http.StatusBadRequest)
				continue
			}
			if widget.ValuePointerBool == nil {
				http.Error(w, "Invalid value pointer for dynamic variable", http.StatusBadRequest)
				continue
			}
			*widget.ValuePointerBool = boolValue
		} else {
			if widget.ValuePointerString == nil {
				http.Error(w, "Invalid value pointer for dynamic variable", http.StatusBadRequest)
				continue
			}
			*widget.ValuePointerString = value
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

func (a *Api) widgetsHandler(w http.ResponseWriter, r *http.Request) {
	a.funcMutext.Lock()
	defer a.funcMutext.Unlock()

	if a.currentModel == nil {
		http.Error(w, "Model not instantiated", http.StatusNotFound)
		return
	}

	widgets := a.currentModel.Widgets()
	stats := a.currentModel.Stats()

	// Create JSON-serializable widget data
	widgetData := make([]map[string]interface{}, 0)

	// Add tick stat as first widget
	tickData := map[string]interface{}{
		"prettyName":      "Ticks",
		"id":              "stats-ticks",
		"widgetType":      "stat",
		"widgetValueType": "int",
		"currentValue":    fmt.Sprintf("%d", a.currentModel.Model().Ticks),
		"index":           0,
	}
	widgetData = append(widgetData, tickData)

	// Add other stats
	index := 1
	for key, value := range stats {
		if value == nil {
			value = "null"
		}

		// if value is type GraphWidget, marshal it differently
		valueStr := ""
		widgetType := "stat"

		// Check for both GraphWidget value and *GraphWidget pointer
		var graphWidget *GraphWidget
		if gw, ok := value.(GraphWidget); ok {
			fmt.Printf("Found GraphWidget (value) for key %s\n", key)
			graphWidget = &gw
		} else if gw, ok := value.(*GraphWidget); ok {
			fmt.Printf("Found GraphWidget (pointer) for key %s\n", key)
			graphWidget = gw
		}

		if graphWidget != nil {
			valueBytes, err := json.Marshal(graphWidget)
			if err != nil {
				http.Error(w, "error marshaling graph widget", http.StatusInternalServerError)
				return
			}
			valueStr = string(valueBytes)
			widgetType = "graph"
			fmt.Printf("Marshaled graph widget for %s: %s\n", key, valueStr)
		} else {
			valueStr = fmt.Sprintf("%v", value)
		}

		statData := map[string]interface{}{
			"prettyName":   key,
			"id":           fmt.Sprintf("stats-%s", key),
			"widgetType":   widgetType,
			"currentValue": valueStr,
			"index":        index,
		}
		widgetData = append(widgetData, statData)
		index++
	}

	// Add model widgets
	for _, widget := range widgets {
		if widget.WidgetType == "background" {
			continue
		}

		data := map[string]interface{}{
			"prettyName":      widget.PrettyName,
			"id":              widget.Id,
			"widgetType":      widget.WidgetType,
			"widgetValueType": widget.WidgetValueType,
			"minValue":        widget.MinValue,
			"maxValue":        widget.MaxValue,
			"defaultValue":    widget.DefaultValue,
			"stepAmount":      widget.StepAmount,
			"currentValue":    widget.getCurrentValue(),
			"index":           index,
		}
		widgetData = append(widgetData, data)
		index++
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(widgetData)
}

// Lightweight handler that only returns widget IDs and current values for syncing
func (a *Api) widgetValuesHandler(w http.ResponseWriter, r *http.Request) {
	a.funcMutext.Lock()
	defer a.funcMutext.Unlock()

	if a.currentModel == nil {
		http.Error(w, "Model not instantiated", http.StatusNotFound)
		return
	}

	widgets := a.currentModel.Widgets()
	stats := a.currentModel.Stats()

	// Create minimal JSON data with only ID and current value
	valueData := make([]map[string]interface{}, 0)

	// Add tick stat
	valueData = append(valueData, map[string]interface{}{
		"id":           "stats-ticks",
		"currentValue": fmt.Sprintf("%d", a.currentModel.Model().Ticks),
		"widgetType":   "stat",
	})

	// Add other stats
	for key, value := range stats {
		if value == nil {
			value = "null"
		}

		valueStr := ""
		widgetType := "stat"

		// Check for GraphWidget types
		var graphWidget *GraphWidget
		if gw, ok := value.(GraphWidget); ok {
			graphWidget = &gw
		} else if gw, ok := value.(*GraphWidget); ok {
			graphWidget = gw
		}

		if graphWidget != nil {
			valueBytes, err := json.Marshal(graphWidget)
			if err != nil {
				http.Error(w, "error marshaling graph widget", http.StatusInternalServerError)
				return
			}
			valueStr = string(valueBytes)
			widgetType = "graph"
		} else {
			valueStr = fmt.Sprintf("%v", value)
		}

		valueData = append(valueData, map[string]interface{}{
			"id":           fmt.Sprintf("stats-%s", key),
			"currentValue": valueStr,
			"widgetType":   widgetType,
		})
	}

	// Add model widgets
	for _, widget := range widgets {
		if widget.WidgetType == "background" {
			continue
		}

		valueData = append(valueData, map[string]interface{}{
			"id":           widget.Id,
			"currentValue": widget.getCurrentValue(),
			"widgetType":   widget.WidgetType,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(valueData)
}
