package api

import (
	"fmt"

	"github.com/nlatham1999/go-agent/pkg/model"
)

func convertModelToApiModel(model *model.Model) *Model {
	apiModel := Model{
		Patches:     convertPatchSetToApiPatchSet(model.Patches),
		Turtles:     convertTurtleSetToApiTurtleSet(model.Turtles()),
		Links:       convertLinkSetToApiLinkSet(model.Links()),
		Ticks:       model.Ticks,
		WorldWidth:  model.WorldWidth(),
		WorldHeight: model.WorldHeight(),
		MinPxCor:    model.MinPxCor(),
		MaxPxCor:    model.MaxPxCor(),
		MinPyCor:    model.MinPyCor(),
		MaxPyCor:    model.MaxPyCor(),
	}
	return &apiModel
}

func convertPatchSetToApiPatchSet(patches *model.PatchAgentSet) []Patch {
	apiPatches := make([]Patch, 0, patches.Count())
	patches.Ask(func(patch *model.Patch) {
		apiPatch := Patch{
			X:     patch.PXCor(),
			Y:     patch.PYCor(),
			Color: convertColorToApiColor(patch.Color),
		}
		apiPatches = append(apiPatches, apiPatch)
	})
	return apiPatches
}

func convertColorToApiColor(color model.Color) Color {
	apiColor := Color{
		Red:   color.Red,
		Green: color.Green,
		Blue:  color.Blue,
		Alpha: color.Alpha,
	}
	return apiColor
}

func convertTurtleSetToApiTurtleSet(turtles *model.TurtleAgentSet) []Turtle {
	apiTurtles := make([]Turtle, 0, turtles.Count())
	turtles.Ask(func(turtle *model.Turtle) {
		apiTurtle := Turtle{
			X:          turtle.XCor(),
			Y:          turtle.YCor(),
			Color:      convertColorToApiColor(turtle.Color),
			Size:       turtle.GetSize(),
			Who:        turtle.Who(),
			Shape:      turtle.Shape,
			Heading:    turtle.GetHeading(),
			Label:      turtle.GetLabel(),
			LabelColor: convertColorToApiColor(turtle.LabelColor),
		}
		apiTurtles = append(apiTurtles, apiTurtle)
	})
	return apiTurtles
}

func convertLinkSetToApiLinkSet(links *model.LinkAgentSet) []Link {
	apiLinks := make([]Link, 0, links.Count())
	links.Ask(func(link *model.Link) {
		if link.End1() == nil || link.End2() == nil {
			fmt.Println("Link has nil ends")
			return
		}
		apiLink := Link{
			End1:       link.End1().Who(),
			End2:       link.End2().Who(),
			Directed:   link.Directed(),
			End1X:      link.End1().XCor(),
			End1Y:      link.End1().YCor(),
			End2X:      link.End2().XCor(),
			End2Y:      link.End2().YCor(),
			End1Size:   link.End1().GetSize(),
			End2Size:   link.End2().GetSize(),
			Color:      convertColorToApiColor(link.Color),
			Label:      link.Label,
			LabelColor: convertColorToApiColor(link.LabelColor),
			Size:       link.Size,
			Hidden:     link.Hidden,
		}
		apiLinks = append(apiLinks, apiLink)
	})
	return apiLinks
}
