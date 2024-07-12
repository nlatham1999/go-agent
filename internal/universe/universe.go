package universe

import (
	"errors"
	"math"
	"math/rand"
)

type Universe struct {
	Ticks   int
	TicksOn bool

	LinksOwn        map[string]interface{}            //additional variables for each link
	LinkBreedsOwn   map[string]map[string]interface{} //additional variables for each link breed. The first key is the breed name
	PatchesOwn      map[string]interface{}            //additional variables for each patch
	TurtlesOwn      map[string]interface{}            //additional variables for each turtle
	TurtleBreedsOwn map[string]map[string]interface{} //additional variables for each turtle breed. The first key is the breed name

	Patches             *PatchAgentSet
	Turtles             *TurtleAgentSet         //all the turtles
	Breeds              map[string]*TurtleBreed //turtles that are part of specific breeds
	Links               *LinkAgentSet           //all the links
	DirectedLinkBreeds  map[string]*LinkBreed
	UndirectedLinkBreed map[string]*LinkBreed

	PosOfPatches map[int]*Patch //map of patches by their index

	MaxPxCor    int
	MaxPyCor    int
	MinPxCor    int
	MinPyCor    int
	WorldWidth  int
	WorldHeight int
	wrapping    bool

	DefaultShapeTurtles string //the default shape for all turtles
	DefaultShapeLinks   string //the default shape for links

	turtlesWhoNumber int //who number of the next turtle to be created

	ColorHueMap map[string]float64

	GlobalFloats map[string]float64
	GlobalBools  map[string]bool
}

func NewUniverse(
	patchesOwn map[string]interface{},
	turtlesOwn map[string]interface{},
	turtleBreedsOwn map[string]map[string]interface{},
	turtleBreeds []string,
	directedLinkBreeds []string,
	undirectedLinkBreeds []string,
	wrapping bool,
) *Universe {
	maxPxCor := 15
	maxPyCor := 15
	minPxCor := -15
	minPyCor := -15

	universe := &Universe{
		MaxPxCor:        maxPxCor,
		MaxPyCor:        maxPyCor,
		MinPxCor:        minPxCor,
		MinPyCor:        minPyCor,
		WorldWidth:      maxPxCor - minPxCor + 1,
		WorldHeight:     maxPyCor - minPyCor + 1,
		PatchesOwn:      patchesOwn,
		TurtlesOwn:      turtlesOwn,
		TurtleBreedsOwn: turtleBreedsOwn,
		wrapping:        wrapping,
	}

	//construct turtle breeds
	turtleBreedsMap := make(map[string]*TurtleBreed)
	for i := 0; i < len(turtleBreeds); i++ {
		turtleBreedsMap[turtleBreeds[i]] = &TurtleBreed{
			Turtles: &TurtleAgentSet{
				turtles:      make(map[*Turtle]interface{}),
				whoToTurtles: map[int]*Turtle{},
			},
			DefaultShape: "",
		}
	}
	universe.Breeds = turtleBreedsMap

	//construct directed link breeds
	directedLinkBreedsMap := make(map[string]*LinkBreed)
	for i := 0; i < len(directedLinkBreeds); i++ {
		directedLinkBreedsMap[directedLinkBreeds[i]] = &LinkBreed{
			Links:        []*Link{},
			Directed:     true,
			DefaultShape: "",
		}
	}
	universe.DirectedLinkBreeds = directedLinkBreedsMap

	//construct undirected link breeds
	undirectedLinkBreedsMap := make(map[string]*LinkBreed)
	for i := 0; i < len(undirectedLinkBreeds); i++ {
		undirectedLinkBreedsMap[undirectedLinkBreeds[i]] = &LinkBreed{
			Links:        []*Link{},
			Directed:     false,
			DefaultShape: "",
		}
	}
	universe.UndirectedLinkBreed = undirectedLinkBreedsMap

	//construct general turtle set
	universe.Turtles = &TurtleAgentSet{
		turtles:      make(map[*Turtle]interface{}),
		whoToTurtles: map[int]*Turtle{},
	}

	//construct general link set
	universe.Links = &LinkAgentSet{
		links: []*Link{},
	}

	universe.buildPatches()

	return universe
}

