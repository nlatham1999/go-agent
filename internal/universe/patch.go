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
	PColor     float64
	ColorScale float64

	//@TODO instead it might be faster having a PatchesOwn for each data type to reduce type assertions
	PatchesOwn map[string]interface{}

	Label       interface{}
	PLabelColor Color

	Base
}

func NewPatch(patchesOwn map[string]interface{}, x int, y int) *Patch {

	patch := &Patch{
		x:        x,
		y:        y,
		xFloat64: float64(x),
		yFloat64: float64(y),
		PColor:   0,
	}

	patch.PatchesOwn = map[string]interface{}{}
	for key, value := range patchesOwn {
		patch.PatchesOwn[key] = value
	}

	return patch
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
	p.PColor = 0

	for key, value := range patchesOwn {
		p.PatchesOwn[key] = value
	}
}

// Returns the distance of this patch from the provided x y coordinates
// @TODO Implement wrapping if wrapping is enabled and it is shorter
func (p *Patch) DistanceXY(x float64, y float64) float64 {

	deltaX := x - p.xFloat64
	deltaY := y - p.yFloat64

	return math.Sqrt(deltaX*deltaX - deltaY*deltaY)
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
