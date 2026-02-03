package tests

import (
	"math"
	"testing"

	"github.com/nlatham1999/go-agent/pkg/model"
)

// Test TurtlesInRadiusXYZ finds turtles in 3D space
func TestTurtlesInRadiusXYZ(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor: -5,
		MaxPxCor: 5,
		MinPyCor: -5,
		MaxPyCor: 5,
		MinPzCor: -5,
		MaxPzCor: 5,
	}
	m := model.NewModel(settings)

	// Create turtles at specific 3D positions
	m.CreateTurtles(1, func(turtle *model.Turtle) {
		turtle.SetXYZ(0, 0, 0) // Center
	})
	m.CreateTurtles(1, func(turtle *model.Turtle) {
		turtle.SetXYZ(1, 0, 0) // 1 unit away in X
	})
	m.CreateTurtles(1, func(turtle *model.Turtle) {
		turtle.SetXYZ(0, 1, 0) // 1 unit away in Y
	})
	m.CreateTurtles(1, func(turtle *model.Turtle) {
		turtle.SetXYZ(0, 0, 1) // 1 unit away in Z
	})
	m.CreateTurtles(1, func(turtle *model.Turtle) {
		turtle.SetXYZ(3, 3, 3) // Far away
	})

	// Find turtles within radius 1.5 of origin
	nearbyTurtles := m.TurtlesInRadiusXYZ(0, 0, 0, 1.5)

	// Should find center turtle and 3 neighbors (not the far one)
	if nearbyTurtles.Count() != 4 {
		t.Errorf("Expected 4 turtles within radius, got %d", nearbyTurtles.Count())
	}

	// Verify the far turtle is not included
	farTurtle := m.Turtle(4)
	if nearbyTurtles.Contains(farTurtle) {
		t.Errorf("Expected far turtle to not be in radius")
	}
}

// Test TurtlesInRadiusXYZ vs TurtlesInRadiusXY
func TestTurtlesInRadiusXYZ_vs_XY(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor: -5,
		MaxPxCor: 5,
		MinPyCor: -5,
		MaxPyCor: 5,
		MinPzCor: -5,
		MaxPzCor: 5,
	}
	m := model.NewModel(settings)

	// Create turtles at different Z levels
	m.CreateTurtles(1, func(turtle *model.Turtle) {
		turtle.SetXYZ(0, 0, 0)
	})
	m.CreateTurtles(1, func(turtle *model.Turtle) {
		turtle.SetXYZ(0.5, 0, 0)
	})
	m.CreateTurtles(1, func(turtle *model.Turtle) {
		turtle.SetXYZ(0, 0, 2) // Same XY but different Z
	})

	// 2D search should find all 3 (ignores Z)
	turtles2D := m.TurtlesInRadiusXY(0, 0, 1)
	if turtles2D.Count() != 3 {
		t.Errorf("Expected 3 turtles in 2D radius, got %d", turtles2D.Count())
	}

	// 3D search should only find 2 (Z=2 is too far)
	turtles3D := m.TurtlesInRadiusXYZ(0, 0, 0, 1)
	if turtles3D.Count() != 2 {
		t.Errorf("Expected 2 turtles in 3D radius, got %d", turtles3D.Count())
	}
}

// Test 3D distance calculation
func TestDistanceBetweenPointsXYZ(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor: -10,
		MaxPxCor: 10,
		MinPyCor: -10,
		MaxPyCor: 10,
		MinPzCor: -10,
		MaxPzCor: 10,
	}
	m := model.NewModel(settings)

	// Test known 3D distance: 3-4-5 triangle (5 units)
	distance := m.DistanceBetweenPointsXYZ(0, 0, 0, 0, 3, 4)
	expected := 5.0
	if math.Abs(distance-expected) > 0.001 {
		t.Errorf("Expected distance %f, got %f", expected, distance)
	}

	// Test unit distance in each axis
	distX := m.DistanceBetweenPointsXYZ(0, 0, 0, 1, 0, 0)
	if math.Abs(distX-1.0) > 0.001 {
		t.Errorf("Expected X distance 1.0, got %f", distX)
	}

	distZ := m.DistanceBetweenPointsXYZ(0, 0, 0, 0, 0, 1)
	if math.Abs(distZ-1.0) > 0.001 {
		t.Errorf("Expected Z distance 1.0, got %f", distZ)
	}
}

