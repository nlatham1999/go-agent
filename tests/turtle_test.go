package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/internal/model"
)

func TestTurtleBack(t *testing.T) {

	//create a basic model
	m := model.NewModel(nil, nil, nil, nil, nil, nil, false, false)

	//create a turtle
	turtle := model.NewTurtle(m, 0, "", 0, 0)

	//set the turtle's heading
	turtle.SetHeading(90)

	//move back 10
	turtle.Back(10)

	//assert that the turtle's y position is now -10
	if turtle.YCor() != -10 && turtle.XCor() != 0 {
		t.Errorf("Expected turtle to move back 10")
	}

	turtle.Back(10)

	//assert that the turtle's y position is now -15 since that is the edge of the map
	if turtle.YCor() != -15 && turtle.XCor() != 0 {
		t.Errorf("Expected turtle to move back 10")
	}
}

func TestTurtleBreedName(t *testing.T) {

	//create a basic model
	m := model.NewModel(nil, nil, nil, []string{"ants"}, nil, nil, false, false)

	//create a turtle
	m.CreateTurtles(1, "ants", nil)

	turtle := m.Turtle("ants", 0)
	//assert that the turtle's breed is "ants"
	if turtle.BreedName() != "ants" {
		t.Errorf("Expected turtle to have breed 'ants'")
	}
}

func TestTurtleBreed(t *testing.T) {

	//create a basic model
	m := model.NewModel(nil, nil, nil, []string{"ants"}, nil, nil, false, false)

	//create a turtle
	m.CreateTurtles(1, "ants", nil)

	turtle := m.Turtle("ants", 0)

	if turtle.Breed() == nil {
		t.Errorf("Expected turtle to have a breed")
	}

	if turtle.Breed().Name() != "ants" {
		t.Errorf("Expected turtle to have breed 'ants'")
	}
}

func TestTurtleSetBreed(t *testing.T) {

	//create a basic model
	m := model.NewModel(nil, nil, nil, []string{"ants", "beetles"}, nil, nil, false, false)

	//create a turtle
	m.CreateTurtles(5, "ants", nil)

	turtle := m.Turtle("ants", 0)

	turtle.SetBreed("beetles")

	if turtle.BreedName() != "beetles" {
		t.Errorf("Expected turtle to have breed 'beetles'")
	}

	turtle = m.Turtle("ants", 0)

	if turtle != nil {
		t.Errorf("Expected turtle to not exist in breed 'ants'")
	}

	turtle = m.Turtle("beetles", 0)

	if turtle == nil {
		t.Errorf("Expected turtle to exist in breed 'beetles'")
	}

	m.CreateTurtles(1, "", nil)

	turtle = m.Turtle("", 5)

	if turtle.BreedName() != "" {
		t.Errorf("Expected turtle to have breed ''")
	}

	turtle.SetBreed("ants")

	if turtle.BreedName() != "ants" {
		t.Errorf("Expected turtle to have breed 'ants'")
	}

	turtle = m.Turtle("ants", 5)

	if turtle == nil {
		t.Errorf("Expected turtle to exist in breed 'ants'")
	}
}

func TestTurtlesOwn(t *testing.T) {

	// create a basic model
	m := model.NewModel(nil, nil, nil, nil, nil, nil, false, false)

	// create a turtle
	m.CreateTurtles(1, "", nil)

	turtle := m.Turtle("", 0)

	// get the turtle's own
	mood := turtle.GetOwn("mood")

	if mood != nil {
		t.Errorf("Expected turtle to not have own 'mood'")
	}

	// create a second model with ants breed
	turtlesOwn := map[string]interface{}{
		"mood": "happy",
		"age":  5,
	}

	antsOwn := map[string]interface{}{
		"mood": 0,
	}

	breedsOwn := map[string]map[string]interface{}{
		"ants": antsOwn,
	}

	m2 := model.NewModel(nil, turtlesOwn, breedsOwn, []string{"ants"}, nil, nil, false, false)

	// create a turtle
	m2.CreateTurtles(1, "ants", nil)

	turtle = m2.Turtle("", 0)

	// get the turtle's own
	mood = turtle.GetOwn("mood")

	if mood != 0 {
		t.Errorf("Expected turtle to have own 'mood' with value 0, got %v", mood)
	}

	age := turtle.GetOwn("age")

	if age != 5 {
		t.Errorf("Expected turtle to have own 'age' with value 5")
	}

	// test changing breed
	// create a second model with ants and beetles breeds
	beetlesOwn := map[string]interface{}{
		"wingspan": 10,
	}

	breedsOwn["beetles"] = beetlesOwn

	m3 := model.NewModel(nil, turtlesOwn, breedsOwn, []string{"ants", "beetles"}, nil, nil, false, false)

	// create a turtle
	m3.CreateTurtles(1, "beetles", nil)

	turtle = m3.Turtle("beetles", 0)

	//change the breed
	turtle.SetBreed("ants")

	wingspan := turtle.GetOwn("wingspan")

	if wingspan != nil {
		t.Errorf("Expected turtle to not have own 'wingspan'")
	}

	mood = turtle.GetOwn("mood")

	if mood == nil {
		t.Errorf("Expected turtle to have own 'mood'")
	}

	if mood != 0 {
		t.Errorf("Expected turtle to have own 'mood' with value 0, got %v", mood)
	}

}

