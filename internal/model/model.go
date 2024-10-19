package model

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/nlatham1999/sortedset"
)

type Model struct {
	Ticks int

	patchesOwnTemplate map[string]interface{} //additional variables for each patch

	Patches              *PatchAgentSet          //all the patches
	turtles              *TurtleAgentSet         //all the turtles
	links                *LinkAgentSet           //all the links
	breeds               map[string]*turtleBreed //turtles split into breeds
	directedLinkBreeds   map[string]*linkBreed   //directed link breeds
	undirectedLinkBreeds map[string]*linkBreed   //undirected link breeds

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

	modelStart time.Time

	Globals map[string]interface{}
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

	patchesOwn := make(map[string]interface{})
	if settings.PatchesOwn != nil {
		for key, value := range settings.PatchesOwn {
			patchesOwn[key] = value
		}
	}

	model := &Model{
		maxPxCor:           settings.MaxPxCor,
		maxPyCor:           settings.MaxPyCor,
		minPxCor:           settings.MinPxCor,
		minPyCor:           settings.MinPyCor,
		maxXCor:            float64(settings.MaxPxCor) + .5,
		maxYCor:            float64(settings.MaxPyCor) + .5,
		minXCor:            float64(settings.MinPxCor) - .5,
		minYCor:            float64(settings.MinPyCor) - .5,
		worldWidth:         settings.MaxPxCor - settings.MinPxCor + 1,
		worldHeight:        settings.MaxPyCor - settings.MinPyCor + 1,
		patchesOwnTemplate: settings.PatchesOwn,
		wrappingX:          settings.WrappingX,
		wrappingY:          settings.WrappingY,
		whoToTurtles:       make(map[int]*Turtle),
		randomGenerator:    rand.New(rand.NewSource(0)),
		modelStart:         time.Now(),
		Globals:            make(map[string]interface{}),
	}

	//construct turtle breeds
	turtleBreedsMap := make(map[string]*turtleBreed)
	for i := 0; i < len(settings.TurtleBreeds); i++ {
		turtleBreedsMap[settings.TurtleBreeds[i]] = &turtleBreed{
			turtles: &TurtleAgentSet{
				turtles: make(map[*Turtle]interface{}),
			},
			name:         settings.TurtleBreeds[i],
			defaultShape: "",
		}

		//copy the turtles breeds own template
		if settings.TurtleBreedsOwn != nil && settings.TurtleBreedsOwn[settings.TurtleBreeds[i]] != nil {
			turtleBreedsMap[settings.TurtleBreeds[i]].turtlesOwnTemplate = make(map[string]interface{})
			for key, value := range settings.TurtleBreedsOwn[settings.TurtleBreeds[i]] {
				turtleBreedsMap[settings.TurtleBreeds[i]].turtlesOwnTemplate[key] = value
			}
		} else {
			turtleBreedsMap[settings.TurtleBreeds[i]].turtlesOwnTemplate = make(map[string]interface{})
		}
	}
	model.breeds = turtleBreedsMap

	//construct directed link breeds
	directedLinkBreedsMap := make(map[string]*linkBreed)
	settings.DirectedLinkBreeds = append(settings.DirectedLinkBreeds, "") // add the general population
	for i := 0; i < len(settings.DirectedLinkBreeds); i++ {
		directedLinkBreedsMap[settings.DirectedLinkBreeds[i]] = &linkBreed{
			links: &LinkAgentSet{
				links: *sortedset.NewSortedSet(),
			},
			Directed:     true,
			DefaultShape: "",
			name:         settings.DirectedLinkBreeds[i],
		}
	}
	model.directedLinkBreeds = directedLinkBreedsMap

	//construct undirected link breeds
	undirectedLinkBreedsMap := make(map[string]*linkBreed)
	settings.UndirectedLinkBreeds = append(settings.UndirectedLinkBreeds, "") // add the general population
	for i := 0; i < len(settings.UndirectedLinkBreeds); i++ {
		undirectedLinkBreedsMap[settings.UndirectedLinkBreeds[i]] = &linkBreed{
			links: &LinkAgentSet{
				links: *sortedset.NewSortedSet(),
			},
			Directed:     false,
			DefaultShape: "",
			name:         settings.UndirectedLinkBreeds[i],
		}
	}
	model.undirectedLinkBreeds = undirectedLinkBreedsMap

	//construct general turtle set
	model.turtles = &TurtleAgentSet{
		turtles: make(map[*Turtle]interface{}),
	}

	// create a breed with no name for the general population
	model.breeds[""] = &turtleBreed{
		turtles:            model.turtles,
		name:               "",
		defaultShape:       "",
		turtlesOwnTemplate: make(map[string]interface{}),
	}
	if settings.TurtlesOwn != nil {
		for key, value := range settings.TurtlesOwn {
			model.breeds[""].turtlesOwnTemplate[key] = value
		}
	} else {
		model.breeds[""].turtlesOwnTemplate = make(map[string]interface{})
	}

	//construct general link set
	model.links = &LinkAgentSet{
		links: *sortedset.NewSortedSet(),
	}

	// build patches
	model.buildPatches()

	// build globals
	for key, value := range settings.Globals {
		model.Globals[key] = value
	}

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
			p := newPatch(m, m.patchesOwnTemplate, x, y)
			m.Patches.Add(p)
			index := y*m.worldHeight + x
			m.posOfPatches[index] = p
			p.index = index
		}
	}

	for _, p := range m.Patches.List() {
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
	return y*m.worldHeight + x
}

