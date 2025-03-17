// playground for quick testing

package playgound

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

	dogs := model.NewTurtleBreed("dogs", "", map[string]interface{}{
		"earshape": "round",
	})
	cats := model.NewTurtleBreed("cats", "", map[string]interface{}{
		"earshape": "pointy",
	})

	parentChildrenLinks := model.NewLinkBreed("parent-children")
	siblingLinks := model.NewLinkBreed("siblings")

	settings := model.ModelSettings{
		TurtleBreeds:         []*model.TurtleBreed{dogs, cats},
		DirectedLinkBreeds:   []*model.LinkBreed{parentChildrenLinks}, //parent chilred are directed
		UndirectedLinkBreeds: []*model.LinkBreed{siblingLinks},        //siblings are undirected
		WrappingX:            false,                                   // wrapping x off
		WrappingY:            false,                                   // wrapping y off
		MinPxCor:             -10,                                     // min x patch coordinate
		MaxPxCor:             10,                                      // max x patch coordinate
		MinPyCor:             -10,                                     // min y patch coordinate
		MaxPyCor:             10,                                      // max y patch coordinate
		RandomSeed:           10,                                      // random seed
	}

	s.model = model.NewModel(settings) // create a new model with the settings
}

func (s *Sim) SetUp() error {

	s.model.ClearAll()

	s.model.Patches.Ask(func(p *model.Patch) {

		p.PColor.SetColor(model.Color{
			Red:   int(math.Abs(float64(p.PXCor() * p.PYCor() * 8))),
			Green: int(math.Abs(float64(p.PXCor() * p.PYCor() * 8))),
			Blue:  int(math.Abs(float64(p.PXCor() * p.PYCor() * 8))),
			Alpha: 1,
		})
	})

	s.model.Patch(0, 0).PColor.SetColor(model.Green)

	s.model.CreateTurtles(1, func(t *model.Turtle) {
		t.SetXY(0, 0)
		t.SetSize(1)
		t.SetHeading(90)
		t.Shape = "triangle"
		t.Color.SetColor(model.Red)
	})

	return nil
}

func (s *Sim) Go() {

	// s.model.Patches.Ask(func(p *model.Patch) {
	// 	p.PColor.SetColor(s.model.RandomColor())
	// })
	// t1 := s.model.Turtle("", 2)
	// t1.Forward(10)

	s.model.Tick()

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
