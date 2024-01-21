package universe

import (
	"github.com/nlatham1999/go-agent/internal/breed"
	patch "github.com/nlatham1999/go-agent/internal/patches"
	turtle "github.com/nlatham1999/go-agent/internal/turtles"
)

type Universe struct {
	TickCounter int
	TicksOn     bool

	PatchesOwn map[string]interface{}  //additional variables for each patch
	Breeds     map[string]*breed.Breed //the different types of breeds

	MaxPxCor    int
	MaxPyCor    int
	MinPxCor    int
	MinPyCor    int
	WorldWidth  int
	WorldHeight int

	DefaultShapeTurtles string //the default shape for all turtles
	DefaultShapeLinks   string //the default shape for links

	Turtles          map[int]*turtle.Turtle //all the turtles
	TurtlesWhoNumber int

	Patches      []*patch.Patch //all the patches
	PatchesArray [][]*patch.Patch
}

func NewUniverse(patchesOwn map[string]interface{}) *Universe {
	maxPxCor := 15
	maxPyCor := 15
	minPxCor := -15
	minPyCor := -15
	universe := &Universe{
		MaxPxCor:    maxPxCor,
		MaxPyCor:    maxPyCor,
		MinPxCor:    minPxCor,
		MinPyCor:    minPyCor,
		WorldWidth:  maxPxCor - minPxCor + 1,
		WorldHeight: maxPyCor - minPyCor + 1,
		PatchesOwn:  patchesOwn,
	}

	universe.buildPatches()

	return universe
}

//builds an array of patches and links them togethor
func (u *Universe) buildPatches() {
	u.PatchesArray = [][]*patch.Patch{}
	for i := 0; i < u.WorldHeight; i++ {
		row := []*patch.Patch{}
		for j := 0; j < u.WorldWidth; j++ {
			p := patch.NewPatch(u.PatchesOwn)
			row = append(row, p)
		}
		u.PatchesArray = append(u.PatchesArray, row)
	}
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

func (u *Universe) ClearTicks() {
	u.TicksOn = false
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

func (u *Universe) ResetTicks() {
	u.TicksOn = true
	u.TickCounter = 0
}
