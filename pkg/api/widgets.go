package api

import "fmt"

type Widget struct {
	PrettyName         string   `json:"prettyName"`
	TargetVariable     string   `json:"targetVariable"`
	WidgetType         string   `json:"widgetType"`
	WidgetValueType    string   `json:"widgetValueType"`
	MinValue           string   `json:"minValue"`
	MaxValue           string   `json:"maxValue"`
	DefaultValue       string   `json:"defaultValue"`
	StepAmount         string   `json:"stepAmount"`
	Target             func()   `json:"target"` // this is a function that will be called when the widget is interacted with if the type is a button
	ValuePointerInt    *int     `json:"valuePointerInt"`
	ValuePointerFloat  *float64 `json:"valuePointerFloat"`
	ValuePointerString *string  `json:"valuePointerString"`
}

func NewFloatSliderWidget(prettyName, targetVariable, minValue, maxValue, defaultValue, stepAmount string, valuePointer *float64) Widget {
	return Widget{
		PrettyName:        prettyName,
		TargetVariable:    targetVariable,
		WidgetType:        "slider",
		WidgetValueType:   "float",
		MinValue:          minValue,
		MaxValue:          maxValue,
		DefaultValue:      defaultValue,
		StepAmount:        stepAmount,
		ValuePointerFloat: valuePointer,
	}
}

func (w *Widget) render(offset int) string {

	id := w.TargetVariable + "-widget"

	html := "<div class='widget' style='top:" + fmt.Sprintf("%d", offset*65) + "px;'>"
	//label for id
	if w.WidgetType == "text" {
		html += `<label for="` + id + `">` + w.PrettyName + `</label>`
		//input type text with id and dynamic name (TargetVariable as the name)
		html += `<input type="text" id="` + id + `" name="` + w.TargetVariable + `" hx-get="/updatedynamic" hx-trigger="change" hx-include="#` + id + `">`
	}
	if w.WidgetType == "slider" {
		labelId := id + "-label"
		html += `<div>`
		html += `<label for="` + id + `">` + w.PrettyName + `: <span id="` + labelId + `">` + w.DefaultValue + `</span></label>`
		html += `</div>`
		//input type range with id and dynamic name (TargetVariable as the name)
		html += `<input type="range" id="` + id + `" name="` + w.TargetVariable + `" 
		min="` + w.MinValue + `" max="` + w.MaxValue + `" value="` + w.DefaultValue + `"`
		if w.StepAmount != "" {
			html += `step="` + w.StepAmount + `"`
		}
		html += `hx-get="/updatedynamic" hx-trigger="change" hx-include="#` + id + `"
		oninput="document.getElementById('` + labelId + `').innerText = this.value;">`
	}
	if w.WidgetType == "button" {
		html += `<button id="` + id + `" hx-swap="none" hx-get="/updatedynamic" hx-trigger="click" hx-vals='{"` + w.TargetVariable + `": "test"}'>` + w.PrettyName + `</button>`
	}

	html += "</div>"

	return html
}
