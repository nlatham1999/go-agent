package model

import (
	"github.com/nlatham1999/sortedset"
)

type PatchAgentSet struct {
	patches *sortedset.SortedSet
}

func NewPatchAgentSet(patches []*Patch) *PatchAgentSet {
	patchSet := &PatchAgentSet{
		patches: sortedset.NewSortedSet(),
	}
	for _, patch := range patches {
		patchSet.patches.Add(patch)
	}
	return patchSet
}

func (p *PatchAgentSet) Add(patch *Patch) {
	p.patches.Add(patch)
}

func (p *PatchAgentSet) All(operation PatchBoolOperation) bool {
	patch := p.patches.First()
	for patch != nil {
		if !operation(patch.(*Patch)) {
			return false
		}
		patch, _ = p.patches.Next()
	}
	return true
}

func (p *PatchAgentSet) Any(operation PatchBoolOperation) bool {
	patch := p.patches.First()
	for patch != nil {
		if operation(patch.(*Patch)) {
			return true
		}
		patch, _ = p.patches.Next()
	}
	return false
}

func (p *PatchAgentSet) Ask(operations []PatchOperation) {
	patches := p.patches.List()

	for _, patch := range patches {
		for j := 0; j < len(operations); j++ {
			operations[j](patch.(*Patch))
		}
	}
}

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

func (p *PatchAgentSet) Contains(patch *Patch) bool {
	return p.patches.Contains(patch)
}

func (p *PatchAgentSet) Count() int {
	return p.patches.Len()
}

func (p PatchAgentSet) InRadiusPatch(radius float64, patch *Patch) *PatchAgentSet {
	patchSet := sortedset.NewSortedSet()

	patchIter := p.patches.First()
	for patchIter != nil {
		distance := patchIter.(*Patch).DistancePatch(patch)
		if distance <= radius {
			patchSet.Add(patchIter)
		}
		patchIter, _ = p.patches.Next()
	}

	return &PatchAgentSet{
		patches: patchSet,
	}
}

func (p PatchAgentSet) InRadiusTurtle(radius float64, turtle *Turtle) *PatchAgentSet {
	patchSet := sortedset.NewSortedSet()

	patchIter := p.patches.First()
	for patchIter != nil {
		if patchIter.(*Patch).DistanceTurtle(turtle) <= radius {
			patchSet.Add(patchIter)
		}
		patchIter, _ = p.patches.Next()
	}

	return &PatchAgentSet{
		patches: patchSet,
	}
}

func (p *PatchAgentSet) List() []*Patch {
	v := []*Patch{}
	patch := p.patches.First()
	for patch != nil {
		v = append(v, patch.(*Patch))
		patch, _ = p.patches.Next()
	}
	return v
}

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

func (p *PatchAgentSet) First() (*Patch, error) {
	patch := p.patches.First()
	if patch == nil {
		return nil, ErrNoLinksInAgentSet
	}
	return patch.(*Patch), nil
}

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

func (p *PatchAgentSet) Last(operation PatchFloatOperation) (*Patch, error) {
	patch := p.patches.Last()
	if patch == nil {
		return nil, ErrNoLinksInAgentSet
	}
	return patch.(*Patch), nil
}

func (p *PatchAgentSet) Next() (*Patch, error) {
	patch, err := p.patches.Next()
	if err != nil {
		return nil, ErrNoLinksInAgentSet
	}
	return patch.(*Patch), nil
}

// @TODO make this random
func (p *PatchAgentSet) OneOf() (*Patch, error) {
	for _, patch := range p.patches.List() {
		return patch.(*Patch), nil
	}

	return nil, ErrNoLinksInAgentSet
}

func (p *PatchAgentSet) Remove(patch *Patch) {
	p.patches.Remove(patch)
}

func (p *PatchAgentSet) SortAsc(operation PatchFloatOperation) {
	p.patches.SortAsc(func(a interface{}) interface{} {
		return operation(a.(*Patch))
	})
}

func (p *PatchAgentSet) SortDesc(operation PatchFloatOperation) {
	p.patches.SortDesc(func(a interface{}) interface{} {
		return operation(a.(*Patch))
	})
}

func (p *PatchAgentSet) UpToNOf(n int) *PatchAgentSet {
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

// returns a new PatchAgentSet with all the patches that are not in the given PatchAgentSet
func (p *PatchAgentSet) WhoAreNot(patches *PatchAgentSet) *PatchAgentSet {
	patchSet := sortedset.NewSortedSet()

	for patch := p.patches.First(); patch != nil; patch, _ = p.patches.Next() {
		if !patches.Contains(patch.(*Patch)) {
			patchSet.Add(patch)
		}
	}

	return &PatchAgentSet{
		patches: patchSet,
	}
}

// returns a new PatchAgentSet with all the patches that are not the given patch
func (p *PatchAgentSet) WhoAreNotPatch(patch *Patch) *PatchAgentSet {
	patchSet := sortedset.NewSortedSet()

	for p1 := p.patches.First(); p1 != nil; p1, _ = p.patches.Next() {
		if p1.(*Patch) != patch {
			patchSet.Add(p1)
		}
	}

	return &PatchAgentSet{
		patches: patchSet,
	}
}

func (p *PatchAgentSet) With(operation PatchBoolOperation) *PatchAgentSet {
	patchSet := sortedset.NewSortedSet()
	for patch := p.patches.First(); patch != nil; patch, _ = p.patches.Next() {
		if operation(patch.(*Patch)) {
			patchSet.Add(patch)
		}
	}
	return &PatchAgentSet{
		patches: patchSet,
	}
}
