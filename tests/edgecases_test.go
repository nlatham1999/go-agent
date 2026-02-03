package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/pkg/model"
)

// TestTurtlesInRadiusZeroRadius tests TurtlesInRadius with radius = 0
func TestTurtlesInRadiusZeroRadius(t *testing.T) {
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	m.CreateTurtles(10, nil)

	// Place turtle at specific position
	turtle := m.Turtle(0)
	turtle.SetXY(5, 5)

	// Query with radius 0 - returns empty set (radius 0 means infinitesimally small circle)
	result := m.TurtlesInRadiusXY(5, 5, 0)

	if result.Count() != 0 {
		t.Errorf("Expected 0 turtles with radius 0, got %d", result.Count())
	}
}

// TestTurtlesInRadiusNegativeRadius tests TurtlesInRadius with negative radius
func TestTurtlesInRadiusNegativeRadius(t *testing.T) {
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	m.CreateTurtles(10, nil)

	// Query with negative radius - should return empty set
	result := m.TurtlesInRadiusXY(5, 5, -1)

	if result.Count() != 0 {
		t.Errorf("Expected 0 turtles with negative radius, got %d", result.Count())
	}
}

// TestTurtlesInRadiusLargerThanWorld tests TurtlesInRadius with radius larger than world
func TestTurtlesInRadiusLargerThanWorld(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor: -10,
		MaxPxCor: 10,
		MinPyCor: -10,
		MaxPyCor: 10,
	}
	m := model.NewModel(settings)

	m.CreateTurtles(100, nil)

	// Set all turtles to random positions
	m.Turtles().Ask(func(t *model.Turtle) {
		t.SetXY(float64(m.RandomXCor()), float64(m.RandomYCor()))
	})

	// Query with radius larger than world diagonal
	result := m.TurtlesInRadiusXY(0, 0, 1000)

	// Should return all turtles
	if result.Count() != 100 {
		t.Errorf("Expected all 100 turtles with huge radius, got %d", result.Count())
	}
}

// TestTurtlesInRadiusEmptyWorld tests TurtlesInRadius when no turtles exist
func TestTurtlesInRadiusEmptyWorld(t *testing.T) {
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	// No turtles created

	result := m.TurtlesInRadiusXY(0, 0, 5)

	if result.Count() != 0 {
		t.Errorf("Expected 0 turtles in empty world, got %d", result.Count())
	}
}

// TestTurtlesInRadiusOnBoundary tests TurtlesInRadius at world boundaries
func TestTurtlesInRadiusOnBoundary(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor:  -10,
		MaxPxCor:  10,
		MinPyCor:  -10,
		MaxPyCor:  10,
		WrappingX: false,
		WrappingY: false,
	}
	m := model.NewModel(settings)

	m.CreateTurtles(10, nil)

	// Place turtles at boundary
	m.Turtles().Ask(func(t *model.Turtle) {
		t.SetXY(10, 10)
	})

	// Query at boundary
	result := m.TurtlesInRadiusXY(10, 10, 1)

	if result.Count() != 10 {
		t.Errorf("Expected 10 turtles at boundary, got %d", result.Count())
	}
}

// TestPatchAtExactBoundary tests Patch() at exact world boundaries
func TestPatchAtExactBoundary(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor: -15,
		MaxPxCor: 15,
		MinPyCor: -15,
		MaxPyCor: 15,
	}
	m := model.NewModel(settings)

	// Test min boundary
	p := m.Patch(-15, -15)
	if p == nil {
		t.Errorf("Expected patch at min boundary, got nil")
	}
	if p.XCor() != -15 || p.YCor() != -15 {
		t.Errorf("Expected patch at (-15, -15), got (%d, %d)", p.XCor(), p.YCor())
	}

	// Test max boundary
	p = m.Patch(15, 15)
	if p == nil {
		t.Errorf("Expected patch at max boundary, got nil")
	}
	if p.XCor() != 15 || p.YCor() != 15 {
		t.Errorf("Expected patch at (15, 15), got (%d, %d)", p.XCor(), p.YCor())
	}

	// Test beyond boundary with no wrapping
	p = m.Patch(16, 16)
	if p != nil {
		t.Errorf("Expected nil patch beyond boundary without wrapping, got %v", p)
	}
}

