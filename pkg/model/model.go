package model

import (
	"errors"
	"math"
	"math/rand/v2"
	"time"

	"github.com/nlatham1999/sortedset"
)

// Model holds all the agents and the world
type Model struct {
	Ticks int

	DefaultPatchProperties map[string]interface{} //additional variables for each patch

	Patches              *PatchAgentSet          //all the patches
	turtles              *TurtleAgentSet         //all the turtles
	links                *LinkAgentSet           //all the links
	breeds               map[string]*TurtleBreed //turtles split into breeds
	directedLinkBreeds   map[string]*LinkBreed   //directed link breeds
	undirectedLinkBreeds map[string]*LinkBreed   //undirected link breeds

	posOfPatches map[int]*Patch  //map of patches by their index
	whoToTurtles map[int]*Turtle //map of turtles by their who number

	maxPxCor    int     //the maximum x coordinate
	maxPyCor    int     //the maximum y coordinate
	minPxCor    int     //the minimum x coordinate
	minPyCor    int     //the minimum y coordinate
	maxXCor     float64 //the maximum x coordinate as a float, adds .5 to the max x cor
	maxYCor     float64 //the maximum y coordinate as a float, adds .5 to the max y cor
	minXCor     float64 //the minimum x coordinate as a float, subtracts .5 from the min x cor
	minYCor     float64 //the minimum y coordinate as a float, subtracts .5 from the min y cor
	worldWidth  int     //the width of the world
	worldHeight int     //the height of the world
	wrappingX   bool    //if the world wraps around in the x direction
	wrappingY   bool    //if the world wraps around in the y direction

	DefaultShapeTurtles string //the default shape for all turtles
	DefaultShapeLinks   string //the default shape for links

	turtlesWhoNumber int //who number of the next turtle to be created

	randomGenerator *rand.Rand
	seedValue       uint64
	seedValue2      uint64
	randomSrc       *rand.PCG

	modelStart time.Time

	// turtles the current turtle is linked to/by/with
	linkedTurtles map[*Turtle]*turtleLinks

	// set of shown links
	// used for efficient rendering
	ShownLinks *LinkAgentSet
}

// Create a new model
func NewModel(
	settings ModelSettings,
) *Model {

	useDefaultPxCor := false

	// if the user did not set the max and min px cor then use the default
	if settings.MaxPxCor == 0 && settings.MaxPyCor == 0 && settings.MinPxCor == 0 && settings.MinPyCor == 0 {
		useDefaultPxCor = true
	}

	// if the user set the max to be less or equal to the min then use the default
	if settings.MaxPxCor <= settings.MinPxCor || settings.MaxPyCor <= settings.MinPyCor {
		useDefaultPxCor = true
	}

	// if the user set the min to greater than 0 or the max to be less than 0 then use the default
	if settings.MinPxCor > 0 || settings.MinPyCor > 0 || settings.MaxPxCor < 0 || settings.MaxPyCor < 0 {
		useDefaultPxCor = true
	}

	if useDefaultPxCor {
		settings.MaxPxCor = 15
		settings.MaxPyCor = 15
		settings.MinPxCor = -15
		settings.MinPyCor = -15
	}

	patchProperties := make(map[string]interface{})
	if settings.PatchProperties != nil {
		for key, value := range settings.PatchProperties {
			patchProperties[key] = value
		}
	}

	model := &Model{
		maxPxCor:               settings.MaxPxCor,
		maxPyCor:               settings.MaxPyCor,
		minPxCor:               settings.MinPxCor,
		minPyCor:               settings.MinPyCor,
		maxXCor:                float64(settings.MaxPxCor) + .5,
		maxYCor:                float64(settings.MaxPyCor) + .5,
		minXCor:                float64(settings.MinPxCor) - .5,
		minYCor:                float64(settings.MinPyCor) - .5,
		worldWidth:             settings.MaxPxCor - settings.MinPxCor + 1,
		worldHeight:            settings.MaxPyCor - settings.MinPyCor + 1,
		DefaultPatchProperties: settings.PatchProperties,
		wrappingX:              settings.WrappingX,
		wrappingY:              settings.WrappingY,
		whoToTurtles:           make(map[int]*Turtle),
		seedValue:              settings.RandomSeed,
		seedValue2:             settings.RandomSeed2,
		modelStart:             time.Now(),
		linkedTurtles:          make(map[*Turtle]*turtleLinks),
	}

	model.randomSrc = rand.NewPCG(model.seedValue, model.seedValue2)
	model.randomGenerator = rand.New(model.randomSrc)

	//construct turtle breeds
	turtleBreedsMap := make(map[string]*TurtleBreed)
	for _, breed := range settings.TurtleBreeds {
		breed.model = model
		turtleBreedsMap[breed.name] = breed
	}
	model.breeds = turtleBreedsMap

	//construct directed link breeds
	model.directedLinkBreeds = make(map[string]*LinkBreed)
	settings.DirectedLinkBreeds = append(settings.DirectedLinkBreeds, NewLinkBreed("")) // add the general population
	for _, directedLink := range settings.DirectedLinkBreeds {
		model.directedLinkBreeds[directedLink.name] = directedLink
		directedLink.model = model
		directedLink.directed = true
	}

	//construct undirected link breeds
	model.undirectedLinkBreeds = make(map[string]*LinkBreed)
	settings.UndirectedLinkBreeds = append(settings.UndirectedLinkBreeds, NewLinkBreed("")) // add the general population
	for _, undirectedLink := range settings.UndirectedLinkBreeds {
		model.undirectedLinkBreeds[undirectedLink.name] = undirectedLink
		undirectedLink.model = model
		undirectedLink.directed = false
	}

	//construct general turtle set
	model.turtles = NewTurtleAgentSet([]*Turtle{})

	// create a breed with no name for the general population
	model.breeds[BreedNone] = NewTurtleBreed("", "", make(map[string]interface{}))
	model.breeds[BreedNone].model = model

	if settings.TurtleProperties != nil {
		for key, value := range settings.TurtleProperties {
			model.breeds[BreedNone].defaultProperties[key] = value
		}
	} else {
		model.breeds[BreedNone].defaultProperties = make(map[string]interface{})
	}

	//construct general link set
	model.links = NewLinkAgentSet([]*Link{})

	// create shown link sets
	model.ShownLinks = NewLinkAgentSet([]*Link{})

	// build patches
	model.buildPatches()

	return model
}

