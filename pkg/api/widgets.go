package api

import "fmt"

type Widget struct {
	PrettyName         string   `json:"prettyName"`
	Id                 string   `json:"targetVariable"`
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
	ValuePointerBool   *bool    `json:"valuePointerBool"`
}

func NewFloatSliderWidget(prettyName, id, minValue, maxValue, defaultValue, stepAmount string, valuePointer *float64) Widget {
	return Widget{
		PrettyName:        prettyName,
		Id:                id,
		WidgetType:        "slider",
		WidgetValueType:   "float",
		MinValue:          minValue,
		MaxValue:          maxValue,
		DefaultValue:      defaultValue,
		StepAmount:        stepAmount,
		ValuePointerFloat: valuePointer,
	}
}

func NewIntSliderWidget(prettyName, id, minValue, maxValue, defaultValue, stepAmount string, valuePointer *int) Widget {
	return Widget{
		PrettyName:      prettyName,
		Id:              id,
		WidgetType:      "slider",
		WidgetValueType: "int",
		MinValue:        minValue,
		MaxValue:        maxValue,
		DefaultValue:    defaultValue,
		StepAmount:      stepAmount,
		ValuePointerInt: valuePointer,
	}
}

func NewButtonWidget(prettyName, id string, target func()) Widget {
	return Widget{
		PrettyName: prettyName,
		Id:         id,
		WidgetType: "button",
		Target:     target,
	}
}

func NewMouseXClickedHook(valuePointer *float64) Widget {
	return Widget{
		WidgetType:        "background",
		Id:                "mouse-x-clicked",
		WidgetValueType:   "float",
		ValuePointerFloat: valuePointer,
	}
}

func NewMouseYClickedHook(valuePointer *float64) Widget {
	return Widget{
		WidgetType:        "background",
		Id:                "mouse-y-clicked",
		WidgetValueType:   "float",
		ValuePointerFloat: valuePointer,
	}
}

func NewMouseClickedHook(valuePointer *bool) Widget {
	return Widget{
		WidgetType:       "background",
		Id:               "mouse-clicked",
		WidgetValueType:  "bool",
		ValuePointerBool: valuePointer,
	}
}

func NewMouseXMovedHook(valuePointer *float64) Widget {
	return Widget{
		WidgetType:        "background",
		Id:                "mouse-x-moved",
		WidgetValueType:   "float",
		ValuePointerFloat: valuePointer,
	}
}

func NewMouseYMovedHook(valuePointer *float64) Widget {
	return Widget{
		WidgetType:        "background",
		Id:                "mouse-y-moved",
		WidgetValueType:   "float",
		ValuePointerFloat: valuePointer,
	}
}

func NewMouseMovedHook(valuePointer *bool) Widget {
	return Widget{
		WidgetType:       "background",
		Id:               "mouse-moved",
		WidgetValueType:  "bool",
		ValuePointerBool: valuePointer,
	}
}

func (w *Widget) render(offset int) string {

	if w.WidgetType == "background" {
		return ""
	}

	id := w.Id + "-widget"

	html := "<div class='widget widget-" + w.WidgetType + "' style='top:" + fmt.Sprintf("%d", offset*65) + "px;'>"
	//label for id
	if w.WidgetType == "text" {
		html += `<label for="` + id + `">` + w.PrettyName + `</label>`
		//input type text with id and dynamic name (TargetVariable as the name)
		html += `<input type="text" id="` + id + `" name="` + w.Id + `" hx-get="/updatedynamic" hx-trigger="change" hx-include="#` + id + `">`
	}
	if w.WidgetType == "slider" {

		initialValue := w.DefaultValue
		if w.ValuePointerInt != nil {
			initialValue = fmt.Sprintf("%d", *w.ValuePointerInt)
		}
		if w.ValuePointerFloat != nil {
			initialValue = fmt.Sprintf("%f", *w.ValuePointerFloat)
			// remove trailing zeros
			for initialValue[len(initialValue)-1] == '0' {
				initialValue = initialValue[:len(initialValue)-1]
			}
		}

		labelId := id + "-label"
		html += `<div>`
		html += `<label for="` + id + `">` + w.PrettyName + `: <span id="` + labelId + `">` + initialValue + `</span></label>`
		html += `</div>`
		//input type range with id and dynamic name (TargetVariable as the name)
		html += `<input type="range" id="` + id + `" name="` + w.Id + `" 
		min="` + w.MinValue + `" max="` + w.MaxValue + `" value="` + initialValue + `"`
		if w.StepAmount != "" {
			html += `step="` + w.StepAmount + `"`
		}
		html += `hx-get="/updatedynamic" hx-trigger="change" hx-include="#` + id + `"
		oninput="document.getElementById('` + labelId + `').innerText = this.value;">`
	}
	if w.WidgetType == "button" {
		html += `<button id="` + id + `" hx-swap="none" hx-get="/updatedynamic" hx-trigger="click" hx-vals='{"` + w.Id + `": "test"}'>` + w.PrettyName + `</button>`
	}

	html += "</div>"

	return html
}