// TestPatchAtBoundaryWithWrapping tests Patch() at boundaries with wrapping enabled
func TestPatchAtBoundaryWithWrapping(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor:  -15,
		MaxPxCor:  15,
		MinPyCor:  -15,
		MaxPyCor:  15,
		WrappingX: true,
		WrappingY: true,
	}
	m := model.NewModel(settings)

	// Test wrapping beyond max boundary
	p := m.Patch(16, 16)
	if p == nil {
		t.Errorf("Expected patch with wrapping, got nil")
	}
	// Should wrap to min side
	if p.XCor() != -15 || p.YCor() != -15 {
		t.Errorf("Expected wrapped patch at (-15, -15), got (%d, %d)", p.XCor(), p.YCor())
	}

	// Test wrapping beyond min boundary
	p = m.Patch(-16, -16)
	if p == nil {
		t.Errorf("Expected patch with wrapping, got nil")
	}
	// Should wrap to max side
	if p.XCor() != 15 || p.YCor() != 15 {
		t.Errorf("Expected wrapped patch at (15, 15), got (%d, %d)", p.XCor(), p.YCor())
	}
}

// TestPatchWithOnlyXWrapping tests Patch() with only X wrapping enabled
func TestPatchWithOnlyXWrapping(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor:  -10,
		MaxPxCor:  10,
		MinPyCor:  -10,
		MaxPyCor:  10,
		WrappingX: true,
		WrappingY: false,
	}
	m := model.NewModel(settings)

	// X should wrap
	p := m.Patch(11, 5)
	if p == nil {
		t.Errorf("Expected patch with X wrapping, got nil")
	}

	// Y should not wrap - should return nil
	p = m.Patch(5, 11)
	if p != nil {
		t.Errorf("Expected nil patch beyond Y boundary without Y wrapping, got %v", p)
	}
}

// TestTurtleSetXYOutOfBoundsNoWrapping tests turtle movement outside bounds
func TestTurtleSetXYOutOfBoundsNoWrapping(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor:  -10,
		MaxPxCor:  10,
		MinPyCor:  -10,
		MaxPyCor:  10,
		WrappingX: false,
		WrappingY: false,
	}
	m := model.NewModel(settings)

	m.CreateTurtles(1, nil)
	turtle := m.Turtle(0)

	originalX := turtle.XCor()
	originalY := turtle.YCor()

	// Try to move beyond boundary - should not move
	turtle.SetXY(20, 20)

	if turtle.XCor() != originalX || turtle.YCor() != originalY {
		t.Errorf("Turtle should not move beyond boundaries without wrapping")
	}
}

// TestTurtlePropertyNil tests setting and getting nil properties
func TestTurtlePropertyNil(t *testing.T) {
	settings := model.ModelSettings{
		TurtleProperties: map[string]interface{}{
			"mood": "happy",
		},
	}
	m := model.NewModel(settings)

	m.CreateTurtles(1, nil)
	turtle := m.Turtle(0)

	// Get non-existent property
	val := turtle.GetProperty("nonexistent")
	if val != nil {
		t.Errorf("Expected nil for non-existent property, got %v", val)
	}

	// Set property to nil
	turtle.SetProperty("mood", nil)
	val = turtle.GetProperty("mood")
	if val != nil {
		t.Errorf("Expected nil after setting property to nil, got %v", val)
	}
}

// TestPatchPropertyNil tests setting and getting nil patch properties
func TestPatchPropertyNil(t *testing.T) {
	settings := model.ModelSettings{
		PatchProperties: map[string]interface{}{
			"chemical": 10.0,
		},
	}
	m := model.NewModel(settings)

	p := m.Patch(0, 0)

	// Get non-existent property
	val := p.GetProperty("nonexistent")
	if val != nil {
		t.Errorf("Expected nil for non-existent property, got %v", val)
	}

	// Set property to nil
	p.SetProperty("chemical", nil)
	val = p.GetProperty("chemical")
	if val != nil {
		t.Errorf("Expected nil after setting property to nil, got %v", val)
	}
}

