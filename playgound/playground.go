package playgound

import (
	"fmt"

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
	settings := model.ModelSettings{}

	s.model = model.NewModel(settings)
}

func (s *Sim) SetUp() error {

	fmt.Println("hitting setup")

	s.model.ClearAll()

	s.model.CreateTurtles(4, "", []model.TurtleOperation{
		func(t *model.Turtle) {
			t.FaceXY(0, 0)
		},
	})

	t1 := s.model.Turtle("", 0)
	t1.SetXY(5, 5)
	t1.Color = model.Red

	t2 := s.model.Turtle("", 1)
	t2.SetXY(-5, 5)
	t2.Color = model.Blue

	t3 := s.model.Turtle("", 2)
	t3.SetXY(-5, -5)
	t3.Color = model.Green

	t4 := s.model.Turtle("", 3)
	t4.SetXY(5, -5)
	t4.Color = model.Yellow

	t1.CreateLinkToTurtle("", t2, nil)
	t2.CreateLinkToTurtle("", t3, nil)
	t3.CreateLinkToTurtle("", t4, nil)
	t4.CreateLinkToTurtle("", t1, nil)

	for link, _ := s.model.Links().First(); link != nil; link, _ = s.model.Links().Next() {
		link.Color = model.White
		// link.TieMode = model.TieModeFixed
	}

	return nil
}

func (s *Sim) Go() {

	// t1 := s.model.Turtle("", 2)
	// t1.Forward(10)
	fmt.Println("starting go")

	for turtle, _ := s.model.Turtles("").First(); turtle != nil; turtle, _ = s.model.Turtles("").Next() {
		fmt.Println("test", turtle.Who())
		allTurtlesInRadius := s.model.Turtles("").All(func(t *model.Turtle) bool {
			return t.DistanceTurtle(turtle) < 12
		})
		if !allTurtlesInRadius {
			turtle.Forward(1)
		}
	}
	fmt.Println("end go")
}

func (s *Sim) Stats() map[string]interface{} {
	return nil
}

func (s *Sim) Stop() bool {
	return false
}

func (s *Sim) Widgets() []api.Widget {
	return nil
}
