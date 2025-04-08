package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/pkg/model" // Add the import statement for the package that contains the NewModel function
)

func TestPatchDistanceToTurtle(t *testing.T) {

	//create a basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings) // Use the fully qualified function name

	//get a patch
	patch := m.Patch(-2, -2)

	//create a turtle
	m.CreateTurtles(1, nil)
	turtle := m.Turtle(0)

	//get the distance between the patch and the turtle
	distance := patch.DistanceTurtle(turtle)

	//make sure the distance is correct
	if distance != 2.8284271247461903 {
		t.Errorf("Expected 2.8284271247461903, got %v", distance)
	}

	//create model that has wrapping
	settings = model.ModelSettings{
		WrappingX: true,
		WrappingY: true,
	}
	m = model.NewModel(settings)

	//get a patch
	patch = m.Patch(-15, -15)

	//create a turtle
	m.CreateTurtles(1, nil)
	turtle = m.Turtle(0)

	//move the turtle to the other side of the world
	turtle.SetXY(14, 14)

	//get the distance between the patch and the turtle
	distance = patch.DistanceTurtle(turtle)

	//make sure the distance is correct
	if distance != 2.8284271247461903 {
		t.Errorf("Expected 2.8284271247461903, got %v", distance)
	}

	turtle.SetXY(0, 12)

	patch = m.Patch(0, -15)

	distance = patch.DistanceTurtle(turtle)

	if distance != 4 {
		t.Errorf("Expected 4, got %v", distance)
	}
}

func TestPatchDistanceToPatch(t *testing.T) {

	//create a basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	//get a patch
	patch1 := m.Patch(0, 0)
	patch2 := m.Patch(0, 1)

	//get the distance between the patches
	distance := patch1.DistancePatch(patch2)

	//make sure the distance is correct
	if distance != 1 {
		t.Errorf("Expected 1, got %v", distance)
	}

	//create model that has wrapping
	settings = model.ModelSettings{
		WrappingX: true,
		WrappingY: true,
	}
	m = model.NewModel(settings)

	//get a patch
	patch1 = m.Patch(-15, -15)
	patch2 = m.Patch(14, 14)

	//get the distance between the patches
	distance = patch1.DistancePatch(patch2)

	//make sure the distance is correct
	if distance != 2.8284271247461903 {
		t.Errorf("Expected 2.8284271247461903, got %v", distance)
	}

	patch1 = m.Patch(0, -15)
	patch2 = m.Patch(0, 12)

	distance = patch1.DistancePatch(patch2)

	if distance != 4 {
		t.Errorf("Expected 4, got %v", distance)
	}
}

func TestPatchDistanceToXY(t *testing.T) {

	//create a basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	//get a patch
	patch := m.Patch(0, 0)

	//get the distance between the patch and the xy coordinates
	distance := patch.DistanceXY(0, 1)

	//make sure the distance is correct
	if distance != 1 {
		t.Errorf("Expected 1, got %v", distance)
	}

	//create model that has wrapping
	settings = model.ModelSettings{
		WrappingX: true,
		WrappingY: true,
	}
	m = model.NewModel(settings)

	//get a patch
	patch = m.Patch(-15, -15)

	//get the distance between the patch and the xy coordinates
	distance = patch.DistanceXY(14, 14)

	//make sure the distance is correct
	if distance != 2.8284271247461903 {
		t.Errorf("Expected 2.8284271247461903, got %v", distance)
	}

	patch = m.Patch(0, -15)

	distance = patch.DistanceXY(0, 12)

	if distance != 4 {
		t.Errorf("Expected 4, got %v", distance)
	}
}