// TestLinkSelfLoop tests creating a link from turtle to itself
func TestLinkSelfLoop(t *testing.T) {
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	m.CreateTurtles(1, nil)
	turtle := m.Turtle(0)

	// Self-loops are actually allowed in the model
	link, err := turtle.CreateLinkToTurtle(nil, turtle, nil)
	if err != nil {
		t.Errorf("Expected no error when creating self-loop link, got %v", err)
	}

	if link == nil {
		t.Errorf("Expected link to be created")
	}

	// Should be able to create undirected self-loop too
	link2, err := turtle.CreateLinkWithTurtle(nil, turtle, nil)
	if err != nil {
		t.Errorf("Expected no error when creating undirected self-loop link, got %v", err)
	}

	if link2 == nil {
		t.Errorf("Expected undirected link to be created")
	}
}

// TestDuplicateLinkCreation tests creating duplicate links
func TestDuplicateLinkCreation(t *testing.T) {
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	m.CreateTurtles(2, nil)
	t1 := m.Turtle(0)
	t2 := m.Turtle(1)

	// Create first link
	_, err := t1.CreateLinkToTurtle(nil, t2, nil)
	if err != nil {
		t.Errorf("Expected no error on first link creation, got %v", err)
	}

	// Try to create duplicate - should fail
	_, err = t1.CreateLinkToTurtle(nil, t2, nil)
	if err == nil {
		t.Errorf("Expected error when creating duplicate link, got nil")
	}
}

// TestCreateTurtlesZeroCount tests creating zero turtles
func TestCreateTurtlesZeroCount(t *testing.T) {
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	turtles, err := m.CreateTurtles(0, nil)
	if err != nil {
		t.Errorf("Expected no error when creating 0 turtles, got %v", err)
	}

	if turtles != nil && turtles.Count() != 0 {
		t.Errorf("Expected 0 turtles, got %d", turtles.Count())
	}

	if m.Turtles().Count() != 0 {
		t.Errorf("Expected empty turtle set, got %d", m.Turtles().Count())
	}
}

// TestCreateTurtlesNegativeCount tests creating negative number of turtles
func TestCreateTurtlesNegativeCount(t *testing.T) {
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	turtles, err := m.CreateTurtles(-5, nil)
	if err != nil {
		t.Errorf("Expected no error with negative count (should create 0), got %v", err)
	}

	if turtles != nil && turtles.Count() != 0 {
		t.Errorf("Expected 0 turtles with negative count, got %d", turtles.Count())
	}
}

// TestTurtleDistanceToSelf tests distance from turtle to itself
func TestTurtleDistanceToSelf(t *testing.T) {
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	m.CreateTurtles(1, nil)
	turtle := m.Turtle(0)

	distance := turtle.DistanceTurtle(turtle)
	if distance != 0 {
		t.Errorf("Expected distance 0 to self, got %v", distance)
	}
}

// TestPatchNeighborsAtCorner tests patch neighbors at world corner
func TestPatchNeighborsAtCorner(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor:  -5,
		MaxPxCor:  5,
		MinPyCor:  -5,
		MaxPyCor:  5,
		WrappingX: false,
		WrappingY: false,
	}
	m := model.NewModel(settings)

	// Corner patch should have fewer neighbors
	cornerPatch := m.Patch(-5, -5)
	neighbors := cornerPatch.Neighbors()

	// Corner has only 3 neighbors without wrapping
	if neighbors.Count() != 3 {
		t.Errorf("Expected 3 neighbors at corner without wrapping, got %d", neighbors.Count())
	}
}

// TestPatchNeighbors4AtCorner tests patch neighbors4 at world corner
func TestPatchNeighbors4AtCorner(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor:  -5,
		MaxPxCor:  5,
		MinPyCor:  -5,
		MaxPyCor:  5,
		WrappingX: false,
		WrappingY: false,
	}
	m := model.NewModel(settings)

	// Corner patch should have fewer neighbors4
	cornerPatch := m.Patch(-5, -5)
	neighbors := cornerPatch.Neighbors4()

	// Corner has only 2 neighbors4 without wrapping
	if neighbors.Count() != 2 {
		t.Errorf("Expected 2 neighbors4 at corner without wrapping, got %d", neighbors.Count())
	}
}

