package state

import (
	"fmt"

	"github.com/nlatham1999/go-agent/pkg/model"
)

func FromModel(model *model.Model) *Model {
	stateModel := Model{
		Patches:     convertPatchSet(model.Patches),
		Turtles:     convertTurtleSet(model.Turtles()),
		Links:       convertLinkSet(model.ShownLinks),
		Ticks:       model.Ticks,
		WorldWidth:  model.WorldWidth(),
		WorldHeight: model.WorldHeight(),
		MinPxCor:    model.MinPxCor(),
		MaxPxCor:    model.MaxPxCor(),
		MinPyCor:    model.MinPyCor(),
		MaxPyCor:    model.MaxPyCor(),
		MinPzCor:    model.MinPzCor(),
		MaxPzCor:    model.MaxPzCor(),
		Is3D:        model.Is3D(),
	}
	return &stateModel
}

func FromPatch(patch *model.Patch) *Patch {
	if patch == nil {
		return nil
	}

	statePatch := Patch{
		X:     patch.XCor(),
		Y:     patch.YCor(),
		Z:     patch.ZCor(),
		Color: FromColor(patch.Color),
	}
	return &statePatch
}

func FromColor(color model.Color) Color {
	stateColor := Color{
		Red:   color.Red,
		Green: color.Green,
		Blue:  color.Blue,
		Alpha: color.Alpha,
	}
	return stateColor
}

func convertPatchSet(patches *model.PatchAgentSet) []Patch {
	statePatches := make([]Patch, 0, patches.Count())
	patches.Ask(func(patch *model.Patch) {
		statePatches = append(statePatches, *FromPatch(patch))
	})
	return statePatches
}

func convertColor(color model.Color) Color {
	return FromColor(color)
}

func convertTurtleSet(turtles *model.TurtleAgentSet) []Turtle {
	stateTurtles := make([]Turtle, 0, turtles.Count())
	turtles.Ask(func(turtle *model.Turtle) {
		stateTurtle := Turtle{
			X:          turtle.XCor(),
			Y:          turtle.YCor(),
			Z:          turtle.ZCor(),
			Color:      convertColor(turtle.Color),
			Size:       turtle.GetSize(),
			Who:        turtle.Who(),
			Shape:      turtle.Shape,
			Heading:    turtle.GetHeading(),
			Label:      turtle.GetLabel(),
			LabelColor: convertColor(turtle.LabelColor),
		}
		stateTurtles = append(stateTurtles, stateTurtle)
	})
	return stateTurtles
}

func convertLinkSet(links *model.LinkAgentSet) []Link {
	stateLinks := make([]Link, 0, links.Count())
	links.Ask(func(link *model.Link) {
		if link.IsHidden() {
			return
		}

		if link.End1() == nil || link.End2() == nil {
			fmt.Println("Link has nil ends")
			return
		}

		stateLink := Link{
			End1:       link.End1().Who(),
			End2:       link.End2().Who(),
			Directed:   link.Directed(),
			End1X:      link.End1().XCor(),
			End1Y:      link.End1().YCor(),
			End1Z:      link.End1().ZCor(),
			End2X:      link.End2().XCor(),
			End2Y:      link.End2().YCor(),
			End2Z:      link.End2().ZCor(),
			End1Size:   link.End1().GetSize(),
			End2Size:   link.End2().GetSize(),
			Color:      convertColor(link.Color),
			Label:      link.Label,
			LabelColor: convertColor(link.LabelColor),
			Size:       link.Size,
			Hidden:     link.IsHidden(),
		}
		stateLinks = append(stateLinks, stateLink)
	})
	return stateLinks
}
