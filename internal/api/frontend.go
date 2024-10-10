package api

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed html/index.html
var indexHTML string

var (
	statsKeys = []string{}
)

func (a *Api) getFrontend(width int, height int, model *Model) string {

	patchSize := a.getPatchSize(width, height, model.WorldWidth, model.WorldHeight)

	var tmpl strings.Builder
	tmpl.WriteString(`<div class="grid-container patch-grid" style="position: absolute; left: 50%; top: 5px;">`)

	// Render patches
	for _, patch := range model.Patches {
		a.renderPatch(&tmpl, patch, model, patchSize)
	}

	// Render turtles
	turtleOffset := patchSize / 2
	for _, turtle := range model.Turtles {
		a.renderTurtle(&tmpl, turtle, model, patchSize, turtleOffset)
	}

	// Render links
	for _, link := range model.Links {
		a.renderLink(&tmpl, link, model, patchSize, turtleOffset)
	}

	tmpl.WriteString(`</div>`)

	// Render stats
	stats := a.Model.Stats()
	if len(statsKeys) == 0 {
		for key := range stats {
			statsKeys = append(statsKeys, key)
		}
	}
	tmpl.WriteString(`<div id="statsContainer">`)
	tmpl.WriteString(`<div>Ticks: ` + fmt.Sprintf("%d", model.Ticks) + `</div>`)
	for _, key := range statsKeys {
		tmpl.WriteString(`<div>` + key + `: ` + fmt.Sprintf("%v", stats[key]) + `</div>`)
	}
	tmpl.WriteString(`</div>`)

	return tmpl.String()
}

func (a *Api) renderLink(tmpl *strings.Builder, link Link, model *Model, patchSize int, turtleOffset int) {

	relative1X := link.End1X - float64(model.MinPxCor)
	relative1Y := link.End1Y - float64(model.MinPyCor)
	relative2X := link.End2X - float64(model.MinPxCor)
	relative2Y := link.End2Y - float64(model.MinPyCor)
	point1X := relative1X*float64(patchSize) + float64(turtleOffset)
	point1Y := relative1Y*float64(patchSize) + float64(turtleOffset)
	point2X := relative2X*float64(patchSize) + float64(turtleOffset)
	point2Y := relative2Y*float64(patchSize) + float64(turtleOffset)

	// Calculate the distance between the two points (line length)
	distance := math.Sqrt(math.Pow(point2X-point1X, 2) + math.Pow(point2Y-point1Y, 2))

	// Calculate the angle of the line in degrees
	angle := math.Atan2(point2Y-point1Y, point2X-point1X) * (180 / math.Pi)

	// Create the line div
	tmpl.WriteString(fmt.Sprintf(`
			<div 
				class="line" 
				style="
					width: %fpx;
					height: 2px; /* Line thickness */
					position: absolute;
					left: %fpx;
					top: %fpx;
					transform: rotate(%fdeg);
					transform-origin: 0 0; /* Ensure the rotation starts from the first point */
					background-color: black; /* Line color */
				"
			></div>
			`, distance, point1X, point1Y, angle))

}

func (a *Api) renderTurtle(tmpl *strings.Builder, turtle Turtle, model *Model, patchSize int, turtleOffset int) {
	turtleSize := int(float64(patchSize) * turtle.Size)
	if turtleSize < 1 {
		turtleSize = 1
	}
	relativeX := turtle.X - float64(model.MinPxCor)
	relativeY := turtle.Y - float64(model.MinPyCor)
	tmpl.WriteString(fmt.Sprintf(`
			<div 
				class="turtle" 
				style="
					width: %d;
					height: %d;
					left: %vpx; 
					top: %vpx;
					background-color: rgba(%d,%d,%d,%d);
					color: rgba(%d,%d,%d,%d);
					"
			>
			`+fmt.Sprintf("%v", turtle.Label)+`
			</div>
		`, turtleSize, turtleSize, relativeX*float64(patchSize)+float64(turtleOffset), relativeY*float64(patchSize)+float64(turtleOffset),
		turtle.Color.Red, turtle.Color.Green, turtle.Color.Blue, turtle.Color.Alpha,
		turtle.LabelColor.Red, turtle.LabelColor.Green, turtle.LabelColor.Blue, turtle.LabelColor.Alpha,
	))
}

func (a *Api) renderPatch(tmpl *strings.Builder, patch Patch, model *Model, patchSize int) {
	relativeX := patch.X - model.MinPxCor
	relativeY := patch.Y - model.MinPyCor
	tmpl.WriteString(fmt.Sprintf(`
			<div 
				class="patch" 
				style="
					width: %dpx;
					height: %dpx;
					left: %dpx; 
					top: %dpx;
					background-color: rgba(%d,%d,%d,%d);
				"
			>
			</div>
		`, patchSize, patchSize, relativeX*patchSize, relativeY*patchSize,
		patch.Color.Red, patch.Color.Green, patch.Color.Blue, patch.Color.Alpha,
	))
}

func (a *Api) getPatchSize(screenWidth int, screenHeight int, worldWidth int, worldHeight int) int {

	// add 1 to worldWidth and worldHeight to account for the extra .5 on each side
	worldWidth++
	worldHeight++

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

func (a *Api) buildWidgets() string {
	html := `<div id="widgetContainer">`

	// Add widgets here
	for _, widget := range a.Model.Widgets() {
		html += widget.Render()
	}

	html += `</div>`
	return html
}
