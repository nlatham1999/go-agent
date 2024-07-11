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
