package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/internal/universe"
)

func TestAllTurtle(t *testing.T) {
	turtle1 := universe.NewTurtle(0)
	turtle2 := universe.NewTurtle(1)
	turtle3 := universe.NewTurtle(2)

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

	turtle1 := universe.NewTurtle(0)
	turtle2 := universe.NewTurtle(1)
	turtle3 := universe.NewTurtle(2)

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