// Test turtle movement in 3D
func TestTurtleMovement3D(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor: -10,
		MaxPxCor: 10,
		MinPyCor: -10,
		MaxPyCor: 10,
		MinPzCor: -10,
		MaxPzCor: 10,
	}
	m := model.NewModel(settings)

	m.CreateTurtles(1, func(turtle *model.Turtle) {
		turtle.SetXYZ(0, 0, 0)
	})

	turtle := m.Turtle(0)

	// Test SetXYZ
	turtle.SetXYZ(5, 3, 2)
	if turtle.XCor() != 5 || turtle.YCor() != 3 || turtle.ZCor() != 2 {
		t.Errorf("Expected position (5,3,2), got (%f,%f,%f)",
			turtle.XCor(), turtle.YCor(), turtle.ZCor())
	}

	// Test out of bounds - turtle should stay in place
	turtle.SetXYZ(100, 100, 100)
	if turtle.XCor() != 5 || turtle.YCor() != 3 || turtle.ZCor() != 2 {
		t.Errorf("Turtle moved out of bounds when it shouldn't have")
	}
}

// Test FaceXYZ sets heading and pitch correctly
func TestTurtleFaceXYZ(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor: -10,
		MaxPxCor: 10,
		MinPyCor: -10,
		MaxPyCor: 10,
		MinPzCor: -10,
		MaxPzCor: 10,
	}
	m := model.NewModel(settings)

	m.CreateTurtles(1, func(turtle *model.Turtle) {
		turtle.SetXYZ(0, 0, 0)
	})

	turtle := m.Turtle(0)

	// Face directly up (positive Z)
	turtle.FaceXYZ(0, 0, 1)
	if math.Abs(turtle.GetPitch()-90) > 0.1 {
		t.Errorf("Expected pitch 90 degrees when facing up, got %f", turtle.GetPitch())
	}

	// Face directly down (negative Z)
	turtle.FaceXYZ(0, 0, -1)
	pitch := turtle.GetPitch()
	// Pitch should be -90 or 270 (same angle)
	if math.Abs(pitch+90) > 0.1 && math.Abs(pitch-270) > 0.1 {
		t.Errorf("Expected pitch -90 or 270 degrees when facing down, got %f", pitch)
	}

	// Face horizontally (east)
	turtle.FaceXYZ(1, 0, 0)
	if math.Abs(turtle.GetPitch()) > 0.1 {
		t.Errorf("Expected pitch 0 degrees when facing horizontally, got %f", turtle.GetPitch())
	}
	// In standard math convention, 0 degrees is east (positive X)
	if math.Abs(turtle.GetHeading()) > 0.1 {
		t.Errorf("Expected heading 0 degrees (east), got %f", turtle.GetHeading())
	}
}

// Test Patch3D retrieval
func TestPatch3D(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor: -5,
		MaxPxCor: 5,
		MinPyCor: -5,
		MaxPyCor: 5,
		MinPzCor: -5,
		MaxPzCor: 5,
	}
	m := model.NewModel(settings)

	// Test valid patch retrieval
	patch := m.Patch3D(0, 0, 0)
	if patch == nil {
		t.Errorf("Expected patch at (0,0,0), got nil")
	}
	if patch.XCor() != 0 || patch.YCor() != 0 || patch.ZCor() != 0 {
		t.Errorf("Patch has wrong coordinates: (%d,%d,%d)",
			patch.XCor(), patch.YCor(), patch.ZCor())
	}

	// Test corner patch
	cornerPatch := m.Patch3D(5, 5, 5)
	if cornerPatch == nil {
		t.Errorf("Expected patch at corner (5,5,5), got nil")
	}

	// Test out of bounds
	outOfBounds := m.Patch3D(100, 100, 100)
	if outOfBounds != nil {
		t.Errorf("Expected nil for out of bounds patch, got patch")
	}

	// Test negative coordinates
	negativePatch := m.Patch3D(-3, -3, -3)
	if negativePatch == nil {
		t.Errorf("Expected patch at (-3,-3,-3), got nil")
	}
}