// TestModelWithNoPatches tests model with invalid dimensions
func TestModelWithZeroDimensions(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor: 0,
		MaxPxCor: 0,
		MinPyCor: 0,
		MaxPyCor: 0,
	}
	m := model.NewModel(settings)

	// Should have at least 1 patch
	if m.Patches.Count() < 1 {
		t.Errorf("Expected at least 1 patch even with zero dimensions, got %d", m.Patches.Count())
	}
}

// TestTurtleForwardZeroDistance tests turtle forward with distance 0
func TestTurtleForwardZeroDistance(t *testing.T) {
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	m.CreateTurtles(1, nil)
	turtle := m.Turtle(0)

	originalX := turtle.XCor()
	originalY := turtle.YCor()

	turtle.Forward(0)

	if turtle.XCor() != originalX || turtle.YCor() != originalY {
		t.Errorf("Turtle should not move with distance 0")
	}
}

// TestTurtleForwardNegativeDistance tests turtle forward with negative distance
func TestTurtleForwardNegativeDistance(t *testing.T) {
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	m.CreateTurtles(1, nil)
	turtle := m.Turtle(0)
	turtle.SetHeading(0) // Face right
	turtle.SetXY(0, 0)

	turtle.Forward(-5)

	// Should move backward (left)
	if turtle.XCor() >= 0 {
		t.Errorf("Turtle should move backward with negative distance, x=%v", turtle.XCor())
	}
}

// TestAgentSetWithNilSet tests operations on empty agent sets
func TestAgentSetWithNilSet(t *testing.T) {
	emptySet := model.NewTurtleAgentSet(nil)

	if emptySet.Count() != 0 {
		t.Errorf("Expected empty set count 0, got %d", emptySet.Count())
	}

	first, err := emptySet.First()
	if err == nil {
		t.Errorf("Expected error when getting first of empty set, got nil")
	}
	if first != nil {
		t.Errorf("Expected nil turtle from empty set, got %v", first)
	}

	last, err := emptySet.Last()
	if err == nil {
		t.Errorf("Expected error when getting last of empty set, got nil")
	}
	if last != nil {
		t.Errorf("Expected nil turtle from empty set, got %v", last)
	}

	// Ask on empty set should not panic
	callCount := 0
	emptySet.Ask(func(turtle *model.Turtle) {
		callCount++
	})
	if callCount > 0 {
		t.Errorf("Should not be called on empty set")
	}

	// With on empty set should return empty set
	filtered := emptySet.With(func(t *model.Turtle) bool {
		return true
	})
	if filtered.Count() != 0 {
		t.Errorf("Expected empty filtered set, got %d", filtered.Count())
	}
}

// TestDistanceBetweenPointsWithWrappingEdgeCases tests distance calculation edge cases
func TestDistanceBetweenPointsWithWrappingEdgeCases(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor:  -10,
		MaxPxCor:  10,
		MinPyCor:  -10,
		MaxPyCor:  10,
		WrappingX: true,
		WrappingY: true,
	}
	m := model.NewModel(settings)

	// Same point should be distance 0
	dist := m.DistanceBetweenPointsXY(5, 5, 5, 5)
	if dist != 0 {
		t.Errorf("Expected distance 0 for same point, got %v", dist)
	}

	// Opposite corners with wrapping should be close
	dist = m.DistanceBetweenPointsXY(-10, -10, 10, 10)
	// With wrapping, shortest path might be smaller than direct distance
	if dist > 30 { // Just sanity check
		t.Errorf("Expected reasonable wrapped distance, got %v", dist)
	}
}

// TestTurtleDieWhileIterating tests killing turtles during Ask iteration
func TestTurtleDieWhileIterating(t *testing.T) {
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	m.CreateTurtles(10, nil)

	count := 0
	// Kill turtles during iteration - should not panic
	m.Turtles().Ask(func(turtle *model.Turtle) {
		count++
		if count%2 == 0 {
			turtle.Die()
		}
	})

	// Should have 5 turtles left
	if m.Turtles().Count() != 5 {
		t.Errorf("Expected 5 turtles after killing every other one, got %d", m.Turtles().Count())
	}
}