// clear all patches and turtles and set the ticks to zero
func (m *Model) ClearAll() {
	m.ClearTicks()
	m.ClearPatches()
	m.ClearTurtles()
}

// clear all links
func (m *Model) ClearLinks() {
	m.links.links = *sortedset.NewSortedSet()
	m.links = &LinkAgentSet{
		links: *sortedset.NewSortedSet(),
	}
	for breed := range m.directedLinkBreeds {
		m.directedLinkBreeds[breed].links = &LinkAgentSet{
			links: *sortedset.NewSortedSet(),
		}
	}
	for breed := range m.undirectedLinkBreeds {
		m.undirectedLinkBreeds[breed].links = &LinkAgentSet{
			links: *sortedset.NewSortedSet(),
		}
	}
	for turtle := range m.turtles.turtles {
		turtle.linkedTurtles = newTurtleLinks()
	}
}

// set the ticks to zero
func (m *Model) ClearTicks() {
	m.Ticks = 0
}

// clear all patches
func (m *Model) ClearPatches() {
	for _, patch := range m.Patches.List() {
		patch.Reset(m.patchesOwnTemplate)
	}
}

// kills all turtles
func (m *Model) ClearTurtles() {
	// delete all links since they are linked to turtles
	m.links.links = *sortedset.NewSortedSet()
	for breed := range m.directedLinkBreeds {
		m.directedLinkBreeds[breed].links.links = *sortedset.NewSortedSet()
	}
	for breed := range m.undirectedLinkBreeds {
		m.undirectedLinkBreeds[breed].links.links = *sortedset.NewSortedSet()
	}

	// remove all turtles from patches
	for _, patch := range m.Patches.List() {
		patch.turtles = make(map[string]*TurtleAgentSet)
	}

	// clear all turtles
	for turtle := range m.turtles.turtles {
		*turtle = Turtle{}
	}

	m.turtles.Clear()
	for breed := range m.breeds {
		m.breeds[breed].turtles.turtles = make(map[*Turtle]interface{})
	}

	m.whoToTurtles = make(map[int]*Turtle)

	m.turtlesWhoNumber = 0
}

