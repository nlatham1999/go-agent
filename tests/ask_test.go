package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/internal/universe"
)

func TestAskLink(t *testing.T) {
	link := &universe.Link{
		End1: nil,
		End2: nil,
	}
	link.Color.SetColorScale(universe.Blue)

	universe.AskLink(link, []universe.LinkOperation{
		func(link *universe.Link) {
			link.Color.SetColorScale(universe.Red)
		},
	})

	//make sure that the color is red
	if link.Color.GetColorScale() != universe.Red {
		t.Errorf("Expected red, got %v", link.Color.GetColorScale())
	}
}

func TestAskLinks(t *testing.T) {
	//create a list of *Links
	link1 := &universe.Link{}
	link2 := &universe.Link{}
	link3 := &universe.Link{}

	//create an agentset
	agentset := universe.LinkSet([]*universe.Link{link1, link2, link3})

	// ask to change all colors to red
	universe.AskLinks(*agentset, []universe.LinkOperation{
		func(link *universe.Link) {
			link.Color.SetColorScale(universe.Red)
		},
	})

	//make sure that all colors are red
	if !agentset.All(func(l *universe.Link) bool {
		return l.Color.GetColorScale() == universe.Red
	}) {
		t.Errorf("Expected linkset to have all links with shape 'circle'")
	}
}

func TestAskPatch(t *testing.T) {
	patch := &universe.Patch{}
	patch.PColor.SetColorScale(universe.Blue)

	universe.AskPatch(patch, []universe.PatchOperation{
		func(patch *universe.Patch) {
			patch.PColor.SetColorScale(universe.Red)
		},
	})

	//make sure that the color is red
	if patch.PColor.GetColorScale() != universe.Red {
		t.Errorf("Expected red, got %v", patch.PColor.GetColorScale())
	}
}

func TestAskPatches(t *testing.T) {
	//create a list of *Patches
	patch1 := &universe.Patch{}
	patch2 := &universe.Patch{}
	patch3 := &universe.Patch{}

	//create an agentset
	agentset := universe.PatchSet([]*universe.Patch{patch1, patch2, patch3})

	// ask to change all colors to red
	universe.AskPatches(agentset, []universe.PatchOperation{
		func(patch *universe.Patch) {
			patch.PColor.SetColorScale(universe.Red)
		},
	})

	//make sure that all colors are red
	if !agentset.All(func(p *universe.Patch) bool {
		return p.PColor.GetColorScale() == universe.Red
	}) {
		t.Errorf("Expected patchset to have all patches with color 'red'")
	}
}

func TestAskTurtle(t *testing.T) {
	turtle := &universe.Turtle{
		Heading: 0,
	}
	turtle.Color.SetColorScale(universe.Blue)

	universe.AskTurtle(turtle, []universe.TurtleOperation{
		func(turtle *universe.Turtle) {
			turtle.Color.SetColorScale(universe.Red)
		},
	})

	//make sure that the color is red
	if turtle.Color.GetColorScale() != universe.Red {
		t.Errorf("Expected red, got %v", turtle.Color.GetColorScale())
	}
}

func TestAskTurtles(t *testing.T) {
	//create a list of *Turtles
	turtle1 := &universe.Turtle{}
	turtle2 := &universe.Turtle{}
	turtle3 := &universe.Turtle{}

	//create an agentset
	agentset := universe.TurtleSet([]*universe.Turtle{turtle1, turtle2, turtle3})

	// ask to change all colors to red
	universe.AskTurtles(agentset, []universe.TurtleOperation{
		func(turtle *universe.Turtle) {
			turtle.Color.SetColorScale(universe.Red)
		},
	})

	//make sure that all colors are red
	if !agentset.All(func(t *universe.Turtle) bool {
		return t.Color.GetColorScale() == universe.Red
	}) {
		t.Errorf("Expected turtleset to have all turtles with color 'red'")
	}
}