func TestTurtleCreateLinkToTurtle(t *testing.T) {

	// create a new model
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false, false)

	// create some turtles
	m.CreateTurtles(2, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	// create a new link
	t1.CreateLinkToTurtle("parent-children", t2, nil)

	// make sure the link exists from t1 to t2
	l := t2.LinkFrom("parent-children", t1)

	if l == nil {
		t.Errorf("Link should have been created")
	}
}

func TestTurtleCreateLinkToSet(t *testing.T) {

	// create a new model
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false, false)

	// create some turtles
	m.CreateTurtles(3, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)
	t3 := m.Turtle("", 2)

	agentSet := model.TurtleSet([]*model.Turtle{t2, t3})

	t1.CreateLinksToSet("parent-children", agentSet, nil)

	// make sure the link exists from t1 to t2
	l := t2.LinkFrom("parent-children", t1)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link exists from t1 to t3
	l = t3.LinkFrom("parent-children", t1)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link does not exist from t2 to t3
	l = t3.LinkFrom("parent-children", t2)

	if l != nil {
		t.Errorf("Link should not have been created")
	}
}

func TestTurtleCreateLinkWithTurtle(t *testing.T) {

	// create a new model
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false, false)

	// create some turtles
	m.CreateTurtles(2, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	// create a new link
	t1.CreateLinkWithTurtle("coworkers", t2, nil)

	// make sure the link exists from t1 to t2
	l := t2.LinkFrom("coworkers", t1)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link exists from t2 to t1
	l = t1.LinkFrom("coworkers", t2)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link does not exist from t1 to t2 for parent-children
	l = t2.LinkFrom("parent-children", t1)

	if l != nil {
		t.Errorf("Link should not have been created")
	}
}

func TestTurtleCreateLinkWithSet(t *testing.T) {

	// create a new model
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false, false)

	// create some turtles
	m.CreateTurtles(3, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)
	t3 := m.Turtle("", 2)

	agentSet := model.TurtleSet([]*model.Turtle{t2, t3})

	t1.CreateLinksWithSet("coworkers", agentSet, nil)

	// make sure the link exists from t1 to t2
	l := t2.LinkFrom("coworkers", t1)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link exists from t1 to t3
	l = t3.LinkFrom("coworkers", t1)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link exists from t2 to t1
	l = t1.LinkFrom("coworkers", t2)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link exists from t3 to t1
	l = t1.LinkFrom("coworkers", t3)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link does not exist from t1 to t2 for parent-children
	l = t2.LinkFrom("parent-children", t1)

	if l != nil {
		t.Errorf("Link should not have been created")
	}
}

func TestTurtleCreateLinkFromTurtle(t *testing.T) {

	// create a new model
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false, false)

	// create some turtles
	m.CreateTurtles(2, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	// create a new link
	t1.CreateLinkFromTurtle("parent-children", t2, nil)

	// make sure the link exists from t2 to t1
	l := t1.LinkFrom("parent-children", t2)
	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link does not exist from t1 to t2
	l = t2.LinkFrom("parent-children", t1)

	if l != nil {
		t.Errorf("Link should not have been created")
	}
}

