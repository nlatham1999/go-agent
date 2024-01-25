package patch

import "math"

type Patch struct {
	x          int
	y          int
	PatchesOwn map[string]interface{}
}

func NewPatch(patchesOwn map[string]interface{}, x int, y int) *Patch {

	patch := &Patch{
		x: x,
		y: y,
	}

	patch.PatchesOwn = map[string]interface{}{}
	for key, value := range patchesOwn {
		patch.PatchesOwn[key] = value
	}

	return patch
}

// Returns the distance of this patch from the provided x y coordinates
//@TODO Implement wrapping if wrapping is enabled and it is shorter
func (p *Patch) DistanceXY(x int, y int) float64 {
	return math.Sqrt(float64((x - p.x) ^ 2 + (y - p.y) ^ 2))
}
