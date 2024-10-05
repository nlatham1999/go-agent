package api

type Widget struct {
	PrettyName      string `json:"prettyName"`
	TargetVariable  string `json:"targetVariable"`
	WidgetType      string `json:"widgetType"`
	WidgetValueType string `json:"widgetValueType"`
	MinValue        string `json:"minValue"`
	MaxValue        string `json:"maxValue"`
	DefaultValue    string `json:"defaultValue"`
	StepAmount      string `json:"stepAmount"`
}

func (w *Widget) Render() string {

	id := w.TargetVariable + "-widget"

	html := "<div>"
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
	html += "</div>"
	return html
}