func TestTurtleCreateLinkFromSet(t *testing.T) {

	// create a new model
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false, false)

	// create some turtles
	m.CreateTurtles(3, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)
	t3 := m.Turtle("", 2)

	agentSet := model.TurtleSet([]*model.Turtle{t2, t3})

	t1.CreateLinksFromSet("parent-children", agentSet, nil)

	// make sure the link exists from t2 to t1
	l := t1.LinkFrom("parent-children", t2)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link exists from t3 to t1
	l = t1.LinkFrom("parent-children", t3)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link does not exist from t1 to t2
	l = t2.LinkFrom("parent-children", t1)

	if l != nil {
		t.Errorf("Link should not have been created")
	}

	// make sure the link does not exist from t1 to t3
	l = t3.LinkFrom("parent-children", t1)

	if l != nil {
		t.Errorf("Link should not have been created")
	}

	// make sure the link does not exist from t1 to t2 for parent-children
	l = t2.LinkFrom("coworkers", t1)

	if l != nil {
		t.Errorf("Link should not have been created")
	}
}

// test the Turtle.DistanceTurtle method
func TestTurtleDistanceTurtle(t *testing.T) {

	// create a basic model
	m := model.NewModel(nil, nil, nil, nil, nil, nil, false, false)

	// create two turtles
	m.CreateTurtles(2, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	// set the position of the turtles
	t1.SetXY(0, 0)
	t2.SetXY(0, 10)

	// get the distance between the turtles
	distance := t1.DistanceTurtle(t2)

	// assert that the distance is 10
	if distance != 10 {
		t.Errorf("Expected distance to be 10, got %v", distance)
	}
}

func TestTurtleDistancePatch(t *testing.T) {

	// create a basic model
	m := model.NewModel(nil, nil, nil, nil, nil, nil, false, false)

	// create a turtle
	m.CreateTurtles(1, "", nil)

	t1 := m.Turtle("", 0)

	// set the position of the turtle
	t1.SetXY(0, 0)

	// get the distance between the turtle and the patch (0, 10)
	distance := t1.DistancePatch(m.Patch(0, 10))

	// assert that the distance is 10
	if distance != 10 {
		t.Errorf("Expected distance to be 10, got %v", distance)
	}
}

func TestTurtleDistanceXY(t *testing.T) {

	// create a basic model
	m := model.NewModel(nil, nil, nil, nil, nil, nil, false, false)

	// create a turtle
	m.CreateTurtles(1, "", nil)

	t1 := m.Turtle("", 0)

	// set the position of the turtle
	t1.SetXY(0, 0)

	// get the distance between the turtle and the point (0, 10)
	distance := t1.DistanceXY(0, 10)

	// assert that the distance is 10
	if distance != 10 {
		t.Errorf("Expected distance to be 10, got %v", distance)
	}
}

func TestTurtleOtherEnd(t *testing.T) {

	// create a new model
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false, false)

	// create some turtles
	m.CreateTurtles(2, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	// create a new link
	l, err := model.NewLink(m, "parent-children", t1, t2, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// get the other end of the link
	otherEnd := t1.OtherEnd(l)

	// make sure the other end is t2
	if otherEnd != t2 {
		t.Errorf("Expected other end to be t2")
	}

	// get the other end of the link
	otherEnd = t2.OtherEnd(l)

	// make sure the other end is t1
	if otherEnd != t1 {
		t.Errorf("Expected other end to be t1")
	}
}

func TestTurtleSetXY(t *testing.T) {

	// create a model with wrapping
	m := model.NewModel(nil, nil, nil, nil, nil, nil, true, true)

	// create a turtle
	m.CreateTurtles(1, "", nil)

	turtle := m.Turtle("", 0)

	turtle.SetXY(15.4999, 15.4999)

	if turtle.XCor() != 15.4999 || turtle.YCor() != 15.4999 {
		t.Errorf("Expected turtle to be at position (15.4999, 15.4999)")
	}

	// set the position of the turtle
	turtle.SetXY(15.5, 15.5)

	if turtle.XCor() != -15.5 || turtle.YCor() != -15.5 {
		t.Errorf("Expected turtle to be at position (15.5, 15.5)")
	}

	turtle.SetXY(15.51, 15.51)

	if turtle.XCor()+15.49 > .0001 || turtle.YCor()+15.49 > .0001 {
		t.Errorf("Expected turtle to be at position (-15.49, -15.49), got (%v, %v)", turtle.XCor(), turtle.YCor())
	}
}

func TestTurtleDownhill(t *testing.T) {

	// create a model with patches own of chemical

	patchesOwn := map[string]interface{}{
		"chemical": 0.0,
	}

	m := model.NewModel(patchesOwn, nil, nil, nil, nil, nil, false, false)

	// create a turtle
	m.CreateTurtles(1, "", nil)

	// get the 9 patches around the turtle
	p1 := m.Patch(0, 0)
	p2 := m.Patch(0, 1)
	p3 := m.Patch(0, -1)
	p4 := m.Patch(1, 0)
	p5 := m.Patch(1, 1)
	p6 := m.Patch(1, -1)
	p7 := m.Patch(-1, 0)
	p8 := m.Patch(-1, 1)
	p9 := m.Patch(-1, -1)

	// set the chemical value of the patches
	p1.SetOwn("chemical", 0.0)
	p2.SetOwn("chemical", 1.0)
	p3.SetOwn("chemical", 2.0)
	p4.SetOwn("chemical", 3.0)
	p5.SetOwn("chemical", 4.0)
	p6.SetOwn("chemical", 5.0)
	p7.SetOwn("chemical", 6.0)
	p8.SetOwn("chemical", 7.0)
	p9.SetOwn("chemical", 8.0)

	turtle := m.Turtle("", 0)
	turtle.Downhill("chemical")

	// make sure that the turtle's position has not changed since the patch it is on has the lowest chemical value
	if turtle.XCor() != 0 || turtle.YCor() != 0 {
		t.Errorf("Expected turtle to stay in place")
	}

	p1.SetOwn("chemical", 9.0)
	turtle.Downhill("chemical")

	// make sure that the turtle's position has changed to the patch with the lowest chemical value
	if turtle.XCor() != 0 || turtle.YCor() != 1 {
		t.Errorf("Expected turtle to move to patch (0, 1), got (%v, %v)", turtle.XCor(), turtle.YCor())
	}
}

func TestTurtleDownhill4(t *testing.T) {

	// create a model with patches own of chemical

	patchesOwn := map[string]interface{}{
		"chemical": 0.0,
	}

	m := model.NewModel(patchesOwn, nil, nil, nil, nil, nil, false, false)

	// create a turtle
	m.CreateTurtles(1, "", nil)

	// get the 5 patches around the turtle
	p1 := m.Patch(0, 0)
	p2 := m.Patch(0, 1)
	p3 := m.Patch(0, -1)
	p4 := m.Patch(1, 0)
	p5 := m.Patch(-1, 0)

	// set the chemical value of the patches
	p1.SetOwn("chemical", 0.0)
	p2.SetOwn("chemical", 1.0)
	p3.SetOwn("chemical", 2.0)
	p4.SetOwn("chemical", 3.0)
	p5.SetOwn("chemical", 4.0)

	turtle := m.Turtle("", 0)
	turtle.Downhill4("chemical")

	// make sure that the turtle's position has not changed since the patch it is on has the lowest chemical value
	if turtle.XCor() != 0 || turtle.YCor() != 0 {
		t.Errorf("Expected turtle to stay in place")
	}

	p1.SetOwn("chemical", 4.0)
	turtle.Downhill4("chemical")

	// make sure that the turtle's position has changed to the patch with the lowest chemical value
	if turtle.XCor() != 0 || turtle.YCor() != 1 {
		t.Errorf("Expected turtle to move to patch (0, 1), got (%v, %v)", turtle.XCor(), turtle.YCor())
	}
}

func TestTurtleFaceTurtle(t *testing.T) {

	// create a basic model with no wrapping

	m := model.NewModel(nil, nil, nil, nil, nil, nil, false, false)

	// create two turtles
	m.CreateTurtles(2, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	t1.SetXY(-14, 14)
	t2.SetXY(14, -14)

	t1.FaceTurtle(t2)

	// make sure that t1 is facing t2
	if t1.GetHeading() != 135 {
		t.Errorf("Expected turtle to face 45 degrees, got %v", t1.GetHeading())
	}

	m.WrappingXOn()
	m.WrappingYOn()

	t1.FaceTurtle(t2)

	// expect the turtle to face -45 degrees because of wrapping
	if t1.GetHeading() != -45 {
		t.Errorf("Expected turtle to face -45 degrees, got %v", t1.GetHeading())
	}

	m.WrappingXOff()
	m.WrappingYOff()

	t1.SetXY(-14, -5)
	t2.SetXY(14, 5)

	t1.FaceTurtle(t2)

	if t1.GetHeading()-70.3461759419467 > .00001 {
		t.Errorf("Expected turtle to face 0 degrees, got %v", t1.GetHeading())
	}

	m.WrappingXOn()
	m.WrappingYOn()

	t1.FaceTurtle(t2)

	if t1.GetHeading()+16.69924423399362 > .000001 {
		t.Errorf("Expected turtle to face -16.69924423399362 degrees, got %v", t1.GetHeading())
	}

}

func TestTurtleInLinkNeighbor(t *testing.T) {

	// create a new model
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false, false)

	// create some turtles
	m.CreateTurtles(8, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)
	t3 := m.Turtle("", 2)
	t4 := m.Turtle("", 3)
	t5 := m.Turtle("", 4)
	t6 := m.Turtle("", 5)
	t7 := m.Turtle("", 6)
	t8 := m.Turtle("", 7)

	// create a directed link between t1 and t2
	t1.CreateLinkToTurtle("parent-children", t2, nil)

	// create an undirected link between t3 and t4
	t3.CreateLinkWithTurtle("coworkers", t4, nil)

	// create a directed link between t5 and t6 that has no breed
	t5.CreateLinkToTurtle("", t6, nil)

	// create an undirected link between t7 and t8 that has no breed
	t7.CreateLinkWithTurtle("", t8, nil)

	v := t1.InLinkNeighbor("parent-children", t2)
	if v {
		t.Errorf("Expected turtle to not be a neighbor")
	}

	v = t2.InLinkNeighbor("parent-children", t1)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t3.InLinkNeighbor("coworkers", t4)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t4.InLinkNeighbor("coworkers", t3)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t5.InLinkNeighbor("", t6)
	if v {
		t.Errorf("Expected turtle to not be a neighbor")
	}

	v = t6.InLinkNeighbor("", t5)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t7.InLinkNeighbor("", t8)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t8.InLinkNeighbor("", t7)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.InLinkNeighbor("", t2)
	if v {
		t.Errorf("Expected turtle to not be a neighbor")
	}

	v = t2.InLinkNeighbor("", t1)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

}

// tests the Turtle.InLinkNeighbors function which returns a list of turtles that either have a directed link to the turtle or an undirected link
func TestTurtleInLinkNeighbors(t *testing.T) {

	// create a new model
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false, false)

	// create some turtles
	m.CreateTurtles(9, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)
	t3 := m.Turtle("", 2)
	t4 := m.Turtle("", 3)
	t5 := m.Turtle("", 4)
	t6 := m.Turtle("", 5)
	t7 := m.Turtle("", 6)
	t8 := m.Turtle("", 7)
	t9 := m.Turtle("", 8)

	/// breeded directed link to t1
	t2.CreateLinkToTurtle("parent-children", t1, nil)
	t3.CreateLinkToTurtle("parent-children", t1, nil)

	// breeded undirected link to t1
	t4.CreateLinkWithTurtle("coworkers", t1, nil)
	t5.CreateLinkWithTurtle("coworkers", t1, nil)

	// directed link to t1
	t6.CreateLinkToTurtle("", t1, nil)
	t7.CreateLinkToTurtle("", t1, nil)

	// undirected link to t1
	t8.CreateLinkWithTurtle("", t1, nil)
	t9.CreateLinkWithTurtle("", t1, nil)

	neighbors := t1.InLinkNeighbors("parent-children")
	if neighbors.Count() != 2 {
		t.Errorf("Expected 2 neighbors, got %d", neighbors.Count())
	}
	if !neighbors.Contains(t2) && !neighbors.Contains(t3) {
		t.Errorf("Expected neighbors to contain t2 and t3")
	}

	neighbors = t1.InLinkNeighbors("coworkers")
	if neighbors.Count() != 2 {
		t.Errorf("Expected 2 neighbors, got %d", neighbors.Count())
	}
	if !neighbors.Contains(t4) && !neighbors.Contains(t5) {
		t.Errorf("Expected neighbors to contain t4 and t5")
	}

	neighbors = t1.InLinkNeighbors("")
	if neighbors.Count() != 8 {
		t.Errorf("Expected 8 neighbors, got %d", neighbors.Count())
	}

}

// tests the Turtle.LinkNeighbor which returns whether there is any link connecting the two turtles
// link can be directed or undirected
func TestTurtleLinkNeighbor(t *testing.T) {

	// create a new model
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false, false)

	// create some turtles
	m.CreateTurtles(9, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)
	t3 := m.Turtle("", 2)
	t4 := m.Turtle("", 3)
	t5 := m.Turtle("", 4)
	t6 := m.Turtle("", 5)
	t7 := m.Turtle("", 6)
	t8 := m.Turtle("", 7)
	t9 := m.Turtle("", 8)

	/// breeded directed link to t1
	t2.CreateLinkToTurtle("parent-children", t1, nil)
	t3.CreateLinkToTurtle("parent-children", t1, nil)

	// breeded undirected link to t1
	t4.CreateLinkWithTurtle("coworkers", t1, nil)
	t5.CreateLinkWithTurtle("coworkers", t1, nil)

	// directed link to t1
	t6.CreateLinkToTurtle("", t1, nil)
	t7.CreateLinkToTurtle("", t1, nil)

	// undirected link to t1
	t8.CreateLinkWithTurtle("", t1, nil)
	t9.CreateLinkWithTurtle("", t1, nil)

	v := t1.LinkNeighbor("", t2)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.LinkNeighbor("parent-children", t3)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.LinkNeighbor("", t4)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.LinkNeighbor("parent-childres", t5)
	if v {
		t.Errorf("Expected turtle to not be a neighbor for that breed")
	}

	v = t1.LinkNeighbor("", t6)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.LinkNeighbor("parent-children", t7)
	if v {
		t.Errorf("Expected turtle to not be a neighbor for that breed")
	}

	v = t1.LinkNeighbor("", t8)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.LinkNeighbor("coworkers", t9)
	if v {
		t.Errorf("Expected turtle to not be a neighbor for that breed")
	}
}

