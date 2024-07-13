package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/internal/model"
)

func TestAskLink(t *testing.T) {
	link := &model.Link{
		End1: nil,
		End2: nil,
	}
	link.Color.SetColorScale(model.Blue)

	model.AskLink(link, []model.LinkOperation{
		func(link *model.Link) {
			link.Color.SetColorScale(model.Red)
		},
	})

	//make sure that the color is red
	if link.Color.GetColorScale() != model.Red {
		t.Errorf("Expected red, got %v", link.Color.GetColorScale())
	}
}

func TestAskLinks(t *testing.T) {
	//create a list of *Links
	link1 := &model.Link{}
	link2 := &model.Link{}
	link3 := &model.Link{}

	//create an agentset
	agentset := model.LinkSet([]*model.Link{link1, link2, link3})

	// ask to change all colors to red
	model.AskLinks(*agentset, []model.LinkOperation{
		func(link *model.Link) {
			link.Color.SetColorScale(model.Red)
		},
	})

	//make sure that all colors are red
	if !agentset.All(func(l *model.Link) bool {
		return l.Color.GetColorScale() == model.Red
	}) {
		t.Errorf("Expected linkset to have all links with shape 'circle'")
	}
}

func TestAskPatch(t *testing.T) {
	patch := &model.Patch{}
	patch.PColor.SetColorScale(model.Blue)

	model.AskPatch(patch, []model.PatchOperation{
		func(patch *model.Patch) {
			patch.PColor.SetColorScale(model.Red)
		},
	})

	//make sure that the color is red
	if patch.PColor.GetColorScale() != model.Red {
		t.Errorf("Expected red, got %v", patch.PColor.GetColorScale())
	}
}

func TestAskPatches(t *testing.T) {
	//create a list of *Patches
	patch1 := &model.Patch{}
	patch2 := &model.Patch{}
	patch3 := &model.Patch{}

	//create an agentset
	agentset := model.PatchSet([]*model.Patch{patch1, patch2, patch3})

	// ask to change all colors to red
	model.AskPatches(agentset, []model.PatchOperation{
		func(patch *model.Patch) {
			patch.PColor.SetColorScale(model.Red)
		},
	})

	//make sure that all colors are red
	if !agentset.All(func(p *model.Patch) bool {
		return p.PColor.GetColorScale() == model.Red
	}) {
		t.Errorf("Expected patchset to have all patches with color 'red'")
	}
}

func TestAskTurtle(t *testing.T) {
	turtle := &model.Turtle{}
	turtle.Color.SetColorScale(model.Blue)

	model.AskTurtle(turtle, []model.TurtleOperation{
		func(turtle *model.Turtle) {
			turtle.Color.SetColorScale(model.Red)
		},
	})

	//make sure that the color is red
	if turtle.Color.GetColorScale() != model.Red {
		t.Errorf("Expected red, got %v", turtle.Color.GetColorScale())
	}
}

func TestAskTurtles(t *testing.T) {
	//create a list of *Turtles
	turtle1 := &model.Turtle{}
	turtle2 := &model.Turtle{}
	turtle3 := &model.Turtle{}

	//create an agentset
	agentset := model.TurtleSet([]*model.Turtle{turtle1, turtle2, turtle3})

	// ask to change all colors to red
	model.AskTurtles(agentset, []model.TurtleOperation{
		func(turtle *model.Turtle) {
			turtle.Color.SetColorScale(model.Red)
		},
	})

	//make sure that all colors are red
	if !agentset.All(func(t *model.Turtle) bool {
		return t.Color.GetColorScale() == model.Red
	}) {
		t.Errorf("Expected turtleset to have all turtles with color 'red'")
	}
}
