package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/model"
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

func TestPatchesInRadiusPatch(t *testing.T) {
	//create basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	//get some random patches from the model
	patch1 := m.Patch(-15, -15)
	patch2 := m.Patch(-15, -14)
	patch3 := m.Patch(-14, -15)
	patch4 := m.Patch(-14, -14)
	patch5 := m.Patch(15, 15)
	patch6 := m.Patch(15, 14)
	patch7 := m.Patch(14, 15)
	patch8 := m.Patch(14, 14)

	//create a patchset
	patchSet := model.NewPatchAgentSet([]*model.Patch{patch1, patch2, patch3, patch4, patch5, patch6, patch7, patch8})

	//get the patches in radius
	patchSetInRadius := patchSet.InRadiusPatch(1, patch2)

	if patchSetInRadius.Count() != 3 {
		t.Errorf("Expected 3 patches, got %d", patchSetInRadius.Count())
	}

	// create a model that has wrapping
	settings = model.ModelSettings{
		WrappingX: true,
		WrappingY: true,
	}
	m = model.NewModel(settings)

	// get some random patches from the model
	patch1 = m.Patch(-15, -15)
	patch2 = m.Patch(-15, -14)
	patch3 = m.Patch(-14, -15)
	patch4 = m.Patch(-14, -14)
	patch5 = m.Patch(-15, 15)
	patch6 = m.Patch(-15, 14)
	patch7 = m.Patch(-14, 15)
	patch8 = m.Patch(-14, 14)

	//create a patchset
	patchSet = model.NewPatchAgentSet([]*model.Patch{patch1, patch2, patch3, patch4, patch5, patch6, patch7, patch8})

	//get the patches in radius
	patchSetInRadius = patchSet.InRadiusPatch(1, patch1)

	if patchSetInRadius.Count() != 4 {
		t.Errorf("Expected 4 patches, got %d", patchSetInRadius.Count())
	}
}

func TestPatchesInRadiusTurtle(t *testing.T) {
	//create basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	//get some random patches from the model
	patch1 := m.Patch(-15, -15)
	patch2 := m.Patch(-15, -14)
	patch3 := m.Patch(-14, -15)
	patch4 := m.Patch(-14, -14)
	patch5 := m.Patch(15, 15)
	patch6 := m.Patch(15, 14)
	patch7 := m.Patch(14, 15)
	patch8 := m.Patch(14, 14)

	//create a patchset
	patchSet := model.NewPatchAgentSet([]*model.Patch{patch1, patch2, patch3, patch4, patch5, patch6, patch7, patch8})

	m.CreateTurtles(1, "", []model.TurtleOperation{
		func(t *model.Turtle) {
			t.SetXY(-15, -14)
		},
	})

	turtle := m.Turtle("", 0)

	//get the patches in radius
	patchSetInRadius := patchSet.InRadiusTurtle(1, turtle)

	if patchSetInRadius.Count() != 3 {
		t.Errorf("Expected 3 patches, got %d", patchSetInRadius.Count())
	}

	// create a model that has wrapping
	settings = model.ModelSettings{
		WrappingX: true,
		WrappingY: true,
	}
	m = model.NewModel(settings)

	// get some random patches from the model
	patch1 = m.Patch(-15, -15)
	patch2 = m.Patch(-15, -14)
	patch3 = m.Patch(-14, -15)
	patch4 = m.Patch(-14, -14)
	patch5 = m.Patch(-15, 15)
	patch6 = m.Patch(-15, 14)
	patch7 = m.Patch(-14, 15)
	patch8 = m.Patch(-14, 14)

	//create a patchset
	patchSet = model.NewPatchAgentSet([]*model.Patch{patch1, patch2, patch3, patch4, patch5, patch6, patch7, patch8})

	m.CreateTurtles(1, "", []model.TurtleOperation{
		func(t *model.Turtle) {
			t.SetXY(-15, -15)
		},
	})

	turtle = m.Turtle("", 0)

	//get the patches in radius
	patchSetInRadius = patchSet.InRadiusTurtle(1, turtle)

	if patchSetInRadius.Count() != 4 {
		t.Errorf("Expected 4 patches, got %d", patchSetInRadius.Count())
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
