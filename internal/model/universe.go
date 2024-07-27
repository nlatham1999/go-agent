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
	MaxXCor     float64
	MaxYCor     float64
	MinXCor     float64
	MinYCor     float64
	WorldWidth  int
	WorldHeight int
	wrappingX   bool
	wrappingY   bool

	DefaultShapeTurtles string //the default shape for all turtles
	DefaultShapeLinks   string //the default shape for links

	turtlesWhoNumber int //who number of the next turtle to be created

	GlobalFloats map[string]float64
	GlobalBools  map[string]bool

	Seed int64 //seed for the random number generator
}

func NewModel(
	patchesOwn map[string]interface{},
	turtlesOwn map[string]interface{},
	turtleBreedsOwn map[string]map[string]interface{},
	turtleBreeds []string,
	directedLinkBreeds []string,
	undirectedLinkBreeds []string,
	wrappingX bool,
	wrappingY bool,
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
		MaxXCor:            float64(maxPxCor) + .5,
		MaxYCor:            float64(maxPyCor) + .5,
		MinXCor:            float64(minPxCor) - .5,
		MinYCor:            float64(minPyCor) - .5,
		WorldWidth:         maxPxCor - minPxCor + 1,
		WorldHeight:        maxPyCor - minPyCor + 1,
		patchesOwnTemplate: patchesOwn,
		wrappingX:          wrappingX,
		wrappingY:          wrappingY,
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
			x := j + m.MinPxCor
			y := (i + m.MinPyCor) * -1
			p := NewPatch(m, m.patchesOwnTemplate, x, y)
			m.Patches.patches[p] = nil
			index := y*m.WorldHeight + x
			m.posOfPatches[index] = p
			p.index = index
		}
	}

	p0 := m.posOfPatches[0]
	l := m.leftNeighbor(p0)
	if l == nil {
	}

	for p := range m.Patches.patches {
		p.patchNeighborsMap = map[*Patch]string{}
		p.neighborsPatchMap = map[string]*Patch{}

		left := m.leftNeighbor(p)
		p.patchNeighborsMap[left] = "left"
		p.neighborsPatchMap["left"] = left

		topLeft := m.topLeftNeighbor(p)
		p.patchNeighborsMap[topLeft] = "topLeft"
		p.neighborsPatchMap["topLeft"] = topLeft

		top := m.topNeighbor(p)
		p.patchNeighborsMap[top] = "top"
		p.neighborsPatchMap["top"] = top

		topRight := m.topRightNeighbor(p)
		p.patchNeighborsMap[topRight] = "topRight"
		p.neighborsPatchMap["topRight"] = topRight

		right := m.rightNeighbor(p)
		p.patchNeighborsMap[right] = "right"
		p.neighborsPatchMap["right"] = right

		bottomRight := m.bottomRightNeighbor(p)
		p.patchNeighborsMap[bottomRight] = "bottomRight"
		p.neighborsPatchMap["bottomRight"] = bottomRight

		bottom := m.bottomNeighbor(p)
		p.patchNeighborsMap[bottom] = "bottom"
		p.neighborsPatchMap["bottom"] = bottom

		bottomLeft := m.bottomLeftNeighbor(p)
		p.patchNeighborsMap[bottomLeft] = "bottomLeft"
		p.neighborsPatchMap["bottomLeft"] = bottomLeft
	}
}

func (m *Model) patchIndex(x int, y int) int {
	return y*m.WorldHeight + x
}

func (m *Model) ClearAll() {
	m.ClearGlobals()
	m.ClearTicks()
	m.ClearPatches()
}

func (m *Model) ClearGlobals() {
	for g := range m.GlobalBools {
		m.GlobalBools[g] = false
	}
	for g := range m.GlobalFloats {
		m.GlobalFloats[g] = 0
	}
}

func (m *Model) ClearLinks() {
	for link := range m.Links.links {
		*link = Link{}
	}
	m.Links = &LinkAgentSet{
		links: make(map[*Link]interface{}),
	}
	for breed := range m.DirectedLinkBreeds {
		m.DirectedLinkBreeds[breed].Links = &LinkAgentSet{
			links: make(map[*Link]interface{}),
		}
	}
	for breed := range m.UndirectedLinkBreeds {
		m.UndirectedLinkBreeds[breed].Links = &LinkAgentSet{
			links: make(map[*Link]interface{}),
		}
	}
	for turtle := range m.Turtles.turtles {
		turtle.linkedTurtles = make(map[linkedTurtle]*Link)
		turtle.linkedTurtlesConnectedFrom = make(map[*Turtle]*Link)
	}
}