// builds an array of patches and links them togethor
func (u *Universe) buildPatches() {
	u.Patches = &PatchAgentSet{
		patches: map[*Patch]interface{}{},
	}
	u.PosOfPatches = make(map[int]*Patch)
	for i := 0; i < u.WorldHeight; i++ {
		for j := 0; j < u.WorldWidth; j++ {
			p := NewPatch(u.PatchesOwn, j+u.MinPxCor, i+u.MinPyCor)
			u.Patches.patches[p] = nil
			u.PosOfPatches[j*u.WorldWidth+i] = p
		}
	}
}

// @TODO implement
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

// @TODO implement
func (u *Universe) ClearLinks() {

}

func (u *Universe) ClearTicks() {
	u.TicksOn = false
}

func (u *Universe) ClearPatches() {
	for patch := range u.Patches.patches {
		patch.Reset(u.PatchesOwn)
	}
}

// @TODO Implement
func (u *Universe) ClearDrawing() {

}

// @TODO Implement
func (u *Universe) ClearAllPlots() {

}

// @TODO Implement
func (u *Universe) ClearOutput() {

}

// @TODO Implement
func (u *Universe) ClearTurtles() {

}

// @TODO Implement
// idea is that if an empty string is passed then it will be for the general population
func (u *Universe) CreateOrderedTurtles(breed string, amount float64, operations []TurtleOperation) {

}

func (u *Universe) CreateTurtles(amount int, breed string, operations []TurtleOperation) error {

	agentSet := u.Turtles
	var agentSet2 *TurtleAgentSet = nil
	if breed != "" {
		breed, found := u.Breeds[breed]
		if !found {
			return errors.New("breed not found")
		}
		agentSet2 = breed.Turtles
	}

	end := amount + u.turtlesWhoNumber
	for u.turtlesWhoNumber < end {
		newTurtle := NewTurtle(u, u.turtlesWhoNumber, breed)

		agentSet.turtles[newTurtle] = nil
		agentSet.whoToTurtles[u.turtlesWhoNumber] = newTurtle

		if agentSet2 != nil {
			agentSet2.turtles[newTurtle] = nil
			agentSet2.whoToTurtles[u.turtlesWhoNumber] = newTurtle
		}

		for i := 0; i < len(operations); i++ {
			operations[i](newTurtle)
		}

		u.turtlesWhoNumber++
	}

	return nil
}

// @TODO implement
func (u *Universe) DieTurtle(turtle *Turtle) {
}

// @TODO implement
func (u *Universe) DieLink(link *Link) {
}

func (u *Universe) Diffuse(patchVariable string, percent float64) error {

	if percent > 1 || percent < 0 {
		return errors.New("percent amount was outside bounds")
	}

	diffusions := make(map[*Patch]float64)

	//go through each patch and calculate the diffusion amount
	for patch := range u.Patches.patches {
		patchAmount := patch.PatchesOwn[patchVariable].(float64)
		amountToGive := patchAmount * percent / 8
		diffusions[patch] = amountToGive
	}

	//go through each patch and get the new amount
	for patch := range u.Patches.patches {

		amountFromNeighbors := 0.0
		x := u.WorldHeight*patch.x + patch.y
		neighbors := u.Neighbors(x)
		if len(neighbors) > 8 || len(neighbors) < 3 {
			return errors.New("invalid amount of neighbors")
		}
		for n := range neighbors {
			amountFromNeighbors += diffusions[neighbors[n]]
		}

		patchAmount := patch.PatchesOwn[patchVariable].(float64)
		amountToKeep := 1 - (patchAmount * percent) + (float64(8-len(neighbors)) * (patchAmount * percent / 8))

		patch.PatchesOwn[patchVariable] = amountToKeep + amountFromNeighbors
	}

	return nil
}