// returns the turtles that are linked to the turtle
// can be directed or undirected, in either incoming or outgoing
func TestTurtleLinkNeighbors(t *testing.T) {

	// create a new model
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false, false)

	// create some turtles
	m.CreateTurtles(9, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)
	t3 := m.Turtle("", 2)
	t4 := m.Turtle("", 3)
	t5 := m.Turtle("", 4)
	t6 := m.Turtle("", 5)
	t7 := m.Turtle("", 6)
	t8 := m.Turtle("", 7)
	t9 := m.Turtle("", 8)

	/// breeded directed link to t1
	t2.CreateLinkToTurtle("parent-children", t1, nil)
	t3.CreateLinkToTurtle("parent-children", t1, nil)

	// breeded undirected link to t1
	t4.CreateLinkWithTurtle("coworkers", t1, nil)
	t5.CreateLinkWithTurtle("coworkers", t1, nil)

	// directed link to t1
	t6.CreateLinkToTurtle("", t1, nil)
	t7.CreateLinkToTurtle("", t1, nil)

	// undirected link to t1
	t8.CreateLinkWithTurtle("", t1, nil)
	t9.CreateLinkWithTurtle("", t1, nil)

	neighbors := t1.LinkNeighbors("")
	if neighbors.Count() != 8 {
		t.Errorf("Expected 4 neighbors, got %d", neighbors.Count())
	}

	neighbors = t1.LinkNeighbors("parent-children")
	if neighbors.Count() != 2 {
		t.Errorf("Expected 2 neighbors, got %d", neighbors.Count())
	}

	neighbors = t1.LinkNeighbors("coworkers")
	if neighbors.Count() != 2 {
		t.Errorf("Expected 2 neighbors, got %d", neighbors.Count())
	}

	neighbors = t2.LinkNeighbors("parent-children")
	if neighbors.Count() != 1 {
		t.Errorf("Expected 1 neighbor, got %d", neighbors.Count())
	}
	if !neighbors.Contains(t1) {
		t.Errorf("Expected neighbors to contain t1")
	}
}

