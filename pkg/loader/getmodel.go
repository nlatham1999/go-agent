package loader

import (
	"encoding/base64"
	"fmt"

	"github.com/nlatham1999/go-agent/pkg/model"
)

func GetModel(model *model.Model) *Model {
	modelJson := Model{
		TurtleBreeds:         convertTurtleBreeds(model),
		DirectedLinkBreeds:   convertLinkBreeds(model.DirectedLinkBreeds()),
		UndirectedLinkBreeds: convertLinkBreeds(model.UndirectedLinkBreeds()),
		PatchProperties:      model.DefaultPatchProperties,
		TurtleProperties:     model.TurtleBreed("").DefaultProperties(),
		WrappingX:            model.WrappingX(),
		WrappingY:            model.WrappingY(),
		WorldWidth:           model.WorldWidth(),
		WorldHeight:          model.WorldHeight(),
		MinPxCor:             model.MinPxCor(),
		MaxPxCor:             model.MaxPxCor(),
		MinPyCor:             model.MinPyCor(),
		MaxPyCor:             model.MaxPyCor(),
		Patches:              convertPatchSet(model.Patches),
		Turtles:              convertTurtleSet(model.Turtles()),
		Links:                convertLinkSet(model.Links()),
		Ticks:                model.Ticks,
	}

	seed1, seed2, state := model.GetRandomState()
	modelJson.RandomSeed1 = seed1
	modelJson.RandomSeed2 = seed2
	modelJson.RandomState = base64.StdEncoding.EncodeToString(state)

	return &modelJson
}

func convertPatchSet(patches *model.PatchAgentSet) []Patch {
	apiPatches := make([]Patch, 0, patches.Count())
	patches.Ask(func(patch *model.Patch) {
		apiPatch := Patch{
			X:     patch.PXCor(),
			Y:     patch.PYCor(),
			Color: convertColor(patch.Color),
		}
		apiPatches = append(apiPatches, apiPatch)
	})
	return apiPatches
}

func convertColor(color model.Color) Color {
	apiColor := Color{
		Red:   color.Red,
		Green: color.Green,
		Blue:  color.Blue,
		Alpha: color.Alpha,
	}
	return apiColor
}

func convertTurtleSet(turtles *model.TurtleAgentSet) []Turtle {
	apiTurtles := make([]Turtle, 0, turtles.Count())
	turtles.Ask(func(turtle *model.Turtle) {
		apiTurtle := Turtle{
			X:          turtle.XCor(),
			Y:          turtle.YCor(),
			Color:      convertColor(turtle.Color),
			Size:       turtle.GetSize(),
			Who:        turtle.Who(),
			Shape:      turtle.Shape,
			Heading:    turtle.GetHeading(),
			Label:      turtle.GetLabel(),
			LabelColor: convertColor(turtle.LabelColor),
			Breed:      turtle.BreedName(),
		}
		apiTurtles = append(apiTurtles, apiTurtle)
	})
	return apiTurtles
}

func convertLinkSet(links *model.LinkAgentSet) []Link {
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
			Color:      convertColor(link.Color),
			Label:      link.Label,
			LabelColor: convertColor(link.LabelColor),
			Size:       link.Size,
			Hidden:     link.Hidden,
			Breed:      link.BreedName(),
		}
		apiLinks = append(apiLinks, apiLink)
	})
	return apiLinks
}

func convertTurtleBreeds(model *model.Model) []TurtleBreed {

	arr := []TurtleBreed{}

	breeds := model.TurtleBreeds()
	for _, breed := range breeds {
		if breed.Name() != "" {
			b := TurtleBreed{
				Name:         breed.Name(),
				Properties:   breed.DefaultProperties(),
				DefaultShape: breed.GetDefaultShape(),
			}
			arr = append(arr, b)
		}
	}

	return arr
}

func convertLinkBreeds(modelLinks []*model.LinkBreed) []LinkBreed {

	arr := []LinkBreed{}

	for _, breed := range modelLinks {
		if breed.Name() != "" {
			b := LinkBreed{
				Name:         breed.Name(),
				DefaultShape: breed.GetDefaultShape(),
				Directed:     breed.Directed(),
			}
			arr = append(arr, b)
		}
	}

	return arr
}
