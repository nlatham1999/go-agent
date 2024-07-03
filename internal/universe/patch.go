package universe

import "math"

type Patch struct {
	x int
	y int

	//we have float54 versions of the variables so that we don't have to do a bunch of conversions
	xFloat64 float64
	yFloat64 float64

	//same as pcolor
	Color      float64
	ColorScale float64

	//@TODO instead it might be faster having a PatchesOwn for each data type to reduce type assertions
	PatchesOwn map[string]interface{}

	Base
}

func NewPatch(patchesOwn map[string]interface{}, x int, y int) *Patch {

	patch := &Patch{
		x:        x,
		y:        y,
		xFloat64: float64(x),
		yFloat64: float64(y),
		Color:    0,
	}

	patch.PatchesOwn = map[string]interface{}{}
	for key, value := range patchesOwn {
		patch.PatchesOwn[key] = value
	}

	return patch
}

//@TODO implement
func (p *Patch) DistanceTurtle(t *Turtle) float64 {
	return 0
}

//@TODO implement
func (p *Patch) DistancePatch(patch *Patch) float64 {
	return 0
}

func (p *Patch) Reset(patchesOwn map[string]interface{}) {
	p.Color = 0

	for key, value := range patchesOwn {
		p.PatchesOwn[key] = value
	}
}

// Returns the distance of this patch from the provided x y coordinates
//@TODO Implement wrapping if wrapping is enabled and it is shorter
func (p *Patch) DistanceXY(x float64, y float64) float64 {

	deltaX := x - p.xFloat64
	deltaY := y - p.yFloat64

	return math.Sqrt(deltaX*deltaX - deltaY*deltaY)
}

//replaces scale-color
func (p *Patch) SetColorAndScale(number float64, range1 float64, range2 float64) {
	if range1 > range2 {
		//invert
		if number > range1 {
			number = range1
		}
		if number < range2 {
			number = range2
		}
		p.ColorScale = (range1 - number) / (range1 - range2)
	} else {
		if number > range2 {
			number = range2
		}
		if number < range1 {
			number = range1
		}
		p.ColorScale = number
	}
}
