package loader

import (
	"github.com/nlatham1999/go-agent/pkg/model"
)

func SetModel(modelJson *Model) *model.Model {

	// build the turtle breeds
	turtleBreeds := []*model.TurtleBreed{}
	for _, breed := range modelJson.TurtleBreeds {
		newBreed := model.NewTurtleBreed(breed.Name, breed.DefaultShape, breed.Properties)
		turtleBreeds = append(turtleBreeds, newBreed)
	}

	// build the directed link breeds
	directedLinkBreeds := []*model.LinkBreed{}
	for _, breed := range modelJson.DirectedLinkBreeds {
		newBreed := model.NewLinkBreed(breed.Name)
		directedLinkBreeds = append(directedLinkBreeds, newBreed)
	}

	// build the undirected link breeds
	undirectedLinkBreeds := []*model.LinkBreed{}
	for _, breed := range modelJson.UndirectedLinkBreeds {
		newBreed := model.NewLinkBreed(breed.Name)
		undirectedLinkBreeds = append(undirectedLinkBreeds, newBreed)
	}

	// build the model
	modelSettings := model.ModelSettings{
		PatchProperties:      modelJson.PatchProperties,
		TurtleProperties:     modelJson.TurtleProperties,
		TurtleBreeds:         turtleBreeds,
		DirectedLinkBreeds:   directedLinkBreeds,
		UndirectedLinkBreeds: undirectedLinkBreeds,
		WrappingX:            modelJson.WrappingX,
		WrappingY:            modelJson.WrappingY,
		MinPxCor:             modelJson.MinPxCor,
		MaxPxCor:             modelJson.MaxPxCor,
		MinPyCor:             modelJson.MinPyCor,
		MaxPyCor:             modelJson.MaxPyCor,
		RandomSeed:           modelJson.RandomSeed1,
		RandomSeed2:          modelJson.RandomSeed2,
	}
	builtModel := model.NewModel(modelSettings)

	//load the random state
	builtModel.SetRandomState(modelJson.RandomSeed1, modelJson.RandomSeed2, []byte(modelJson.RandomState))

	// set all the patches
	for _, patch := range modelJson.Patches {
		p := builtModel.Patch(float64(patch.X), float64(patch.Y))
		p.Color.SetColorRGBA(patch.Color.Red, patch.Color.Green, patch.Color.Blue, patch.Color.Alpha)
		for key, val := range patch.Properties {
			p.SetProperty(key, val)
		}
	}

	// create basic turtles
	builtModel.CreateTurtles(len(modelJson.Turtles), nil)

	// set all the turtles
	for _, turtle := range modelJson.Turtles {
		t := builtModel.Turtle(turtle.Who)
		t.SetXY(turtle.X, turtle.Y)
		t.Color.SetColorRGBA(turtle.Color.Red, turtle.Color.Green, turtle.Color.Blue, turtle.Color.Alpha)
		t.SetSize(turtle.Size)
		t.Shape = turtle.Shape
		t.SetHeading(turtle.Heading)
		t.SetLabel(turtle.Label)
		t.LabelColor = model.Color{
			Red:   turtle.LabelColor.Red,
			Green: turtle.LabelColor.Green,
			Blue:  turtle.LabelColor.Blue,
			Alpha: turtle.LabelColor.Alpha,
		}
		for key, val := range turtle.Properties {
			t.SetProperty(key, val)
		}
		if turtle.Breed != "" {
			breed := builtModel.TurtleBreed(turtle.Breed)
			t.SetBreed(breed)
		}
	}

	// set all the links
	for _, link := range modelJson.Links {
		end1 := builtModel.Turtle(link.End1)
		end2 := builtModel.Turtle(link.End2)

		if link.Directed {
			breed := builtModel.DirectedLinkBreed(link.Breed)
			end1.CreateLinkToTurtle(breed, end2, func(l *model.Link) {
				l.Color.SetColorRGBA(link.Color.Red, link.Color.Green, link.Color.Blue, link.Color.Alpha)
				l.Label = link.Label
				l.LabelColor.SetColorRGBA(link.LabelColor.Red, link.LabelColor.Green, link.LabelColor.Blue, link.LabelColor.Alpha)
				l.Size = link.Size
				l.Hidden = link.Hidden
			})
		} else {
			breed := builtModel.UndirectedLinkBreed(link.Breed)
			end1.CreateLinkWithTurtle(breed, end2, func(l *model.Link) {
				l.Color.SetColorRGBA(link.Color.Red, link.Color.Green, link.Color.Blue, link.Color.Alpha)
				l.Label = link.Label
				l.LabelColor.SetColorRGBA(link.LabelColor.Red, link.LabelColor.Green, link.LabelColor.Blue, link.LabelColor.Alpha)
				l.Size = link.Size
				l.Hidden = link.Hidden
			})
		}
	}

	return builtModel
}
