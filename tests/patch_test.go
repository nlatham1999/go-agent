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
		t.Errorf("Expected 4, got %v", distance)
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
