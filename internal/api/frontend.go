package api

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/nlatham1999/go-agent/internal/model"
)

//go:embed html/index.html
var indexHTML string

var (
	statsKeys = []string{}
)

func (a *Api) getFrontend(width int, height int) string {

	model := a.Sim.Model()

	patchSize := a.getPatchSize(width, height, model.WorldWidth(), model.WorldHeight())

	var tmpl strings.Builder
	tmpl.WriteString(`<div class="grid-container patch-grid" style="position: absolute; left: 50%; top: 5px;">`)

	// Render patches
	for _, patch := range model.Patches.List() {
		tmpl.WriteString(a.renderPatch(patch, model, patchSize))
	}

	// Render turtles
	turtleOffset := patchSize / 2
	for _, turtle := range model.Turtles("").ListSorted() {
		tmpl.WriteString(a.renderTurtle(turtle, model, patchSize, turtleOffset))
	}

	tmpl.WriteString(`</div>`)

	// Render stats
	stats := a.Sim.Stats()
	if len(statsKeys) == 0 {
		for key := range stats {
			statsKeys = append(statsKeys, key)
		}
	}
	tmpl.WriteString(`<div style="position: absolute; left: 5px; top: 20%;">`)
	tmpl.WriteString(`<div>Ticks: ` + fmt.Sprintf("%d", model.Ticks) + `</div>`)
	for _, key := range statsKeys {
		tmpl.WriteString(`<div>` + key + `: ` + fmt.Sprintf("%v", stats[key]) + `</div>`)
	}
	tmpl.WriteString(`</div>`)

	return tmpl.String()
}

func (a *Api) renderTurtle(turtle *model.Turtle, model *model.Model, patchSize int, turtleOffset int) string {
	turtleSize := int(float64(patchSize) * turtle.GetSize())
	if turtleSize < 1 {
		turtleSize = 1
	}
	relativeX := turtle.XCor() - float64(model.MinPxCor())
	relativeY := turtle.YCor() - float64(model.MinPyCor())
	tmpl := `
			<div 
				class="turtle" 
				style="
					width:` + fmt.Sprintf("%dpx", turtleSize) + `;
					height:` + fmt.Sprintf("%dpx", turtleSize) + `;
					left:` + fmt.Sprintf("%vpx", relativeX*float64(patchSize)+float64(turtleOffset)) + `; 
					top:` + fmt.Sprintf("%vpx", relativeY*float64(patchSize)+float64(turtleOffset)) + `;
					background-color: rgba(` + fmt.Sprintf("%d", turtle.Color.Red) + `, ` + fmt.Sprintf("%d", turtle.Color.Green) + `, ` + fmt.Sprintf("%d", turtle.Color.Blue) + `, ` + fmt.Sprintf("%d", turtle.Color.Alpha) + `);
					color: rgba(` + fmt.Sprintf("%d", turtle.GetLabelColor().Red) + `, ` + fmt.Sprintf("%d", turtle.GetLabelColor().Green) + `, ` + fmt.Sprintf("%d", turtle.GetLabelColor().Blue) + `, ` + fmt.Sprintf("%d", turtle.GetLabelColor().Alpha) + `);
					"
			>
			` + fmt.Sprintf("%v", turtle.GetLabel()) + `
			</div>
		`
	return tmpl
}

func (a *Api) renderPatch(patch *model.Patch, model *model.Model, patchSize int) string {
	relativeX := patch.PXCor() - model.MinPxCor()
	relativeY := patch.PYCor() - model.MinPyCor()
	tmpl := `
			<div 
				class="patch" 
				style="
					width:` + fmt.Sprintf("%dpx", patchSize) + `;
					height:` + fmt.Sprintf("%dpx", patchSize) + `;
					left:` + fmt.Sprintf("%dpx", relativeX*patchSize) + `; 
					top:` + fmt.Sprintf("%dpx", relativeY*patchSize) + `;
					background-color: rgba(` + fmt.Sprintf("%d", patch.PColor.Red) + `, ` + fmt.Sprintf("%d", patch.PColor.Green) + `, ` + fmt.Sprintf("%d", patch.PColor.Blue) + `, ` + fmt.Sprintf("%d", patch.PColor.Alpha) + `);
				"
			>
			</div>
		`
	return tmpl
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
	html := `<div class="grid-container widgets" style="position: absolute; left: 1%; top: 50%;">`

	// Add widgets here
	for _, widget := range a.Sim.Widgets() {
		html += widget.Render()
	}

	html += `</div>`
	return html
}