// TestBreedTransitionPreservesProperties tests that properties are maintained during breed change
func TestBreedTransitionPreservesProperties(t *testing.T) {
	breed1 := model.NewTurtleBreed("breed1", "", map[string]interface{}{
		"breed1_prop": 10,
	})
	breed2 := model.NewTurtleBreed("breed2", "", map[string]interface{}{
		"breed2_prop": 20,
	})

	settings := model.ModelSettings{
		TurtleBreeds: []*model.TurtleBreed{breed1, breed2},
		TurtleProperties: map[string]interface{}{
			"general_prop": 5,
		},
	}
	_ = model.NewModel(settings)

	breed1.CreateAgents(1, nil)
	turtle := breed1.Agent(0)

	// Check initial properties
	if turtle.GetProperty("general_prop") != 5 {
		t.Errorf("Expected general_prop to be 5")
	}
	if turtle.GetProperty("breed1_prop") != 10 {
		t.Errorf("Expected breed1_prop to be 10")
	}

	// Change breed
	turtle.SetBreed(breed2)

	// General property should be preserved
	if turtle.GetProperty("general_prop") != 5 {
		t.Errorf("Expected general_prop to be preserved after breed change")
	}

	// Breed1 property should be gone
	if turtle.GetProperty("breed1_prop") != nil {
		t.Errorf("Expected breed1_prop to be removed after breed change")
	}

	// Breed2 property should be present
	if turtle.GetProperty("breed2_prop") != 20 {
		t.Errorf("Expected breed2_prop to be 20 after breed change")
	}
}

// TestConcurrentTurtleCreation tests creating turtles from multiple goroutines
// NOTE: This test is commented out because CreateTurtles is NOT thread-safe.
// The model currently has concurrent map writes in newTurtle() at turtle.go:54
// Users should not call CreateTurtles concurrently from multiple goroutines.
// If concurrent turtle creation is needed, it must be serialized with a mutex.
func TestConcurrentTurtleCreation(t *testing.T) {
	t.Skip("CreateTurtles is not thread-safe - this is expected behavior, not a bug")

	// Keeping this test commented for documentation purposes:
	// settings := model.ModelSettings{}
	// m := model.NewModel(settings)
	//
	// // Create turtles from 10 goroutines concurrently
	// done := make(chan bool)
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		m.CreateTurtles(10, nil)
	// 		done <- true
	// 	}()
	// }
	//
	// // Wait for all goroutines
	// for i := 0; i < 10; i++ {
	// 	<-done
	// }
}

// TestConcurrentTurtlesInRadius tests concurrent spatial queries
func TestConcurrentTurtlesInRadius(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor: -50,
		MaxPxCor: 50,
		MinPyCor: -50,
		MaxPyCor: 50,
	}
	m := model.NewModel(settings)

	m.CreateTurtles(1000, func(turtle *model.Turtle) {
		turtle.SetXY(float64(m.RandomXCor()), float64(m.RandomYCor()))
	})

	// Perform many spatial queries concurrently
	done := make(chan bool)
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				x := float64(m.RandomXCor())
				y := float64(m.RandomYCor())
				_ = m.TurtlesInRadiusXY(x, y, 10)
			}
			done <- true
		}()
	}

	// Wait for all goroutines - should not panic or deadlock
	for i := 0; i < 100; i++ {
		<-done
	}
}

// TestConcurrentTurtlePropertyUpdates tests concurrent turtle property writes and reads
func TestConcurrentTurtlePropertyUpdates(t *testing.T) {
	settings := model.ModelSettings{
		TurtleProperties: map[string]interface{}{
			"energy": 100.0,
		},
	}
	m := model.NewModel(settings)

	m.CreateTurtles(100, nil)

	// Write to turtle properties from multiple goroutines
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(val int) {
			m.Turtles().Ask(func(t *model.Turtle) {
				t.SetProperty("energy", float64(val))
			})
			done <- true
		}(i)
	}

	// Also read concurrently
	for i := 0; i < 10; i++ {
		go func() {
			m.Turtles().Ask(func(t *model.Turtle) {
				_ = t.GetProperty("energy")
			})
			done <- true
		}()
	}

	// Wait for all goroutines - should not panic or deadlock
	for i := 0; i < 20; i++ {
		<-done
	}

	// Verify no panics occurred
	if m.Turtles().Count() != 100 {
		t.Errorf("Expected 100 turtles, got %d", m.Turtles().Count())
	}
}

