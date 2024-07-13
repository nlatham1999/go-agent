package model

import (
	"errors"
	"math"
	"math/rand"
)

type Model struct {
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

	posOfPatches map[int]*Patch  //map of patches by their index
	whoToTurtles map[int]*Turtle //map of turtles by their who number

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

func NewModel(
	patchesOwn map[string]interface{},
	turtlesOwn map[string]interface{},
	turtleBreedsOwn map[string]map[string]interface{},
	turtleBreeds []string,
	directedLinkBreeds []string,
	undirectedLinkBreeds []string,
	wrapping bool,
) *Model {
	maxPxCor := 15
	maxPyCor := 15
	minPxCor := -15
	minPyCor := -15

	model := &Model{
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
		whoToTurtles:    make(map[int]*Turtle),
	}

	//construct turtle breeds
	turtleBreedsMap := make(map[string]*TurtleBreed)
	for i := 0; i < len(turtleBreeds); i++ {
		turtleBreedsMap[turtleBreeds[i]] = &TurtleBreed{
			Turtles: &TurtleAgentSet{
				turtles: make(map[*Turtle]interface{}),
			},
			DefaultShape: "",
		}
	}
	model.Breeds = turtleBreedsMap

	//construct directed link breeds
	directedLinkBreedsMap := make(map[string]*LinkBreed)
	for i := 0; i < len(directedLinkBreeds); i++ {
		directedLinkBreedsMap[directedLinkBreeds[i]] = &LinkBreed{
			Links:        []*Link{},
			Directed:     true,
			DefaultShape: "",
		}
	}
	model.DirectedLinkBreeds = directedLinkBreedsMap

	//construct undirected link breeds
	undirectedLinkBreedsMap := make(map[string]*LinkBreed)
	for i := 0; i < len(undirectedLinkBreeds); i++ {
		undirectedLinkBreedsMap[undirectedLinkBreeds[i]] = &LinkBreed{
			Links:        []*Link{},
			Directed:     false,
			DefaultShape: "",
		}
	}
	model.UndirectedLinkBreed = undirectedLinkBreedsMap

	//construct general turtle set
	model.Turtles = &TurtleAgentSet{
		turtles: make(map[*Turtle]interface{}),
	}

	//construct general link set
	model.Links = &LinkAgentSet{
		links: make(map[*Link]interface{}),
	}

	model.buildPatches()

	return model
}

// builds an array of patches and links them togethor
func (m *Model) buildPatches() {
	m.Patches = &PatchAgentSet{
		patches: map[*Patch]interface{}{},
	}
	m.posOfPatches = make(map[int]*Patch)
	for i := 0; i < m.WorldHeight; i++ {
		for j := 0; j < m.WorldWidth; j++ {
			p := NewPatch(m.PatchesOwn, j+m.MinPxCor, i+m.MinPyCor)
			m.Patches.patches[p] = nil
			m.posOfPatches[j*m.WorldWidth+i] = p
		}
	}
}

// @TODO implement
func (m *Model) BothEnds(link *Link) []*Turtle {
	return nil
}

func (m *Model) ClearAll() {
	m.ClearGlobals()
	m.ClearTicks()
	m.ClearPatches()
	m.ClearDrawing()
	m.ClearAllPlots()
	m.ClearOutput()
}

func (m *Model) ClearGlobals() {
	for g := range m.GlobalBools {
		m.GlobalBools[g] = false
	}
	for g := range m.GlobalFloats {
		m.GlobalFloats[g] = 0
	}
}

// @TODO implement
func (m *Model) ClearLinks() {

}

func (m *Model) ClearTicks() {
	m.TicksOn = false
}

func (m *Model) ClearPatches() {
	for patch := range m.Patches.patches {
		patch.Reset(m.PatchesOwn)
	}
}

// @TODO Implement
func (m *Model) ClearDrawing() {

}

// @TODO Implement
func (m *Model) ClearAllPlots() {

}

// @TODO Implement
func (m *Model) ClearOutput() {

}

// @TODO Implement
func (m *Model) ClearTurtles() {

}

// @TODO Implement
// idea is that if an empty string is passed then it will be for the general population
func (m *Model) CreateOrderedTurtles(breed string, amount float64, operations []TurtleOperation) {

}

func (m *Model) CreateTurtles(amount int, breed string, operations []TurtleOperation) error {

	agentSet := m.Turtles
	var agentSet2 *TurtleAgentSet = nil
	if breed != "" {
		breed, found := m.Breeds[breed]
		if !found {
			return errors.New("breed not found")
		}
		agentSet2 = breed.Turtles
	}

	end := amount + m.turtlesWhoNumber
	for m.turtlesWhoNumber < end {
		newTurtle := NewTurtle(m, m.turtlesWhoNumber, breed)

		agentSet.turtles[newTurtle] = nil
		m.whoToTurtles[m.turtlesWhoNumber] = newTurtle

		if agentSet2 != nil {
			agentSet2.turtles[newTurtle] = nil
		}

		for i := 0; i < len(operations); i++ {
			operations[i](newTurtle)
		}

		m.turtlesWhoNumber++
	}

	return nil
}

// @TODO implement
func (m *Model) DieTurtle(turtle *Turtle) {
}

// @TODO implement
func (m *Model) DieLink(link *Link) {
}

func (m *Model) Diffuse(patchVariable string, percent float64) error {

	if percent > 1 || percent < 0 {
		return errors.New("percent amount was outside bounds")
	}

	diffusions := make(map[*Patch]float64)

	//go through each patch and calculate the diffusion amount
	for patch := range m.Patches.patches {
		patchAmount := patch.PatchesOwn[patchVariable].(float64)
		amountToGive := patchAmount * percent / 8
		diffusions[patch] = amountToGive
	}

	//go through each patch and get the new amount
	for patch := range m.Patches.patches {

		amountFromNeighbors := 0.0
		x := m.WorldHeight*patch.x + patch.y
		neighbors := m.Neighbors(x)
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
func (m *Model) Diffuse4(patchVariable string, percent float64) error {
	return nil
}

func (m *Model) LayoutCircle(turtles []*Turtle, radius float64) {
	amount := len(turtles)
	for i := 0; i < amount; i++ {
		agent := turtles[i]
		agent.SetXY(radius*math.Cos(2*math.Pi*float64(i)/float64(amount)), radius*math.Sin(2*math.Pi*float64(i)/float64(amount)))
		agent.heading = 2 * math.Pi * float64(i) / float64(amount)
	}
}

// @TODO implement
func (m *Model) LayoutRadial(turtles []*Turtle, links []*Link, root *Turtle) {

}

// @TODO implement
func (m *Model) LayoutSpring(turtles []*Turtle, links []*Link, springConstant float64, springLength float64, repulsionConstant float64) {

}

// @TODO implement
func (m *Model) LayoutTutte(turtles []*Turtle, links []*Link, radius float64) {

}

// @TODO implement
func (m *Model) Link(breed string, turtle1 int, turtle2 int) *Link {
	return nil
}

// @TODO implement
func (m *Model) LinkDirected(breed string, turtle1 int, turtle2 int) *Link {
	return nil
}

// @TODO implement
func (m *Model) LinkShapes() []string {
	return []string{}
}

func (m *Model) getPatchAtCoords(x int, y int) *Patch {
	if x < m.MinPxCor || x > m.MaxPxCor || y < m.MinPyCor || y > m.MaxPyCor {
		return nil
	}

	offsetX := x - m.MinPxCor
	offsetY := y - m.MinPyCor

	pos := offsetY*m.WorldWidth + offsetX

	return m.posOfPatches[pos]
}

func (m *Model) OneOfInt(arr []int) interface{} {

	return arr[rand.Intn(len(arr))-1]
}

func (m *Model) RandomAmount(n int) int {
	return rand.Intn(n)
}

func (m *Model) topLeftNeighbor(x int) *Patch {
	return m.safeGetPatch(x - m.WorldWidth - 1)
}

func (m *Model) topNeighbor(x int) *Patch {
	return m.safeGetPatch(x - m.WorldWidth)
}

func (m *Model) topRightNeighbor(x int) *Patch {
	return m.safeGetPatch(x - m.WorldWidth + 1)
}

func (m *Model) leftNeighbor(x int) *Patch {
	return m.safeGetPatch(x - 1)
}

func (m *Model) rightNeighbor(x int) *Patch {
	return m.safeGetPatch(x + 1)
}

func (m *Model) bottomLeftNeighbor(x int) *Patch {
	return m.safeGetPatch(x + m.WorldWidth - 1)
}

func (m *Model) bottomNeighbor(x int) *Patch {
	return m.safeGetPatch(x + m.WorldWidth)
}

func (m *Model) bottomRightNeighbor(x int) *Patch {
	return m.safeGetPatch(x + m.WorldWidth + 1)
}

// @TODO check to see if we are wrapping around
func (m *Model) Neighbors(x int) []*Patch {
	n := []*Patch{}

	topLeft := m.topLeftNeighbor(x)
	if topLeft != nil {
		n = append(n, topLeft)
	}

	bottomLeft := m.bottomLeftNeighbor(x)
	if bottomLeft != nil {
		n = append(n, bottomLeft)
	}

	top := m.topNeighbor(x)
	if top != nil {
		n = append(n, top)
	}

	topRight := m.topRightNeighbor(x)
	if topRight != nil {
		n = append(n, topRight)
	}

	right := m.rightNeighbor(x)
	if right != nil {
		n = append(n, right)
	}

	bottomRight := m.bottomRightNeighbor(x)
	if bottomRight != nil {
		n = append(n, bottomRight)
	}

	bottom := m.bottomNeighbor(x)
	if bottom != nil {
		n = append(n, bottom)
	}

	return n
}

// @TODO check to see if we are wrapping around
func (m *Model) Neighbors4(x int) []*Patch {
	n := []*Patch{}

	top := m.topNeighbor(x)
	if top != nil {
		n = append(n, top)
	}

	left := m.leftNeighbor(x)
	if left != nil {
		n = append(n, left)
	}

	right := m.rightNeighbor(x)
	if right != nil {
		n = append(n, right)
	}

	bottom := m.bottomNeighbor(x)
	if bottom != nil {
		n = append(n, bottom)
	}

	return n
}

func (m *Model) safeGetPatch(x int) *Patch {
	if x < 0 || x > len(m.Patches.patches) {
		return nil
	}

	return m.posOfPatches[x]
}

func (m *Model) Patch(pxcor float64, pycor float64) *Patch {
	//round to get x and y
	x := int(math.Round(pxcor))
	y := int(math.Round(pycor))

	return m.getPatchAtCoords(x, y)
}

func (m *Model) ResetTicks() {
	m.Ticks = 0
}

// @TODO implement
func (m *Model) ResetTimer() {

}

func (m *Model) ResizeWorld(minPxcor int, maxPxcor int, minPycor int, maxPycor int) {
	m.MinPxCor = minPxcor
	m.MaxPxCor = maxPxcor
	m.MinPyCor = minPycor
	m.MaxPyCor = maxPycor
	m.WorldWidth = maxPxcor - minPxcor + 1
	m.WorldHeight = maxPycor - minPycor + 1

	m.buildPatches()
}

func (m *Model) SetDefaultShapeLinks(shape string) {
	m.DefaultShapeLinks = shape
}

func (m *Model) SetDefaultShapeTurtles(shape string) {
	m.DefaultShapeTurtles = shape
}

func (m *Model) SetDefaultShapeLinkBreed(breed string, shape string) {
	m.DirectedLinkBreeds[breed].DefaultShape = shape
}

func (m *Model) SetDefaultShapeTurtleBreed(breed string, shape string) {
	m.Breeds[breed].DefaultShape = shape
}

func (m *Model) Tick() {
	if m.TicksOn {
		m.Ticks++
	}
}

func (m *Model) TickAdvance(amount int) {
	if m.TicksOn {
		m.Ticks += amount
	}
}

// provides a turtle from the model given a breed and who number
// if the breed is empty then selects from the general population
// if the breed or who number is not found then returns nil
func (m *Model) Turtle(breed string, who int) *Turtle {

	t := m.whoToTurtles[who]
	if t == nil {
		return nil //turtle not found
	}
	if breed == "" {
		return t
	} else {
		if m.Breeds[breed] == nil {
			return nil //breed not found
		}
		if t.breed != breed {
			return nil //turtle not found for that breed
		}
		return t
	}
}

// @TODO implement
func (m *Model) TurtlesAt(breed string, pxcor float64, pycor float64) *TurtleAgentSet {
	x := int(math.Round(pxcor))
	y := int(math.Round(pycor))

	patch := m.getPatchAtCoords(x, y)

	if patch == nil {
		return nil
	}

	return nil
}

// @TODO implement
func (m *Model) TurtlesOnPatch(patch *Patch) *TurtleAgentSet {
	return nil
}

// @TODO implement
func (m *Model) TurtlesOnPatches(patches *PatchAgentSet) *TurtleAgentSet {
	return nil
}

// @TODO implement
func (m *Model) TurtlesWithTurtle(turtle *Turtle) *TurtleAgentSet {
	return nil
}

// @TODO implement
func (m *Model) TurtlesWithTurtles(turtles *TurtleAgentSet) *TurtleAgentSet {
	return nil
}
