package model

import (
	"errors"
	"math"
	"math/rand"
)

type Model struct {
	Ticks   int
	TicksOn bool

	linksOwnTemplate      map[string]interface{}            //additional variables for each link
	linkBreedsOwnTemplate map[string]map[string]interface{} //additional variables for each link breed. The first key is the breed name
	patchesOwnTemplate    map[string]interface{}            //additional variables for each patch

	Patches              *PatchAgentSet
	Turtles              *TurtleAgentSet         //all the turtles
	Breeds               map[string]*TurtleBreed //turtles that are part of specific breeds
	Links                *LinkAgentSet           //all the links
	DirectedLinkBreeds   map[string]*LinkBreed
	UndirectedLinkBreeds map[string]*LinkBreed

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
		MaxPxCor:           maxPxCor,
		MaxPyCor:           maxPyCor,
		MinPxCor:           minPxCor,
		MinPyCor:           minPyCor,
		WorldWidth:         maxPxCor - minPxCor + 1,
		WorldHeight:        maxPyCor - minPyCor + 1,
		patchesOwnTemplate: patchesOwn,
		wrapping:           wrapping,
		whoToTurtles:       make(map[int]*Turtle),
	}

	//construct turtle breeds
	turtleBreedsMap := make(map[string]*TurtleBreed)
	for i := 0; i < len(turtleBreeds); i++ {
		turtleBreedsMap[turtleBreeds[i]] = &TurtleBreed{
			Turtles: &TurtleAgentSet{
				turtles: make(map[*Turtle]interface{}),
			},
			name:         turtleBreeds[i],
			defaultShape: "",
		}

		//copy the turtles breeds own template
		if turtleBreedsOwn != nil && turtleBreedsOwn[turtleBreeds[i]] != nil {
			turtleBreedsMap[turtleBreeds[i]].turtlesOwnTemplate = make(map[string]interface{})
			for key, value := range turtleBreedsOwn[turtleBreeds[i]] {
				turtleBreedsMap[turtleBreeds[i]].turtlesOwnTemplate[key] = value
			}
		} else {
			turtleBreedsMap[turtleBreeds[i]].turtlesOwnTemplate = make(map[string]interface{})
		}
	}
	model.Breeds = turtleBreedsMap

	//construct directed link breeds
	directedLinkBreedsMap := make(map[string]*LinkBreed)
	directedLinkBreeds = append(directedLinkBreeds, "") // add the general population
	for i := 0; i < len(directedLinkBreeds); i++ {
		directedLinkBreedsMap[directedLinkBreeds[i]] = &LinkBreed{
			Links: &LinkAgentSet{
				links: make(map[*Link]interface{}),
			},
			Directed:     true,
			DefaultShape: "",
			name:         directedLinkBreeds[i],
		}
	}
	model.DirectedLinkBreeds = directedLinkBreedsMap

	//construct undirected link breeds
	undirectedLinkBreedsMap := make(map[string]*LinkBreed)
	undirectedLinkBreeds = append(undirectedLinkBreeds, "") // add the general population
	for i := 0; i < len(undirectedLinkBreeds); i++ {
		undirectedLinkBreedsMap[undirectedLinkBreeds[i]] = &LinkBreed{
			Links: &LinkAgentSet{
				links: make(map[*Link]interface{}),
			},
			Directed:     false,
			DefaultShape: "",
			name:         undirectedLinkBreeds[i],
		}
	}
	model.UndirectedLinkBreeds = undirectedLinkBreedsMap

	//construct general turtle set
	model.Turtles = &TurtleAgentSet{
		turtles: make(map[*Turtle]interface{}),
	}

	// create a breed with no name for the general population
	model.Breeds[""] = &TurtleBreed{
		Turtles:            model.Turtles,
		name:               "",
		defaultShape:       "",
		turtlesOwnTemplate: make(map[string]interface{}),
	}
	if turtlesOwn != nil {
		for key, value := range turtlesOwn {
			model.Breeds[""].turtlesOwnTemplate[key] = value
		}
	} else {
		model.Breeds[""].turtlesOwnTemplate = make(map[string]interface{})
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
			p := NewPatch(m, m.patchesOwnTemplate, j+m.MinPxCor, i+m.MinPyCor)
			m.Patches.patches[p] = nil
			index := i*m.WorldHeight + j
			m.posOfPatches[index] = p
			p.index = index
		}
	}
}

