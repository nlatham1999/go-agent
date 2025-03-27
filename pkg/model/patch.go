package model

import (
	"math"
)

// Patches are agents that resemble the physical space
type Patch struct {
	// should never be changed
	x int // x coordinate of the patch
	y int // y coordinate of the patch

	// this corresponds to the position in the patches array
	// set as x*m.worldWidth + y
	// maps to parent.posOfPatches[index]
	index int

	parent *Model

	//we have float54 versions of the variables so that we don't have to do a bunch of conversions
	xFloat64 float64
	yFloat64 float64

	PColor Color

	//instead it might be faster having a PatchesOwn for each data type to reduce type assertions
	patchProperties map[string]interface{}

	Label       interface{}
	PlabelColor Color

	turtles map[*TurtleBreed]*TurtleAgentSet // sets of turtles keyed by breed

	// patch to string and string to patch for the neighbors
	patchNeighborsMap map[*Patch]string
	neighborsPatchMap map[string]*Patch
}

func newPatch(m *Model, patchProperties map[string]interface{}, x int, y int) *Patch {

	patch := &Patch{
		x:        x,
		y:        y,
		xFloat64: float64(x),
		yFloat64: float64(y),
		PColor:   Color{},
		turtles:  make(map[*TurtleBreed]*TurtleAgentSet),
		parent:   m,
	}

	patch.PColor.SetColor(Black)

	patch.patchProperties = map[string]interface{}{}
	for key, value := range patchProperties {
		patch.patchProperties[key] = value
	}

	return patch
}

// links a turtle to this patch
func (p *Patch) addTurtle(t *Turtle) {
	if _, ok := p.turtles[t.breed]; !ok {
		p.turtles[t.breed] = NewTurtleAgentSet([]*Turtle{})
	}
	p.turtles[t.breed].Add(t)

	// if the breed is provided, add it to the general set of turtles as well
	if t.breed.name != "" {
		generalBreed := p.parent.breeds[BreedNone]
		if _, ok := p.turtles[generalBreed]; !ok {
			p.turtles[generalBreed] = NewTurtleAgentSet([]*Turtle{})
		}
		p.turtles[generalBreed].Add(t)
	}
}

// unlinks a turtle from this patch
func (p *Patch) removeTurtle(t *Turtle) {
	if _, ok := p.turtles[t.breed]; ok {
		p.turtles[t.breed].Remove(t)
	}

	// if the breed is provided, remove it from the general set of turtles as well
	if t.breed.name != "" {
		generalBreed := p.parent.breeds[BreedNone]
		if _, ok := p.turtles[generalBreed]; ok {
			p.turtles[generalBreed].Remove(t)
		}
	}
}

// returns the distance of this patch to the provided turtle
func (p *Patch) DistanceTurtle(t *Turtle) float64 {
	return p.parent.DistanceBetweenPoints(p.xFloat64, p.yFloat64, t.xcor, t.ycor)
}

// returns the distance of this patch to the provided patch
func (p *Patch) DistancePatch(patch *Patch) float64 {
	return p.parent.DistanceBetweenPoints(p.xFloat64, p.yFloat64, patch.xFloat64, patch.yFloat64)
}

// Returns the distance of this patch from the provided x y coordinates
func (p *Patch) DistanceXY(x float64, y float64) float64 {
	return p.parent.DistanceBetweenPoints(p.xFloat64, p.yFloat64, x, y)
}

// returns the neighbors of this patch
func (p *Patch) Neighbors() *PatchAgentSet {
	neighbors := p.parent.neighbors(p)

	return neighbors
}

// returns the neighbors of this patch that are to the top, bottom, left, and right of this patch
func (p *Patch) Neighbors4() *PatchAgentSet {
	neighbors := p.parent.neighbors4(p)

	return neighbors
}

// returns a set of patches that do not include this patch
func (p *Patch) Other(patches *PatchAgentSet) *PatchAgentSet {
	return patches.WhoAreNotPatch(p)
}

// gets the patch relavitve to this patch at the given dx dy
func (p *Patch) PatchAt(dx float64, dy float64) *Patch {

	x := p.xFloat64 + dx
	y := p.yFloat64 + dy

	return p.parent.Patch(x, y)
}

