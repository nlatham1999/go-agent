package universe

import (
	"github.com/nlatham1999/go-agent/internal/breed"
	turtle "github.com/nlatham1999/go-agent/internal/turtles"
)

type Universe struct {
	TickCounter int
	PatchesOwn  map[string]interface{}  //additional variables for each patch
	Breeds      map[string]*breed.Breed //the different types of breeds

	DefaultShapeTurtles string //the default shape for all turtles
	DefaultShapeLinks   string //the default shape for links

	Turtles map[int]*turtle.Turtle //all the turtles

}

func NewUniverse() *Universe {
	return &Universe{}
}

func (u *Universe) ClearAll() {
	u.ClearGlobals()
	u.ClearTicks()
	u.ClearPatches()
	u.ClearDrawing()
	u.ClearAllPlots()
	u.ClearOutput()
}

//@TODO Implement
func (u *Universe) ClearGlobals() {

}

//@TODO Implement
func (u *Universe) ClearTicks() {

}

//@TODO Implement
func (u *Universe) ClearPatches() {

}

//@TODO Implement
func (u *Universe) ClearDrawing() {

}

//@TODO Implement
func (u *Universe) ClearAllPlots() {

}

//@TODO Implement
func (u *Universe) ClearOutput() {

}

func (u *Universe) SetDefaultShapeTurtles(shape string) {
	u.DefaultShapeTurtles = shape
}

func (u *Universe) SetDefaultShapeLinks(shape string) {
	u.DefaultShapeLinks = shape
}

func (u *Universe) SetDefaultShapeBreed(shape string, breedType string) {
	u.Breeds[breedType].DefaultShape = shape
}

func (u *Universe) CreateTurtles(amount int, operations []turtle.TurtleOperation) {
	startIndex := len(u.Turtles)
	end := amount + startIndex
	for startIndex < end {
		newTurtle := turtle.NewTurtle(startIndex)

		for i := 0; i < len(operations); i++ {
			operations[i](newTurtle)
		}

		u.Turtles[startIndex] = newTurtle

		startIndex++
	}
}

func (u *Universe) ResetTickCounter() {
	u.TickCounter = 0
}