// Test diffusion in 2D (baseline)
func TestDiffuse2D(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor: -2,
		MaxPxCor: 2,
		MinPyCor: -2,
		MaxPyCor: 2,
		PatchProperties: map[string]interface{}{
			"chemical": 0.0,
		},
	}
	m := model.NewModel(settings)

	// Set center patch to 100
	centerPatch := m.Patch(0, 0)
	centerPatch.SetProperty("chemical", 100.0)

	// Diffuse 50%
	err := m.Diffuse("chemical", 0.5)
	if err != nil {
		t.Errorf("Diffuse failed: %v", err)
	}

	// Center should have less than 100
	centerAmount := centerPatch.GetProperty("chemical").(float64)
	if centerAmount >= 100 {
		t.Errorf("Center patch should have diffused, still has %f", centerAmount)
	}

	// Neighbors should have more than 0
	neighbors := centerPatch.Neighbors()
	hasNonZero := false
	neighbors.Ask(func(p *model.Patch) {
		amount := p.GetProperty("chemical").(float64)
		if amount > 0 {
			hasNonZero = true
		}
	})
	if !hasNonZero {
		t.Errorf("Expected neighbors to receive diffused chemical")
	}
}

// Test diffuse4 in 2D
func TestDiffuse4_2D(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor: -2,
		MaxPxCor: 2,
		MinPyCor: -2,
		MaxPyCor: 2,
		PatchProperties: map[string]interface{}{
			"chemical": 0.0,
		},
	}
	m := model.NewModel(settings)

	// Set center patch to 100
	centerPatch := m.Patch(0, 0)
	centerPatch.SetProperty("chemical", 100.0)

	// Diffuse 50%
	err := m.Diffuse4("chemical", 0.5)
	if err != nil {
		t.Errorf("Diffuse4 failed: %v", err)
	}

	// Center should have less than 100
	centerAmount := centerPatch.GetProperty("chemical").(float64)
	if centerAmount >= 100 {
		t.Errorf("Center patch should have diffused, still has %f", centerAmount)
	}

	// 4-neighbors should have more than 0
	neighbors := centerPatch.Neighbors4()
	if neighbors.Count() != 4 {
		t.Errorf("Expected 4 neighbors, got %d", neighbors.Count())
	}

	hasNonZero := false
	neighbors.Ask(func(p *model.Patch) {
		amount := p.GetProperty("chemical").(float64)
		if amount > 0 {
			hasNonZero = true
		}
	})
	if !hasNonZero {
		t.Errorf("Expected 4-neighbors to receive diffused chemical")
	}
}

// Test that 3D model is properly initialized
func TestIs3D(t *testing.T) {
	settings2D := model.ModelSettings{
		MinPxCor: -5,
		MaxPxCor: 5,
		MinPyCor: -5,
		MaxPyCor: 5,
	}
	m2D := model.NewModel(settings2D)

	if m2D.Is3D() {
		t.Errorf("Expected 2D model, got 3D")
	}

	settings3D := model.ModelSettings{
		MinPxCor: -5,
		MaxPxCor: 5,
		MinPyCor: -5,
		MaxPyCor: 5,
		MinPzCor: -5,
		MaxPzCor: 5,
	}
	m3D := model.NewModel(settings3D)

	if !m3D.Is3D() {
		t.Errorf("Expected 3D model, got 2D")
	}
}