func TestPatchNeighbors(t *testing.T) {

	//create a basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	//get a patch
	patch := m.Patch(0, 0)

	//get the neighbors of the patch
	neighbors := patch.Neighbors()

	//make sure the neighbors are correct
	if neighbors.Count() != 8 {
		t.Errorf("Expected 8 neighbors, got %d", neighbors.Count())
	}

	neighbors.Ask(
		func(p *model.Patch) {
			p.Color.SetColor(model.Red)
		},
	)

	//make sure that the patches around the patch are red
	p1 := m.Patch(1, 0)
	p2 := m.Patch(1, 1)
	p3 := m.Patch(0, 1)
	p4 := m.Patch(-1, 1)
	p5 := m.Patch(-1, 0)
	p6 := m.Patch(-1, -1)
	p7 := m.Patch(0, -1)
	p8 := m.Patch(1, -1)
	patchSet := model.NewPatchAgentSet([]*model.Patch{p1, p2, p3, p4, p5, p6, p7, p8})

	if !patchSet.All(func(p *model.Patch) bool {
		return p.Color == model.Red
	}) {
		t.Errorf("Expected all neighbors to be red")
	}

	// get the top left patch and make sure we only get 3 neighbors
	patch = m.Patch(-15, -15)
	neighbors = patch.Neighbors()

	if neighbors.Count() != 3 {
		t.Errorf("Expected 3 neighbors, got %d", neighbors.Count())
	}

	// get the bottom right patch and make sure we only get 3 neighbors
	patch = m.Patch(15, 15)
	neighbors = patch.Neighbors()

	if neighbors.Count() != 3 {
		t.Errorf("Expected 3 neighbors, got %d", neighbors.Count())
	}

	// get the top right patch and make sure we only get 3 neighbors
	patch = m.Patch(15, -15)
	neighbors = patch.Neighbors()

	if neighbors.Count() != 3 {
		t.Errorf("Expected 3 neighbors, got %d", neighbors.Count())
	}

	// get the bottom left patch and make sure we only get 3 neighbors
	patch = m.Patch(-15, 15)
	neighbors = patch.Neighbors()

	if neighbors.Count() != 3 {
		t.Errorf("Expected 3 neighbors, got %d", neighbors.Count())
	}

	//create model that has wrapping
	settings = model.ModelSettings{
		WrappingX: true,
		WrappingY: true,
	}
	m = model.NewModel(settings)

	//get the top left patch and make sure we get 8 neighbors
	patch = m.Patch(-15, -15)
	neighbors = patch.Neighbors()

	if neighbors.Count() != 8 {
		t.Errorf("Expected 8 neighbors, got %d", neighbors.Count())
	}

	p1 = m.Patch(-14, -15)
	p2 = m.Patch(-14, -14)
	p3 = m.Patch(-15, -14)
	p4 = m.Patch(-15, 15)
	p5 = m.Patch(-14, 15)
	p6 = m.Patch(15, -15)
	p7 = m.Patch(15, -14)
	p8 = m.Patch(15, 15)

	neighbors.Ask(
		func(p *model.Patch) {
			p.Color.SetColor(model.Red)
		},
	)

	patchSet = model.NewPatchAgentSet([]*model.Patch{p1, p2, p3, p4, p5, p6, p7, p8})

	if !patchSet.All(func(p *model.Patch) bool {
		return p.Color == model.Red
	}) {
		t.Errorf("Expected all neighbors to be red")
	}
}

func TestPatchNeighbors4(t *testing.T) {

	//create a basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	//get a patch
	patch := m.Patch(0, 0)

	//get the neighbors of the patch
	neighbors := patch.Neighbors4()

	//make sure the neighbors are correct
	if neighbors.Count() != 4 {
		t.Errorf("Expected 4 neighbors, got %d", neighbors.Count())
	}

	neighbors.Ask(
		func(p *model.Patch) {
			p.Color.SetColor(model.Red)
		},
	)

	//make sure that the patches around the patch are red
	p1 := m.Patch(1, 0)
	p2 := m.Patch(0, 1)
	p3 := m.Patch(-1, 0)
	p4 := m.Patch(0, -1)
	patchSet := model.NewPatchAgentSet([]*model.Patch{p1, p2, p3, p4})

	if !patchSet.All(func(p *model.Patch) bool {
		return p.Color == model.Red
	}) {
		t.Errorf("Expected all neighbors to be red")
	}

	// get the top left patch and make sure we only get 2 neighbors
	patch = m.Patch(-15, -15)
	neighbors = patch.Neighbors4()

	if neighbors.Count() != 2 {
		t.Errorf("Expected 2 neighbors, got %d", neighbors.Count())
	}

	// get the bottom right patch and make sure we only get 2 neighbors
	patch = m.Patch(15, 15)
	neighbors = patch.Neighbors4()

	if neighbors.Count() != 2 {
		t.Errorf("Expected 2 neighbors, got %d", neighbors.Count())
	}

	// get the top right patch and make sure we only get 2 neighbors
	patch = m.Patch(15, -15)
	neighbors = patch.Neighbors4()

	if neighbors.Count() != 2 {
		t.Errorf("Expected 2 neighbors, got %d", neighbors.Count())
	}

	// create model that has wrapping
	settings = model.ModelSettings{
		WrappingX: true,
		WrappingY: true,
	}
	m = model.NewModel(settings)

	// get the top left patch and make sure we get 4 neighbors
	patch = m.Patch(-15, -15)
	neighbors = patch.Neighbors4()

	if neighbors.Count() != 4 {
		t.Errorf("Expected 4 neighbors, got %d", neighbors.Count())
	}

	p1 = m.Patch(-14, -15)
	p2 = m.Patch(-15, -14)
	p3 = m.Patch(-15, 15)
	p4 = m.Patch(15, -15)

	neighbors.Ask(
		func(p *model.Patch) {
			p.Color.SetColor(model.Red)
		},
	)

	patchSet = model.NewPatchAgentSet([]*model.Patch{p1, p2, p3, p4})

	if !patchSet.All(func(p *model.Patch) bool {
		return p.Color == model.Red
	}) {
		t.Errorf("Expected all neighbors to be red")
	}
}

