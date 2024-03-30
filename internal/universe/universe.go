package universe

import (
	"errors"
	"math/rand"
)

type Universe struct {
	Ticks   int
	TicksOn bool

	PatchesOwn map[string]interface{}            //additional variables for each patch
	TurtlesOwn map[string]interface{}            //additional variables for each turtle
	BreedsOwn  map[string]map[string]interface{} //additional variables for each breed. The first key is the breed name

	PatchesArray2D [][]*Patch
	Turtles        map[int]*Turtle            //all the turtles
	Breeds         map[string]map[int]*Turtle //turtles that are part of specific breeds

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

func (u *Universe) buildColors() {
	u.ColorHueMap = map[string]float64{
		"black":     0,
		"white":     9.9,
		"grey":      5,
		"gray":      5,
		"red":       15,
		"orange":    25,
		"brown":     35,
		"yellow":    45,
		"green":     55,
		"lime":      65,
		"turquoise": 75,
		"cyan":      85,
		"sky":       95,
		"blue":      105,
		"violet":    115,
		"magenta":   125,
		"pink":      135,
	}
}

//builds an array of patches and links them togethor
func (u *Universe) buildPatches() {
	u.PatchesArray2D = [][]*Patch{}
	for i := 0; i < u.WorldHeight; i++ {
		row := []*Patch{}
		for j := 0; j < u.WorldWidth; j++ {
			p := NewPatch(u.PatchesOwn, j+u.MinPxCor, i+u.MinPyCor)
			row = append(row, p)
		}
		u.PatchesArray2D = append(u.PatchesArray2D, row)
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

func (u *Universe) ClearGlobals() {
	for g := range u.GlobalBools {
		u.GlobalBools[g] = false
	}
	for g := range u.GlobalFloats {
		u.GlobalFloats[g] = 0
	}
}

func (u *Universe) ClearTicks() {
	u.TicksOn = false
}

func (u *Universe) ClearPatches() {
	for y := range u.PatchesArray2D {
		for x := range u.PatchesArray2D[y] {
			u.PatchesArray2D[y][x].Reset(u.PatchesOwn)
		}
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

func (u *Universe) SetDefaultShapeTurtles(shape string) {
	u.DefaultShapeTurtles = shape
}

func (u *Universe) SetDefaultShapeLinks(shape string) {
	u.DefaultShapeLinks = shape
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

	return u.PatchesArray2D[y-u.MinPyCor][x-u.MinPxCor]
}

func (u *Universe) OneOfInt(arr []int) interface{} {

	return arr[rand.Intn(len(arr))-1]
}

func (u *Universe) RandomAmount(n int) int {
	return rand.Intn(n)
}

func (u *Universe) Diffuse(patchVariable string, percent float64) error {

	if percent > 1 || percent < 0 {
		return errors.New("percent amount was outside bounds")
	}

	diffusions := make(map[*Patch]float64)

	//go through each patch and calculate the diffusion amount
	for y := range u.PatchesArray2D {
		for x := range u.PatchesArray2D[y] {
			currentPatch := u.PatchesArray2D[y][x]
			patchAmount := currentPatch.PatchesOwn[patchVariable].(float64)
			amountToGive := patchAmount * percent / 8
			diffusions[currentPatch] = amountToGive
		}
	}

	//go through each patch and get the new amount
	for y := range u.PatchesArray2D {
		for x := range u.PatchesArray2D[y] {
			currentPatch := u.PatchesArray2D[y][x]

			amountFromNeighbors := 0.0
			neighbors := u.getNeighbors(x, y)
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
	}

	return nil
}

//@TODO if we are wrapping around then it will always be 8
func (u *Universe) howManyNeighbors(x int, y int) int {

	neighborCount := 0

	hasNeighborsLeft := x > 0
	hasNeighborsRight := x < len(u.PatchesArray2D[0])-1
	hasNeighborsAbove := y > 0
	hasNeighborsBelow := y < len(u.PatchesArray2D)-1

	if hasNeighborsAbove {
		neighborCount++
	}

	if hasNeighborsBelow {
		neighborCount++
	}

	if hasNeighborsLeft {
		neighborCount++
	}

	if hasNeighborsRight {
		neighborCount++
	}

	if hasNeighborsAbove && hasNeighborsLeft {
		neighborCount++
	}

	if hasNeighborsAbove && hasNeighborsRight {
		neighborCount++
	}

	if hasNeighborsBelow && hasNeighborsLeft {
		neighborCount++
	}

	if hasNeighborsBelow && hasNeighborsRight {
		neighborCount++
	}

	return neighborCount
}

//@TODO check to see if we are wrapping around
func (u *Universe) getNeighbors(x int, y int) []*Patch {
	n := []*Patch{}

	left := u.safeGetPatch(x-1, y)
	if left != nil {
		n = append(n, left)
	}

	topLeft := u.safeGetPatch(x-1, y-1)
	if topLeft != nil {
		n = append(n, topLeft)
	}

	bottomLeft := u.safeGetPatch(x-1, y+1)
	if bottomLeft != nil {
		n = append(n, bottomLeft)
	}

	top := u.safeGetPatch(x, y-1)
	if top != nil {
		n = append(n, top)
	}

	topRight := u.safeGetPatch(x+1, y+1)
	if topRight != nil {
		n = append(n, topRight)
	}

	right := u.safeGetPatch(x+1, y)
	if right != nil {
		n = append(n, right)
	}

	bottomRight := u.safeGetPatch(x+1, y+1)
	if bottomRight != nil {
		n = append(n, bottomRight)
	}

	bottom := u.safeGetPatch(x, y+1)
	if bottom != nil {
		n = append(n, bottom)
	}

	return n
}

func (u *Universe) safeGetPatch(x int, y int) *Patch {
	if x < 0 || y < 0 || x > len(u.PatchesArray2D[0]) || y > len(u.PatchesArray2D) {
		return nil
	}

	return u.PatchesArray2D[y][x]
}