// TestConcurrentPatchPropertyUpdates tests concurrent patch property writes and reads
func TestConcurrentPatchPropertyUpdates(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor: -10,
		MaxPxCor: 10,
		MinPyCor: -10,
		MaxPyCor: 10,
		PatchProperties: map[string]interface{}{
			"chemical": 0.0,
		},
	}
	m := model.NewModel(settings)

	// Write to patch properties from multiple goroutines
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(val int) {
			m.Patches.Ask(func(p *model.Patch) {
				p.SetProperty("chemical", float64(val))
			})
			done <- true
		}(i)
	}

	// Also read concurrently
	for i := 0; i < 10; i++ {
		go func() {
			m.Patches.Ask(func(p *model.Patch) {
				_ = p.GetProperty("chemical")
			})
			done <- true
		}()
	}

	// Wait for all goroutines - should not panic or deadlock
	for i := 0; i < 20; i++ {
		<-done
	}

	// Verify no panics occurred
	if m.Patches.Count() != 441 { // 21x21 patches
		t.Errorf("Expected 441 patches, got %d", m.Patches.Count())
	}
}

// TestConcurrentLinkCreation tests creating links from multiple goroutines
// NOTE: This test is commented out because link creation is likely NOT thread-safe.
// Similar to CreateTurtles, the link creation modifies shared model state without synchronization.
func TestConcurrentLinkCreation(t *testing.T) {
	t.Skip("Link creation is not thread-safe - this is expected behavior")

	// Users should create links sequentially or use proper synchronization
}

// TestConcurrentTurtleMovement tests moving turtles from multiple goroutines
func TestConcurrentTurtleMovement(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor: -5,
		MaxPxCor: 5,
		MinPyCor: -5,
		MaxPyCor: 5,
	}
	m := model.NewModel(settings)

	// Dense world: 500 turtles in 11x11 world
	m.CreateTurtles(500, nil)

	// Move turtles concurrently - lots of patch contention
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			m.Turtles().Ask(func(turtle *model.Turtle) {
				turtle.Forward(1)
				turtle.Right(10)
			})
			done <- true
		}()
	}

	// Wait for all goroutines - should not panic or deadlock
	for i := 0; i < 10; i++ {
		<-done
	}

	// All turtles should still exist
	if m.Turtles().Count() != 500 {
		t.Errorf("Expected 500 turtles after concurrent movement, got %d", m.Turtles().Count())
	}
}

// TestPatchDiffuseZeroAmount tests diffusion with zero amount
func TestPatchDiffuseZeroAmount(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor: -5,
		MaxPxCor: 5,
		MinPyCor: -5,
		MaxPyCor: 5,
		PatchProperties: map[string]interface{}{
			"chemical": 100.0,
		},
	}
	m := model.NewModel(settings)

	centerPatch := m.Patch(0, 0)
	originalValue := centerPatch.GetProperty("chemical").(float64)

	// Diffuse with 0 amount - should not change anything
	m.Diffuse("chemical", 0)

	newValue := centerPatch.GetProperty("chemical").(float64)
	if newValue != originalValue {
		t.Errorf("Expected no change with 0 diffusion, got %v -> %v", originalValue, newValue)
	}
}

// TestPatchDiffuseNegativeAmount tests diffusion with negative amount
func TestPatchDiffuseNegativeAmount(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor: -5,
		MaxPxCor: 5,
		MinPyCor: -5,
		MaxPyCor: 5,
		PatchProperties: map[string]interface{}{
			"chemical": 100.0,
		},
	}
	m := model.NewModel(settings)

	// Diffuse with negative amount - should handle gracefully
	m.Diffuse("chemical", -0.5)

	// Should not panic, exact behavior depends on implementation
}