// builds an array of patches and links them togethor
func (m *Model) buildPatches() {
	m.Patches = NewPatchAgentSet([]*Patch{})
	m.posOfPatches = make(map[int]*Patch)
	for i := m.minPyCor; i <= m.maxPyCor; i++ {
		for j := m.minPxCor; j <= m.maxPxCor; j++ {
			x := j
			y := i
			p := newPatch(m, m.DefaultPatchProperties, x, y)
			m.Patches.Add(p)
			index := y*m.worldHeight + x
			m.posOfPatches[index] = p
			p.index = index
		}
	}

	m.Patches.Ask(func(p *Patch) {
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
	})
}

func (m *Model) patchIndex(x int, y int) int {
	return y*m.worldHeight + x
}

func (m *Model) TurtleBreed(breedName string) *TurtleBreed {
	return m.breeds[breedName]
}

// returns a list of the turtle breeds
func (m *Model) TurtleBreeds() []*TurtleBreed {
	breeds := make([]*TurtleBreed, 0, len(m.breeds))
	for _, breed := range m.breeds {
		breeds = append(breeds, breed)
	}
	return breeds
}

// clear all patches and turtles and set the ticks to zero
func (m *Model) ClearAll() {
	m.ClearTicks()
	m.ClearPatches()
	m.ClearTurtles()
}

// clear all links
func (m *Model) ClearLinks() {
	m.links = NewLinkAgentSet([]*Link{})
	for breed := range m.directedLinkBreeds {
		m.directedLinkBreeds[breed].links = NewLinkAgentSet([]*Link{})
	}
	for breed := range m.undirectedLinkBreeds {
		m.undirectedLinkBreeds[breed].links = NewLinkAgentSet([]*Link{})
	}
	m.turtles.Ask(func(turtle *Turtle) {
		m.linkedTurtles[turtle] = newTurtleLinks()
	})
}

// set the ticks to zero
func (m *Model) ClearTicks() {
	m.Ticks = 0
}

// clear all patches
func (m *Model) ClearPatches() {
	m.Patches.Ask(func(p *Patch) {
		p.Reset(m.DefaultPatchProperties)
	})
}

// kills all turtles
func (m *Model) ClearTurtles() {
	// delete all links since they are linked to turtles
	m.links = NewLinkAgentSet([]*Link{})
	for breed := range m.directedLinkBreeds {
		m.directedLinkBreeds[breed].links = NewLinkAgentSet([]*Link{})
	}
	for breed := range m.undirectedLinkBreeds {
		m.undirectedLinkBreeds[breed].links = NewLinkAgentSet([]*Link{})
	}

	// remove all turtles from patches
	m.Patches.Ask(func(p *Patch) {
		p.turtles = make(map[*TurtleBreed]*TurtleAgentSet)
	})

	// clear all turtles
	m.turtles.Ask(func(turtle *Turtle) {
		*turtle = Turtle{}
	})

	m.turtles = NewTurtleAgentSet([]*Turtle{})
	for breed := range m.breeds {
		m.breeds[breed].turtles = NewTurtleAgentSet([]*Turtle{})
	}

	m.whoToTurtles = make(map[int]*Turtle)

	m.turtlesWhoNumber = 0
}

// CreateTurtles creates the specified amount of turtles with the specified operation.
//
// WARNING: Not thread-safe. Do not call concurrently from multiple goroutines.
func (m *Model) CreateTurtles(amount int, operation TurtleOperation) (*TurtleAgentSet, error) {

	generalBreed := m.breeds[BreedNone]

	return m.createTurtlesBreeded(amount, generalBreed, operation)
}