func (m *Model) ClearTicks() {
	m.TicksOn = false
}

func (m *Model) ClearPatches() {
	for patch := range m.Patches.patches {
		patch.Reset(m.patchesOwnTemplate)
	}
}

// kills all turtles
func (m *Model) ClearTurtles() {
	// delete all links since they are linked to turtles
	m.Links.links = make(map[*Link]interface{})
	for breed := range m.DirectedLinkBreeds {
		m.DirectedLinkBreeds[breed].Links.links = make(map[*Link]interface{})
	}
	for breed := range m.UndirectedLinkBreeds {
		m.UndirectedLinkBreeds[breed].Links.links = make(map[*Link]interface{})
	}

	// remove all turtles from patches
	for patch := range m.Patches.patches {
		patch.turtles = make(map[string]*TurtleAgentSet)
	}

	// clear all turtles
	for turtle := range m.Turtles.turtles {
		*turtle = Turtle{}
	}

	m.Turtles.turtles = make(map[*Turtle]interface{})
	for breed := range m.Breeds {
		m.Breeds[breed].Turtles.turtles = make(map[*Turtle]interface{})
	}

	m.whoToTurtles = make(map[int]*Turtle)

	m.turtlesWhoNumber = 0
}

// @TODO Implement
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

// if the topology allows it then convert the x y to within bounds if it is outside of the world
// returns the new x y and if it is in bounds
// returns false if the x y is not in bounds and the topology does not allow it
func (m *Model) convertXYToInBounds(x float64, y float64) (float64, float64, bool) {

	if x < m.MinXCor {
		if m.wrappingX {
			x = m.MaxXCor - math.Mod(m.MinXCor-x, float64(m.WorldWidth)) + 1
		} else {
			return x, y, false
		}
	}

	if x >= m.MaxXCor {
		if m.wrappingX {
			x = m.MinXCor + math.Mod(x-m.MinXCor, float64(m.WorldWidth))
		} else {
			return x, y, false
		}
	}

	if y < m.MinYCor {
		if m.wrappingY {
			y = m.MaxYCor - math.Mod(m.MinYCor-y, float64(m.WorldHeight)) + 1
		} else {
			return x, y, false
		}
	}

	if y >= m.MaxYCor {
		if m.wrappingY {
			y = m.MinYCor + math.Mod(y-m.MinYCor, float64(m.WorldHeight))
		} else {
			return x, y, false
		}

	}

	return x, y, true
}

// kills a turtle
func (m *Model) KillTurtle(turtle *Turtle) {

	delete(m.Turtles.turtles, turtle)
	if turtle.breed != "" {
		delete(m.Breeds[turtle.breed].Turtles.turtles, turtle)
	}
	delete(m.whoToTurtles, turtle.who)

	p := turtle.PatchHere()
	if p != nil {
		delete(p.turtles[""].turtles, turtle)
		if turtle.breed != "" {
			delete(p.turtles[turtle.breed].turtles, turtle)
		}
	}

	for _, link := range turtle.linkedTurtles {
		delete(m.Links.links, link)
		if link.breed != "" {
			if link.Directed {
				delete(m.DirectedLinkBreeds[link.breed].Links.links, link)
			} else {
				delete(m.UndirectedLinkBreeds[link.breed].Links.links, link)
			}
		}
	}

	for _, link := range turtle.linkedTurtlesConnectedFrom {
		delete(m.Links.links, link)
		if link.breed != "" {
			if link.Directed {
				delete(m.DirectedLinkBreeds[link.breed].Links.links, link)
			} else {
				delete(m.UndirectedLinkBreeds[link.breed].Links.links, link)
			}
		}
	}

	*turtle = Turtle{}
}

