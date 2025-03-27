// playground for quick testing
package playground

import (
	"fmt"
	"math"

	"github.com/nlatham1999/go-agent/pkg/api"
	"github.com/nlatham1999/go-agent/pkg/model"
)

// Init()        // runs at the very beginning
// SetUp() error // sets up the model
// Go()          // runs the model

// Model() *model.Model           // returns the model
// Stats() map[string]interface{} //returns the stats of the model
// Stop() bool                    // on whether to stop the model
// Widgets() []Widget             // returns the widgets of the model

var _ api.ModelInterface = &Sim{}

type Sim struct {
	model *model.Model
}

func NewSim() *Sim {
	return &Sim{}
}

func (s *Sim) Model() *model.Model {
	return s.model
}

func (s *Sim) Init() {
	settings := model.ModelSettings{
		WrappingX:  false,
		WrappingY:  false,
		MinPxCor:   -100, // min x patch coordinate
		MaxPxCor:   100,  // max x patch coordinate
		MinPyCor:   -100, // min y patch coordinate
		MaxPyCor:   100,  // max y patch coordinate
		RandomSeed: 10,   // random seed
		TurtleProperties: map[string]interface{}{
			"newHeading": nil,
		},
	}

	s.model = model.NewModel(settings) // create a new model with the settings
}

func (s *Sim) SetUp() error {

	s.model.ClearAll()

	s.model.Patches.Ask(func(p *model.Patch) {
		p.PColor.SetColor(model.Green)
	})

	s.model.Patch(float64(s.model.MaxPxCor())-2, float64(s.model.MaxPyCor())-2).PColor.SetColor(model.White)
	s.model.Patch(float64(s.model.MaxPxCor())-2, float64(s.model.MaxPyCor())-1).PColor.SetColor(model.White)
	s.model.Patch(float64(s.model.MaxPxCor())-2, float64(s.model.MaxPyCor())).PColor.SetColor(model.White)
	s.model.Patch(float64(s.model.MaxPxCor())-1, float64(s.model.MaxPyCor())-2).PColor.SetColor(model.White)
	s.model.Patch(float64(s.model.MaxPxCor())-1, float64(s.model.MaxPyCor())-1).PColor.SetColor(model.White)
	s.model.Patch(float64(s.model.MaxPxCor())-1, float64(s.model.MaxPyCor())).PColor.SetColor(model.White)
	s.model.Patch(float64(s.model.MaxPxCor()), float64(s.model.MaxPyCor())-2).PColor.SetColor(model.White)
	s.model.Patch(float64(s.model.MaxPxCor()), float64(s.model.MaxPyCor())-1).PColor.SetColor(model.White)
	s.model.Patch(float64(s.model.MaxPxCor()), float64(s.model.MaxPyCor())).PColor.SetColor(model.White)

	return nil
}

func (s *Sim) Go() {

	s.model.Patches.Ask(func(p *model.Patch) {

		if p.PColor == model.White {
			p.PColor.SetColor(model.Black)
			return
		}

		neighborsTotal := p.Neighbors()

		neighborsColored := neighborsTotal.With(func(p *model.Patch) bool {
			return p.PColor == model.White || p.PColor == model.Black
		})

		if s.model.RandomFloat(1) < float64(neighborsColored.Count())/float64(neighborsTotal.Count()) {
			p.PColor.SetColor(model.White)
		} else {
			p.PColor.SetColor(model.Green)
		}
	})
}

func (s *Sim) Stats() map[string]interface{} {
	return nil
}

func (s *Sim) Stop() bool {
	return false
}

func (s *Sim) MoveForward() {
	t1 := s.model.Turtle(0)
	fmt.Println("Moving forward")
	t1.Forward(1)
}

func (s *Sim) Rotate() {
	t1 := s.model.Turtle(0)
	fmt.Println("Rotating")
	t1.Right(10)
}

/*
type Widget struct {
	PrettyName         string   `json:"prettyName"`
	TargetVariable     string   `json:"targetVariable"`
	WidgetType         string   `json:"widgetType"`
	WidgetValueType    string   `json:"widgetValueType"`
	MinValue           string   `json:"minValue"`
	MaxValue           string   `json:"maxValue"`
	DefaultValue       string   `json:"defaultValue"`
	StepAmount         string   `json:"stepAmount"`
	Target             func()   `json:"target"`
	ValuePointerInt    *int     `json:"valuePointerInt"`
	ValuePointerFloat  *float64 `json:"valuePointerFloat"`
	ValuePointerString *string  `json:"valuePointerString"`
}
*/

func (s *Sim) Widgets() []api.Widget {
	return []api.Widget{
		{
			PrettyName:     "Move Forward",
			TargetVariable: "move-forward",
			WidgetType:     "button",
			Target:         s.MoveForward,
		},
		{
			PrettyName:     "Rotate",
			TargetVariable: "rotate",
			WidgetType:     "button",
			Target:         s.Rotate,
		},
	}
}

func elasticCollisionHeadings(x1, y1, x2, y2, h1, h2 float64) (newH1, newH2 float64) {
	// convert headings to velocity vectors (unit speed assumed)
	v1x := math.Cos(h1)
	v1y := math.Sin(h1)
	v2x := math.Cos(h2)
	v2y := math.Sin(h2)

	// compute normal vector (from p1 to p2) and normalize it
	nx := x2 - x1
	ny := y2 - y1

	// snap near-zero values to exactly 0
	const epsilon = 1e-8
	if math.Abs(nx) < epsilon {
		nx = 0
	}
	if math.Abs(ny) < epsilon {
		ny = 0
	}
	nLen := math.Hypot(nx, ny)
	if nLen == 0 {
		return h1, h2
	}
	nx /= nLen
	ny /= nLen

	// relative velocity
	dvx := v1x - v2x
	dvy := v1y - v2y

	// project relative velocity onto normal
	dot := dvx*nx + dvy*ny
	if math.Abs(dot) < epsilon {
		// they're not moving toward each other along the collision normal
		return h1, h2
	}

	// exchange velocity components along the normal
	v1xNew := v1x - dot*nx
	v1yNew := v1y - dot*ny
	v2xNew := v2x + dot*nx
	v2yNew := v2y + dot*ny

	// compute new headings
	newH1 = math.Atan2(v1yNew, v1xNew)
	newH2 = math.Atan2(v2yNew, v2xNew)
	return
}