func (m *Model) createTurtlesBreeded(amount int, breed *TurtleBreed, operation TurtleOperation) (*TurtleAgentSet, error) {

	turtles := NewTurtleAgentSet([]*Turtle{})

	end := amount + m.turtlesWhoNumber
	for m.turtlesWhoNumber < end {
		newTurtle := newTurtle(m, m.turtlesWhoNumber, breed, 0, 0)

		// set heading to be random
		newTurtle.setHeadingRadians(m.randomGenerator.Float64() * 2 * math.Pi)

		// get a random color from the base colors and set it
		newTurtle.Color.SetColor(baseColorsList[m.randomGenerator.IntN(len(baseColorsList))])

		turtles.Add(newTurtle)

		m.turtlesWhoNumber++
	}

	turtles.Ask(operation)

	return turtles, nil
}

// if the topology allows it then convert the x y to within bounds if it is outside of the world
// returns the new x y and if it is in bounds
// returns false if the x y is not in bounds and the topology does not allow it
func (m *Model) convertXYToInBounds(x float64, y float64) (float64, float64, bool) {

	if x < m.minXCor {
		if m.wrappingX {
			x = m.maxXCor - math.Mod(m.minXCor-x, float64(m.worldWidth))
		} else {
			return x, y, false
		}
	}

	if x >= m.maxXCor {
		if m.wrappingX {
			x = m.minXCor + math.Mod(x-m.minXCor, float64(m.worldWidth))
		} else {
			return x, y, false
		}
	}

	if y < m.minYCor {
		if m.wrappingY {
			modu := math.Mod(m.minYCor-y, float64(m.worldHeight))
			y = m.maxYCor - modu
		} else {
			return x, y, false
		}
	}

	if y >= m.maxYCor {
		if m.wrappingY {
			y = m.minYCor + math.Mod(y-m.minYCor, float64(m.worldHeight))
		} else {
			return x, y, false
		}

	}

	return x, y, true
}

// kills a turtle
func (m *Model) KillTurtle(turtle *Turtle) {

	m.turtles.Remove(turtle)
	if turtle.breed != nil {
		m.breeds[turtle.breed.name].turtles.Remove(turtle)
	}
	delete(m.whoToTurtles, turtle.who)

	p := turtle.PatchHere()
	if p != nil {
		generalBreed := m.breeds[BreedNone]
		p.turtles[generalBreed].Remove(turtle)
		if turtle.breed != nil {
			p.turtles[turtle.breed].Remove(turtle)
		}
	}

	// kill all directed out links
	for link := range m.linkedTurtles[turtle].getAllDirectedOutLinks() {
		m.KillLink(link)
	}

	// kill all directed in links
	for link := range m.linkedTurtles[turtle].getAllDirectedInLinks() {
		m.KillLink(link)
	}

	// kill all undirected links
	for link := range m.linkedTurtles[turtle].getAllUndirectedLinks() {
		m.KillLink(link)
	}

	*turtle = Turtle{}
}

// kills a link
func (m *Model) KillLink(link *Link) {

	m.links.links.Remove(link)
	m.ShownLinks.Remove(link)

	if link.breed.name != BreedNone {
		if link.directed {
			m.directedLinkBreeds[link.breed.name].links.links.Remove(link)
		} else {
			m.undirectedLinkBreeds[link.breed.name].links.links.Remove(link)
		}
	}

	// remove the link from the turtles
	if link.directed {
		m.linkedTurtles[link.end1].removeDirectedOutBreed(link.breed, link.end2, link)
		m.linkedTurtles[link.end2].removeDirectedInBreed(link.breed, link.end1, link)
	} else {
		m.linkedTurtles[link.end1].removeUndirectedBreed(link.breed, link.end2, link)
		m.linkedTurtles[link.end2].removeUndirectedBreed(link.breed, link.end1, link)
	}

	*link = Link{}
}

// diffuse the patch variable of each patch to its neighbors
func (m *Model) Diffuse(patchVariable string, percent float64) error {

	if percent > 1 || percent < 0 {
		return errors.New("percent amount was outside bounds")
	}

	diffusions := make(map[*Patch]float64)

	//go through each patch and calculate the diffusion amount
	m.Patches.Ask(func(patch *Patch) {
		patchAmount := patch.patchProperties[patchVariable].(float64)
		amountToGive := patchAmount * percent / 8
		diffusions[patch] = amountToGive
	})

	//go through each patch and get the new amount
	m.Patches.Ask(func(patch *Patch) {
		amountFromNeighbors := 0.0
		neighbors := m.neighbors(patch)
		if neighbors.Count() > 8 || neighbors.Count() < 3 {
			return
		}
		neighbors.Ask(func(n *Patch) {
			amountFromNeighbors += diffusions[n]
		})

		patchAmount := patch.patchProperties[patchVariable].(float64)
		amountToKeep := (patchAmount * (1 - percent)) + (float64(8-neighbors.Count()) * (patchAmount * percent / 8))

		patch.patchProperties[patchVariable] = amountToKeep + amountFromNeighbors
	})

	return nil
}