// tests the Turtle.OutLinkNeighbor function which returns whether there is a link from the turtle to another turtle
// link can be directed or undirected
// but unlike LinkNeighbor, this function only checks for outgoing links
func TestTurtleOutLinkNeighbor(t *testing.T) {

	// create a new model
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false, false)

	// create some turtles
	m.CreateTurtles(9, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)
	t3 := m.Turtle("", 2)
	t4 := m.Turtle("", 3)
	t5 := m.Turtle("", 4)
	t6 := m.Turtle("", 5)
	t7 := m.Turtle("", 6)
	t8 := m.Turtle("", 7)
	t9 := m.Turtle("", 8)

	/// breeded directed link to t1
	t1.CreateLinkToTurtle("parent-children", t2, nil)
	t1.CreateLinkToTurtle("parent-children", t3, nil)

	// breeded undirected link to t1
	t1.CreateLinkWithTurtle("coworkers", t4, nil)
	t1.CreateLinkWithTurtle("coworkers", t5, nil)

	// directed link to t1
	t1.CreateLinkToTurtle("", t6, nil)
	t1.CreateLinkToTurtle("", t7, nil)

	// undirected link to t1
	t1.CreateLinkWithTurtle("", t8, nil)
	t1.CreateLinkWithTurtle("", t9, nil)

	v := t1.OutLinkNeighbor("", t2)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.OutLinkNeighbor("parent-children", t3)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.OutLinkNeighbor("", t4)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.OutLinkNeighbor("parent-childres", t5)
	if v {
		t.Errorf("Expected turtle to not be a neighbor for that breed")
	}

	v = t1.OutLinkNeighbor("", t6)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.OutLinkNeighbor("parent-children", t7)
	if v {
		t.Errorf("Expected turtle to not be a neighbor for that breed")
	}

	v = t1.OutLinkNeighbor("", t8)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.OutLinkNeighbor("coworkers", t9)
	if v {
		t.Errorf("Expected turtle to not be a neighbor for that breed")
	}

	v = t2.OutLinkNeighbor("parent-children", t1)
	if v {
		t.Errorf("Expected turtle to not be a neighbor since it is a directed link")
	}
}