// like create turtles but goes through the list of colors and evenly spaces out the headings
func (m *Model) CreateOrderedTurtles(breed string, amount int, operations []TurtleOperation) error {
	if breed != "" {
		_, found := m.breeds[breed]
		if !found {
			return errors.New("breed not found")
		}
	}

	end := amount + m.turtlesWhoNumber
	count := 0
	headingAmount := 2 * math.Pi / float64(amount)
	turtles := []*Turtle{}
	for m.turtlesWhoNumber < end {
		newTurtle := NewTurtle(m, m.turtlesWhoNumber, breed, 0, 0)

		// set heading to be random
		newTurtle.setHeadingRadians(headingAmount * float64(count))

		newTurtle.Color.SetColor(baseColorsList[count%len(baseColorsList)])

		count++

		turtles = append(turtles, newTurtle)

		m.turtlesWhoNumber++
	}

	for _, turtle := range turtles {
		for i := 0; i < len(operations); i++ {
			operations[i](turtle)
		}
	}

	return nil
}

// create the specified amount of turtles with the specified breed and operations
// if the breed is empty then it will add it to the general population
func (m *Model) CreateTurtles(amount int, breed string, operations []TurtleOperation) error {

	if breed != "" {
		_, found := m.breeds[breed]
		if !found {
			return errors.New("breed not found")
		}
	}

	turtles := []*Turtle{}

	end := amount + m.turtlesWhoNumber
	for m.turtlesWhoNumber < end {
		newTurtle := NewTurtle(m, m.turtlesWhoNumber, breed, 0, 0)

		// set heading to be random
		newTurtle.setHeadingRadians(m.randomGenerator.Float64() * 2 * math.Pi)

		// get a random color from the base colors and set it
		newTurtle.Color.SetColor(baseColorsList[rand.Intn(len(baseColorsList))])

		turtles = append(turtles, newTurtle)

		m.turtlesWhoNumber++
	}

	for _, turtle := range turtles {
		for i := 0; i < len(operations); i++ {
			operations[i](turtle)
		}
	}

	return nil
}