// diffuse the patch variable of each patch to its neighbors at the top, bottom, left, and right
func (m *Model) Diffuse4(patchVariable string, percent float64) error {

	if percent > 1 || percent < 0 {
		return errors.New("percent amount was outside bounds")
	}

	diffusions := make(map[*Patch]float64)

	//go through each patch and calculate the diffusion amount
	m.Patches.Ask(func(patch *Patch) {
		patchAmount := patch.patchProperties[patchVariable].(float64)
		amountToGive := patchAmount * percent / 4
		diffusions[patch] = amountToGive
	})

	//go through each patch and get the new amount
	m.Patches.Ask(func(patch *Patch) {
		amountFromNeighbors := 0.0
		neighbors := m.neighbors4(patch)
		if neighbors.Count() > 4 || neighbors.Count() < 2 {
			return
		}
		neighbors.Ask(func(n *Patch) {
			amountFromNeighbors += diffusions[n]
		})

		patchAmount := patch.patchProperties[patchVariable].(float64)
		amountToKeep := (patchAmount * (1 - percent)) + (float64(4-neighbors.Count()) * (patchAmount * percent / 4))

		patch.patchProperties[patchVariable] = amountToKeep + amountFromNeighbors
	})

	return nil
}

// returns the directed link breed associated with the name
func (m *Model) DirectedLinkBreed(name string) *LinkBreed {
	return m.directedLinkBreeds[name]
}

// returns all the directed link breeds
func (m *Model) DirectedLinkBreeds() []*LinkBreed {
	arr := []*LinkBreed{}

	for _, val := range m.directedLinkBreeds {
		arr = append(arr, val)
	}

	return arr
}

// returns the undirected link breed associated with the name
func (m *Model) UndirectedLinkBreed(name string) *LinkBreed {
	return m.undirectedLinkBreeds[name]
}

// returns all the undirected link breeds
func (m *Model) UndirectedLinkBreeds() []*LinkBreed {
	arr := []*LinkBreed{}

	for _, val := range m.undirectedLinkBreeds {
		arr = append(arr, val)
	}

	return arr
}

// returns the linkset containing the directed links
// to get the links for a breed call <linkBreed>.Links()
func (m *Model) DirectedLinks() *LinkAgentSet {
	return m.directedLinkBreeds[BreedNone].links
}

// returns the linkset containing the undirected links
// to get the links for a breed call <linkBreed>.Links()
func (m *Model) UndirectedLinks() *LinkAgentSet {
	return m.directedLinkBreeds[BreedNone].links
}

// returns the link agentset containing all the links
func (m *Model) Links() *LinkAgentSet {
	return m.links
}

// returns the distance between two points
// convertXYToInBounds should be called before this function
func (m *Model) DistanceBetweenPoints(x1 float64, y1 float64, x2 float64, y2 float64) float64 {

	deltaX := x1 - x2
	deltaY := y1 - y2

	distance := math.Sqrt(deltaX*deltaX + deltaY*deltaY)

	if !m.wrappingX && !m.wrappingY {
		return distance
	}

	deltaXInverse := float64(m.worldWidth) - math.Abs(deltaX)
	deltaYInverse := float64(m.worldHeight) - math.Abs(deltaY)

	if m.wrappingX {
		distance = math.Min(distance, math.Sqrt(deltaXInverse*deltaXInverse+deltaY*deltaY))
	}

	if m.wrappingY {
		distance = math.Min(distance, math.Sqrt(deltaX*deltaX+deltaYInverse*deltaYInverse))
	}

	if m.wrappingX && m.wrappingY {
		distance = math.Min(distance, math.Sqrt(deltaXInverse*deltaXInverse+deltaYInverse*deltaYInverse))
	}

	return distance
}

// layout the turtles in a circle with the specified radius
func (m *Model) LayoutCircle(turtles []*Turtle, radius float64) {
	amount := len(turtles)
	for i := 0; i < amount; i++ {
		agent := turtles[i]
		agent.SetXY(radius*math.Cos(2*math.Pi*float64(i)/float64(amount)), radius*math.Sin(2*math.Pi*float64(i)/float64(amount)))
		heading := 2 * math.Pi * float64(i) / float64(amount)
		agent.setHeadingRadians(heading)
	}
}

// returns a link between two turtles that connects from turtle1 to turtle2
func (m *Model) Link(turtle1 int, turtle2 int) *Link {

	generalBreed := m.directedLinkBreeds[BreedNone]

	return m.linkBreeded(generalBreed, turtle1, turtle2)
}

func (m *Model) linkBreeded(breed *LinkBreed, turtle1 int, turtle2 int) *Link {
	t1 := m.whoToTurtles[turtle1]
	t2 := m.whoToTurtles[turtle2]

	if t1 == nil || t2 == nil {
		return nil
	}

	return m.linkedTurtles[t1].getLink(breed, t2)
}

// returns a link that is directed that connects from turtle1 to turtle2
func (m *Model) LinkDirected(turtle1 int, turtle2 int) *Link {

	generalBreed := m.directedLinkBreeds[BreedNone]

	return m.linkDirectedBreed(generalBreed, turtle1, turtle2)
}

