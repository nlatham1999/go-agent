package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/internal/universe"
)

func TestAllPatch(t *testing.T) {
	patch1 := universe.NewPatch(nil, 0, 0)
	patch2 := universe.NewPatch(nil, 0, 0)
	patch3 := universe.NewPatch(nil, 0, 0)

	patchSet := universe.PatchSet([]*universe.Patch{patch1, patch2, patch3})

	patch1.PColor.SetColorScale(universe.Blue)
	patch2.PColor.SetColorScale(universe.Blue)
	patch3.PColor.SetColorScale(universe.Blue)

	// assert that patchset has All of shape "circle"
	if !patchSet.All(func(p *universe.Patch) bool {
		return p.PColor.GetColorScale() == universe.Blue
	}) {
		t.Errorf("Expected patchset to have all patches with color 'blue'")
	}

	patch2.PColor.SetColorScale(universe.Red)

	if patchSet.All(func(p *universe.Patch) bool {
		return p.PColor.GetColorScale() == universe.Blue
	}) {
		t.Errorf("Expected patchset to not have all patches with color 'blue'")
	}

}

func TestAnyPatch(t *testing.T) {

	patch1 := universe.NewPatch(nil, 0, 0)
	patch2 := universe.NewPatch(nil, 0, 0)
	patch3 := universe.NewPatch(nil, 0, 0)

	patchSet := universe.PatchSet([]*universe.Patch{patch1, patch2, patch3})

	patch1.PColor.SetColorScale(universe.Blue)

	// assert that patchset has Any of shape "circle"
	if !patchSet.Any(func(p *universe.Patch) bool {
		return p.PColor.GetColorScale() == universe.Blue
	}) {
		t.Errorf("Expected patchset to have a patch with color 'blue'")
	}

	patch1.PColor.SetColorScale(universe.Red)

	if patchSet.Any(func(p *universe.Patch) bool {
		return p.PColor.GetColorScale() == universe.Blue
	}) {
		t.Errorf("Expected patchset to not have a patch with color 'blue'")
	}
}

func TestAtPointsPatch(t *testing.T) {
	//create basic universe
	u := universe.NewUniverse(nil, nil, nil, nil, nil, nil, false)

	//get some random patches from the universe
	patch1 := u.Patch(0, 0)
	patch2 := u.Patch(1, 1)
	patch3 := u.Patch(2, 2)

	//create a patchset
	patchSet := universe.PatchSet([]*universe.Patch{patch1, patch2, patch3})

	//get the patchset at the points
	patchSetAtPoints := patchSet.AtPoints(u, []universe.Coordinate{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 5, Y: 5}, {X: 100, Y: 0}})
	if patchSetAtPoints.Count() != 2 {
		t.Errorf("Expected 2 patches, got %d", patchSetAtPoints.Count())
	}

}
