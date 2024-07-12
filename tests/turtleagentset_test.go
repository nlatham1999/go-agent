package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/internal/universe"
)

func TestAllTurtle(t *testing.T) {

	//create a basic universe
	u := universe.NewUniverse(nil, nil, nil, nil, nil, nil, false)

	turtle1 := universe.NewTurtle(u, 0, "")
	turtle2 := universe.NewTurtle(u, 1, "")
	turtle3 := universe.NewTurtle(u, 2, "")

	turtleSet := universe.TurtleSet([]*universe.Turtle{turtle1, turtle2, turtle3})

	turtle1.Color.SetColorScale(1.0)
	turtle2.Color.SetColorScale(1.0)
	turtle3.Color.SetColorScale(1.0)

	// assert that turtleset has All of shape "circle"
	if !turtleSet.All(func(t *universe.Turtle) bool {
		return t.Color.GetColorScale() == 1.0
	}) {
		t.Errorf("Expected turtleset to have all turtles with color '1.0'")
	}

	turtle2.Color.SetColorScale(2.0)

	if turtleSet.All(func(t *universe.Turtle) bool {
		return t.Color.GetColorScale() == 1.0
	}) {
		t.Errorf("Expected turtleset to not have all turtles with color '1.0'")
	}
}

func TestAnyTurtle(t *testing.T) {

	//create a basic universe
	u := universe.NewUniverse(nil, nil, nil, nil, nil, nil, false)

	turtle1 := universe.NewTurtle(u, 0, "")
	turtle2 := universe.NewTurtle(u, 1, "")
	turtle3 := universe.NewTurtle(u, 2, "")

	turtleSet := universe.TurtleSet([]*universe.Turtle{turtle1, turtle2, turtle3})

	turtle1.Color.SetColorScale(1.0)

	// assert that turtleset has Any of shape "circle"
	if !turtleSet.Any(func(t *universe.Turtle) bool {
		return t.Color.GetColorScale() == 1.0
	}) {
		t.Errorf("Expected turtleset to have a turtle with color '1.0'")
	}

	turtle1.Color.SetColorScale(2.0)

	if turtleSet.Any(func(t *universe.Turtle) bool {
		return t.Color.GetColorScale() == 1.0
	}) {
		t.Errorf("Expected turtleset to not have a turtle with color '1.0'")
	}
}

func TestAtPointsTurtle(t *testing.T) {

	//create basic universe
	u := universe.NewUniverse(nil, nil, nil, nil, nil, nil, false)

	//create some random turtles from the universe
	u.CreateTurtles(5, "", nil)

	turtle1 := u.Turtle("", 0)
	turtle2 := u.Turtle("", 1)
	turtle3 := u.Turtle("", 2)
	turtle4 := u.Turtle("", 3)
	turtle5 := u.Turtle("", 4)

	turtle1.SetXY(0, 0)
	turtle2.SetXY(1, 1)
	turtle3.SetXY(2, 2)
	turtle4.SetXY(3, 3)
	turtle5.SetXY(4, 4)

	//create a turtleset
	turtleSet := universe.TurtleSet([]*universe.Turtle{turtle1, turtle2, turtle3, turtle4, turtle5})

	//get turtles at the patches
	turtleSetAtPatches := turtleSet.AtPoints(u, []universe.Coordinate{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 2}})

	//make sure that all turtles are at the patches
	if turtleSetAtPatches.Count() != 3 {
		t.Errorf("Expected 3 turtles, got %d", turtleSetAtPatches.Count())
	}

}
