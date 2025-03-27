package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/pkg/model"
)

func TestAllPatch(t *testing.T) {

	//create a basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	patch1 := m.Patch(0, 0)
	patch2 := m.Patch(0, 1)
	patch3 := m.Patch(0, 2)

	patchSet := model.NewPatchAgentSet([]*model.Patch{patch1, patch2, patch3})

	patch1.PColor.SetColor(model.Blue)
	patch2.PColor.SetColor(model.Blue)
	patch3.PColor.SetColor(model.Blue)

	// assert that patchset has All of shape "circle"
	if !patchSet.All(func(p *model.Patch) bool {
		return p.PColor == model.Blue
	}) {
		t.Errorf("Expected patchset to have all patches with color 'blue'")
	}

	patch2.PColor.SetColor(model.Red)

	if patchSet.All(func(p *model.Patch) bool {
		return p.PColor == model.Blue
	}) {
		t.Errorf("Expected patchset to not have all patches with color 'blue'")
	}

}

func TestAnyPatch(t *testing.T) {

	// create a basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	patch1 := m.Patch(0, 0)
	patch2 := m.Patch(0, 1)
	patch3 := m.Patch(0, 2)

	patchSet := model.NewPatchAgentSet([]*model.Patch{patch1, patch2, patch3})

	patch1.PColor.SetColor(model.Blue)

	// assert that patchset has Any of shape "circle"
	if !patchSet.Any(func(p *model.Patch) bool {
		return p.PColor == model.Blue
	}) {
		t.Errorf("Expected patchset to have a patch with color 'blue'")
	}

	patch1.PColor.SetColor(model.Red)

	if patchSet.Any(func(p *model.Patch) bool {
		return p.PColor == model.Blue
	}) {
		t.Errorf("Expected patchset to not have a patch with color 'blue'")
	}
}

func TestAtPointsPatch(t *testing.T) {
	//create basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	//get some random patches from the model
	patch1 := m.Patch(0, 0)
	patch2 := m.Patch(1, 1)
	patch3 := m.Patch(2, 2)

	//create a patchset
	patchSet := model.NewPatchAgentSet([]*model.Patch{patch1, patch2, patch3})

	//get the patchset at the points
	patchSetAtPoints := patchSet.AtPoints(m, []model.Coordinate{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 5, Y: 5}, {X: 100, Y: 0}})
	if patchSetAtPoints.Count() != 2 {
		t.Errorf("Expected 2 patches, got %d", patchSetAtPoints.Count())
	}
}

func TestPatchesWhoAreNotInPatches(t *testing.T) {

	//create a basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	patch1 := m.Patch(0, 0)
	patch2 := m.Patch(0, 1)
	patch3 := m.Patch(0, 2)

	patchSet := model.NewPatchAgentSet([]*model.Patch{patch1, patch2, patch3})

	patchSet2 := model.NewPatchAgentSet([]*model.Patch{patch1, patch2})

	patchSet3 := patchSet.WhoAreNot(patchSet2)

	if patchSet3.Count() != 1 {
		t.Errorf("Expected patchSet3 to have 1 patch")
	}

	if !patchSet3.Contains(patch3) {
		t.Errorf("Expected patchSet3 to have patch3")
	}
}

func TestPatchesWhoAreNotPatch(t *testing.T) {

	//create a basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	patch1 := m.Patch(0, 0)
	patch2 := m.Patch(0, 1)
	patch3 := m.Patch(0, 2)

	patchSet := model.NewPatchAgentSet([]*model.Patch{patch1, patch2, patch3})

	patchSet2 := patchSet.WhoAreNotPatch(patch1)

	if patchSet2.Count() != 2 {
		t.Errorf("Expected patchSet2 to have 2 patches")
	}
}
