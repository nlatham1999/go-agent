package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/internal/model" // Add the import statement for the package that contains the NewModel function
)

func TestPatchDistanceToTurtle(t *testing.T) {

	//create a basic model
	m := model.NewModel(nil, nil, nil, nil, nil, nil, false) // Use the fully qualified function name

	//get a patch
	patch := m.Patch(-2, -2)

	//create a turtle
	m.CreateTurtles(1, "", nil)
	turtle := m.Turtle("", 0)

	//get the distance between the patch and the turtle
	distance := patch.DistanceTurtle(turtle)

	//make sure the distance is correct
	if distance != 2.8284271247461903 {
		t.Errorf("Expected 2.8284271247461903, got %v", distance)
	}

	//create model that has wrapping
	m = model.NewModel(nil, nil, nil, nil, nil, nil, true)

	//get a patch
	patch = m.Patch(-15, -15)

	//create a turtle
	m.CreateTurtles(1, "", nil)
	turtle = m.Turtle("", 0)

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
		t.Errorf("  4, got %v", distance)
	}
}

func TestPatchDistanceToPatch(t *testing.T) {

	//create a basic model
	m := model.NewModel(nil, nil, nil, nil, nil, nil, false)

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
	m = model.NewModel(nil, nil, nil, nil, nil, nil, true)

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
	m := model.NewModel(nil, nil, nil, nil, nil, nil, false)

	//get a patch
	patch := m.Patch(0, 0)

	//get the distance between the patch and the xy coordinates
	distance := patch.DistanceXY(0, 1)

	//make sure the distance is correct
	if distance != 1 {
		t.Errorf("Expected 1, got %v", distance)
	}

	//create model that has wrapping
	m = model.NewModel(nil, nil, nil, nil, nil, nil, true)

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
	m := model.NewModel(nil, nil, nil, nil, nil, nil, false)

	//get a patch
	patch := m.Patch(0, 0)

	//get the neighbors of the patch
	neighbors := patch.Neighbors()

	//make sure the neighbors are correct
	if neighbors.Count() != 8 {
		t.Errorf("Expected 8 neighbors, got %d", neighbors.Count())
	}

	model.AskPatches(neighbors,
		[]model.PatchOperation{
			func(p *model.Patch) {
				p.PColor.SetColorScale(model.Red)
			},
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
	patchSet := model.PatchSet([]*model.Patch{p1, p2, p3, p4, p5, p6, p7, p8})

	if !patchSet.All(func(p *model.Patch) bool {
		return p.PColor.GetColorScale() == model.Red
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
	m = model.NewModel(nil, nil, nil, nil, nil, nil, true)

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

	model.AskPatches(neighbors,
		[]model.PatchOperation{
			func(p *model.Patch) {
				p.PColor.SetColorScale(model.Red)
			},
		},
	)

	patchSet = model.PatchSet([]*model.Patch{p1, p2, p3, p4, p5, p6, p7, p8})

	if !patchSet.All(func(p *model.Patch) bool {
		return p.PColor.GetColorScale() == model.Red
	}) {
		t.Errorf("Expected all neighbors to be red")
	}
}

func TestPatchNeighbors4(t *testing.T) {

	//create a basic model
	m := model.NewModel(nil, nil, nil, nil, nil, nil, false)

	//get a patch
	patch := m.Patch(0, 0)

	//get the neighbors of the patch
	neighbors := patch.Neighbors4()

	//make sure the neighbors are correct
	if neighbors.Count() != 4 {
		t.Errorf("Expected 4 neighbors, got %d", neighbors.Count())
	}

	model.AskPatches(neighbors,
		[]model.PatchOperation{
			func(p *model.Patch) {
				p.PColor.SetColorScale(model.Red)
			},
		},
	)

	//make sure that the patches around the patch are red
	p1 := m.Patch(1, 0)
	p2 := m.Patch(0, 1)
	p3 := m.Patch(-1, 0)
	p4 := m.Patch(0, -1)
	patchSet := model.PatchSet([]*model.Patch{p1, p2, p3, p4})

	if !patchSet.All(func(p *model.Patch) bool {
		return p.PColor.GetColorScale() == model.Red
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

	// get the bottom left patch and make sure we only get 2 neighbors
	patch = m.Patch(-15, 15)
	neighbors = patch.Neighbors4()

	// create model that has wrapping
	m = model.NewModel(nil, nil, nil, nil, nil, nil, true)

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

	model.AskPatches(neighbors,
		[]model.PatchOperation{
			func(p *model.Patch) {
				p.PColor.SetColorScale(model.Red)
			},
		},
	)

	patchSet = model.PatchSet([]*model.Patch{p1, p2, p3, p4})

	if !patchSet.All(func(p *model.Patch) bool {
		return p.PColor.GetColorScale() == model.Red
	}) {
		t.Errorf("Expected all neighbors to be red")
	}
}