func (m *Model) linkDirectedBreed(breed *LinkBreed, turtle1 int, turtle2 int) *Link {
	t1 := m.whoToTurtles[turtle1]
	t2 := m.whoToTurtles[turtle2]

	if t1 == nil || t2 == nil {
		return nil
	}

	return m.linkedTurtles[t1].getLinkDirected(breed, t2)
}

// returns the maximum patch x coordinate
// the maximum x coordinate for a turtle is MaxPxCor() + .5
func (m *Model) MaxPxCor() int {
	return m.maxPxCor
}

// returns the maximum patch y coordinate
// the maximum y coordinate for a turtle is MaxPyCor() + .5
func (m *Model) MaxPyCor() int {
	return m.maxPyCor
}

// returns the minimum patch x coordinate
// the minimum x coordinate for a turtle is MinPxCor() - .5
func (m *Model) MinPxCor() int {
	return m.minPxCor
}

// returns the minimum patch y coordinate
// the minimum y coordinate for a turtle is MinPyCor() - .5
func (m *Model) MinPyCor() int {
	return m.minPyCor
}

func (m *Model) MaxXCor() float64 {
	return m.maxXCor
}

func (m *Model) MaxYCor() float64 {
	return m.maxYCor
}

func (m *Model) MinXCor() float64 {
	return m.minXCor
}

func (m *Model) MinYCor() float64 {
	return m.minYCor
}

// does not implement wrappimg, that is the responsibilty of the caller
// should only be called by Patch()!!! since Patch correctly converts the floats to ints
func (m *Model) getPatchAtCoords(x int, y int) *Patch {
	if x < m.minPxCor || x > m.maxPxCor || y < m.minPyCor || y > m.maxPyCor {
		return nil
	}

	pos := y*m.worldHeight + x

	return m.getPatchAtPos(pos)
}

// returns a random int n the provided list
func (m *Model) OneOfInt(arr []int) interface{} {
	return arr[m.randomGenerator.IntN(len(arr))-1]
}

// returns a random int from (0, n]
func (m *Model) RandomAmount(n int) int {
	return m.randomGenerator.IntN(n)
}

func (m *Model) topLeftNeighbor(p *Patch) *Patch {
	x := p.x - 1
	y := p.y - 1

	if x < m.minPxCor {
		if !m.wrappingX {
			return nil
		} else {
			x = m.maxPxCor
		}
	}

	if y < m.minPyCor {
		if !m.wrappingY {
			return nil
		} else {
			y = m.maxPyCor
		}
	}

	n := m.patchIndex(x, y)

	return m.getPatchAtPos(n)
}

func (m *Model) topNeighbor(p *Patch) *Patch {
	y := p.y - 1

	if y < m.minPyCor {
		if !m.wrappingY {
			return nil
		} else {
			y = m.maxPyCor
		}
	}

	n := m.patchIndex(p.x, y)

	return m.getPatchAtPos(n)
}

func (m *Model) topRightNeighbor(p *Patch) *Patch {
	x := p.x + 1
	y := p.y - 1

	if x > m.maxPxCor {
		if !m.wrappingX {
			return nil
		} else {
			x = m.minPxCor
		}
	}

	if y < m.minPyCor {
		if !m.wrappingY {
			return nil
		} else {
			y = m.maxPyCor
		}
	}

	n := m.patchIndex(x, y)

	return m.getPatchAtPos(n)
}

func (m *Model) leftNeighbor(p *Patch) *Patch {
	x := p.x - 1

	if x < m.minPxCor {
		if !m.wrappingX {
			return nil
		} else {
			x = m.maxPxCor
		}
	}

	n := m.patchIndex(x, p.y)

	return m.getPatchAtPos(n)
}

func (m *Model) rightNeighbor(p *Patch) *Patch {
	x := p.x + 1

	if x > m.maxPxCor {
		if !m.wrappingX {
			return nil
		} else {
			x = m.minPxCor
		}
	}

	n := m.patchIndex(x, p.y)

	return m.getPatchAtPos(n)
}

func (m *Model) bottomLeftNeighbor(p *Patch) *Patch {
	x := p.x - 1
	y := p.y + 1

	if x < m.minPxCor {
		if !m.wrappingX {
			return nil
		} else {
			x = m.maxPxCor
		}
	}

	if y > m.maxPyCor {
		if !m.wrappingX {
			return nil
		} else {
			y = m.minPyCor
		}
	}

	n := m.patchIndex(x, y)

	return m.getPatchAtPos(n)
}

func (m *Model) bottomNeighbor(p *Patch) *Patch {
	y := p.y + 1

	if y > m.maxPyCor {
		if !m.wrappingY {
			return nil
		} else {
			y = m.minPyCor
		}
	}

	n := m.patchIndex(p.x, y)

	return m.getPatchAtPos(n)
}

func (m *Model) bottomRightNeighbor(p *Patch) *Patch {
	x := p.x + 1
	y := p.y + 1

	if x > m.maxPxCor {
		if !m.wrappingX {
			return nil
		} else {
			x = m.minPxCor
		}
	}

	if y > m.maxPyCor {
		if !m.wrappingY {
			return nil
		} else {
			y = m.minPyCor
		}
	}

	n := m.patchIndex(x, y)

	return m.getPatchAtPos(n)
}