// @TODO implement
func (u *Universe) Diffuse4(patchVariable string, percent float64) error {
	return nil
}

func (u *Universe) LayoutCircle(turtles []*Turtle, radius float64) {
	amount := len(turtles)
	for i := 0; i < amount; i++ {
		agent := turtles[i]
		agent.SetXY(radius*math.Cos(2*math.Pi*float64(i)/float64(amount)), radius*math.Sin(2*math.Pi*float64(i)/float64(amount)))
		agent.Heading = 2 * math.Pi * float64(i) / float64(amount)
	}
}

// @TODO implement
func (u *Universe) LayoutRadial(turtles []*Turtle, links []*Link, root *Turtle) {

}

// @TODO implement
func (u *Universe) LayoutSpring(turtles []*Turtle, links []*Link, springConstant float64, springLength float64, repulsionConstant float64) {

}

// @TODO implement
func (u *Universe) LayoutTutte(turtles []*Turtle, links []*Link, radius float64) {

}

// @TODO implement
func (u *Universe) Link(breed string, turtle1 int, turtle2 int) *Link {
	return nil
}

// @TODO implement
func (u *Universe) LinkDirected(breed string, turtle1 int, turtle2 int) *Link {
	return nil
}

// @TODO implement
func (u *Universe) LinkShapes() []string {
	return []string{}
}

func (u *Universe) getPatchAtCoords(x int, y int) *Patch {
	if x < u.MinPxCor || x > u.MaxPxCor || y < u.MinPyCor || y > u.MaxPyCor {
		return nil
	}

	offsetX := x - u.MinPxCor
	offsetY := y - u.MinPyCor

	pos := offsetY*u.WorldWidth + offsetX

	return u.PosOfPatches[pos]
}

func (u *Universe) OneOfInt(arr []int) interface{} {

	return arr[rand.Intn(len(arr))-1]
}

func (u *Universe) RandomAmount(n int) int {
	return rand.Intn(n)
}

func (u *Universe) topLeftNeighbor(x int) *Patch {
	return u.safeGetPatch(x - u.WorldWidth - 1)
}

func (u *Universe) topNeighbor(x int) *Patch {
	return u.safeGetPatch(x - u.WorldWidth)
}

func (u *Universe) topRightNeighbor(x int) *Patch {
	return u.safeGetPatch(x - u.WorldWidth + 1)
}

func (u *Universe) leftNeighbor(x int) *Patch {
	return u.safeGetPatch(x - 1)
}

func (u *Universe) rightNeighbor(x int) *Patch {
	return u.safeGetPatch(x + 1)
}

func (u *Universe) bottomLeftNeighbor(x int) *Patch {
	return u.safeGetPatch(x + u.WorldWidth - 1)
}

func (u *Universe) bottomNeighbor(x int) *Patch {
	return u.safeGetPatch(x + u.WorldWidth)
}

func (u *Universe) bottomRightNeighbor(x int) *Patch {
	return u.safeGetPatch(x + u.WorldWidth + 1)
}

// @TODO check to see if we are wrapping around
func (u *Universe) Neighbors(x int) []*Patch {
	n := []*Patch{}

	topLeft := u.topLeftNeighbor(x)
	if topLeft != nil {
		n = append(n, topLeft)
	}

	bottomLeft := u.bottomLeftNeighbor(x)
	if bottomLeft != nil {
		n = append(n, bottomLeft)
	}

	top := u.topNeighbor(x)
	if top != nil {
		n = append(n, top)
	}

	topRight := u.topRightNeighbor(x)
	if topRight != nil {
		n = append(n, topRight)
	}

	right := u.rightNeighbor(x)
	if right != nil {
		n = append(n, right)
	}

	bottomRight := u.bottomRightNeighbor(x)
	if bottomRight != nil {
		n = append(n, bottomRight)
	}

	bottom := u.bottomNeighbor(x)
	if bottom != nil {
		n = append(n, bottom)
	}

	return n
}

