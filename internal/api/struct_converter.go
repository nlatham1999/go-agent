package api

import (
	"github.com/nlatham1999/go-agent/internal/model"
)

func convertModelToApiModel(model *model.Model) *Model {
	apiModel := Model{
		Patches:          convertPatchSetToApiPatchSet(model.Patches),
		Turtles:          convertTurtleSetToApiTurtleSet(model.Turtles("")),
		Links:            convertLinkSetToApiLinkSet(model.Links),
		DynamicVariables: model.DynamicVariables,
		Ticks:            model.Ticks,
		Width:            model.WorldWidth,
		Height:           model.WorldHeight,
	}
	return &apiModel
}

func convertPatchSetToApiPatchSet(patches *model.PatchAgentSet) []Patch {
	apiPatches := make([]Patch, 0, patches.Count())
	for _, patch := range patches.List() {
		apiPatch := Patch{
			X:     patch.PXCor(),
			Y:     patch.PYCor(),
			Color: convertColorToApiColor(&patch.PColor),
		}
		apiPatches = append(apiPatches, apiPatch)
	}
	return apiPatches
}

func convertColorToApiColor(color *model.Color) Color {
	apiColor := Color{
		R: color.Red,
		G: color.Green,
		B: color.Blue,
		A: color.Alpha,
	}
	return apiColor
}

func convertTurtleSetToApiTurtleSet(turtles *model.TurtleAgentSet) []Turtle {
	apiTurtles := make([]Turtle, 0, turtles.Count())
	for _, turtle := range turtles.List() {
		apiTurtle := Turtle{
			X:       turtle.XCor(),
			Y:       turtle.YCor(),
			Color:   convertColorToApiColor(&turtle.Color),
			Size:    turtle.Size,
			Who:     turtle.Who(),
			Shape:   turtle.Shape,
			Heading: turtle.GetHeading(),
		}
		apiTurtles = append(apiTurtles, apiTurtle)
	}
	return apiTurtles
}

func convertLinkSetToApiLinkSet(links *model.LinkAgentSet) []Link {
	apiLinks := make([]Link, 0, links.Count())
	for _, link := range links.List() {
		apiLink := Link{
			End1:     link.End1().Who(),
			End2:     link.End2().Who(),
			Directed: link.Directed,
		}
		apiLinks = append(apiLinks, apiLink)
	}
	return apiLinks
}