// if the topology allows it then convert the x y to within bounds if it is outside of the world
// returns the new x y and if it is in bounds
// returns false if the x y is not in bounds and the topology does not allow it
func (m *Model) convertXYToInBounds(x float64, y float64) (float64, float64, bool) {

	if x < m.minXCor {
		if m.wrappingX {
			x = m.maxXCor - math.Mod(m.minXCor-x, float64(m.worldWidth)) + 1
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
			y = m.maxYCor - math.Mod(m.minYCor-y, float64(m.worldHeight)) + 1
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
	if turtle.breed != "" {
		m.breeds[turtle.breed].turtles.Remove(turtle)
	}
	delete(m.whoToTurtles, turtle.who)

	p := turtle.PatchHere()
	if p != nil {
		delete(p.turtles[""].turtles, turtle)
		if turtle.breed != "" {
			p.turtles[turtle.breed].Remove(turtle)
		}
	}

	// kill all directed out links
	for link := range turtle.linkedTurtles.getAllDirectedOutLinks() {
		m.KillLink(link)
	}

	// kill all directed in links
	for link := range turtle.linkedTurtles.getAllDirectedInLinks() {
		m.KillLink(link)
	}

	// kill all undirected links
	for link := range turtle.linkedTurtles.getAllUndirectedLinks() {
		m.KillLink(link)
	}

	*turtle = Turtle{}
}

// kills a link
func (m *Model) KillLink(link *Link) {

	m.links.links.Remove(link)

	if link.breed != "" {
		if link.directed {
			m.directedLinkBreeds[link.breed].links.links.Remove(link)
		} else {
			m.undirectedLinkBreeds[link.breed].links.links.Remove(link)
		}
	}

	// remove the link from the turtles
	if link.directed {
		link.end1.linkedTurtles.removeDirectedOutBreed(link.breed, link.end2, link)
		link.end2.linkedTurtles.removeDirectedInBreed(link.breed, link.end1, link)
	} else {
		link.end1.linkedTurtles.removeUndirectedBreed(link.breed, link.end2, link)
		link.end2.linkedTurtles.removeUndirectedBreed(link.breed, link.end1, link)
	}

	*link = Link{}
}

// diffuse the patch variable of each patch to its neighbors
func (m *Model) Diffuse(patchVariable string, percent float64) error {

	if percent > 1 || percent < 0 {
		return errors.New("percent amount was outside bounds")
	}

	diffusions := make(map[*Patch]float64)

	patchList := m.Patches.List()

	//go through each patch and calculate the diffusion amount
	for _, patch := range patchList {
		patchAmount := patch.patchesOwn[patchVariable].(float64)
		amountToGive := patchAmount * percent / 8
		diffusions[patch] = amountToGive
	}

	//go through each patch and get the new amount
	for _, patch := range patchList {

		amountFromNeighbors := 0.0
		neighbors := m.neighbors(patch)
		if neighbors.Count() > 8 || neighbors.Count() < 3 {
			return errors.New("invalid amount of neighbors")
		}
		for _, n := range neighbors.List() {
			amountFromNeighbors += diffusions[n]
		}

		patchAmount := patch.patchesOwn[patchVariable].(float64)
		amountToKeep := (patchAmount * (1 - percent)) + (float64(8-neighbors.Count()) * (patchAmount * percent / 8))

		patch.patchesOwn[patchVariable] = amountToKeep + amountFromNeighbors
	}

	return nil
}

// diffuse the patch variable of each patch to its neighbors at the top, bottom, left, and right
func (m *Model) Diffuse4(patchVariable string, percent float64) error {

	if percent > 1 || percent < 0 {
		return errors.New("percent amount was outside bounds")
	}

	diffusions := make(map[*Patch]float64)

	patchList := m.Patches.List()

	//go through each patch and calculate the diffusion amount
	for _, patch := range patchList {
		patchAmount := patch.patchesOwn[patchVariable].(float64)
		amountToGive := patchAmount * percent / 4
		diffusions[patch] = amountToGive
	}

	//go through each patch and get the new amount
	for _, patch := range patchList {

		amountFromNeighbors := 0.0
		neighbors := m.neighbors4(patch)
		if neighbors.Count() > 4 || neighbors.Count() < 2 {
			return errors.New("invalid amount of neighbors")
		}
		for _, n := range neighbors.List() {
			amountFromNeighbors += diffusions[n]
		}

		patchAmount := patch.patchesOwn[patchVariable].(float64)
		amountToKeep := (patchAmount * (1 - percent)) + (float64(4-neighbors.Count()) * (patchAmount * percent / 4))

		patch.patchesOwn[patchVariable] = amountToKeep + amountFromNeighbors
	}

	return nil
}

// returns the linkset containing the directed links for the specified breed
// if the breed is empty then it will return the general population of directed links
func (m *Model) DirectedLinks(breed string) *LinkAgentSet {
	if breed, ok := m.directedLinkBreeds[breed]; ok {
		return breed.links
	}
	return nil
}

// returns the linkset containing the undirected links for the specified breed
// if the breed is empty then it will return the general population of undirected links
func (m *Model) UndirectedLinks(breed string) *LinkAgentSet {
	if breed, ok := m.undirectedLinkBreeds[breed]; ok {
		return breed.links
	}
	return nil
}

// returns the link agentset containing all the links
func (m *Model) Links() *LinkAgentSet {
	return m.links
}

// returns the distance between two points
func (m *Model) DistanceBetweenPoints(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	deltaX := x1 - x2
	deltaY := y1 - y2

	distance := math.Abs(math.Sqrt(deltaX*deltaX + deltaY*deltaY))

	if !m.wrappingX && !m.wrappingY {
		return distance
	}

	deltaXInverse := float64(m.worldWidth) - math.Abs(deltaX)
	deltaYInverse := float64(m.worldHeight) - math.Abs(deltaY)

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

// gets a gobal variable
func (m *Model) GetGlobal(key string) interface{} {
	return m.Globals[key]
}

// gets a global variable as a bool
func (m *Model) GetGlobalB(key string) (bool, error) {
	v := m.Globals[key]
	if v == nil {
		return false, fmt.Errorf("global %s not found", key)
	}
	switch v := v.(type) {
	case bool:
		return v, nil
	default:
		return false, fmt.Errorf("global %s is not a bool", key)
	}
}

// gets a global variable as an int
func (m *Model) GetGlobalI(key string) (int, error) {
	v := m.Globals[key]
	if v == nil {
		return 0, fmt.Errorf("global %s not found", key)
	}
	switch v := v.(type) {
	case int:
		return v, nil
	case float64:
		return int(v), nil
	default:
		return 0, fmt.Errorf("global %s is not an int", key)
	}
}

// gets a global variable as a float64
func (m *Model) GetGlobalF(key string) (float64, error) {
	v := m.Globals[key]
	if v == nil {
		return 0, fmt.Errorf("global %s not found", key)
	}
	switch v := v.(type) {
	case int:
		return float64(v), nil
	case float64:
		return v, nil
	default:
		return 0, fmt.Errorf("global %s is not a float64", key)
	}
}

// gets a global variable as a string
func (m *Model) GetGlobalS(key string) (string, error) {
	v := m.Globals[key]
	if v == nil {
		return "", fmt.Errorf("global %s not found", key)
	}
	switch v := v.(type) {
	case string:
		return v, nil
	default:
		return "", fmt.Errorf("global %s is not a string", key)
	}
}

// sets a global variable
func (m *Model) SetGlobal(key string, value interface{}) {
	m.Globals[key] = value
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
// if the breed is empty then selects from the general population
func (m *Model) Link(breed string, turtle1 int, turtle2 int) *Link {
	t1 := m.whoToTurtles[turtle1]
	t2 := m.whoToTurtles[turtle2]

	if t1 == nil || t2 == nil {
		return nil
	}

	return t1.linkedTurtles.getLink(breed, t2)
}

// returns a link that is directed that connects from turtle1 to turtle2
// if the breed is empty then selects from the general population
func (m *Model) LinkDirected(breed string, turtle1 int, turtle2 int) *Link {
	t1 := m.whoToTurtles[turtle1]
	t2 := m.whoToTurtles[turtle2]

	if t1 == nil || t2 == nil {
		return nil
	}

	return t1.linkedTurtles.getLinkDirected(breed, t2)
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

// does not implement wrappimg, that is the responsibilty of the caller
func (m *Model) getPatchAtCoords(x int, y int) *Patch {
	if x < m.minPxCor || x > m.maxPxCor || y < m.minPyCor || y > m.maxPyCor {
		return nil
	}

	pos := y*m.worldHeight + x

	return m.posOfPatches[pos]
}

// returns a random int n the provided list
func (m *Model) OneOfInt(arr []int) interface{} {

	return arr[rand.Intn(len(arr))-1]
}

// returns a random int from (0, n]
func (m *Model) RandomAmount(n int) int {
	return rand.Intn(n)
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

	topLeft := p.neighborsPatchMap["topLeft"]
	if topLeft != nil {
		n.Add(topLeft)
	}

	left := p.neighborsPatchMap["left"]
	if left != nil {
		n.Add(left)
	}

	bottomLeft := p.neighborsPatchMap["bottomLeft"]
	if bottomLeft != nil {
		n.Add(bottomLeft)
	}

	top := p.neighborsPatchMap["top"]
	if top != nil {
		n.Add(top)
	}

	topRight := p.neighborsPatchMap["topRight"]
	if topRight != nil {
		n.Add(topRight)
	}

	right := p.neighborsPatchMap["right"]
	if right != nil {
		n.Add(right)
	}

	bottomRight := p.neighborsPatchMap["bottomRight"]
	if bottomRight != nil {
		n.Add(bottomRight)
	}

	bottom := p.neighborsPatchMap["bottom"]
	if bottom != nil {
		n.Add(bottom)
	}

	return &PatchAgentSet{
		patches: *n,
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
		patches: *n,
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

	return m.getPatchAtCoords(x, y)
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

	return m.randomGenerator.Intn(number) * sign
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

// resizes the world to the provided bounds
func (m *Model) ResizeWorld(minPxcor int, maxPxcor int, minPycor int, maxPycor int) {
	m.minPxCor = minPxcor
	m.maxPxCor = maxPxcor
	m.minPyCor = minPycor
	m.maxPyCor = maxPycor
	m.maxXCor = float64(maxPxcor) + .5
	m.maxYCor = float64(maxPycor) + .5
	m.minXCor = float64(minPxcor) - .5
	m.minYCor = float64(minPycor) - .5
	m.worldWidth = maxPxcor - minPxcor + 1
	m.worldHeight = maxPycor - minPycor + 1

	m.buildPatches()

	//@TODO: resize the turtles
}

// sets the default shape for links
func (m *Model) SetDefaultShapeLinks(shape string) {
	m.DefaultShapeLinks = shape
}

// sets the default shape for turtles
func (m *Model) SetDefaultShapeTurtles(shape string) {
	m.DefaultShapeTurtles = shape
}

// sets the default shape for a directed link breed
func (m *Model) SetDefaultShapeLinkBreed(breed string, shape string) {
	m.directedLinkBreeds[breed].DefaultShape = shape
}

// sets the default shape for a turtle breed
func (m *Model) SetDefaultShapeTurtleBreed(breed string, shape string) {
	m.breeds[breed].defaultShape = shape
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
		if m.breeds[breed] == nil {
			return nil //breed not found
		}
		if t.breed != breed {
			return nil //turtle not found for that breed
		}
		return t
	}
}

// returns the turtle agentset for the provided breed
func (m *Model) Turtles(breed string) *TurtleAgentSet {
	if breed == "" {
		return m.turtles
	}
	return m.breeds[breed].turtles
}

// returns the turtle agentset for the provided breed that is on the provided patch
func (m *Model) TurtlesAt(breed string, pxcor float64, pycor float64) *TurtleAgentSet {
	x := int(math.Round(pxcor))
	y := int(math.Round(pycor))

	patch := m.getPatchAtCoords(x, y)

	if patch == nil {
		return nil
	}

	return patch.TurtlesHere(breed)
}

// Returns the turtles on the provided patch
func (m *Model) TurtlesOnPatch(breed string, patch *Patch) *TurtleAgentSet {
	return patch.TurtlesHere(breed)
}

// Returns the turtles on the provided patches
func (m *Model) TurtlesOnPatches(breed string, patches *PatchAgentSet) *TurtleAgentSet {
	turtles := NewTurtleAgentSet(nil)

	for _, patch := range patches.List() {
		s := m.TurtlesOnPatch(breed, patch)
		for turtle := range s.turtles {
			turtles.Add(turtle)
		}
	}

	return turtles
}

// Returns the turtles on the same patch as the provided turtle
func (m *Model) TurtlesWithTurtle(breed string, turtle *Turtle) *TurtleAgentSet {
	p := turtle.PatchHere()
	if p == nil {
		return nil
	}

	return p.TurtlesHere(breed)
}

// Returns the turtles on the same patch as the provided turtle and with the breed provided
func (m *Model) TurtlesWithTurtles(breed string, turtles *TurtleAgentSet) *TurtleAgentSet {
	patches := NewPatchAgentSet(nil)

	for turtle := range turtles.turtles {
		p := turtle.PatchHere()
		if p != nil {
			patches.Add(p)
		}
	}

	return m.TurtlesOnPatches(breed, patches)
}

// returns the world height
func (m *Model) WorldHeight() int {
	return m.worldHeight
}

// returns the world width
func (m *Model) WorldWidth() int {
	return m.worldWidth
}

// sets the x coordinate to wrap
func (m *Model) WrappingXOn() {
	m.wrappingX = true
}

// sets the y coordinate to wrap
func (m *Model) WrappingYOn() {
	m.wrappingY = true
}

// sets the x coordinate to not wrap
func (m *Model) WrappingXOff() {
	m.wrappingX = false
}

// sets the y coordinate to not wrap
func (m *Model) WrappingYOff() {
	m.wrappingY = false
}
