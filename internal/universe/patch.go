package universe

import "math"

type Patch struct {
	x     int
	y     int
	index int

	parent *Universe

	//we have float54 versions of the variables so that we don't have to do a bunch of conversions
	xFloat64 float64
	yFloat64 float64

	//same as pcolor
	PColor Color

	//@TODO instead it might be faster having a PatchesOwn for each data type to reduce type assertions
	PatchesOwn map[string]interface{}

	Label       interface{}
	PLabelColor Color

	turtles map[string]map[*Turtle]interface{} // sets of turtles keyed by breed
}

func NewPatch(patchesOwn map[string]interface{}, x int, y int) *Patch {

	patch := &Patch{
		x:        x,
		y:        y,
		xFloat64: float64(x),
		yFloat64: float64(y),
		PColor:   Color{},
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
		p.turtles[t.breed] = map[*Turtle]interface{}{}
	}
	p.turtles[t.breed][t] = nil

	// if the breed is provided, add it to the general set of turtles as well
	if t.breed != "" {
		if _, ok := p.turtles[""]; !ok {
			p.turtles[""] = map[*Turtle]interface{}{}
		}
		p.turtles[""][t] = nil
	}
}

// unlinks a turtle from this patch
func (p *Patch) removeTurtle(t *Turtle) {
	if _, ok := p.turtles[t.breed]; ok {
		delete(p.turtles[t.breed], t)
	}

	if t.breed != "" {
		if _, ok := p.turtles[""]; ok {
			delete(p.turtles[""], t)
		}
	}
}

// @TODO implement
func (p *Patch) DistanceTurtle(t *Turtle) float64 {
	return 0
}

// @TODO implement
func (p *Patch) DistancePatch(patch *Patch) float64 {
	return 0
}

func (p *Patch) Reset(patchesOwn map[string]interface{}) {
	p.PColor.SetColorScale(Black)

	for key, value := range patchesOwn {
		p.PatchesOwn[key] = value
	}
}

// Returns the distance of this patch from the provided x y coordinates
// @TODO Implement wrapping if wrapping is enabled and it is shorter
func (p *Patch) DistanceXY(x float64, y float64) float64 {

	deltaX := x - p.xFloat64
	deltaY := y - p.yFloat64

	distance := math.Sqrt(deltaX*deltaX - deltaY*deltaY)

	if !p.parent.wrapping {
		return distance
	}

	deltaXInverse := float64(p.parent.WorldWidth) - math.Abs(deltaX)
	deltaYInverse := float64(p.parent.WorldHeight) - math.Abs(deltaY)

	distance = math.Min(distance, math.Sqrt(deltaX*deltaX+deltaYInverse*deltaYInverse))
	distance = math.Min(distance, math.Sqrt(deltaXInverse*deltaXInverse+deltaY*deltaY))
	distance = math.Min(distance, math.Sqrt(deltaXInverse*deltaXInverse+deltaYInverse*deltaYInverse))

	return distance
}

// @TODO implement
func (p *Patch) InRadiusPatches(radius float64) []*Patch {
	return nil
}

// @TODO implement
func (p *Patch) InRadiusTurtles(radius float64) []*Turtle {
	return nil
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
	other := &PatchAgentSet{}

	for _, patch := range patches.patches {
		if patch != p {
			other.patches = append(other.patches, patch)
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
	return &TurtleAgentSet{
		turtles: []*Turtle{},
	}
}
