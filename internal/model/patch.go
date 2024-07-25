package model

import (
	"math"
	"math/rand"
)

type Patch struct {
	x int
	y int

	// this corresponds to the position in the patches array
	// set as x*m.WorldWidth + y
	// maps to parent.posOfPatches[index]
	index int

	parent *Model

	//we have float54 versions of the variables so that we don't have to do a bunch of conversions
	xFloat64 float64
	yFloat64 float64

	//same as pcolor
	PColor Color

	//@TODO instead it might be faster having a PatchesOwn for each data type to reduce type assertions
	PatchesOwn map[string]interface{}

	Label       interface{}
	PLabelColor Color

	turtles map[string]*TurtleAgentSet // sets of turtles keyed by breed
}

func NewPatch(m *Model, patchesOwn map[string]interface{}, x int, y int) *Patch {

	patch := &Patch{
		x:        x,
		y:        y,
		xFloat64: float64(x),
		yFloat64: float64(y),
		PColor:   Color{},
		turtles:  make(map[string]*TurtleAgentSet),
		parent:   m,
	}

	patch.PColor.SetColor(Black)

	patch.PatchesOwn = map[string]interface{}{}
	for key, value := range patchesOwn {
		patch.PatchesOwn[key] = value
	}

	return patch
}

// links a turtle to this patch
func (p *Patch) addTurtle(t *Turtle) {
	if _, ok := p.turtles[t.breed]; !ok {
		p.turtles[t.breed] = TurtleSet([]*Turtle{})
	}
	p.turtles[t.breed].Add(t)

	// if the breed is provided, add it to the general set of turtles as well
	if t.breed != "" {
		if _, ok := p.turtles[""]; !ok {
			p.turtles[""] = TurtleSet([]*Turtle{})
		}
		p.turtles[""].Add(t)
	}
}

// unlinks a turtle from this patch
func (p *Patch) removeTurtle(t *Turtle) {
	if _, ok := p.turtles[t.breed]; ok {
		p.turtles[t.breed].Remove(t)
	}

	if t.breed != "" {
		if _, ok := p.turtles[""]; ok {
			p.turtles[""].Remove(t)
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

func (p *Patch) Neighbors() *PatchAgentSet {
	neighbors := p.parent.neighbors(p)

	return neighbors
}

func (p *Patch) Neighbors4() *PatchAgentSet {
	neighbors := p.parent.neighbors4(p)

	return neighbors
}

func (p *Patch) Other(patches *PatchAgentSet) *PatchAgentSet {
	other := &PatchAgentSet{
		patches: make(map[*Patch]interface{}),
	}

	for patch := range patches.patches {
		if patch != p {
			other.patches[patch] = nil
		}
	}

	return other
}

// gets the patch relavitve to this patch at the given dx dy
func (p *Patch) PatchAt(dx float64, dy float64) *Patch {

	//round the coords
	x := int(math.Round(p.xFloat64 + dx))
	y := int(math.Round(p.yFloat64 + dy))

	if x < p.parent.MinPxCor {
		if p.parent.wrappingX {
			x = p.parent.MaxPxCor + 1 + ((x - p.parent.MinPxCor) % p.parent.WorldWidth)
		} else {
			return nil
		}
	}

	if y < p.parent.MinPyCor {
		if p.parent.wrappingY {
			y = p.parent.MaxPyCor + 1 + ((y - p.parent.MinPyCor) % p.parent.WorldHeight)
		} else {
			return nil
		}
	}

	if x > p.parent.MaxPxCor {
		if p.parent.wrappingX {
			x = (x-p.parent.MaxPxCor)%p.parent.WorldWidth + p.parent.MinPxCor - 1
		} else {
			return nil
		}
	}

	if y > p.parent.MaxPyCor {
		if p.parent.wrappingY {
			y = (y-p.parent.MaxPyCor)%p.parent.WorldHeight + p.parent.MinPyCor - 1
		} else {
			return nil
		}
	}

	return p.parent.getPatchAtCoords(x, y)
}

func (p *Patch) PatchAtHeadingAndDistance(heading float64, distance float64) *Patch {
	dx := distance * math.Cos(heading)
	dy := distance * math.Sin(heading)

	return p.PatchAt(dx, dy)
}

func (p *Patch) PXCor() int {
	return p.x
}

func (p *Patch) PYCor() int {
	return p.y
}

func (p *Patch) Reset(patchesOwn map[string]interface{}) {
	p.PColor.SetColor(Black)

	for key, value := range patchesOwn {
		p.PatchesOwn[key] = value
	}
}

func (p *Patch) Sprout(breed string, number int, operations []TurtleOperation) {

	turtlesAdded := TurtleSet([]*Turtle{})
	for i := 0; i < number; i++ {

		t := NewTurtle(p.parent, p.parent.turtlesWhoNumber, breed, p.xFloat64, p.yFloat64)
		p.parent.turtlesWhoNumber++

		//set the heading to be between 0 and 360
		heading := rand.Intn(360)

		//convert to radians
		t.SetHeading(float64(heading))

		turtlesAdded.Add(t)
	}

	AskTurtles(turtlesAdded, operations)
}

func (p *Patch) TowardsPatch(patch *Patch) float64 {
	//returns heading that points towards the patch
	deltaX := patch.xFloat64 - p.xFloat64
	deltaY := patch.yFloat64 - p.yFloat64

	return math.Atan2(deltaY, deltaX)
}

func (p *Patch) TowardsTurtle(t *Turtle) float64 {
	//returns heading that points towards the turtle
	deltaX := t.xcor - p.xFloat64
	deltaY := t.ycor - p.yFloat64

	return math.Atan2(deltaY, deltaX)
}

func (p *Patch) TowardsXY(x float64, y float64) float64 {
	//returns heading that points towards the x y coordinates
	deltaX := x - p.xFloat64
	deltaY := y - p.yFloat64

	return math.Atan2(deltaY, deltaX)
}

func (p *Patch) TurtlesHere(breed string) *TurtleAgentSet {

	//is the breed valid
	if breed != "" {
		if _, ok := p.turtles[breed]; !ok {
			return TurtleSet([]*Turtle{})
		}
	}

	turtles := p.turtles[breed]
	if turtles == nil {
		return TurtleSet([]*Turtle{})
	}

	return &TurtleAgentSet{}
}
