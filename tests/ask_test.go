package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/pkg/model"
)

func TestAskLinks(t *testing.T) {
	//create a list of *Links
	link1 := &model.Link{}
	link2 := &model.Link{}
	link3 := &model.Link{}

	//create an agentset
	agentset := model.NewLinkAgentSet([]*model.Link{link1, link2, link3})

	// ask to change all colors to red
	agentset.Ask(
		func(link *model.Link) {
			link.Color.SetColor(model.Red)
		},
	)

	//make sure that all colors are red
	if !agentset.All(func(l *model.Link) bool {
		return l.Color == model.Red
	}) {
		t.Errorf("Expected linkset to have all links with shape 'circle'")
	}
}

func TestAskPatches(t *testing.T) {
	//create a list of *Patches
	patch1 := &model.Patch{}
	patch2 := &model.Patch{}
	patch3 := &model.Patch{}

	//create an agentset
	agentset := model.NewPatchAgentSet([]*model.Patch{patch1, patch2, patch3})

	// ask to change all colors to red
	agentset.Ask(
		func(patch *model.Patch) {
			patch.Color.SetColor(model.Red)
		},
	)

	//make sure that all colors are red
	if !agentset.All(func(p *model.Patch) bool {
		return p.Color == model.Red
	}) {
		t.Errorf("Expected patchset to have all patches with color 'red'")
	}
}

func TestAskTurtles(t *testing.T) {
	//create a list of *Turtles
	turtle1 := &model.Turtle{}
	turtle2 := &model.Turtle{}
	turtle3 := &model.Turtle{}

	//create an agentset
	agentset := model.NewTurtleAgentSet([]*model.Turtle{turtle1, turtle2, turtle3})

	// ask to change all colors to red
	agentset.Ask(
		func(turtle *model.Turtle) {
			turtle.Color.SetColor(model.Red)
		},
	)

	//make sure that all colors are red
	if !agentset.All(func(t *model.Turtle) bool {
		return t.Color == model.Red
	}) {
		t.Errorf("Expected turtleset to have all turtles with color 'red'")
	}
}
