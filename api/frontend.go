package api

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed html/index.html
var indexHTML string

//go:embed html/threejs.html
var threejsHTML string

//go:embed html/scripts.html
var scriptsHTML string

//go:embed html/style.html
var styleHTML string

var (
	statsKeys = []string{}
)

func (a *Api) renderStatsAsWidgets() (string, int) {
	html := ""
	stats := a.Model.Stats()

	//tick stat
	html += fmt.Sprintf(`
		<div class="widget widget-stats">
			<div id="stats-ticks">Ticks : %d</div>
		</div>
	`, a.Model.Model().Ticks)

	count := 2
	for key, value := range stats {
		if value == nil {
			value = "null"
		}
		html += fmt.Sprintf(`
		<div class="widget widget-stats" style="top: %dpx;">
			<div id="stats-%s">%s : %v</div>
		</div>
		`, count*20, key, key, value)
		count++
	}
	return html, len(stats) + 1
}

func (a *Api) renderLink(tmpl *strings.Builder, link Link, model *Model, patchSize int, screenHeight float64) {

	if link.Hidden {
		return
	}

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

func (a *Api) buildWidgets() string {

	// Add stats widget
	html, count := a.renderStatsAsWidgets()

	// Add widgets here
	count++
	for _, widget := range a.Model.Widgets() {
		html += widget.Render(count)
		count++
	}

	return html
}