// tests the Turtle.OutLinkNeighbors function which returns a list of turtles that the turtle has a directed or undirected link to
// like LinkNeighbors, but only checks for outgoing links
func TestTurtleOutLinkNeighbors(t *testing.T) {

	// create a new model
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false, false)

	// create some turtles
	m.CreateTurtles(9, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)
	t3 := m.Turtle("", 2)
	t4 := m.Turtle("", 3)
	t5 := m.Turtle("", 4)
	t6 := m.Turtle("", 5)
	t7 := m.Turtle("", 6)
	t8 := m.Turtle("", 7)
	t9 := m.Turtle("", 8)

	/// breeded directed link to t1
	t1.CreateLinkToTurtle("parent-children", t2, nil)
	t1.CreateLinkToTurtle("parent-children", t3, nil)

	// breeded undirected link to t1
	t1.CreateLinkWithTurtle("coworkers", t4, nil)
	t1.CreateLinkWithTurtle("coworkers", t5, nil)

	// directed link to t1
	t1.CreateLinkToTurtle("", t6, nil)
	t1.CreateLinkToTurtle("", t7, nil)

	// undirected link to t1
	t1.CreateLinkWithTurtle("", t8, nil)
	t1.CreateLinkWithTurtle("", t9, nil)

	neighbors := t1.OutLinkNeighbors("")
	if neighbors.Count() != 8 {
		t.Errorf("Expected 4 neighbors, got %d", neighbors.Count())
	}

	neighbors = t1.OutLinkNeighbors("parent-children")
	if neighbors.Count() != 2 {
		t.Errorf("Expected 2 neighbors, got %d", neighbors.Count())
	}

	neighbors = t1.OutLinkNeighbors("coworkers")
	if neighbors.Count() != 2 {
		t.Errorf("Expected 2 neighbors, got %d", neighbors.Count())
	}

	neighbors = t2.OutLinkNeighbors("parent-children")
	if neighbors.Count() != 0 {
		t.Errorf("Expected 0 neighbors, got %d", neighbors.Count())
	}
}

