package model

import (
	"math"
)

type Patch struct {
	x     int
	y     int
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

	patch.PColor.SetColorScale(Black)

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

func (p *Patch) Neighbors() []*Patch {
	neighbors := p.parent.Neighbors(p.index)

	return neighbors
}

func (p *Patch) Neighbors4() []*Patch {
	neighbors := p.parent.Neighbors4(p.index)

	return neighbors
}

// @TODO implement
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

func (p *Patch) PatchAt(dx float64, dy float64) *Patch {

	//round the coords
	rx := int(math.Round(p.xFloat64 + dx))
	ry := int(math.Round(p.yFloat64 + dy))

	x := p.x + rx
	y := p.y + ry

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
	p.PColor.SetColorScale(Black)

	for key, value := range patchesOwn {
		p.PatchesOwn[key] = value
	}
}

// @TODO implement
func (p *Patch) Sprout(breed string, number int, operations []TurtleOperation) {

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

// @TODO implement
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