// kills a link
func (m *Model) KillLink(link *Link) {

	delete(m.Links.links, link)

	if link.breed != "" {
		if link.Directed {
			delete(m.DirectedLinkBreeds[link.breed].Links.links, link)
		} else {
			delete(m.UndirectedLinkBreeds[link.breed].Links.links, link)
		}
	}

	// remove the link from the turtles
	if link.Directed {
		key := linkedTurtle{
			true,
			"",
			link.end2,
		}
		delete(link.end1.linkedTurtles, key)
		delete(link.end2.linkedTurtlesConnectedFrom, link.end1)
		if link.end1.breed != "" {
			key = linkedTurtle{
				true,
				link.breed,
				link.end2,
			}
			delete(link.end1.linkedTurtles, key)
		}
	} else {
		key1 := linkedTurtle{
			false,
			"",
			link.end2,
		}
		key2 := linkedTurtle{
			false,
			"",
			link.end1,
		}
		delete(link.end1.linkedTurtles, key1)
		delete(link.end2.linkedTurtles, key2)
		if link.end1.breed != "" {
			key1 = linkedTurtle{
				false,
				link.breed,
				link.end2,
			}
			key2 = linkedTurtle{
				false,
				link.breed,
				link.end1,
			}
			delete(link.end1.linkedTurtles, key1)
			delete(link.end2.linkedTurtles, key2)
		}
	}

	*link = Link{}
}

