package model

import "math"

type PatchAgentSet struct {
	patches map[*Patch]interface{}
}

func PatchSet(patches []*Patch) *PatchAgentSet {
	newPatches := make(map[*Patch]interface{})
	for _, patch := range patches {
		newPatches[patch] = nil
	}

	return &PatchAgentSet{
		patches: newPatches,
	}
}

func (p *PatchAgentSet) All(operation PatchBoolOperation) bool {
	for patch := range p.patches {
		if !operation(patch) {
			return false
		}
	}
	return true
}

func (p *PatchAgentSet) Any(operation PatchBoolOperation) bool {
	for patch := range p.patches {
		if operation(patch) {
			return true
		}
	}
	return false
}

func (p *PatchAgentSet) AtPoints(m *Model, points []Coordinate) *PatchAgentSet {
	// create a map of the patches
	pointsMap := make(map[*Patch]interface{})
	for _, point := range points {
		patch := m.Patch(point.X, point.Y)
		if patch != nil {
			if _, ok := p.patches[patch]; ok {
				pointsMap[patch] = nil
			}
		}
	}

	return &PatchAgentSet{
		patches: pointsMap,
	}
}

func (p *PatchAgentSet) Contains(patch *Patch) bool {
	_, ok := p.patches[patch]
	return ok
}

func (p *PatchAgentSet) Count() int {
	return len(p.patches)
}

func (p PatchAgentSet) InRadiusPatch(radius float64, patch *Patch) *PatchAgentSet {
	patchMap := make(map[*Patch]interface{})

	for p := range p.patches {
		distance := p.DistancePatch(patch)
		if distance <= radius {
			patchMap[p] = nil
		}
	}

	return &PatchAgentSet{
		patches: patchMap,
	}
}

func (p PatchAgentSet) InRadiusTurtle(radius float64, turtle *Turtle) *PatchAgentSet {
	patchMap := make(map[*Patch]interface{})

	for p := range p.patches {
		if p.DistanceTurtle(turtle) <= radius {
			patchMap[p] = nil
		}
	}

	return &PatchAgentSet{
		patches: patchMap,
	}
}

func (p *PatchAgentSet) MaxNOf(n int, operation PatchFloatOperation) *PatchAgentSet {
	return nil
}

func (p *PatchAgentSet) MaxOneOf(operation PatchFloatOperation) *Patch {
	max := math.MaxFloat64 * -1
	var maxPatch *Patch
	for patch := range p.patches {
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
	for patch := range p.patches {
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
	for patch := range p.patches {
		if operation(patch) {
			patches = append(patches, patch)
		}
	}
	return PatchSet(patches)
}

func (p *PatchAgentSet) WithMax(operation PatchFloatOperation) *PatchAgentSet {
	max := math.MaxFloat64 * -1
	for patch := range p.patches {
		if operation(patch) > max {
			max = operation(patch)
		}
	}

	//get all patches where the float operation is equal to the max
	patches := make([]*Patch, 0)
	for patch := range p.patches {
		if operation(patch) == max {
			patches = append(patches, patch)
		}
	}

	return PatchSet(patches)
}

func (p *PatchAgentSet) WithMin(operation PatchFloatOperation) *PatchAgentSet {
	min := math.MaxFloat64
	for patch := range p.patches {
		if operation(patch) < min {
			min = operation(patch)
		}
	}

	//get all patches where the float operation is equal to the min
	patches := make([]*Patch, 0)
	for patch := range p.patches {
		if operation(patch) == min {
			patches = append(patches, patch)
		}
	}

	return PatchSet(patches)
}
