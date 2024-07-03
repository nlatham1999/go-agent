package universe

import (
	"errors"
	"math/rand"
)

const (
	//color constants
	Black     float64 = 0
	White     float64 = 9.9
	Grey      float64 = 5
	Gray      float64 = 5
	Red       float64 = 15
	Orange    float64 = 25
	Brown     float64 = 35
	Yellow    float64 = 45
	Green     float64 = 55
	Lime      float64 = 65
	Turquoise float64 = 75
	Cyan      float64 = 85
	Sky       float64 = 95
	Blue      float64 = 105
	Violet    float64 = 115
	Magenta   float64 = 125
	Pink      float64 = 135
)

type Universe struct {
	Ticks   int
	TicksOn bool

	PatchesOwn map[string]interface{}            //additional variables for each patch
	TurtlesOwn map[string]interface{}            //additional variables for each turtle
	BreedsOwn  map[string]map[string]interface{} //additional variables for each breed. The first key is the breed name

	Patches []*Patch
	Turtles []*Turtle          //all the turtles
	Breeds  map[string]*Turtle //turtles that are part of specific breeds

	MaxPxCor    int
	MaxPyCor    int
	MinPxCor    int
	MinPyCor    int
	WorldWidth  int
	WorldHeight int

	DefaultShapeTurtles string //the default shape for all turtles
	DefaultShapeLinks   string //the default shape for links

	TurtlesWhoNumber int

	ColorHueMap map[string]float64

	GlobalFloats map[string]float64
	GlobalBools  map[string]bool

	Base
}

func NewUniverse(patchesOwn map[string]interface{}, turtlesOwn map[string]interface{}, breedsOwn map[string]map[string]interface{}) *Universe {
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
		TurtlesOwn:  turtlesOwn,
		BreedsOwn:   breedsOwn,
	}

	universe.buildPatches()

	return universe
}

//builds an array of patches and links them togethor
func (u *Universe) buildPatches() {
	u.Patches = []*Patch{}
	for i := 0; i < u.WorldHeight; i++ {
		row := []*Patch{}
		for j := 0; j < u.WorldWidth; j++ {
			p := NewPatch(u.PatchesOwn, j+u.MinPxCor, i+u.MinPyCor)
			row = append(row, p)
		}
		u.Patches = append(u.Patches, row...)
	}
}

//@TODO implement
func (u *Universe) AllLinks(agentset LinkSet, operation LinkBoolOperation) bool {
	return false
}

//@TODO implement
func (u *Universe) AllPatches(agentset PatchSet, operation PatchBoolOperation) bool {
	return false
}

//@TODO implement
func (u *Universe) AllTurtles(agentset TurtleSet, operation TurtleBoolOperation) bool {
	return false
}

//@TODO implement
func (u *Universe) AnyLinks(agentset LinkSet, operation LinkBoolOperation) bool {
	return false
}

//@TODO implement
func (u *Universe) AnyPatches(agentset PatchSet, operation PatchBoolOperation) bool {
	return false
}

//@TODO implement
func (u *Universe) AnyTurtles(agentset TurtleSet, operation TurtleBoolOperation) bool {
	return false
}

//@TODO implement
func (u *Universe) BothEnds(link *Link) []*Turtle {
	return nil
}

func (u *Universe) ClearAll() {
	u.ClearGlobals()
	u.ClearTicks()
	u.ClearPatches()
	u.ClearDrawing()
	u.ClearAllPlots()
	u.ClearOutput()
}

func (u *Universe) ClearGlobals() {
	for g := range u.GlobalBools {
		u.GlobalBools[g] = false
	}
	for g := range u.GlobalFloats {
		u.GlobalFloats[g] = 0
	}
}

//@TODO implement
func (u *Universe) ClearLinks() {

}

func (u *Universe) ClearTicks() {
	u.TicksOn = false
}

