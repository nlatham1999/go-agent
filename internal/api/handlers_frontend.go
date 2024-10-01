package api

import (
	_ "embed"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

//go:embed html/index.html
var indexHTML string

var (
	statsKeys = []string{}
)

func (a *Api) HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("content").Parse(indexHTML)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title": "Go Agent",
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

	model := convertModelToApiModel(a.Sim.Model())

	patchSize := getPatchSize(width, height, model.Width, model.Height)
	turtleSize := patchSize - 2
	if turtleSize < 1 {
		turtleSize = 1
	}

	// Start the HTML template for rendering
	tmpl := `<div class="grid-container patch-grid" style="position: absolute; left: 50%; top: 5px;">`

	// Render patches
	for _, patch := range model.Patches {
		relativeX := patch.X + 16
		relativeY := patch.Y + 16
		tmpl += `
			<div 
				class="patch" 
				style="
					width:` + fmt.Sprintf("%dpx", patchSize) + `;
					height:` + fmt.Sprintf("%dpx", patchSize) + `;
					left:` + fmt.Sprintf("%dpx", relativeX*patchSize) + `; 
					top:` + fmt.Sprintf("%dpx", relativeY*patchSize) + `;
					background-color: rgba(` + fmt.Sprintf("%d", patch.Color.R) + `, ` + fmt.Sprintf("%d", patch.Color.G) + `, ` + fmt.Sprintf("%d", patch.Color.B) + `, ` + fmt.Sprintf("%d", patch.Color.A) + `);
				"
			>
			</div>
		`
	}

	turtleOffset := patchSize / 2
	// Render turtles
	for _, turtle := range model.Turtles {
		relativeX := turtle.X + 16
		relativeY := turtle.Y + 16
		tmpl += `
			<div 
				class="turtle" 
				style="
					width:` + fmt.Sprintf("%dpx", turtleSize) + `;
					height:` + fmt.Sprintf("%dpx", turtleSize) + `;
					left:` + fmt.Sprintf("%vpx", relativeX*float64(patchSize)+float64(turtleOffset)) + `; 
					top:` + fmt.Sprintf("%vpx", relativeY*float64(patchSize)+float64(turtleOffset)) + `;
					background-color: rgba(` + fmt.Sprintf("%d", turtle.Color.R) + `, ` + fmt.Sprintf("%d", turtle.Color.G) + `, ` + fmt.Sprintf("%d", turtle.Color.B) + `, ` + fmt.Sprintf("%d", turtle.Color.A) + `);
					"
			>
			</div>
		`
	}
	tmpl += `</div>`

	stats := a.Sim.Stats()
	if len(statsKeys) == 0 {
		for key := range stats {
			statsKeys = append(statsKeys, key)
		}
	}

	tmpl += `<div style="position: absolute; left: 5px; top: 20%;">`
	tmpl += `<div>Ticks: ` + fmt.Sprintf("%d", model.Ticks) + `</div>`
	for _, key := range statsKeys {
		tmpl += `<div>` + key + `: ` + fmt.Sprintf("%v", stats[key]) + `</div>`
	}
	tmpl += `</div>`

	// Execute the template
	_, err = w.Write([]byte(tmpl))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getPatchSize(screenWidth int, screenHeight int, worldWidth int, worldHeight int) int {
	//can only take up 50% of the width of the screen
	screenWidth = screenWidth / 2

	// leave a 1% margin on the left and right
	screenWidth = screenWidth - (screenWidth / 50)

	// leave a 1% margin on the top and bottom
	screenHeight = screenHeight - (screenHeight / 50)

	//calculate the max width of the patches
	maxPatchWidth := screenWidth / worldWidth

	//calculate the max height of the patches
	maxPatchHeight := screenHeight / worldHeight

	//return the minimum of the two
	if maxPatchWidth < maxPatchHeight {
		return maxPatchWidth
	}
	return maxPatchHeight
}
