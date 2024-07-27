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
	l := t2.InLinkFrom("parent-children", t1)

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
	l := t2.InLinkFrom("parent-children", t1)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link exists from t1 to t3
	l = t3.InLinkFrom("parent-children", t1)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link does not exist from t2 to t3
	l = t3.InLinkFrom("parent-children", t2)

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
	l := t2.InLinkFrom("coworkers", t1)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link exists from t2 to t1
	l = t1.InLinkFrom("coworkers", t2)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link does not exist from t1 to t2 for parent-children
	l = t2.InLinkFrom("parent-children", t1)

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
	l := t2.InLinkFrom("coworkers", t1)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link exists from t1 to t3
	l = t3.InLinkFrom("coworkers", t1)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link exists from t2 to t1
	l = t1.InLinkFrom("coworkers", t2)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link exists from t3 to t1
	l = t1.InLinkFrom("coworkers", t3)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link does not exist from t1 to t2 for parent-children
	l = t2.InLinkFrom("parent-children", t1)

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
	l := t1.InLinkFrom("parent-children", t2)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link does not exist from t1 to t2
	l = t2.InLinkFrom("parent-children", t1)

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
	l := t1.InLinkFrom("parent-children", t2)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link exists from t3 to t1
	l = t1.InLinkFrom("parent-children", t3)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link does not exist from t1 to t2
	l = t2.InLinkFrom("parent-children", t1)

	if l != nil {
		t.Errorf("Link should not have been created")
	}

	// make sure the link does not exist from t1 to t3
	l = t3.InLinkFrom("parent-children", t1)

	if l != nil {
		t.Errorf("Link should not have been created")
	}

	// make sure the link does not exist from t1 to t2 for parent-children
	l = t2.InLinkFrom("coworkers", t1)

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
	l := model.NewLink(m, "parent-children", t1, t2, true)

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
