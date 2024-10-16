package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/internal/model"
)

func TestAllTurtle(t *testing.T) {

	//create a basic model
	u := model.NewModel(model.ModelSettings{})

	turtle1 := model.NewTurtle(u, 0, "", 0, 0)
	turtle2 := model.NewTurtle(u, 1, "", 0, 0)
	turtle3 := model.NewTurtle(u, 2, "", 0, 0)

	turtleSet := model.NewTurtleAgentSet([]*model.Turtle{turtle1, turtle2, turtle3})

	turtle1.Color.SetColor(model.Lime)
	turtle2.Color.SetColor(model.Lime)
	turtle3.Color.SetColor(model.Lime)

	// assert that turtleset has All of shape "circle"
	if !turtleSet.All(func(t *model.Turtle) bool {
		return t.Color == model.Lime
	}) {
		t.Errorf("Expected turtleset to have all turtles with color '1.0'")
	}

	turtle2.Color.SetColor(model.Blue)

	if turtleSet.All(func(t *model.Turtle) bool {
		return t.Color == model.Lime
	}) {
		t.Errorf("Expected turtleset to not have all turtles with color '1.0'")
	}
}

func TestAnyTurtle(t *testing.T) {

	//create a basic model
	u := model.NewModel(model.ModelSettings{})

	turtle1 := model.NewTurtle(u, 0, "", 0, 0)
	turtle2 := model.NewTurtle(u, 1, "", 0, 0)
	turtle3 := model.NewTurtle(u, 2, "", 0, 0)

	turtleSet := model.NewTurtleAgentSet([]*model.Turtle{turtle1, turtle2, turtle3})

	turtle1.Color.SetColor(model.Lime)

	// assert that turtleset has Any of shape "circle"
	if !turtleSet.Any(func(t *model.Turtle) bool {
		return t.Color == model.Lime
	}) {
		t.Errorf("Expected turtleset to have a turtle with color '1.0'")
	}

	turtle1.Color.SetColor(model.Blue)

	if turtleSet.Any(func(t *model.Turtle) bool {
		return t.Color == model.Lime
	}) {
		t.Errorf("Expected turtleset to not have a turtle with color '1.0'")
	}
}

func TestAtPointsTurtle(t *testing.T) {

	//create basic model
	m := model.NewModel(model.ModelSettings{})

	//create some random turtles from the model
	m.CreateTurtles(5, "", nil)

	turtle1 := m.Turtle("", 0)
	turtle2 := m.Turtle("", 1)
	turtle3 := m.Turtle("", 2)
	turtle4 := m.Turtle("", 3)
	turtle5 := m.Turtle("", 4)

	turtle1.SetXY(0, 0)
	turtle2.SetXY(1, 1)
	turtle3.SetXY(2, 2)
	turtle4.SetXY(3, 3)
	turtle5.SetXY(4, 4)

	//create a turtleset
	turtleSet := model.NewTurtleAgentSet([]*model.Turtle{turtle1, turtle2, turtle3, turtle4, turtle5})

	//get turtles at the patches
	turtleSetAtPatches := turtleSet.AtPoints(m, []model.Coordinate{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 2}})

	//make sure that all turtles are at the patches
	if turtleSetAtPatches.Count() != 3 {
		t.Errorf("Expected 3 turtles, got %d", turtleSetAtPatches.Count())
	}

}

func TestTurtlesWhoAreNotInTurtles(t *testing.T) {

	//create a basic model
	u := model.NewModel(model.ModelSettings{})

	turtle1 := model.NewTurtle(u, 0, "", 0, 0)
	turtle2 := model.NewTurtle(u, 1, "", 0, 0)
	turtle3 := model.NewTurtle(u, 2, "", 0, 0)

	turtleSet := model.NewTurtleAgentSet([]*model.Turtle{turtle1, turtle2, turtle3})

	turtleSet2 := model.NewTurtleAgentSet([]*model.Turtle{turtle1, turtle2})

	turtleSet3 := turtleSet.WhoAreNot(turtleSet2)

	if turtleSet3.Count() != 1 {
		t.Errorf("Expected turtleSet3 to have 1 turtle")
	}

	if !turtleSet3.Contains(turtle3) {
		t.Errorf("Expected turtleSet3 to have turtle3")
	}
}

func TestTurtlesWhoAreNotTurtle(t *testing.T) {

	//create a basic model
	u := model.NewModel(model.ModelSettings{})

	turtle1 := model.NewTurtle(u, 0, "", 0, 0)
	turtle2 := model.NewTurtle(u, 1, "", 0, 0)
	turtle3 := model.NewTurtle(u, 2, "", 0, 0)

	turtleSet := model.NewTurtleAgentSet([]*model.Turtle{turtle1, turtle2, turtle3})

	turtleSet2 := turtleSet.WhoAreNotTurtle(turtle1)

	if turtleSet2.Count() != 2 {
		t.Errorf("Expected turtleSet2 to have 2 turtles")
	}
}

func TestTurtlesMaxNOf(t *testing.T) {

	//create a basic model
	settings := model.ModelSettings{}
	u := model.NewModel(settings)

	turtle1 := model.NewTurtle(u, 0, "", 0, 0)
	turtle2 := model.NewTurtle(u, 1, "", 0, 0)
	turtle3 := model.NewTurtle(u, 2, "", 0, 0)
	turtle4 := model.NewTurtle(u, 3, "", 0, 0)

	turtleSet := model.NewTurtleAgentSet([]*model.Turtle{turtle1, turtle2, turtle3, turtle4})

	turtleSet2 := turtleSet.MaxNOf(2, func(t *model.Turtle) float64 {
		return float64(t.Who())
	})

	if turtleSet2.Count() != 2 {
		t.Errorf("Expected turtleSet2 to have 2 turtles")
	}
}
