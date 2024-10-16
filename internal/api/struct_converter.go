package api

import (
	"fmt"

	"github.com/nlatham1999/go-agent/internal/model"
)

func convertModelToApiModel(model *model.Model) *Model {
	apiModel := Model{
		Patches:          convertPatchSetToApiPatchSet(model.Patches),
		Turtles:          convertTurtleSetToApiTurtleSet(model.Turtles("")),
		Links:            convertLinkSetToApiLinkSet(model.Links()),
		DynamicVariables: model.Globals,
		Ticks:            model.Ticks,
		WorldWidth:       model.WorldWidth(),
		WorldHeight:      model.WorldHeight(),
		MinPxCor:         model.MinPxCor(),
		MaxPxCor:         model.MaxPxCor(),
		MinPyCor:         model.MinPyCor(),
		MaxPyCor:         model.MaxPyCor(),
	}
	return &apiModel
}

func convertPatchSetToApiPatchSet(patches *model.PatchAgentSet) []Patch {
	apiPatches := make([]Patch, 0, patches.Count())
	for _, patch := range patches.List() {
		apiPatch := Patch{
			X:     patch.PXCor(),
			Y:     patch.PYCor(),
			Color: convertColorToApiColor(patch.PColor),
		}
		apiPatches = append(apiPatches, apiPatch)
	}
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
	for _, turtle := range turtles.ListSorted() {
		apiTurtle := Turtle{
			X:          turtle.XCor(),
			Y:          turtle.YCor(),
			Color:      convertColorToApiColor(turtle.Color),
			Size:       turtle.GetSize(),
			Who:        turtle.Who(),
			Shape:      turtle.Shape,
			Heading:    turtle.GetHeading(),
			Label:      turtle.GetLabel(),
			LabelColor: convertColorToApiColor(turtle.GetLabelColor()),
		}
		apiTurtles = append(apiTurtles, apiTurtle)
	}
	return apiTurtles
}

func convertLinkSetToApiLinkSet(links *model.LinkAgentSet) []Link {
	apiLinks := make([]Link, 0, links.Count())
	for _, link := range links.List() {
		if link.End1() == nil || link.End2() == nil {
			fmt.Println("Link has nil ends")
			return nil
		}
		apiLink := Link{
			End1:       link.End1().Who(),
			End2:       link.End2().Who(),
			Directed:   link.Directed,
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
		}
		apiLinks = append(apiLinks, apiLink)
	}
	return apiLinks
}