func TestTurtleMyLinks(t *testing.T) {

	// create a basic model with breeds for undirected and directed links
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false, false)

	// create some turtles
	m.CreateTurtles(9, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)
	t3 := m.Turtle("", 2)
	t4 := m.Turtle("", 3)
	t5 := m.Turtle("", 4)
	t6 := m.Turtle("", 5)
	t7 := m.Turtle("", 6)
	t8 := m.Turtle("", 7)
	t9 := m.Turtle("", 8)

	// directed unbreeded link from t1 to t2
	t1.CreateLinkToTurtle("", t2, nil)

	// directed breeded link from t1 to t3
	t1.CreateLinkToTurtle("parent-children", t3, nil)

	// undirected unbreeded link from t1 to t4
	t1.CreateLinkWithTurtle("", t4, nil)

	// undirected breeded link from t1 to t5
	t1.CreateLinkWithTurtle("coworkers", t5, nil)

	// directed unbreeded link from t6 to t1
	t1.CreateLinkFromTurtle("", t6, nil)

	// directed breeded link from t7 to t1
	t1.CreateLinkFromTurtle("parent-children", t7, nil)

	// undirected unbreeded link from t8 to t1
	t1.CreateLinkWithTurtle("", t8, nil)

	// undirected breeded link from t9 to t1
	t1.CreateLinkWithTurtle("coworkers", t9, nil)

	l1 := t1.LinkTo("", t2)
	l2 := t1.LinkTo("parent-children", t3)
	l3 := t1.LinkWith("", t4)
	l4 := t1.LinkWith("coworkers", t5)
	l5 := t1.LinkFrom("", t6)
	l6 := t1.LinkFrom("parent-children", t7)
	l7 := t1.LinkWith("", t8)
	l8 := t1.LinkWith("coworkers", t9)

	links := t1.MyLinks("")
	if links.Count() != 8 {
		t.Errorf("Expected 8 links, got %d", links.Count())
	}
	if !links.Contains(l1) || !links.Contains(l2) || !links.Contains(l3) || !links.Contains(l4) || !links.Contains(l5) || !links.Contains(l6) || !links.Contains(l7) || !links.Contains(l8) {
		t.Errorf("Expected links to contain all links")
	}

	links = t1.MyLinks("parent-children")
	if links.Count() != 2 {
		t.Errorf("Expected 2 links, got %d", links.Count())
	}
	if !links.Contains(l2) || !links.Contains(l6) {
		t.Errorf("Expected links to contain l2 and l6")
	}

	links = t1.MyLinks("coworkers")
	if links.Count() != 2 {
		t.Errorf("Expected 2 links, got %d", links.Count())
	}
	if !links.Contains(l4) || !links.Contains(l8) {
		t.Errorf("Expected links to contain l4 and l8")
	}

	links = t1.MyInLinks("")
	if links.Count() != 6 {
		t.Errorf("Expected 6 links, got %d", links.Count())
	}
	if !links.Contains(l3) || !links.Contains(l4) || !links.Contains(l5) || !links.Contains(l6) || !links.Contains(l7) || !links.Contains(l8) {
		t.Errorf("Expected links to contain all links")
	}

	links = t1.MyInLinks("parent-children")
	if links.Count() != 1 {
		t.Errorf("Expected 1 link, got %d", links.Count())
	}
	if !links.Contains(l6) {
		t.Errorf("Expected links to contain l6")
	}

	links = t1.MyInLinks("coworkers")
	if links.Count() != 2 {
		t.Errorf("Expected 2 links, got %d", links.Count())
	}
	if !links.Contains(l4) || !links.Contains(l8) {
		t.Errorf("Expected links to contain l4 and l8")
	}

	links = t1.MyOutLinks("")
	if links.Count() != 6 {
		t.Errorf("Expected 6 links, got %d", links.Count())
	}
	if !links.Contains(l1) || !links.Contains(l2) || !links.Contains(l3) || !links.Contains(l4) || !links.Contains(l7) || !links.Contains(l8) {
		t.Errorf("Expected links to contain all links")
	}

	links = t1.MyOutLinks("parent-children")
	if links.Count() != 1 {
		t.Errorf("Expected 1 link, got %d", links.Count())
	}
	if !links.Contains(l2) {
		t.Errorf("Expected links to contain l2")
	}

	links = t1.MyOutLinks("coworkers")
	if links.Count() != 2 {
		t.Errorf("Expected 2 links, got %d", links.Count())
	}
	if !links.Contains(l4) || !links.Contains(l8) {
		t.Errorf("Expected links to contain l4 and l8")
	}
}