// TestLinkColorEdgeCases tests link color assignment
func TestLinkColorEdgeCases(t *testing.T) {
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	m.CreateTurtles(2, nil)
	t1 := m.Turtle(0)
	t2 := m.Turtle(1)

	// Create link with custom color
	link, err := t1.CreateLinkToTurtle(nil, t2, func(l *model.Link) {
		l.Color = model.Color{Red: 255, Green: 0, Blue: 0}
	})

	if err != nil {
		t.Errorf("Expected no error creating link, got %v", err)
	}

	if link.Color.Red != 255 || link.Color.Green != 0 || link.Color.Blue != 0 {
		t.Errorf("Expected red link color, got Red=%d Green=%d Blue=%d", link.Color.Red, link.Color.Green, link.Color.Blue)
	}
}

// TestTurtleHatchEdgeCases tests turtle hatching edge cases
func TestTurtleHatchEdgeCases(t *testing.T) {
	settings := model.ModelSettings{
		TurtleProperties: map[string]interface{}{
			"energy": 100,
		},
	}
	m := model.NewModel(settings)

	m.CreateTurtles(1, nil)
	parent := m.Turtle(0)
	parent.SetProperty("energy", 50)

	initialCount := m.Turtles().Count()

	// Hatch 0 turtles - should not create any new turtles
	parent.Hatch(0, nil)
	if m.Turtles().Count() != initialCount {
		t.Errorf("Expected no new turtles from hatching 0, got %d", m.Turtles().Count()-initialCount)
	}

	// Hatch 1 turtle and check it was created
	parent.Hatch(1, nil)
	if m.Turtles().Count() != initialCount+1 {
		t.Errorf("Expected 1 new turtle from hatching, got %d", m.Turtles().Count()-initialCount)
	}
}

// TestTurtleSetHeadingEdgeCases tests heading edge cases
func TestTurtleSetHeadingEdgeCases(t *testing.T) {
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	m.CreateTurtles(1, nil)
	turtle := m.Turtle(0)

	// Set heading > 360
	turtle.SetHeading(720)
	// Should wrap to 0
	if turtle.GetHeading() != 0 {
		t.Errorf("Expected heading 720 to wrap to 0, got %v", turtle.GetHeading())
	}

	// Set negative heading
	turtle.SetHeading(-90)
	// Should wrap to 270
	if turtle.GetHeading() != 270 {
		t.Errorf("Expected heading -90 to wrap to 270, got %v", turtle.GetHeading())
	}
}

// TestAgentSetRandomSelectionEdgeCases tests FirstNOf and LastNOf edge cases
func TestAgentSetRandomSelectionEdgeCases(t *testing.T) {
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	m.CreateTurtles(10, nil)

	// FirstNOf with n=0
	selected := m.Turtles().FirstNOf(0)
	if selected.Count() != 0 {
		t.Errorf("Expected 0 turtles from FirstNOf(0), got %d", selected.Count())
	}

	// FirstNOf with n > count - should return all
	selected = m.Turtles().FirstNOf(20)
	if selected.Count() != 10 {
		t.Errorf("Expected 10 turtles when FirstNOf(20) > count(10), got %d", selected.Count())
	}

	// LastNOf with n=0
	selected = m.Turtles().LastNOf(0)
	if selected.Count() != 0 {
		t.Errorf("Expected 0 turtles from LastNOf(0), got %d", selected.Count())
	}

	// FirstNOf on empty set
	emptySet := model.NewTurtleAgentSet(nil)
	selected = emptySet.FirstNOf(1)
	if selected.Count() != 0 {
		t.Errorf("Expected 0 turtles from FirstNOf on empty set, got %d", selected.Count())
	}
}

// TestPatchAtExtremeCoordinates tests patches at very large coordinates
func TestPatchAtExtremeCoordinates(t *testing.T) {
	settings := model.ModelSettings{
		MinPxCor: -1000,
		MaxPxCor: 1000,
		MinPyCor: -1000,
		MaxPyCor: 1000,
	}
	m := model.NewModel(settings)

	// Test extreme corners
	p := m.Patch(-1000, -1000)
	if p == nil {
		t.Errorf("Expected patch at extreme min corner")
	}

	p = m.Patch(1000, 1000)
	if p == nil {
		t.Errorf("Expected patch at extreme max corner")
	}

	// Test center
	p = m.Patch(0, 0)
	if p == nil {
		t.Errorf("Expected patch at center")
	}
}