// gets the patch relavitve to this patch at the given heading and distance
func (p *Patch) PatchAtHeadingAndDistance(heading float64, distance float64) *Patch {
	dx := distance * math.Cos(heading)
	dy := distance * math.Sin(heading)

	return p.PatchAt(dx, dy)
}

// returns the x coordinate of this patch
func (p *Patch) PXCor() int {
	return p.x
}

// returns the y coordinate of this patch
func (p *Patch) PYCor() int {
	return p.y
}

// resest the patch to the default values
func (p *Patch) Reset(patchProperties map[string]interface{}) {
	p.PColor.SetColor(Black)

	for key, value := range patchProperties {
		p.patchProperties[key] = value
	}
}

// creates new turtles on this patch
func (p *Patch) Sprout(number int, operation TurtleOperation) {

	generalBreed := p.parent.breeds[BreedNone]

	p.sproutBreeded(generalBreed, number, operation)
}

func (p *Patch) sproutBreeded(breed *TurtleBreed, number int, operation TurtleOperation) {

	turtlesAdded := NewTurtleAgentSet([]*Turtle{})
	for i := 0; i < number; i++ {

		t := newTurtle(p.parent, p.parent.turtlesWhoNumber, breed, p.xFloat64, p.yFloat64)
		p.parent.turtlesWhoNumber++

		//set the heading to be between 0 and 360
		heading := p.parent.randomGenerator.Intn(360)

		//convert to radians
		t.SetHeading(float64(heading))

		turtlesAdded.Add(t)
	}

	turtlesAdded.Ask(operation)
}

// returns the heading that points towards the provided patch
func (p *Patch) TowardsPatch(patch *Patch) float64 {
	//returns heading that points towards the patch
	return p.TowardsXY(patch.xFloat64, patch.yFloat64)
}

// returns the heading that points towards the provided turtle
func (p *Patch) TowardsTurtle(t *Turtle) float64 {
	//returns heading that points towards the turtle
	return p.TowardsXY(t.xcor, t.ycor)
}

// returns the heading that points towards the provided x y coordinates
func (p *Patch) TowardsXY(x float64, y float64) float64 {
	//returns heading that points towards the x y coordinates
	deltaX := x - p.xFloat64
	deltaY := y - p.yFloat64

	return radiansToDegrees(math.Atan2(deltaY, deltaX))
}

// returns the turtles that are on this patch regardless of breed
// if you want to get the turtles of a specific breed, use turtleBreed.TurtlesOnPatch(patch)
// the agentset returned is a pointer to the agentset in the patch, to get a copy of the agentset call .Copy()
func (p *Patch) TurtlesHere() *TurtleAgentSet {

	generalBreed := p.parent.breeds[BreedNone]

	return p.turtlesHereBreeded(generalBreed)
}

func (p *Patch) turtlesHereBreeded(breed *TurtleBreed) *TurtleAgentSet {

	//is the breed valid
	if _, ok := p.turtles[breed]; !ok {
		return NewTurtleAgentSet([]*Turtle{})
	}

	turtles := p.turtles[breed]
	if turtles == nil {
		return NewTurtleAgentSet([]*Turtle{})
	}

	return turtles
}

func (p *Patch) GetProperty(key string) interface{} {
	return p.patchProperties[key]
}

func (t *Patch) GetPropI(key string) int {
	v := t.GetProperty(key)
	if v == nil {
		return 0
	}
	switch v := v.(type) {
	case int:
		return v
	case float64:
		return int(v)
	default:
		return 0
	}
}

func (t *Patch) GetPropF(key string) float64 {
	v := t.GetProperty(key)
	if v == nil {
		return 0
	}
	switch v := v.(type) {
	case int:
		return float64(v)
	case float64:
		return v
	default:
		return 0
	}
}

func (t *Patch) GetPropS(key string) string {
	v := t.GetProperty(key)
	if v == nil {
		return ""
	}
	switch v := v.(type) {
	case string:
		return v
	default:
		return ""
	}
}

func (t *Patch) GetPropB(key string) bool {
	v := t.GetProperty(key)
	if v == nil {
		return false
	}
	switch v := v.(type) {
	case bool:
		return v
	default:
		return false
	}
}

func (p *Patch) SetProperty(key string, value interface{}) {
	// if the value is an int, convert it to a float64
	if _, ok := value.(int); ok {
		value = float64(value.(int))
	}

	if _, ok := p.patchProperties[key]; !ok {
		return
	}
	p.patchProperties[key] = value
}