// @TODO check to see if we are wrapping around
func (u *Universe) Neighbors4(x int) []*Patch {
	n := []*Patch{}

	top := u.topNeighbor(x)
	if top != nil {
		n = append(n, top)
	}

	left := u.leftNeighbor(x)
	if left != nil {
		n = append(n, left)
	}

	right := u.rightNeighbor(x)
	if right != nil {
		n = append(n, right)
	}

	bottom := u.bottomNeighbor(x)
	if bottom != nil {
		n = append(n, bottom)
	}

	return n
}

func (u *Universe) safeGetPatch(x int) *Patch {
	if x < 0 || x > len(u.Patches.patches) {
		return nil
	}

	return u.PosOfPatches[x]
}

func (u *Universe) Patch(pxcor float64, pycor float64) *Patch {
	//round to get x and y
	x := int(math.Round(pxcor))
	y := int(math.Round(pycor))

	return u.getPatchAtCoords(x, y)
}

func (u *Universe) ResetTicks() {
	u.Ticks = 0
}

// @TODO implement
func (u *Universe) ResetTimer() {

}

func (u *Universe) ResizeWorld(minPxcor int, maxPxcor int, minPycor int, maxPycor int) {
	u.MinPxCor = minPxcor
	u.MaxPxCor = maxPxcor
	u.MinPyCor = minPycor
	u.MaxPyCor = maxPycor
	u.WorldWidth = maxPxcor - minPxcor + 1
	u.WorldHeight = maxPycor - minPycor + 1

	u.buildPatches()
}

func (u *Universe) SetDefaultShapeLinks(shape string) {
	u.DefaultShapeLinks = shape
}

func (u *Universe) SetDefaultShapeTurtles(shape string) {
	u.DefaultShapeTurtles = shape
}

func (u *Universe) SetDefaultShapeLinkBreed(breed string, shape string) {
	u.DirectedLinkBreeds[breed].DefaultShape = shape
}

func (u *Universe) SetDefaultShapeTurtleBreed(breed string, shape string) {
	u.Breeds[breed].DefaultShape = shape
}

func (u *Universe) Tick() {
	if u.TicksOn {
		u.Ticks++
	}
}

func (u *Universe) TickAdvance(amount int) {
	if u.TicksOn {
		u.Ticks += amount
	}
}

// provides a turtle from the universe given a breed and who number
// if the breed is empty then selects from the general population
// if the breed or who number is not found then returns nil
func (u *Universe) Turtle(breed string, who int) *Turtle {
	if breed == "" {
		return u.Turtles.whoToTurtles[who]
	} else {
		if u.Breeds[breed] == nil {
			return nil //breed not found
		}
		return u.Breeds[breed].Turtles.whoToTurtles[who]
	}
}

// @TODO implement
func (u *Universe) TurtlesAt(breed string, pxcor float64, pycor float64) *TurtleAgentSet {
	x := int(math.Round(pxcor))
	y := int(math.Round(pycor))

	patch := u.getPatchAtCoords(x, y)

	if patch == nil {
		return nil
	}

	return nil
}

// @TODO implement
func (u *Universe) TurtlesOnPatch(patch *Patch) *TurtleAgentSet {
	return nil
}

// @TODO implement
func (u *Universe) TurtlesOnPatches(patches *PatchAgentSet) *TurtleAgentSet {
	return nil
}

// @TODO implement
func (u *Universe) TurtlesWithTurtle(turtle *Turtle) *TurtleAgentSet {
	return nil
}

// @TODO implement
func (u *Universe) TurtlesWithTurtles(turtles *TurtleAgentSet) *TurtleAgentSet {
	return nil
}