func (m *Model) neighbors(p *Patch) *PatchAgentSet {
	n := sortedset.NewSortedSet()

	for _, neighbor := range p.neighborsPatchMap {
		if neighbor != nil {
			n.Add(neighbor)
		}
	}

	return &PatchAgentSet{
		patches: n,
	}
}

func (m *Model) neighbors4(p *Patch) *PatchAgentSet {
	n := sortedset.NewSortedSet()

	top := p.neighborsPatchMap["top"]
	if top != nil {
		n.Add(top)
	}

	left := p.neighborsPatchMap["left"]
	if left != nil {
		n.Add(left)
	}

	right := p.neighborsPatchMap["right"]
	if right != nil {
		n.Add(right)
	}

	bottom := p.neighborsPatchMap["bottom"]
	if bottom != nil {
		n.Add(bottom)
	}

	return &PatchAgentSet{
		patches: n,
	}
}

func (m *Model) getPatchAtPos(x int) *Patch {
	return m.posOfPatches[x]
}

// returns the patch at the provided x y coordinates
func (m *Model) Patch(pxcor float64, pycor float64) *Patch {

	// round the x and y except in cases where the x or y is the min value
	// since the min value will be -*.5 and we want to round up in that case
	var x int
	var y int
	if pxcor == m.minXCor {
		x = int(math.Ceil(pxcor))
	} else {
		x = int(math.Round(pxcor))
	}

	if pycor == m.minYCor {
		y = int(math.Ceil(pycor))
	} else {
		y = int(math.Round(pycor))
	}

	// check if the x and y are within the world bounds
	// if wrapping is enabled then adjust the x and y to be within the world bounds if needed
	if x < m.minPxCor {
		if m.wrappingX {
			x = m.maxPxCor + 1 + ((x - m.minPxCor) % m.worldWidth)
		} else {
			return nil
		}
	}

	if y < m.minPyCor {
		if m.wrappingY {
			y = m.maxPyCor + 1 + ((y - m.minPyCor) % m.worldHeight)
		} else {
			return nil
		}
	}

	if x > m.maxPxCor {
		if m.wrappingX {
			x = (x-m.maxPxCor)%m.worldWidth + m.minPxCor - 1
		} else {
			return nil
		}
	}

	if y > m.maxPyCor {
		if m.wrappingY {
			y = (y-m.maxPyCor)%m.worldHeight + m.minPyCor - 1
		} else {
			return nil
		}
	}

	return m.getPatchAtCoords(x, y)
}

func (m *Model) RandomColor() Color {
	return baseColorsList[m.randomGenerator.IntN(len(baseColorsList))]
}

// If number is positive, reports a random floating point number greater than or equal to 0 but strictly less than number.
// If number is negative, reports a random floating point number less than or equal to 0, but strictly greater than number.
// If number is zero, the result is always 0.
func (m *Model) RandomFloat(number float64) float64 {
	return m.randomGenerator.Float64() * number
}

// If number is positive, reports a random integer greater than or equal to 0, but strictly less than number.
// If number is negative, reports a random integer less than or equal to 0, but strictly greater than number.
// If number is zero, the result is always 0 as well.
func (m *Model) RandomInt(number int) int {
	if number == 0 {
		return 0
	}

	sign := 1
	if number < 0 {
		sign = -1
		number = number * -1
	}

	return m.randomGenerator.IntN(number) * sign
}

// returns a random x cor that is within the world bounds
func (m *Model) RandomXCor() float64 {
	return m.RandomFloat(m.maxXCor-m.minXCor) + m.minXCor
}

// returns a random y cor that is within the world bounds
func (m *Model) RandomYCor() float64 {
	return m.RandomFloat(m.maxYCor-m.minYCor) + m.minYCor
}

// sets the tick counter to zero
func (m *Model) ResetTicks() {
	m.Ticks = 0
}

// resets the timer
func (m *Model) ResetTimer() {
	m.modelStart = time.Now()
}

func (m *Model) GetRandomState() (uint64, uint64, []byte) {

	bin, err := m.randomSrc.MarshalBinary()
	if err != nil {
		return 0, 0, nil
	}

	return m.seedValue, m.seedValue2, bin
}

func (m *Model) SetRandomState(seed1 uint64, seed2 uint64, state []byte) error {
	m.seedValue = seed1
	m.seedValue2 = seed2

	err := m.randomSrc.UnmarshalBinary(state)
	if err != nil {
		return err
	}

	m.randomGenerator = rand.New(m.randomSrc)

	return nil
}

// sets the default shape for links
func (m *Model) SetDefaultShapeLinks(shape string) {
	m.DefaultShapeLinks = shape
}

// sets the default shape for turtles
func (m *Model) SetDefaultShapeTurtles(shape string) {
	m.DefaultShapeTurtles = shape
}

// increments the tick counter by one
func (m *Model) Tick() {
	m.Ticks++
}

// increments the tick counter by the provided amount
func (m *Model) TickAdvance(amount int) {
	m.Ticks += amount
}