func TestPatchPatchAt(t *testing.T) {

	//create a basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	//get a patch
	patch := m.Patch(0, 0)

	//get the patch at the given dx dy
	patchAt := patch.PatchAt(1, 1)

	//make sure the patch is correct
	if patchAt.PXCor() != 1 || patchAt.PYCor() != 1 {
		t.Errorf("Expected patch at 1, 1, got %v, %v", patchAt.PXCor(), patchAt.PYCor())
	}

	// get patch outside of world
	patchAt = patch.PatchAt(16, 16)

	if patchAt != nil {
		t.Errorf("Expected patch to be nil")
	}

	//create model that has wrapping
	settings = model.ModelSettings{
		WrappingX: true,
		WrappingY: true,
	}
	m = model.NewModel(settings)

	//get a patch
	patch = m.Patch(-15, -15)

	//get the patch at the given dx dy
	patchAt = patch.PatchAt(-1, -1)

	//make sure the patch is correct
	if patchAt.PXCor() != 15 || patchAt.PYCor() != 15 {
		t.Errorf("Expected patch at 15, 15, got %v, %v", patchAt.PXCor(), patchAt.PYCor())
	}

	// get patch outside of world
	patchAt = patch.PatchAt(16, 16)

	if patchAt.PXCor() != 1 || patchAt.PYCor() != 1 {
		t.Errorf("Expected patch at 1, 1, got %v, %v", patchAt.PXCor(), patchAt.PYCor())
	}

	patchAt = patch.PatchAt(-32, -32)

	if patchAt.PXCor() != 15 || patchAt.PYCor() != 15 {
		t.Errorf("Expected patch at 15, 15, got %v, %v", patchAt.PXCor(), patchAt.PYCor())
	}
}

func TestPatchSprout(t *testing.T) {

	//create a basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	//get a patch
	patch := m.Patch(0, 0)

	//sprout a patch
	patch.Sprout(5,
		func(t *model.Turtle) {
			switch t.Who() {
			case 0:
				t.Color.SetColor(model.Red)
			case 1:
				t.Color.SetColor(model.Blue)
			case 2:
				t.Color.SetColor(model.Green)
			case 3:
				t.Color.SetColor(model.Yellow)
			case 4:
				t.Color.SetColor(model.Pink)
			}
		},
	)

	if m.Turtle(0).Color != model.Red {
		t.Errorf("Expected red turtle")
	}

	if m.Turtle(1).Color != model.Blue {
		t.Errorf("Expected blue turtle")
	}

	if m.Turtle(2).Color != model.Green {
		t.Errorf("Expected green turtle")
	}

	if m.Turtle(3).Color != model.Yellow {
		t.Errorf("Expected yellow turtle")
	}

	if m.Turtle(4).Color != model.Pink {
		t.Errorf("Expected pink turtle")
	}
}

func TestPatchTowardsXY(t *testing.T) {

	// create a basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	// get a patch
	patch := m.Patch(0, 0)

	// get the heading towards the xy coordinates
	heading := patch.TowardsXY(1, -1)

	// make sure the heading is correct
	if heading != 315 {
		t.Errorf("Expected 315, got %v", heading)
	}
}
