package universe

import "math/rand"

type Universe struct {
	Ticks   int
	TicksOn bool

	PatchesOwn map[string]interface{} //additional variables for each patch
	Breeds     map[string]*Breed      //the different types of breeds

	MaxPxCor    int
	MaxPyCor    int
	MinPxCor    int
	MinPyCor    int
	WorldWidth  int
	WorldHeight int

	DefaultShapeTurtles string //the default shape for all turtles
	DefaultShapeLinks   string //the default shape for links

	Turtles          map[int]*Turtle //all the turtles
	TurtlesWhoNumber int

	Patches      []*Patch //all the patches
	PatchesArray [][]*Patch

	ColorHueMap map[string]float64
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
	u.PatchesArray = [][]*Patch{}
	for i := 0; i < u.WorldHeight; i++ {
		row := []*Patch{}
		for j := 0; j < u.WorldWidth; j++ {
			p := NewPatch(u.PatchesOwn, j+u.MinPxCor, i+u.MinPyCor)
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

func (u *Universe) getPatchAtCoords(x int, y int) *Patch {
	if x < u.MinPxCor || x > u.MaxPxCor || y < u.MinPyCor || y > u.MaxPyCor {
		return nil
	}

	return u.PatchesArray[y-u.MinPyCor][x-u.MinPxCor]
}

func (u *Universe) OneOfInt(arr []int) interface{} {

	return arr[rand.Intn(len(arr))-1]
}

func (u *Universe) RandomAmount(n int) int {
	return rand.Intn(n)
}