// returns the time since the model was started in milliseconds
func (m *Model) Timer() int64 {
	return time.Since(m.modelStart).Milliseconds()
}

// provides a turtle from the model given a who number
func (m *Model) Turtle(who int) *Turtle {

	t := m.whoToTurtles[who]
	if t == nil {
		return nil //turtle not found
	}

	return t
}

func (m *Model) TurtleWillCollide(turtle *Turtle, distance float64, biggestSize float64) bool {

	dx := distance * math.Cos(turtle.heading)
	dy := distance * math.Sin(turtle.heading)

	// radius of the given turtle
	radius := turtle.GetSize() / 2

	tX := turtle.XCor() + dx
	tY := turtle.YCor() + dy

	tX, tY, inBounds := m.convertXYToInBounds(tX, tY)
	if !inBounds {
		return false
	}

	// get the turtles that are in the radius of the turtle and who might be given the biggest possible size
	potentialTurtles := m.TurtlesInRadius(tX, tY, radius+(biggestSize/2))

	// remove the turtle from the potential turtles
	potentialTurtles.Remove(turtle)

	// if there are no turtles in the radius then there is no collision
	if potentialTurtles.Count() == 0 {
		return false
	}
	// check if any of the potential turtles actually do collide with the given turtle
	if potentialTurtles.Any(func(t *Turtle) bool {
		return m.TurtlesCollide(turtle, t, distance, 0, 0)
	}) {
		return true
	}

	return false
}

// returns the turtle agentset
func (m *Model) Turtles() *TurtleAgentSet {
	return m.turtles
}

// returns true if the two turtles are within the provided distanc
func (m *Model) TurtlesCollide(t1 *Turtle, t2 *Turtle, t1dist float64, t2dist float64, difference float64) bool {

	t1dx := t1dist * math.Cos(t1.heading)
	t1dy := t1dist * math.Sin(t1.heading)

	t2dx := t2dist * math.Cos(t2.heading)
	t2dy := t2dist * math.Sin(t2.heading)

	t1X := t1.XCor() + t1dx
	t1Y := t1.YCor() + t1dy

	t2X := t2.XCor() + t2dx
	t2Y := t2.YCor() + t2dy

	t1X, t1Y, inBounds1 := m.convertXYToInBounds(t1X, t1Y)
	t2X, t2Y, inBounds2 := m.convertXYToInBounds(t2X, t2Y)

	if !inBounds1 || !inBounds2 {
		return false
	}

	distance := m.DistanceBetweenPoints(t1X, t1Y, t2X, t2Y)

	radius1 := t1.GetSize() / 2

	radius2 := t2.GetSize() / 2

	return distance <= radius1+radius2+difference
}

// returns the turtle agentset that is on patch of the proviced x y coordinates
// same as TurtlesOnPatch(breed, Patch(x, y))
// the agentset returned is a pointer to the agentset in the patch, to get a copy of the agentset call .Copy()
func (m *Model) TurtlesAtCoords(pxcor float64, pycor float64) *TurtleAgentSet {

	generalBreed := m.breeds[BreedNone]

	return m.turtlesAtCoordsBreeded(generalBreed, pxcor, pycor)
}

func (m *Model) turtlesAtCoordsBreeded(breed *TurtleBreed, pxcor float64, pycor float64) *TurtleAgentSet {
	x := math.Round(pxcor)
	y := math.Round(pycor)

	patch := m.Patch(x, y)

	if patch == nil {
		return nil
	}

	return patch.turtlesHereBreeded(breed)
}

// returns an agentset of turtles that are within the provided radius of the provided x y coordinates
func (m *Model) TurtlesInRadius(xCor float64, yCor float64, radius float64) *TurtleAgentSet {
	// xMin, xMax := int(math.Floor(cx-r)), int(math.Ceil(cx+r))
	// yMin, yMax := int(math.Floor(cy-r)), int(math.Ceil(cy+r))

	xMin := int(math.Floor(xCor - radius))
	xMax := int(math.Ceil(xCor + radius))
	yMin := int(math.Floor(yCor - radius))
	yMax := int(math.Ceil(yCor + radius))

	//are we going to be looping through more patches than there are turtles?
	//if so, just loop through all the turtles
	if (xMax-xMin+1)*(yMax-yMin+1) > m.turtles.Count() {
		return m.Turtles().With(func(t *Turtle) bool {
			return m.DistanceBetweenPoints(xCor, yCor, t.XCor(), t.YCor()) <= radius
		})
	}

	patchesFullyInsideRadius := make([]*Patch, 0)
	patchesPartiallyInsideRadius := make([]*Patch, 0)

	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {

			// get the patch, this will handle wrapping
			// if it is nil, then that probably means the world is not wrapping and the x y is out of bounds
			patch := m.Patch(float64(x), float64(y))
			if patch == nil {
				continue
			}

			// if there are no turtles on the patch then we can skip it
			if patch.TurtlesHere().Count() == 0 {
				continue
			}

			// get the center of the patch from the patch
			// this makes sure that this is actually the center of the patch, just in case the coords shifted due to world wrapping
			px := float64(patch.XCor())
			py := float64(patch.YCor())

			// patch corners
			corners := [4][2]float64{
				{px - 0.5, py - 0.5}, // bottom-left
				{px + 0.5, py - 0.5}, // bottom-right
				{px - 0.5, py + 0.5}, // top-left
				{px + 0.5, py + 0.5}, // top-right
			}

			// check how many corners are inside the circle
			insideCount := 0
			for _, corner := range corners {
				if m.DistanceBetweenPoints(xCor, yCor, corner[0], corner[1]) <= radius {
					insideCount++
				}
			}

			// classify
			if insideCount == 4 {
				patchesFullyInsideRadius = append(patchesFullyInsideRadius, patch)
			} else if insideCount > 0 {
				patchesPartiallyInsideRadius = append(patchesPartiallyInsideRadius, patch)
			}
		}
	}

	turtles := NewTurtleAgentSet(nil)

	// add turtles from patches that are fully inside the radius
	for _, patch := range patchesFullyInsideRadius {
		patch.TurtlesHere().Ask(func(t *Turtle) {
			turtles.Add(t)
		})
	}

	// add turtles from patches that are partially inside the radius provided they are within the radius
	for _, patch := range patchesPartiallyInsideRadius {
		patch.TurtlesHere().Ask(func(t *Turtle) {
			if m.DistanceBetweenPoints(xCor, yCor, t.XCor(), t.YCor()) <= radius {
				turtles.Add(t)
			}
		})
	}

	return turtles
}

