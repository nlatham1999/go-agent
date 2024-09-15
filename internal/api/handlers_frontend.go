package api

import (
	_ "embed"
	"net/http"
	"text/template"
)

//go:embed html/index.html
var indexHTML string

//go:embed html/script.js
var scriptJS string

func (a *Api) HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("content").Parse(indexHTML)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":  "Go Agent",
		"Script": scriptJS,
	}

	tmpl.Execute(w, data)
}

func (a *Api) loadHandler(w http.ResponseWriter, r *http.Request) {
	// Load the partial template with the new content
	tmpl, err := template.New("content").Parse(`<p>This content was loaded dynamically!</p>`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
