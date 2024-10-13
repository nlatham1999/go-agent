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

	// Render links
	for _, link := range model.Links {
		a.renderLink(&tmpl, link, model, patchSize, float64(height))
	}

	// Render turtles
	for _, turtle := range model.Turtles {
		a.renderTurtle(&tmpl, turtle, model, patchSize)
	}

	tmpl.WriteString(`</div>`)

	return tmpl.String()
}

func (a *Api) renderStats() string {

	var tmpl strings.Builder

	// Render stats
	stats := a.Model.Stats()
	if len(statsKeys) == 0 || len(statsKeys) != len(stats) {
		for key := range stats {
			statsKeys = append(statsKeys, key)
		}
	}
	tmpl.WriteString(`<div>Ticks: ` + fmt.Sprintf("%d", a.Model.Model().Ticks) + `</div>`)
	for _, key := range statsKeys {
		tmpl.WriteString(`<div>` + key + `: ` + fmt.Sprintf("%v", stats[key]) + `</div>`)
	}

	return tmpl.String()
}

func (a *Api) renderLink(tmpl *strings.Builder, link Link, model *Model, patchSize int, screenHeight float64) {

	// turtleSize1 := int(float64(patchSize) * link.End1Size)
	// if turtleSize1 < 1 {
	// 	turtleSize1 = 1
	// }
	// turtleOffset1 := (patchSize - turtleSize1) / 2
	turtleOffset1 := patchSize / 2

	// turtleSize2 := int(float64(patchSize) * link.End2Size)
	// if turtleSize2 < 1 {
	// 	turtleSize2 = 1
	// }
	// turtleOffset2 := (patchSize - turtleSize2) / 2
	turtleOffset2 := patchSize / 2

	// fmt.Println(turtleSize1, turtleSize2, turtleOffset1, turtleOffset2)

	relative1X := link.End1X - float64(model.MinPxCor)
	relative1Y := link.End1Y - float64(model.MinPyCor)
	relative2X := link.End2X - float64(model.MinPxCor)
	relative2Y := link.End2Y - float64(model.MinPyCor)
	point1X := relative1X*float64(patchSize) + float64(turtleOffset1)
	point1Y := relative1Y*float64(patchSize) + float64(turtleOffset1)
	point2X := relative2X*float64(patchSize) + float64(turtleOffset2)
	point2Y := relative2Y*float64(patchSize) + float64(turtleOffset2)

	// Calculate the distance between the two points (line length)
	distance := math.Sqrt(math.Pow(point2X-point1X, 2) + math.Pow(point2Y-point1Y, 2))

	// Calculate the angle of the line in degrees
	angle := math.Atan2(point2Y-point1Y, point2X-point1X) * (180 / math.Pi)

	offset := screenHeight * .01 * 4

	// Create the line div
	tmpl.WriteString(fmt.Sprintf(`
			<div 
				class="line" 
				style="
					width: %fpx;
					height: %dpx; /* Line thickness */
					position: absolute;
					left: %fpx;
					top: %fpx;
					transform: rotate(%fdeg);
					transform-origin: 0 0; /* Ensure the rotation starts from the first point */
					background-color: rgba(%d,%d,%d,%d);
					color: rgba(%d,%d,%d,%d);
				"
			>%v</div>
			`, distance, link.Size, point1X, point1Y+offset, angle, link.Color.Red, link.Color.Green, link.Color.Blue, link.Color.Alpha, link.LabelColor.Red, link.LabelColor.Green, link.LabelColor.Blue, link.LabelColor.Alpha, link.Label))

}

func (a *Api) renderTurtle(tmpl *strings.Builder, turtle Turtle, model *Model, patchSize int) {
	turtleSize := int(float64(patchSize) * turtle.Size)
	if turtleSize < 1 {
		turtleSize = 1
	}
	turtleOffset := (patchSize - turtleSize) / 2
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
	screenHeight = screenHeight - int(float64(screenHeight)*.02)

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