// returns the turtle agentset that is on the provided patch
// the agentset returned is a pointer to the agentset in the patch, to get a copy of the agentset call .Copy()
func (m *Model) TurtlesOnPatch(patch *Patch) *TurtleAgentSet {

	generalBreed := m.breeds[BreedNone]

	return m.turtlesOnPatchBreeded(generalBreed, patch)
}

func (m *Model) turtlesOnPatchBreeded(breed *TurtleBreed, patch *Patch) *TurtleAgentSet {
	return patch.turtlesHereBreeded(breed)
}

// Returns the turtles on the provided patches
// the agentset returned is a pointer to the agentset in the patch, to get a copy of the agentset call .Copy()
func (m *Model) TurtlesOnPatches(patches *PatchAgentSet) *TurtleAgentSet {

	generalBreed := m.breeds[BreedNone]

	return m.turtlesOnPatchesBreeded(generalBreed, patches)
}

func (m *Model) turtlesOnPatchesBreeded(breed *TurtleBreed, patches *PatchAgentSet) *TurtleAgentSet {
	turtles := NewTurtleAgentSet(nil)

	patches.Ask(func(patch *Patch) {
		s := m.turtlesOnPatchBreeded(breed, patch)
		s.Ask(func(turtle *Turtle) {
			turtles.Add(turtle)
		})
	})

	return turtles
}

// Returns the turtles on the same patch as the provided turtle
// the agentset returned is a pointer to the agentset in the patch, to get a copy of the agentset call .Copy()
func (m *Model) TurtlesWithTurtle(turtle *Turtle) *TurtleAgentSet {

	generalBreed := m.breeds[BreedNone]

	return m.turtlesWithTurtleBreeded(generalBreed, turtle)
}

func (m *Model) turtlesWithTurtleBreeded(breed *TurtleBreed, turtle *Turtle) *TurtleAgentSet {
	p := turtle.PatchHere()
	if p == nil {
		return nil
	}

	return p.turtlesHereBreeded(breed)
}

// Returns the turtles on the same patch as the provided turtle
// the agentset returned is a pointer to the agentset in the patch, to get a copy of the agentset call .Copy()
func (m *Model) TurtlesWithTurtles(turtles *TurtleAgentSet) *TurtleAgentSet {

	generalBreed := m.breeds[BreedNone]

	return m.turtlesWithTurtlesBreeded(generalBreed, turtles)
}

func (m *Model) turtlesWithTurtlesBreeded(breed *TurtleBreed, turtles *TurtleAgentSet) *TurtleAgentSet {
	patches := NewPatchAgentSet(nil)

	turtles.Ask(func(turtle *Turtle) {
		p := turtle.PatchHere()
		if p != nil {
			patches.Add(p)
		}
	})

	return m.turtlesOnPatchesBreeded(breed, patches)
}

// returns the world height
func (m *Model) WorldHeight() int {
	return m.worldHeight
}

// returns the world width
func (m *Model) WorldWidth() int {
	return m.worldWidth
}

// returns if wrapping x is on
func (m *Model) WrappingX() bool {
	return m.wrappingX
}

// sets the x coordinate to wrap
func (m *Model) WrappingXOn() {
	m.wrappingX = true
}

// sets the x coordinate to not wrap
func (m *Model) WrappingXOff() {
	m.wrappingX = false
}

// returns if wrapping y is on
func (m *Model) WrappingY() bool {
	return m.wrappingY
}

// sets the y coordinate to wrap
func (m *Model) WrappingYOn() {
	m.wrappingY = true
}

// sets the y coordinate to not wrap
func (m *Model) WrappingYOff() {
	m.wrappingY = false
}
