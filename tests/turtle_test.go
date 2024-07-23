package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/internal/model"
)

func TestTurtleBack(t *testing.T) {

	//create a basic model
	m := model.NewModel(nil, nil, nil, nil, nil, nil, false)

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
	m := model.NewModel(nil, nil, nil, []string{"ants"}, nil, nil, false)

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
	m := model.NewModel(nil, nil, nil, []string{"ants"}, nil, nil, false)

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
	m := model.NewModel(nil, nil, nil, []string{"ants", "beetles"}, nil, nil, false)

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
	m := model.NewModel(nil, nil, nil, nil, nil, nil, false)

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

	m2 := model.NewModel(nil, turtlesOwn, breedsOwn, []string{"ants"}, nil, nil, false)

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

	m3 := model.NewModel(nil, turtlesOwn, breedsOwn, []string{"ants", "beetles"}, nil, nil, false)

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
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false)

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
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false)

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
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false)

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
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false)

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
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false)

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
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false)

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
	m := model.NewModel(nil, nil, nil, nil, nil, nil, false)

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
	m := model.NewModel(nil, nil, nil, nil, nil, nil, false)

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
	m := model.NewModel(nil, nil, nil, nil, nil, nil, false)

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