func (u *Universe) ClearPatches() {
	for x := range u.Patches {
		u.Patches[x].Reset(u.PatchesOwn)
	}
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

//@TODO Implement
func (u *Universe) ClearTurtles() {

}

//@TODO Implement
//idea is that if an empty string is passed then it will be for the general population
func (u *Universe) CreateOrderedTurtles(breed string, amount float64, operations []TurtleOperation) {

}

func (u *Universe) CreateTurtles(amount int, operations []TurtleOperation) {
	startIndex := len(u.Turtles)
	end := amount + startIndex
	for startIndex < end {
		newTurtle := NewTurtle(startIndex)

		for i := 0; i < len(operations); i++ {
			operations[i](newTurtle)
		}

		u.Turtles[startIndex] = newTurtle

		startIndex++
	}
}

//@TODO implement
func (u *Universe) DieTurtle(turtle *Turtle) {
}

//@TODO implement
func (u *Universe) DieLink(link *Link) {
}

func (u *Universe) Diffuse(patchVariable string, percent float64) error {

	if percent > 1 || percent < 0 {
		return errors.New("percent amount was outside bounds")
	}

	diffusions := make(map[*Patch]float64)

	//go through each patch and calculate the diffusion amount
	for x := range u.Patches {
		currentPatch := u.Patches[x]
		patchAmount := currentPatch.PatchesOwn[patchVariable].(float64)
		amountToGive := patchAmount * percent / 8
		diffusions[currentPatch] = amountToGive
	}

	//go through each patch and get the new amount
	for x := range u.Patches {
		currentPatch := u.Patches[x]

		amountFromNeighbors := 0.0
		neighbors := u.getNeighbors(x)
		if len(neighbors) > 8 || len(neighbors) < 3 {
			return errors.New("invalid amount of neighbors")
		}
		for n := range neighbors {
			amountFromNeighbors += diffusions[neighbors[n]]
		}

		patchAmount := currentPatch.PatchesOwn[patchVariable].(float64)
		amountToKeep := 1 - (patchAmount * percent) + (float64(8-len(neighbors)) * (patchAmount * percent / 8))

		currentPatch.PatchesOwn[patchVariable] = amountToKeep + amountFromNeighbors
	}

	return nil
}

//@TODO implement
func (u *Universe) Diffuse4(patchVariable string, percent float64) error {
	return nil
}

//@TODO implement
func (u *Universe) SetDefaultShapeTurtles(shape string) {
	u.DefaultShapeTurtles = shape
}

func (u *Universe) SetDefaultShapeLinks(shape string) {
	u.DefaultShapeLinks = shape
}
func (u *Universe) ResetTicks() {
	u.TicksOn = true
	u.Ticks = 0
}

func (u *Universe) Tick() {
	if u.TicksOn {
		u.Ticks++
	}
}

func (u *Universe) getPatchAtCoords(x int, y int) *Patch {
	if x < u.MinPxCor || x > u.MaxPxCor || y < u.MinPyCor || y > u.MaxPyCor {
		return nil
	}

	offsetX := x - u.MinPxCor
	offsetY := y - u.MinPyCor

	pos := offsetY*u.WorldWidth + offsetX

	return u.Patches[pos]
}

func (u *Universe) OneOfInt(arr []int) interface{} {

	return arr[rand.Intn(len(arr))-1]
}

func (u *Universe) RandomAmount(n int) int {
	return rand.Intn(n)
}

//@TODO check to see if we are wrapping around
func (u *Universe) getNeighbors(x int) []*Patch {
	n := []*Patch{}

	topLeftPos := x - u.WorldWidth - 1
	topPos := x - u.WorldWidth
	topRightPos := x - u.WorldWidth + 1
	leftPos := x - 1
	rightPos := x + 1
	bottomLeftPos := x + u.WorldWidth - 1
	bottomPos := x + u.WorldWidth
	bottomRightPos := x + u.WorldWidth + 1
	left := u.safeGetPatch(leftPos)
	if left != nil {
		n = append(n, left)
	}

	topLeft := u.safeGetPatch(topLeftPos)
	if topLeft != nil {
		n = append(n, topLeft)
	}

	bottomLeft := u.safeGetPatch(bottomLeftPos)
	if bottomLeft != nil {
		n = append(n, bottomLeft)
	}

	top := u.safeGetPatch(topPos)
	if top != nil {
		n = append(n, top)
	}

	topRight := u.safeGetPatch(topRightPos)
	if topRight != nil {
		n = append(n, topRight)
	}

	right := u.safeGetPatch(rightPos)
	if right != nil {
		n = append(n, right)
	}

	bottomRight := u.safeGetPatch(bottomRightPos)
	if bottomRight != nil {
		n = append(n, bottomRight)
	}

	bottom := u.safeGetPatch(bottomPos)
	if bottom != nil {
		n = append(n, bottom)
	}

	return n
}

func (u *Universe) safeGetPatch(x int) *Patch {
	if x < 0 || x > len(u.Patches) {
		return nil
	}

	return u.Patches[x]
}
