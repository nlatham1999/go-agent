package universe

import "math"

type PatchAgentSet struct {
	patches []*Patch
}

func PatchSet(patches []*Patch) *PatchAgentSet {
	newPatches := make([]*Patch, len(patches))
	copy(newPatches, patches)

	return &PatchAgentSet{
		patches: newPatches,
	}
}

func (p *PatchAgentSet) All(operation PatchBoolOperation) bool {
	for _, patch := range p.patches {
		if !operation(patch) {
			return false
		}
	}
	return true
}

func (p *PatchAgentSet) Any(operation PatchBoolOperation) bool {
	for _, patch := range p.patches {
		if operation(patch) {
			return true
		}
	}
	return false
}

func (p *PatchAgentSet) AtPoints(u *Universe, points []Coordinate) *PatchAgentSet {
	// create a map of the patches
	pointsMap := make(map[*Patch]interface{})
	for _, point := range points {
		patch := u.Patch(point.X, point.Y)
		if patch != nil {
			pointsMap[patch] = nil
		}
	}

	// get the patches that are in the map
	patches := make([]*Patch, 0)
	for _, patch := range p.patches {
		if _, ok := pointsMap[patch]; ok {
			patches = append(patches, patch)
		}
	}

	return PatchSet(patches)
}

func (p *PatchAgentSet) Count() int {
	return len(p.patches)
}

func (p *PatchAgentSet) MaxNOf(n int, operation PatchFloatOperation) *PatchAgentSet {
	return nil
}

func (p *PatchAgentSet) MaxOneOf(operation PatchFloatOperation) *Patch {
	max := math.MaxFloat64 * -1
	var maxPatch *Patch
	for _, patch := range p.patches {
		if operation(patch) > max {
			max = operation(patch)
			maxPatch = patch
		}
	}
	return maxPatch
}

func (p *PatchAgentSet) MinNOf(n int, operation PatchFloatOperation) *PatchAgentSet {
	return nil
}

func (p *PatchAgentSet) MinOneOf(operation PatchFloatOperation) *Patch {
	min := math.MaxFloat64
	var minPatch *Patch
	for _, patch := range p.patches {
		if operation(patch) < min {
			min = operation(patch)
			minPatch = patch
		}
	}
	return minPatch
}

// @TODO implement
func (p *PatchAgentSet) OneOf() *Patch {
	return nil
}

func (p *PatchAgentSet) UpToNOf(n int) *PatchAgentSet {
	return nil
}

// @TODO implement
func (p *PatchAgentSet) WhoAreNot(patches *PatchAgentSet) *PatchAgentSet {
	return nil
}

// @TODO implement
func (p *PatchAgentSet) WhoAreNotPatch(patch *Patch) *PatchAgentSet {
	return nil
}

func (p *PatchAgentSet) With(operation PatchBoolOperation) *PatchAgentSet {
	patches := make([]*Patch, 0)
	for _, patch := range p.patches {
		if operation(patch) {
			patches = append(patches, patch)
		}
	}
	return PatchSet(patches)
}

func (p *PatchAgentSet) WithMax(operation PatchFloatOperation) *PatchAgentSet {
	max := math.MaxFloat64 * -1
	for _, patch := range p.patches {
		if operation(patch) > max {
			max = operation(patch)
		}
	}

	//get all patches where the float operation is equal to the max
	patches := make([]*Patch, 0)
	for _, patch := range p.patches {
		if operation(patch) == max {
			patches = append(patches, patch)
		}
	}

	return PatchSet(patches)
}

func (p *PatchAgentSet) WithMin(operation PatchFloatOperation) *PatchAgentSet {
	min := math.MaxFloat64
	for _, patch := range p.patches {
		if operation(patch) < min {
			min = operation(patch)
		}
	}

	//get all patches where the float operation is equal to the min
	patches := make([]*Patch, 0)
	for _, patch := range p.patches {
		if operation(patch) == min {
			patches = append(patches, patch)
		}
	}

	return PatchSet(patches)
}