func (m *Model) Diffuse(patchVariable string, percent float64) error {

	if percent > 1 || percent < 0 {
		return errors.New("percent amount was outside bounds")
	}

	diffusions := make(map[*Patch]float64)

	//go through each patch and calculate the diffusion amount
	for patch := range m.Patches.patches {
		patchAmount := patch.patchesOwn[patchVariable].(float64)
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

		patchAmount := patch.patchesOwn[patchVariable].(float64)
		amountToKeep := 1 - (patchAmount * percent) + (float64(8-neighbors.Count()) * (patchAmount * percent / 8))

		patch.patchesOwn[patchVariable] = amountToKeep + amountFromNeighbors
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

	if !m.wrappingX && !m.wrappingY {
		return distance
	}

	deltaXInverse := float64(m.WorldWidth) - math.Abs(deltaX)
	deltaYInverse := float64(m.WorldHeight) - math.Abs(deltaY)

	if m.wrappingX {
		distance = math.Min(distance, math.Abs(math.Sqrt(deltaXInverse*deltaXInverse+deltaY*deltaY)))
	}

	if m.wrappingY {
		distance = math.Min(distance, math.Abs(math.Sqrt(deltaX*deltaX+deltaYInverse*deltaYInverse)))
	}

	if m.wrappingX && m.wrappingY {
		distance = math.Min(distance, math.Abs(math.Sqrt(deltaXInverse*deltaXInverse+deltaYInverse*deltaYInverse)))
	}

	return distance
}

func (m *Model) LayoutCircle(turtles []*Turtle, radius float64) {
	amount := len(turtles)
	for i := 0; i < amount; i++ {
		agent := turtles[i]
		agent.SetXY(radius*math.Cos(2*math.Pi*float64(i)/float64(amount)), radius*math.Sin(2*math.Pi*float64(i)/float64(amount)))
		heading := 2 * math.Pi * float64(i) / float64(amount)
		agent.setHeadingRadians(heading)
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

	pos := y*m.WorldHeight + x

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
		if !m.wrappingX {
			return nil
		} else {
			x = m.MaxPxCor
		}
	}

	if y < m.MinPyCor {
		if !m.wrappingY {
			return nil
		} else {
			y = m.MaxPyCor
		}
	}

	n := m.patchIndex(x, y)

	return m.getPatchAtPos(n)
}

func (m *Model) topNeighbor(p *Patch) *Patch {
	y := p.y - 1

	if y < m.MinPyCor {
		if !m.wrappingY {
			return nil
		} else {
			y = m.MaxPyCor
		}
	}

	n := m.patchIndex(p.x, y)

	return m.getPatchAtPos(n)
}

func (m *Model) topRightNeighbor(p *Patch) *Patch {
	x := p.x + 1
	y := p.y - 1

	if x > m.MaxPxCor {
		if !m.wrappingX {
			return nil
		} else {
			x = m.MinPxCor
		}
	}

	if y < m.MinPyCor {
		if !m.wrappingY {
			return nil
		} else {
			y = m.MaxPyCor
		}
	}

	n := m.patchIndex(x, y)

	return m.getPatchAtPos(n)
}

func (m *Model) leftNeighbor(p *Patch) *Patch {
	x := p.x - 1

	if x < m.MinPxCor {
		if !m.wrappingX {
			return nil
		} else {
			x = m.MaxPxCor
		}
	}

	n := m.patchIndex(x, p.y)

	return m.getPatchAtPos(n)
}

func (m *Model) rightNeighbor(p *Patch) *Patch {
	x := p.x + 1

	if x > m.MaxPxCor {
		if !m.wrappingX {
			return nil
		} else {
			x = m.MinPxCor
		}
	}

	n := m.patchIndex(x, p.y)

	return m.getPatchAtPos(n)
}

func (m *Model) bottomLeftNeighbor(p *Patch) *Patch {
	x := p.x - 1
	y := p.y + 1

	if x < m.MinPxCor {
		if !m.wrappingX {
			return nil
		} else {
			x = m.MaxPxCor
		}
	}

	if y > m.MaxPyCor {
		if !m.wrappingX {
			return nil
		} else {
			y = m.MinPyCor
		}
	}

	n := m.patchIndex(x, y)

	return m.getPatchAtPos(n)
}

func (m *Model) bottomNeighbor(p *Patch) *Patch {
	y := p.y + 1

	if y > m.MaxPyCor {
		if !m.wrappingY {
			return nil
		} else {
			y = m.MinPyCor
		}
	}

	n := m.patchIndex(p.x, y)

	return m.getPatchAtPos(n)
}

func (m *Model) bottomRightNeighbor(p *Patch) *Patch {
	x := p.x + 1
	y := p.y + 1

	if x > m.MaxPxCor {
		if !m.wrappingX {
			return nil
		} else {
			x = m.MinPxCor
		}
	}

	if y > m.MaxPyCor {
		if !m.wrappingY {
			return nil
		} else {
			y = m.MinPyCor
		}
	}

	n := m.patchIndex(x, y)

	return m.getPatchAtPos(n)
}

// @TODO check to see if we are wrapping around
func (m *Model) neighbors(p *Patch) *PatchAgentSet {
	n := make(map[*Patch]interface{})

	topLeft := p.neighborsPatchMap["topLeft"]
	if topLeft != nil {
		n[topLeft] = nil
	}

	left := p.neighborsPatchMap["left"]
	if left != nil {
		n[left] = nil
	}

	bottomLeft := p.neighborsPatchMap["bottomLeft"]
	if bottomLeft != nil {
		n[bottomLeft] = nil
	}

	top := p.neighborsPatchMap["top"]
	if top != nil {
		n[top] = nil
	}

	topRight := p.neighborsPatchMap["topRight"]
	if topRight != nil {
		n[topRight] = nil
	}

	right := p.neighborsPatchMap["right"]
	if right != nil {
		n[right] = nil
	}

	bottomRight := p.neighborsPatchMap["bottomRight"]
	if bottomRight != nil {
		n[bottomRight] = nil
	}

	bottom := p.neighborsPatchMap["bottom"]
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

	top := p.neighborsPatchMap["top"]
	if top != nil {
		n[top] = nil
	}

	left := p.neighborsPatchMap["left"]
	if left != nil {
		n[left] = nil
	}

	right := p.neighborsPatchMap["right"]
	if right != nil {
		n[right] = nil
	}

	bottom := p.neighborsPatchMap["bottom"]
	if bottom != nil {
		n[bottom] = nil
	}

	return &PatchAgentSet{
		patches: n,
	}
}

func (m *Model) getPatchAtPos(x int) *Patch {
	return m.posOfPatches[x]
}

func (m *Model) Patch(pxcor float64, pycor float64) *Patch {

	// round the x and y except in cases where the x or y is the min value
	// since the min value will be -*.5 and we want to round up in that case
	var x int
	var y int
	if pxcor == m.MinXCor {
		x = int(math.Ceil(pxcor))
	} else {
		x = int(math.Round(pxcor))
	}

	if pycor == m.MinYCor {
		y = int(math.Ceil(pycor))
	} else {
		y = int(math.Round(pycor))
	}

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

func (m *Model) WrappingXOn() {
	m.wrappingX = true
}

func (m *Model) WrappingYOn() {
	m.wrappingY = true
}

func (m *Model) WrappingXOff() {
	m.wrappingX = false
}

func (m *Model) WrappingYOff() {
	m.wrappingY = false
}
