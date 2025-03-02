package playgound

import (
	"fmt"
	"time"

	"github.com/nlatham1999/go-agent/api"
	"github.com/nlatham1999/go-agent/model"
)

// Init()        // runs at the very beginning
// SetUp() error // sets up the model
// Go()          // runs the model

// Model() *model.Model           // returns the model
// Stats() map[string]interface{} //returns the stats of the model
// Stop() bool                    // on whether to stop the model
// Widgets() []Widget             // returns the widgets of the model

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
		WrappingX: false,
		WrappingY: false,
		MinPxCor:  -10,
		MaxPxCor:  10,
		MinPyCor:  -10,
		MaxPyCor:  10,
		Globals: map[string]interface{}{
			"max-length":  10.0,
			"num-turtles": 2,
		},
	}

	s.model = model.NewModel(settings)
}

func (s *Sim) SetUp() error {

	s.model.ClearAll()

	// s.model.CreateTurtles(100, "", func(t *model.Turtle) {
	// 	t.SetXY(s.model.RandomXCor(), s.model.RandomYCor())
	// })

	s.model.CreateTurtles(s.model.GetGlobal("num-turtles").(int), "", func(t *model.Turtle) {
		t.SetXY(s.model.RandomXCor(), s.model.RandomYCor())
	})

	s.model.Turtles("").Ask(func(t *model.Turtle) {
		s.model.Turtles("").Ask(func(t2 *model.Turtle) {
			t.CreateLinkWithTurtle("", t2, func(l *model.Link) {
				l.TieMode = model.TieModeAllTied
				// fmt.Println("Link created", l.TieMode)
			})
		})
	})

	return nil
}

func (s *Sim) Go() {

	// t1 := s.model.Turtle("", 2)
	// t1.Forward(10)

	s.RotateTurtleTest()

}

func (s Sim) CreateGraph() {
	start := time.Now()

	s.model.ClearLinks()

	turtles := s.model.Turtles("")
	turtles.Ask(func(t *model.Turtle) {

		tInRadius := s.model.Turtles("").InRadiusTurtle(s.model.GetGlobal("max-length").(float64), t)

		tInRadius.Ask(func(t2 *model.Turtle) {
			t.CreateLinkWithTurtle("", t2, func(l *model.Link) {
				// l.TieMode = model.TieModeAllTied
				// fmt.Println("Link created", l.TieMode)
			})
		})
	})

	turtles.Ask(func(t *model.Turtle) {
		t.SetSize(float64(t.LinkedTurtles("").Count()) / 1000)
	})

	fmt.Println("Time taken: ", time.Since(start))
}

func (s *Sim) Stats() map[string]interface{} {
	return nil
}

func (s *Sim) Stop() bool {
	return false
}

func (s *Sim) ChangeColor() {
	turtles := s.model.Turtles("")
	turtles.Ask(func(t *model.Turtle) {
		t.Color.SetColor(s.model.RandomColor())
	})
}

func (s *Sim) RotateTurtleTest() {
	fmt.Println("Rotating turtle")
	// t1 := s.model.Turtle("", 0)
	// t1.Right(1)
	// t2 := s.model.Turtle("", 1)
	// t2.Right(1)

	s.model.Turtles("").Ask(func(t *model.Turtle) {
		t.Right(1)
	})
}

func (s *Sim) Widgets() []api.Widget {
	return []api.Widget{
		{
			PrettyName:      "Max Length",
			TargetVariable:  "max-length",
			WidgetType:      "slider",
			WidgetValueType: "float",
			MinValue:        ".01",
			MaxValue:        "20",
			DefaultValue:    "10",
			StepAmount:      ".01",
		},
		{
			PrettyName:     "Change Color",
			TargetVariable: "change-color",
			Target:         s.ChangeColor,
			WidgetType:     "button",
		},
		{
			PrettyName:     "Rotate Turtle 1",
			TargetVariable: "rotate-turtle-1",
			Target:         s.RotateTurtleTest,
			WidgetType:     "button",
		},
		{
			PrettyName:      "Number of Turtles",
			TargetVariable:  "num-turtles",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "100",
			DefaultValue:    "2",
		},
	}
}