func (m *Model) patchIndex(x int, y int) int {
	x = x - m.MinPxCor
	y = y - m.MinPyCor
	return y*m.WorldHeight + x
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
		patch.Reset(m.patchesOwnTemplate)
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

	if breed != "" {
		_, found := m.Breeds[breed]
		if !found {
			return errors.New("breed not found")
		}
	}

	end := amount + m.turtlesWhoNumber
	for m.turtlesWhoNumber < end {
		newTurtle := NewTurtle(m, m.turtlesWhoNumber, breed, 0, 0)

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
		neighbors := m.neighbors(patch)
		if neighbors.Count() > 8 || neighbors.Count() < 3 {
			return errors.New("invalid amount of neighbors")
		}
		for n := range neighbors.patches {
			amountFromNeighbors += diffusions[n]
		}

		patchAmount := patch.PatchesOwn[patchVariable].(float64)
		amountToKeep := 1 - (patchAmount * percent) + (float64(8-neighbors.Count()) * (patchAmount * percent / 8))

		patch.PatchesOwn[patchVariable] = amountToKeep + amountFromNeighbors
	}

	return nil
}

// @TODO implement
func (m *Model) Diffuse4(patchVariable string, percent float64) error {
	return nil
}

func (m *Model) DistanceBetweenPoints(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	deltaX := x1 - x2
	deltaY := y1 - y2

	distance := math.Abs(math.Sqrt(deltaX*deltaX + deltaY*deltaY))

	if !m.wrapping {
		return distance
	}

	deltaXInverse := float64(m.WorldWidth) - math.Abs(deltaX)
	deltaYInverse := float64(m.WorldHeight) - math.Abs(deltaY)

	distance = math.Min(distance, math.Abs(math.Sqrt(deltaX*deltaX+deltaYInverse*deltaYInverse)))
	distance = math.Min(distance, math.Abs(math.Sqrt(deltaXInverse*deltaXInverse+deltaY*deltaY)))
	distance = math.Min(distance, math.Abs(math.Sqrt(deltaXInverse*deltaXInverse+deltaYInverse*deltaYInverse)))

	return distance
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

// does not implement wrappimg, that is the responsibilty of the caller
func (m *Model) getPatchAtCoords(x int, y int) *Patch {
	if x < m.MinPxCor || x > m.MaxPxCor || y < m.MinPyCor || y > m.MaxPyCor {
		return nil
	}

	offsetX := x - m.MinPxCor
	offsetY := y - m.MinPyCor

	pos := offsetY*m.WorldHeight + offsetX

	return m.posOfPatches[pos]
}

func (m *Model) OneOfInt(arr []int) interface{} {

	return arr[rand.Intn(len(arr))-1]
}

func (m *Model) RandomAmount(n int) int {
	return rand.Intn(n)
}

func (m *Model) topLeftNeighbor(p *Patch) *Patch {
	x := p.x - 1
	y := p.y - 1

	if x < m.MinPxCor {
		if !m.wrapping {
			return nil
		} else {
			x = m.MaxPxCor
		}
	}

	if y < m.MinPyCor {
		if !m.wrapping {
			return nil
		} else {
			y = m.MaxPyCor
		}
	}

	n := m.patchIndex(x, y)

	return m.safeGetPatch(n)
}

func (m *Model) topNeighbor(p *Patch) *Patch {
	y := p.y - 1

	if y < m.MinPyCor {
		if !m.wrapping {
			return nil
		} else {
			y = m.MaxPyCor
		}
	}

	n := m.patchIndex(p.x, y)

	return m.safeGetPatch(n)
}

func (m *Model) topRightNeighbor(p *Patch) *Patch {
	x := p.x + 1
	y := p.y - 1

	if x > m.MaxPxCor {
		if !m.wrapping {
			return nil
		} else {
			x = m.MinPxCor
		}
	}

	if y < m.MinPyCor {
		if !m.wrapping {
			return nil
		} else {
			y = m.MaxPyCor
		}
	}

	n := m.patchIndex(x, y)

	return m.safeGetPatch(n)
}

func (m *Model) leftNeighbor(p *Patch) *Patch {
	x := p.x - 1

	if x < m.MinPxCor {
		if !m.wrapping {
			return nil
		} else {
			x = m.MaxPxCor
		}
	}

	n := m.patchIndex(x, p.y)

	return m.safeGetPatch(n)
}

func (m *Model) rightNeighbor(p *Patch) *Patch {
	x := p.x + 1

	if x > m.MaxPxCor {
		if !m.wrapping {
			return nil
		} else {
			x = m.MinPxCor
		}
	}

	n := m.patchIndex(x, p.y)

	return m.safeGetPatch(n)
}

func (m *Model) bottomLeftNeighbor(p *Patch) *Patch {
	x := p.x - 1
	y := p.y + 1

	if x < m.MinPxCor {
		if !m.wrapping {
			return nil
		} else {
			x = m.MaxPxCor
		}
	}

	if y > m.MaxPyCor {
		if !m.wrapping {
			return nil
		} else {
			y = m.MinPyCor
		}
	}

	n := m.patchIndex(x, y)

	return m.safeGetPatch(n)
}

func (m *Model) bottomNeighbor(p *Patch) *Patch {
	y := p.y + 1

	if y > m.MaxPyCor {
		if !m.wrapping {
			return nil
		} else {
			y = m.MinPyCor
		}
	}

	n := m.patchIndex(p.x, y)

	return m.safeGetPatch(n)
}

func (m *Model) bottomRightNeighbor(p *Patch) *Patch {
	x := p.x + 1
	y := p.y + 1

	if x > m.MaxPxCor {
		if !m.wrapping {
			return nil
		} else {
			x = m.MinPxCor
		}
	}

	if y > m.MaxPyCor {
		if !m.wrapping {
			return nil
		} else {
			y = m.MinPyCor
		}
	}

	n := m.patchIndex(x, y)

	return m.safeGetPatch(n)
}

// @TODO check to see if we are wrapping around
func (m *Model) neighbors(p *Patch) *PatchAgentSet {
	n := make(map[*Patch]interface{})

	topLeft := m.topLeftNeighbor(p)
	if topLeft != nil {
		n[topLeft] = nil
	}

	left := m.leftNeighbor(p)
	if left != nil {
		n[left] = nil
	}

	bottomLeft := m.bottomLeftNeighbor(p)
	if bottomLeft != nil {
		n[bottomLeft] = nil
	}

	top := m.topNeighbor(p)
	if top != nil {
		n[top] = nil
	}

	topRight := m.topRightNeighbor(p)
	if topRight != nil {
		n[topRight] = nil
	}

	right := m.rightNeighbor(p)
	if right != nil {
		n[right] = nil
	}

	bottomRight := m.bottomRightNeighbor(p)
	if bottomRight != nil {
		n[bottomRight] = nil
	}

	bottom := m.bottomNeighbor(p)
	if bottom != nil {
		n[bottom] = nil
	}

	return &PatchAgentSet{
		patches: n,
	}
}

// @TODO check to see if we are wrapping around
func (m *Model) neighbors4(p *Patch) *PatchAgentSet {
	n := make(map[*Patch]interface{})

	top := m.topNeighbor(p)
	if top != nil {
		n[top] = nil
	}

	left := m.leftNeighbor(p)
	if left != nil {
		n[left] = nil
	}

	right := m.rightNeighbor(p)
	if right != nil {
		n[right] = nil
	}

	bottom := m.bottomNeighbor(p)
	if bottom != nil {
		n[bottom] = nil
	}

	return &PatchAgentSet{
		patches: n,
	}
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
	m.Breeds[breed].defaultShape = shape
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
