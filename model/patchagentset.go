package model

import (
	"github.com/nlatham1999/sortedset"
)

// PatchAgentSet is an ordered set of patches than can be sorted
// implements github.com/nlatham1999/sortedset
type PatchAgentSet struct {
	patches *sortedset.SortedSet
}

// create a new PatchAgentSet
func NewPatchAgentSet(patches []*Patch) *PatchAgentSet {
	patchSet := &PatchAgentSet{
		patches: sortedset.NewSortedSet(),
	}
	for _, patch := range patches {
		patchSet.patches.Add(patch)
	}
	return patchSet
}

// add a patch to the agent set
func (p *PatchAgentSet) Add(patch *Patch) {
	p.patches.Add(patch)
}

// returns true if all the patches in the agent set satisfy the operation
func (p *PatchAgentSet) All(operation PatchBoolOperation) bool {
	if operation == nil {
		return false
	}

	return p.patches.All(func(a interface{}) bool {
		return operation(a.(*Patch))
	})
}

// returns true if any of the patches in the agent set satisfy the operation
func (p *PatchAgentSet) Any(operation PatchBoolOperation) bool {
	if operation == nil {
		return false
	}

	return p.patches.Any(func(a interface{}) bool {
		return operation(a.(*Patch))
	})
}

// perform the operation for all patches in the agent set
func (p *PatchAgentSet) Ask(operation PatchOperation) {
	if operation == nil {
		return
	}

	p.patches.Ask(func(a interface{}) {
		operation(a.(*Patch))
	})
}

// returns a subset of patches that are at the given coordinates
func (p *PatchAgentSet) AtPoints(m *Model, points []Coordinate) *PatchAgentSet {
	// create a map of the patches
	patchesAtPoints := sortedset.NewSortedSet()
	for _, point := range points {
		patch := m.Patch(point.X, point.Y)
		if patch != nil {
			if p.patches.Contains(patch) {
				patchesAtPoints.Add(patch)
			}
		}
	}

	return &PatchAgentSet{
		patches: patchesAtPoints,
	}

}

// returns true if the patch is in the agent set
func (p *PatchAgentSet) Contains(patch *Patch) bool {
	return p.patches.Contains(patch)
}

// returns the number of patches in the agent set
func (p *PatchAgentSet) Count() int {
	return p.patches.Len()
}

// returns a subset of patches that are in the radius of the given patch
func (p PatchAgentSet) InRadiusPatch(radius float64, patch *Patch) *PatchAgentSet {
	return p.With(func(p *Patch) bool {
		return p.DistancePatch(patch) <= radius
	})
}

// returns a subset of patches that are in the radius of the given turtle
func (p PatchAgentSet) InRadiusTurtle(radius float64, turtle *Turtle) *PatchAgentSet {
	return p.With(func(p *Patch) bool {
		return p.DistanceTurtle(turtle) <= radius
	})
}

// returns the agent set as a list
func (p *PatchAgentSet) List() []*Patch {
	patches := make([]*Patch, 0)
	p.patches.Ask(func(a interface{}) {
		patches = append(patches, a.(*Patch))
	})
	return patches
}

// returns the first n patches in the agent set
func (p *PatchAgentSet) FirstNOf(n int) *PatchAgentSet {
	patchSet := sortedset.NewSortedSet()
	patch := p.patches.First()
	for i := 0; i < n && patch != nil; i++ {
		patchSet.Add(patch)
		patch, _ = p.patches.Next()
	}
	return &PatchAgentSet{
		patches: patchSet,
	}
}

// returns the first patch in the agent set
func (p *PatchAgentSet) First() (*Patch, error) {
	patch := p.patches.First()
	if patch == nil {
		return nil, ErrNoLinksInAgentSet
	}
	return patch.(*Patch), nil
}

// returns the last n patches in the agent set
func (p *PatchAgentSet) LastNOf(n int) *PatchAgentSet {
	patchSet := sortedset.NewSortedSet()
	patch := p.patches.Last()
	for i := 0; i < n && patch != nil; i++ {
		patchSet.Add(patch)
		patch, _ = p.patches.Previous()
	}
	return &PatchAgentSet{
		patches: patchSet,
	}
}

// returns the last patch in the agent set
func (p *PatchAgentSet) Last() (*Patch, error) {
	patch := p.patches.Last()
	if patch == nil {
		return nil, ErrNoLinksInAgentSet
	}
	return patch.(*Patch), nil
}

// func (p *PatchAgentSet) Next() (*Patch, error) {
// 	patch, err := p.patches.Next()
// 	if err != nil {
// 		return nil, ErrNoLinksInAgentSet
// 	}
// 	return patch.(*Patch), nil
// }

// remove a patch from the agent set
func (p *PatchAgentSet) Remove(patch *Patch) {
	p.patches.Remove(patch)
}

// sort the patches in the agent set in ascending order based on the float operation
func (p *PatchAgentSet) SortAsc(operation PatchFloatOperation) {
	p.patches.SortAsc(func(a interface{}) interface{} {
		return operation(a.(*Patch))
	})
}

// sort the patches in the agent set in descending order based on the float operation
func (p *PatchAgentSet) SortDesc(operation PatchFloatOperation) {
	p.patches.SortDesc(func(a interface{}) interface{} {
		return operation(a.(*Patch))
	})
}

// returns a new PatchAgentSet with all the patches that are not in the given PatchAgentSet
func (p *PatchAgentSet) WhoAreNot(patches *PatchAgentSet) *PatchAgentSet {
	return &PatchAgentSet{
		patches: p.patches.Difference(patches.patches),
	}
}

// returns a new PatchAgentSet with all the patches that are not the given patch
func (p *PatchAgentSet) WhoAreNotPatch(patch *Patch) *PatchAgentSet {
	return &PatchAgentSet{
		patches: p.patches.Difference(sortedset.NewSortedSet(patch)),
	}
}

// returns a new agent set that contains all the patches that satisfy the operation
func (p *PatchAgentSet) With(operation PatchBoolOperation) *PatchAgentSet {
	patchSet := sortedset.NewSortedSet()

	if operation == nil {
		return nil
	}

	p.patches.Ask(func(a interface{}) {
		if operation(a.(*Patch)) {
			patchSet.Add(a)
		}
	})

	return &PatchAgentSet{
		patches: patchSet,
	}
}
