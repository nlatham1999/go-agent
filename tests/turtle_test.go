package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/pkg/model"
)

func TestTurtleBack(t *testing.T) {

	//create a basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	//create a turtle
	m.CreateTurtles(1, nil)
	turtle := m.Turtle(0)

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

	ants := model.NewTurtleBreed("ants", "", nil)

	//create a basic model
	settings := model.ModelSettings{
		TurtleBreeds: []*model.TurtleBreed{ants},
	}
	_ = model.NewModel(settings)

	//create a turtle
	ants.CreateAgents(1, nil)

	turtle := ants.Agent(0)
	//assert that the turtle's breed is "ants"
	if turtle.BreedName() != "ants" {
		t.Errorf("Expected turtle to have breed 'ants'")
	}
}

func TestTurtleBreed(t *testing.T) {

	ants := model.NewTurtleBreed("ants", "", nil)

	//create a basic model
	settings := model.ModelSettings{
		TurtleBreeds: []*model.TurtleBreed{ants},
	}
	_ = model.NewModel(settings)

	//create a turtle
	ants.CreateAgents(1, nil)

	turtle := ants.Agent(0)

	if turtle.BreedName() == "" {
		t.Errorf("Expected turtle to have a breed")
	}

	if turtle.BreedName() != "ants" {
		t.Errorf("Expected turtle to have breed 'ants'")
	}
}

func TestTurtleSetBreed(t *testing.T) {

	ants := model.NewTurtleBreed("ants", "", nil)
	beetles := model.NewTurtleBreed("beetles", "", nil)

	//create a basic model
	settings := model.ModelSettings{
		TurtleBreeds: []*model.TurtleBreed{ants, beetles},
	}
	m := model.NewModel(settings)

	//create a turtle
	ants.CreateAgents(5, nil)

	turtle := ants.Agent(0)

	turtle.SetBreed(beetles)

	if turtle.BreedName() != "beetles" {
		t.Errorf("Expected turtle to have breed 'beetles'")
	}

	turtle = ants.Agent(0)

	if turtle != nil {
		t.Errorf("Expected turtle to not exist in breed 'ants'")
	}

	turtle = beetles.Agent(0)

	if turtle == nil {
		t.Errorf("Expected turtle to exist in breed 'beetles'")
	}

	m.CreateTurtles(1, nil)

	turtle = m.Turtle(5)

	if turtle.BreedName() != "" {
		t.Errorf("Expected turtle to have breed ''")
	}

	turtle.SetBreed(ants)

	if turtle.BreedName() != "ants" {
		t.Errorf("Expected turtle to have breed 'ants'")
	}

	turtle = ants.Agent(5)

	if turtle == nil {
		t.Errorf("Expected turtle to exist in breed 'ants'")
	}
}

func TestTurtlesOwn(t *testing.T) {

	// create a basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	// create a turtle
	m.CreateTurtles(1, nil)

	turtle := m.Turtle(0)

	// get the turtle's own
	mood := turtle.GetProperty("mood")

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

	ants := model.NewTurtleBreed("ants", "", antsOwn)

	settings = model.ModelSettings{
		TurtleProperties: turtlesOwn,
		TurtleBreeds:     []*model.TurtleBreed{ants},
	}
	_ = model.NewModel(settings)

	// create a turtle
	ants.CreateAgents(1, nil)

	turtle = ants.Agent(0)

	// get the turtle's own
	mood = turtle.GetProperty("mood")

	if mood != 0 {
		t.Errorf("Expected turtle to have own 'mood' with value 0, got %v", mood)
	}

	age := turtle.GetProperty("age")

	if age != 5 {
		t.Errorf("Expected turtle to have own 'age' with value 5")
	}

	// test changing breed
	// create a second model with ants and beetles breeds
	beetlesOwn := map[string]interface{}{
		"wingspan": 10,
	}

	beetles := model.NewTurtleBreed("beetles", "", beetlesOwn)

	settings = model.ModelSettings{
		TurtleProperties: turtlesOwn,
		TurtleBreeds:     []*model.TurtleBreed{ants, beetles},
	}
	_ = model.NewModel(settings)

	// create a turtle
	beetles.CreateAgents(1, nil)

	turtle = beetles.Agent(0)

	//change the breed
	turtle.SetBreed(ants)

	wingspan := turtle.GetProperty("wingspan")

	if wingspan != nil {
		t.Errorf("Expected turtle to not have own 'wingspan'")
	}

	mood = turtle.GetProperty("mood")

	if mood == nil {
		t.Errorf("Expected turtle to have own 'mood'")
	}

	if mood != 0 {
		t.Errorf("Expected turtle to have own 'mood' with value 0, got %v", mood)
	}

}

func TestTurtleCreateLinkToTurtle(t *testing.T) {

	parentChildren := model.NewLinkBreed("parent-children")
	coworkers := model.NewLinkBreed("coworkers")

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []*model.LinkBreed{parentChildren},
		UndirectedLinkBreeds: []*model.LinkBreed{coworkers},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(2, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)

	// create a new link
	t1.CreateLinkToTurtle(parentChildren, t2, nil)

	// make sure the link exists from t1 to t2
	l := t2.LinkFrom(parentChildren, t1)

	if l == nil {
		t.Errorf("Link should have been created")
	}
}

func TestTurtleCreateLinkToSet(t *testing.T) {

	parentChildren := model.NewLinkBreed("parent-children")
	coworkers := model.NewLinkBreed("coworkers")

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []*model.LinkBreed{parentChildren},
		UndirectedLinkBreeds: []*model.LinkBreed{coworkers},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(3, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)
	t3 := m.Turtle(2)

	agentSet := model.NewTurtleAgentSet([]*model.Turtle{t2, t3})

	t1.CreateLinksToSet(parentChildren, agentSet, nil)

	// make sure the link exists from t1 to t2
	l := t2.LinkFrom(parentChildren, t1)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link exists from t1 to t3
	l = t3.LinkFrom(parentChildren, t1)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link does not exist from t2 to t3
	l = t3.LinkFrom(parentChildren, t2)

	if l != nil {
		t.Errorf("Link should not have been created")
	}
}

func TestTurtleCreateLinkWithTurtle(t *testing.T) {

	parentChildren := model.NewLinkBreed("parent-children")
	coworkers := model.NewLinkBreed("coworkers")

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []*model.LinkBreed{parentChildren},
		UndirectedLinkBreeds: []*model.LinkBreed{coworkers},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(2, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)

	// create a new link
	t1.CreateLinkWithTurtle(coworkers, t2, nil)

	// make sure the link exists from t1 to t2
	l := t2.LinkFrom(coworkers, t1)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link exists from t2 to t1
	l = t1.LinkFrom(coworkers, t2)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link does not exist from t1 to t2 for parent-children
	l = t2.LinkFrom(parentChildren, t1)

	if l != nil {
		t.Errorf("Link should not have been created")
	}
}

func TestTurtleCreateLinkWithSet(t *testing.T) {

	parentChildren := model.NewLinkBreed("parent-children")
	coworkers := model.NewLinkBreed("coworkers")

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []*model.LinkBreed{parentChildren},
		UndirectedLinkBreeds: []*model.LinkBreed{coworkers},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(3, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)
	t3 := m.Turtle(2)

	agentSet := model.NewTurtleAgentSet([]*model.Turtle{t2, t3})

	t1.CreateLinksWithSet(coworkers, agentSet, nil)

	// make sure the link exists from t1 to t2
	l := t2.LinkFrom(coworkers, t1)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link exists from t1 to t3
	l = t3.LinkFrom(coworkers, t1)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link exists from t2 to t1
	l = t1.LinkFrom(coworkers, t2)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link exists from t3 to t1
	l = t1.LinkFrom(coworkers, t3)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link does not exist from t1 to t2 for parent-children
	l = t2.LinkFrom(parentChildren, t1)

	if l != nil {
		t.Errorf("Link should not have been created")
	}
}

func TestTurtleCreateLinkFromTurtle(t *testing.T) {

	parentChildren := model.NewLinkBreed("parent-children")
	coworkers := model.NewLinkBreed("coworkers")

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []*model.LinkBreed{parentChildren},
		UndirectedLinkBreeds: []*model.LinkBreed{coworkers},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(2, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)

	// create a new link
	t1.CreateLinkFromTurtle(parentChildren, t2, nil)

	// make sure the link exists from t2 to t1
	l := t1.LinkFrom(parentChildren, t2)
	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link does not exist from t1 to t2
	l = t2.LinkFrom(parentChildren, t1)

	if l != nil {
		t.Errorf("Link should not have been created")
	}
}

func TestTurtleCreateLinkFromSet(t *testing.T) {

	parentChildren := model.NewLinkBreed("parent-children")
	coworkers := model.NewLinkBreed("coworkers")

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []*model.LinkBreed{parentChildren},
		UndirectedLinkBreeds: []*model.LinkBreed{coworkers},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(3, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)
	t3 := m.Turtle(2)

	agentSet := model.NewTurtleAgentSet([]*model.Turtle{t2, t3})

	t1.CreateLinksFromSet(parentChildren, agentSet, nil)

	// make sure the link exists from t2 to t1
	l := t1.LinkFrom(parentChildren, t2)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link exists from t3 to t1
	l = t1.LinkFrom(parentChildren, t3)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure the link does not exist from t1 to t2
	l = t2.LinkFrom(parentChildren, t1)

	if l != nil {
		t.Errorf("Link should not have been created")
	}

	// make sure the link does not exist from t1 to t3
	l = t3.LinkFrom(parentChildren, t1)

	if l != nil {
		t.Errorf("Link should not have been created")
	}

	// make sure the link does not exist from t1 to t2 for parent-children
	l = t2.LinkFrom(coworkers, t1)

	if l != nil {
		t.Errorf("Link should not have been created")
	}
}

// test the Turtle.DistanceTurtle method
func TestTurtleDistanceTurtle(t *testing.T) {

	// create a basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	// create two turtles
	m.CreateTurtles(2, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)

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

// test DistanceTurtle with wrapping
func TestTurtleDistanceTurtleWrappingY(t *testing.T) {

	// create a basic model with wrapping
	settings := model.ModelSettings{
		WrappingX: true,
		WrappingY: true,
		MinPxCor:  -15,
		MinPyCor:  -15,
		MaxPxCor:  15,
		MaxPyCor:  15,
	}

	m := model.NewModel(settings)

	// create two turtles
	m.CreateTurtles(2, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)

	// set the position of the turtles
	t1.SetXY(0, -14)
	t2.SetXY(0, 14)

	// get the distance between the turtles
	distance := t1.DistanceTurtle(t2)

	// assert that the distance is 2
	if distance != 3 {
		t.Errorf("Expected distance to be 2, got %v", distance)
	}
}

// test DistanceTurtle with wrapping X
func TestTurtleDistanceTurtleWrappingX(t *testing.T) {

	// create a basic model with wrapping
	settings := model.ModelSettings{
		WrappingX: true,
		WrappingY: true,
		MinPxCor:  -15,
		MinPyCor:  -15,
		MaxPxCor:  15,
		MaxPyCor:  15,
	}

	m := model.NewModel(settings)

	// create two turtles
	m.CreateTurtles(2, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)

	// set the position of the turtles
	t1.SetXY(-14, 0)
	t2.SetXY(14, 0)

	// get the distance between the turtles
	distance := t1.DistanceTurtle(t2)

	// assert that the distance is 2
	if distance != 3 {
		t.Errorf("Expected distance to be 2, got %v", distance)
	}
}

// test DistanceTurtle with wrapping X and Y
func TestTurtleDistanceTurtleWrappingXY(t *testing.T) {

	// create a basic model with wrapping
	settings := model.ModelSettings{
		WrappingX: true,
		WrappingY: true,
		MinPxCor:  -15,
		MinPyCor:  -15,
		MaxPxCor:  15,
		MaxPyCor:  15,
	}

	m := model.NewModel(settings)

	// create two turtles
	m.CreateTurtles(2, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)

	// set the position of the turtles to be 3,4,5 triangle
	t1.SetXY(-13, -14)
	t2.SetXY(14, 14)

	// get the distance between the turtles
	distance := t1.DistanceTurtle(t2)

	// assert that the distance is 2
	if distance != 5 {
		t.Errorf("Expected distance to be 2, got %v", distance)
	}

	// set the position of the turtles to be 3,4,5 triangle
	t1.SetXY(-14, -14)
	t2.SetXY(13, 14)

	// get the distance between the turtles
	distance = t1.DistanceTurtle(t2)

	// assert that the distance is 2
	if distance != 5 {
		t.Errorf("Expected distance to be 2, got %v", distance)
	}

	// set the position of the turtles to be 3,4,5 triangle
	t1.SetXY(-14, -13)
	t2.SetXY(14, 14)

	// get the distance between the turtles
	distance = t1.DistanceTurtle(t2)

	// assert that the distance is 2
	if distance != 5 {
		t.Errorf("Expected distance to be 2, got %v", distance)
	}

	// set the position of the turtles to be 3,4,5 triangle
	t1.SetXY(-14, -14)
	t2.SetXY(14, 13)

	// get the distance between the turtles
	distance = t1.DistanceTurtle(t2)

	// assert that the distance is 2
	if distance != 5 {
		t.Errorf("Expected distance to be 2, got %v", distance)
	}
}

func TestTurtleDistancePatch(t *testing.T) {

	// create a basic model
	m := model.NewModel(model.ModelSettings{})

	// create a turtle
	m.CreateTurtles(1, nil)

	t1 := m.Turtle(0)

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
	m := model.NewModel(model.ModelSettings{})

	// create a turtle
	m.CreateTurtles(1, nil)

	t1 := m.Turtle(0)

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

	parentChildren := model.NewLinkBreed("parent-children")
	coworkers := model.NewLinkBreed("coworkers")

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []*model.LinkBreed{parentChildren},
		UndirectedLinkBreeds: []*model.LinkBreed{coworkers},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(2, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)

	// create a new link
	l, err := t1.CreateLinkToTurtle(parentChildren, t2, nil)
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
	settings := model.ModelSettings{
		WrappingX: true,
		WrappingY: true,
	}
	m := model.NewModel(settings)

	// create a turtle
	m.CreateTurtles(1, nil)

	turtle := m.Turtle(0)

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

	settings := model.ModelSettings{
		PatchProperties: patchesOwn,
	}
	m := model.NewModel(settings)

	// create a turtle
	m.CreateTurtles(1, nil)

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
	p1.SetProperty("chemical", 0.0)
	p2.SetProperty("chemical", 1.0)
	p3.SetProperty("chemical", 2.0)
	p4.SetProperty("chemical", 3.0)
	p5.SetProperty("chemical", 4.0)
	p6.SetProperty("chemical", 5.0)
	p7.SetProperty("chemical", 6.0)
	p8.SetProperty("chemical", 7.0)
	p9.SetProperty("chemical", 8.0)

	turtle := m.Turtle(0)
	turtle.Downhill("chemical")

	// make sure that the turtle's position has not changed since the patch it is on has the lowest chemical value
	if turtle.XCor() != 0 || turtle.YCor() != 0 {
		t.Errorf("Expected turtle to stay in place")
	}

	p1.SetProperty("chemical", 9.0)
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

	settings := model.ModelSettings{
		PatchProperties: patchesOwn,
	}
	m := model.NewModel(settings)

	// create a turtle
	m.CreateTurtles(1, nil)

	// get the 5 patches around the turtle
	p1 := m.Patch(0, 0)
	p2 := m.Patch(0, 1)
	p3 := m.Patch(0, -1)
	p4 := m.Patch(1, 0)
	p5 := m.Patch(-1, 0)

	// set the chemical value of the patches
	p1.SetProperty("chemical", 0.0)
	p2.SetProperty("chemical", 1.0)
	p3.SetProperty("chemical", 2.0)
	p4.SetProperty("chemical", 3.0)
	p5.SetProperty("chemical", 4.0)

	turtle := m.Turtle(0)
	turtle.Downhill4("chemical")

	// make sure that the turtle's position has not changed since the patch it is on has the lowest chemical value
	if turtle.XCor() != 0 || turtle.YCor() != 0 {
		t.Errorf("Expected turtle to stay in place")
	}

	p1.SetProperty("chemical", 4.0)
	turtle.Downhill4("chemical")

	// make sure that the turtle's position has changed to the patch with the lowest chemical value
	if turtle.XCor() != 0 || turtle.YCor() != 1 {
		t.Errorf("Expected turtle to move to patch (0, 1), got (%v, %v)", turtle.XCor(), turtle.YCor())
	}
}

func TestTurtleFaceTurtle(t *testing.T) {

	// create a basic model with no wrapping

	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	// create two turtles
	m.CreateTurtles(2, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)

	t1.SetXY(-14, 14)
	t2.SetXY(14, -14)

	t1.FaceTurtle(t2)

	// make sure that t1 is facing t2
	if t1.GetHeading() != 315 {
		t.Errorf("Expected turtle to face 315 degrees, got %v", t1.GetHeading())
	}

	m.WrappingXOn()
	m.WrappingYOn()

	t1.FaceTurtle(t2)

	// expect the turtle to face -45 degrees because of wrapping
	if t1.GetHeading() != 135 {
		t.Errorf("Expected turtle to face 135 degrees, got %v", t1.GetHeading())
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

	if t1.GetHeading()-343.3007557660064 > .000001 {
		t.Errorf("Expected turtle to face -16.69924423399362 degrees, got %v", t1.GetHeading())
	}

}

func TestFaceXY(t *testing.T) {

	// create a basic model with no wrapping
	m := model.NewModel(model.ModelSettings{})

	// create a turtle
	m.CreateTurtles(1,
		func(t *model.Turtle) {
			t.SetXY(0, 5)
			t.SetHeading(270)
		},
	)

	if m.Turtle(0).GetHeading() != 270 {
		t.Errorf("Expected turtle to face 270 degrees, got %v", m.Turtle(0).GetHeading())
	}

	// face the turtle towards the point (5, 0)
	m.Turtle(0).FaceXY(5, 0)

	// make sure the turtle is facing 0 degrees
	if m.Turtle(0).GetHeading() != 315 {
		t.Errorf("Expected turtle to face 315 degrees, got %v", m.Turtle(0).GetHeading())
	}

}

func TestTurtleInLinkNeighbor(t *testing.T) {

	parentChildren := model.NewLinkBreed("parent-children")
	coworkers := model.NewLinkBreed("coworkers")

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []*model.LinkBreed{parentChildren},
		UndirectedLinkBreeds: []*model.LinkBreed{coworkers},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(8, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)
	t3 := m.Turtle(2)
	t4 := m.Turtle(3)
	t5 := m.Turtle(4)
	t6 := m.Turtle(5)
	t7 := m.Turtle(6)
	t8 := m.Turtle(7)

	// create a directed link between t1 and t2
	t1.CreateLinkToTurtle(parentChildren, t2, nil)

	// create an undirected link between t3 and t4
	t3.CreateLinkWithTurtle(coworkers, t4, nil)

	// create a directed link between t5 and t6 that has no breed
	t5.CreateLinkToTurtle(nil, t6, nil)

	// create an undirected link between t7 and t8 that has no breed
	t7.CreateLinkWithTurtle(nil, t8, nil)

	v := t1.LinkToTurtleExists(parentChildren, t2)
	if v {
		t.Errorf("Expected turtle to not be a neighbor")
	}

	v = t2.LinkToTurtleExists(parentChildren, t1)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t3.LinkToTurtleExists(coworkers, t4)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t4.LinkToTurtleExists(coworkers, t3)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t5.LinkToTurtleExists(nil, t6)
	if v {
		t.Errorf("Expected turtle to not be a neighbor")
	}

	v = t6.LinkToTurtleExists(nil, t5)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t7.LinkToTurtleExists(nil, t8)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t8.LinkToTurtleExists(nil, t7)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.LinkToTurtleExists(nil, t2)
	if v {
		t.Errorf("Expected turtle to not be a neighbor")
	}

	v = t2.LinkToTurtleExists(nil, t1)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

}

// tests the Turtle.InLinkNeighbors function which returns a list of turtles that either have a directed link to the turtle or an undirected link
func TestTurtleInLinkNeighbors(t *testing.T) {

	parentChildren := model.NewLinkBreed("parent-children")
	coworkers := model.NewLinkBreed("coworkers")

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []*model.LinkBreed{parentChildren},
		UndirectedLinkBreeds: []*model.LinkBreed{coworkers},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(9, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)
	t3 := m.Turtle(2)
	t4 := m.Turtle(3)
	t5 := m.Turtle(4)
	t6 := m.Turtle(5)
	t7 := m.Turtle(6)
	t8 := m.Turtle(7)
	t9 := m.Turtle(8)

	/// breeded directed link to t1
	t2.CreateLinkToTurtle(parentChildren, t1, nil)
	t3.CreateLinkToTurtle(parentChildren, t1, nil)

	// breeded undirected link to t1
	t4.CreateLinkWithTurtle(coworkers, t1, nil)
	t5.CreateLinkWithTurtle(coworkers, t1, nil)

	// directed link to t1
	t6.CreateLinkToTurtle(nil, t1, nil)
	t7.CreateLinkToTurtle(nil, t1, nil)

	// undirected link to t1
	t8.CreateLinkWithTurtle(nil, t1, nil)
	t9.CreateLinkWithTurtle(nil, t1, nil)

	neighbors := t1.LinkedTurtlesToThis(parentChildren)
	if neighbors.Count() != 2 {
		t.Errorf("Expected 2 neighbors, got %d", neighbors.Count())
	}
	if !neighbors.Contains(t2) && !neighbors.Contains(t3) {
		t.Errorf("Expected neighbors to contain t2 and t3")
	}

	neighbors = t1.LinkedTurtlesToThis(coworkers)
	if neighbors.Count() != 2 {
		t.Errorf("Expected 2 neighbors, got %d", neighbors.Count())
	}
	if !neighbors.Contains(t4) && !neighbors.Contains(t5) {
		t.Errorf("Expected neighbors to contain t4 and t5")
	}

	neighbors = t1.LinkedTurtlesToThis(nil)
	if neighbors.Count() != 8 {
		t.Errorf("Expected 8 neighbors, got %d", neighbors.Count())
	}

}

// tests the Turtle.LinkNeighbor which returns whether there is any link connecting the two turtles
// link can be directed or undirected
func TestTurtleLinkNeighbor(t *testing.T) {

	parentChildren := model.NewLinkBreed("parent-children")
	coworkers := model.NewLinkBreed("coworkers")

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []*model.LinkBreed{parentChildren},
		UndirectedLinkBreeds: []*model.LinkBreed{coworkers},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(9, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)
	t3 := m.Turtle(2)
	t4 := m.Turtle(3)
	t5 := m.Turtle(4)
	t6 := m.Turtle(5)
	t7 := m.Turtle(6)
	t8 := m.Turtle(7)
	t9 := m.Turtle(8)

	/// breeded directed link to t1
	t2.CreateLinkToTurtle(parentChildren, t1, nil)
	t3.CreateLinkToTurtle(parentChildren, t1, nil)

	// breeded undirected link to t1
	t4.CreateLinkWithTurtle(coworkers, t1, nil)
	t5.CreateLinkWithTurtle(coworkers, t1, nil)

	// directed link to t1
	t6.CreateLinkToTurtle(nil, t1, nil)
	t7.CreateLinkToTurtle(nil, t1, nil)

	// undirected link to t1
	t8.CreateLinkWithTurtle(nil, t1, nil)
	t9.CreateLinkWithTurtle(nil, t1, nil)

	v := t1.LinkExists(nil, t2)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.LinkExists(parentChildren, t3)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.LinkExists(nil, t4)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.LinkExists(nil, t6)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.LinkExists(parentChildren, t7)
	if v {
		t.Errorf("Expected turtle to not be a neighbor for that breed")
	}

	v = t1.LinkExists(nil, t8)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.LinkExists(coworkers, t9)
	if v {
		t.Errorf("Expected turtle to not be a neighbor for that breed")
	}
}

// returns the turtles that are linked to the turtle
// can be directed or undirected, in either incoming or outgoing
func TestTurtleLinkNeighbors(t *testing.T) {

	parentChildren := model.NewLinkBreed("parent-children")
	coworkers := model.NewLinkBreed("coworkers")

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []*model.LinkBreed{parentChildren},
		UndirectedLinkBreeds: []*model.LinkBreed{coworkers},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(9, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)
	t3 := m.Turtle(2)
	t4 := m.Turtle(3)
	t5 := m.Turtle(4)
	t6 := m.Turtle(5)
	t7 := m.Turtle(6)
	t8 := m.Turtle(7)
	t9 := m.Turtle(8)

	/// breeded directed link to t1
	t2.CreateLinkToTurtle(parentChildren, t1, nil)
	t3.CreateLinkToTurtle(parentChildren, t1, nil)

	// breeded undirected link to t1
	t4.CreateLinkWithTurtle(coworkers, t1, nil)
	t5.CreateLinkWithTurtle(coworkers, t1, nil)

	// directed link to t1
	t6.CreateLinkToTurtle(nil, t1, nil)
	t7.CreateLinkToTurtle(nil, t1, nil)

	// undirected link to t1
	t8.CreateLinkWithTurtle(nil, t1, nil)
	t9.CreateLinkWithTurtle(nil, t1, nil)

	neighbors := t1.LinkedTurtles(nil)
	if neighbors.Count() != 8 {
		t.Errorf("Expected 4 neighbors, got %d", neighbors.Count())
	}

	neighbors = t1.LinkedTurtles(parentChildren)
	if neighbors.Count() != 2 {
		t.Errorf("Expected 2 neighbors, got %d", neighbors.Count())
	}

	neighbors = t1.LinkedTurtles(coworkers)
	if neighbors.Count() != 2 {
		t.Errorf("Expected 2 neighbors, got %d", neighbors.Count())
	}

	neighbors = t2.LinkedTurtles(parentChildren)
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

	parentChildren := model.NewLinkBreed("parent-children")
	coworkers := model.NewLinkBreed("coworkers")

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []*model.LinkBreed{parentChildren},
		UndirectedLinkBreeds: []*model.LinkBreed{coworkers},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(9, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)
	t3 := m.Turtle(2)
	t4 := m.Turtle(3)
	t5 := m.Turtle(4)
	t6 := m.Turtle(5)
	t7 := m.Turtle(6)
	t8 := m.Turtle(7)
	t9 := m.Turtle(8)

	/// breeded directed link to t1
	t1.CreateLinkToTurtle(parentChildren, t2, nil)
	t1.CreateLinkToTurtle(parentChildren, t3, nil)

	// breeded undirected link to t1
	t1.CreateLinkWithTurtle(coworkers, t4, nil)
	t1.CreateLinkWithTurtle(coworkers, t5, nil)

	// directed link to t1
	t1.CreateLinkToTurtle(nil, t6, nil)
	t1.CreateLinkToTurtle(nil, t7, nil)

	// undirected link to t1
	t1.CreateLinkWithTurtle(nil, t8, nil)
	t1.CreateLinkWithTurtle(nil, t9, nil)

	v := t1.LinkFromTurtleExists(nil, t2)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.LinkFromTurtleExists(parentChildren, t3)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.LinkFromTurtleExists(nil, t4)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.LinkFromTurtleExists(nil, t6)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.LinkFromTurtleExists(parentChildren, t7)
	if v {
		t.Errorf("Expected turtle to not be a neighbor for that breed")
	}

	v = t1.LinkFromTurtleExists(nil, t8)
	if !v {
		t.Errorf("Expected turtle to be a neighbor")
	}

	v = t1.LinkFromTurtleExists(coworkers, t9)
	if v {
		t.Errorf("Expected turtle to not be a neighbor for that breed")
	}

	v = t2.LinkFromTurtleExists(parentChildren, t1)
	if v {
		t.Errorf("Expected turtle to not be a neighbor since it is a directed link")
	}
}

// tests the Turtle.OutLinkNeighbors function which returns a list of turtles that the turtle has a directed or undirected link to
// like LinkNeighbors, but only checks for outgoing links
func TestTurtleOutLinkNeighbors(t *testing.T) {

	parentChildren := model.NewLinkBreed("parent-children")
	coworkers := model.NewLinkBreed("coworkers")

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []*model.LinkBreed{parentChildren},
		UndirectedLinkBreeds: []*model.LinkBreed{coworkers},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(9, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)
	t3 := m.Turtle(2)
	t4 := m.Turtle(3)
	t5 := m.Turtle(4)
	t6 := m.Turtle(5)
	t7 := m.Turtle(6)
	t8 := m.Turtle(7)
	t9 := m.Turtle(8)

	/// breeded directed link to t1
	t1.CreateLinkToTurtle(parentChildren, t2, nil)
	t1.CreateLinkToTurtle(parentChildren, t3, nil)

	// breeded undirected link to t1
	t1.CreateLinkWithTurtle(coworkers, t4, nil)
	t1.CreateLinkWithTurtle(coworkers, t5, nil)

	// directed link to t1
	t1.CreateLinkToTurtle(nil, t6, nil)
	t1.CreateLinkToTurtle(nil, t7, nil)

	// undirected link to t1
	t1.CreateLinkWithTurtle(nil, t8, nil)
	t1.CreateLinkWithTurtle(nil, t9, nil)

	neighbors := t1.LinkedTurtlesFromThis(nil)
	if neighbors.Count() != 8 {
		t.Errorf("Expected 4 neighbors, got %d", neighbors.Count())
	}

	neighbors = t1.LinkedTurtlesFromThis(parentChildren)
	if neighbors.Count() != 2 {
		t.Errorf("Expected 2 neighbors, got %d", neighbors.Count())
	}

	neighbors = t1.LinkedTurtlesFromThis(coworkers)
	if neighbors.Count() != 2 {
		t.Errorf("Expected 2 neighbors, got %d", neighbors.Count())
	}

	neighbors = t2.LinkedTurtlesFromThis(parentChildren)
	if neighbors.Count() != 0 {
		t.Errorf("Expected 0 neighbors, got %d", neighbors.Count())
	}
}

func TestTurtleMyLinks(t *testing.T) {

	parentChildren := model.NewLinkBreed("parent-children")
	coworkers := model.NewLinkBreed("coworkers")

	// create a basic model with breeds for undirected and directed links
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []*model.LinkBreed{parentChildren},
		UndirectedLinkBreeds: []*model.LinkBreed{coworkers},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(9, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)
	t3 := m.Turtle(2)
	t4 := m.Turtle(3)
	t5 := m.Turtle(4)
	t6 := m.Turtle(5)
	t7 := m.Turtle(6)
	t8 := m.Turtle(7)
	t9 := m.Turtle(8)

	// directed unbreeded link from t1 to t2
	t1.CreateLinkToTurtle(nil, t2, nil)

	// directed breeded link from t1 to t3
	t1.CreateLinkToTurtle(parentChildren, t3, nil)

	// undirected unbreeded link from t1 to t4
	t1.CreateLinkWithTurtle(nil, t4, nil)

	// undirected breeded link from t1 to t5
	t1.CreateLinkWithTurtle(coworkers, t5, nil)

	// directed unbreeded link from t6 to t1
	t1.CreateLinkFromTurtle(nil, t6, nil)

	// directed breeded link from t7 to t1
	t1.CreateLinkFromTurtle(parentChildren, t7, nil)

	// undirected unbreeded link from t8 to t1
	t1.CreateLinkWithTurtle(nil, t8, nil)

	// undirected breeded link from t9 to t1
	t1.CreateLinkWithTurtle(coworkers, t9, nil)

	l1 := t1.LinkTo(nil, t2)
	l2 := t1.LinkTo(parentChildren, t3)
	l3 := t1.LinkWith(nil, t4)
	l4 := t1.LinkWith(coworkers, t5)
	l5 := t1.LinkFrom(nil, t6)
	l6 := t1.LinkFrom(parentChildren, t7)
	l7 := t1.LinkWith(nil, t8)
	l8 := t1.LinkWith(coworkers, t9)

	links := t1.Links(nil)
	if links.Count() != 8 {
		t.Errorf("Expected 8 links, got %d", links.Count())
	}
	if !links.Contains(l1) || !links.Contains(l2) || !links.Contains(l3) || !links.Contains(l4) || !links.Contains(l5) || !links.Contains(l6) || !links.Contains(l7) || !links.Contains(l8) {
		t.Errorf("Expected links to contain all links")
	}

	links = t1.Links(parentChildren)
	if links.Count() != 2 {
		t.Errorf("Expected 2 links, got %d", links.Count())
	}
	if !links.Contains(l2) || !links.Contains(l6) {
		t.Errorf("Expected links to contain l2 and l6")
	}

	links = t1.Links(coworkers)
	if links.Count() != 2 {
		t.Errorf("Expected 2 links, got %d", links.Count())
	}
	if !links.Contains(l4) || !links.Contains(l8) {
		t.Errorf("Expected links to contain l4 and l8")
	}

	links = t1.InLinks(nil)
	if links.Count() != 6 {
		t.Errorf("Expected 6 links, got %d", links.Count())
	}
	if !links.Contains(l3) || !links.Contains(l4) || !links.Contains(l5) || !links.Contains(l6) || !links.Contains(l7) || !links.Contains(l8) {
		t.Errorf("Expected links to contain all links")
	}

	links = t1.InLinks(parentChildren)
	if links.Count() != 1 {
		t.Errorf("Expected 1 link, got %d", links.Count())
	}
	if !links.Contains(l6) {
		t.Errorf("Expected links to contain l6")
	}

	links = t1.InLinks(coworkers)
	if links.Count() != 2 {
		t.Errorf("Expected 2 links, got %d", links.Count())
	}
	if !links.Contains(l4) || !links.Contains(l8) {
		t.Errorf("Expected links to contain l4 and l8")
	}

	links = t1.OutLinks(nil)
	if links.Count() != 6 {
		t.Errorf("Expected 6 links, got %d", links.Count())
	}
	if !links.Contains(l1) || !links.Contains(l2) || !links.Contains(l3) || !links.Contains(l4) || !links.Contains(l7) || !links.Contains(l8) {
		t.Errorf("Expected links to contain all links")
	}

	links = t1.OutLinks(parentChildren)
	if links.Count() != 1 {
		t.Errorf("Expected 1 link, got %d", links.Count())
	}
	if !links.Contains(l2) {
		t.Errorf("Expected links to contain l2")
	}

	links = t1.OutLinks(coworkers)
	if links.Count() != 2 {
		t.Errorf("Expected 2 links, got %d", links.Count())
	}
	if !links.Contains(l4) || !links.Contains(l8) {
		t.Errorf("Expected links to contain l4 and l8")
	}
}

func TestTurtleTurtlesHere(t *testing.T) {

	ants := model.NewTurtleBreed("ants", "", nil)

	// create a basic model with an ants breed for turtles
	settings := model.ModelSettings{
		TurtleBreeds: []*model.TurtleBreed{ants},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(2, nil)
	ants.CreateAgents(2, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)

	t3 := ants.Agent(2)
	t4 := ants.Agent(3)

	t1.SetXY(1, 1)
	t2.SetXY(2, 2)
	t3.SetXY(2, 2)
	t4.SetXY(3, 3)

	turtles := t1.TurtlesHere()
	if turtles.Count() != 1 {
		t.Errorf("Expected 1 turtle, got %d", turtles.Count())
	}
	if !turtles.Contains(t1) {
		t.Errorf("Expected turtles to contain t1")
	}

	turtles = ants.AgentsOnPatch(t1.PatchHere())
	// turtles = t1.TurtlesHere("ants")
	if turtles.Count() != 0 {
		t.Errorf("Expected 0 turtles, got %d", turtles.Count())
	}

	turtles = t2.TurtlesHere()
	if turtles.Count() != 2 {
		t.Errorf("Expected 2 turtles, got %d", turtles.Count())
	}
	if !turtles.Contains(t2) || !turtles.Contains(t3) {
		t.Errorf("Expected turtles to contain t2 and t3")
	}

	turtles = ants.AgentsWithAgent(t2)
	// turtles = ants.TurtlesOnPatch(t2.PatchHere())
	// turtles = t2.TurtlesHere("ants")
	if turtles.Count() != 1 {
		t.Errorf("Expected 1 turtle, got %d", turtles.Count())
	}
	if !turtles.Contains(t3) {
		t.Errorf("Expected turtles to contain t3")
	}

	turtles = ants.AgentsOnPatch(t3.PatchHere())
	// turtles = t3.TurtlesHere("ants")
	if turtles.Count() != 1 {
		t.Errorf("Expected 1 turtle, got %d", turtles.Count())
	}
	if !turtles.Contains(t3) {
		t.Errorf("Expected turtles to contain t3")
	}

	turtles = t4.TurtlesHere()
	if turtles.Count() != 1 {
		t.Errorf("Expected 1 turtle, got %d", turtles.Count())
	}
	if !turtles.Contains(t4) {
		t.Errorf("Expected turtles to contain t4")
	}
}

func TestTurtleUphill(t *testing.T) {

	// create a model with patches own of chemical

	patchesOwn := map[string]interface{}{
		"chemical": 0.0,
	}

	settings := model.ModelSettings{
		PatchProperties: patchesOwn,
	}
	m := model.NewModel(settings)

	// create a turtle
	m.CreateTurtles(1, nil)

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
	p1.SetProperty("chemical", 10.0)
	p2.SetProperty("chemical", 1.0)
	p3.SetProperty("chemical", 2.0)
	p4.SetProperty("chemical", 3.0)
	p5.SetProperty("chemical", 4.0)
	p6.SetProperty("chemical", 0.0)
	p7.SetProperty("chemical", 6.0)
	p8.SetProperty("chemical", 7.0)
	p9.SetProperty("chemical", 8.0)

	turtle := m.Turtle(0)
	turtle.Uphill("chemical")

	// make sure that the turtle's position has not changed since the patch it is on has the lowest chemical value
	if turtle.XCor() != 0 || turtle.YCor() != 0 {
		t.Errorf("Expected turtle to stay in place")
	}

	p1.SetProperty("chemical", -5.0)
	turtle.Uphill("chemical")

	// make sure that the turtle's position has changed to the patch with the lowest chemical value
	if turtle.XCor() != -1 || turtle.YCor() != -1 {
		t.Errorf("Expected turtle to move to patch (-1, -1), got (%v, %v)", turtle.XCor(), turtle.YCor())
	}
}

func TestTurtleUphill4(t *testing.T) {

	// create a model with patches own of chemical

	patchesOwn := map[string]interface{}{
		"chemical": 0.0,
	}

	settings := model.ModelSettings{
		PatchProperties: patchesOwn,
	}
	m := model.NewModel(settings)

	// create a turtle
	m.CreateTurtles(1, nil)

	// get the 5 patches around the turtle
	p1 := m.Patch(0, 0)
	p2 := m.Patch(0, 1)
	p3 := m.Patch(0, -1)
	p4 := m.Patch(1, 0)
	p5 := m.Patch(-1, 0)

	// set the chemical value of the patches
	p1.SetProperty("chemical", 14.0)
	p2.SetProperty("chemical", 11.0)
	p3.SetProperty("chemical", 2.0)
	p4.SetProperty("chemical", 3.0)
	p5.SetProperty("chemical", 4.0)

	turtle := m.Turtle(0)
	turtle.Uphill4("chemical")

	// make sure that the turtle's position has not changed since the patch it is on has the lowest chemical value
	if turtle.XCor() != 0 || turtle.YCor() != 0 {
		t.Errorf("Expected turtle to stay in place")
	}

	p1.SetProperty("chemical", 4.0)
	turtle.Uphill4("chemical")

	// make sure that the turtle's position has changed to the patch with the lowest chemical value
	if turtle.XCor() != 0 || turtle.YCor() != 1 {
		t.Errorf("Expected turtle to move to patch (0, 1), got (%v, %v)", turtle.XCor(), turtle.YCor())
	}
}

func TestTurtleTowardsXY(t *testing.T) {
	// create a basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	// create a turtle
	m.CreateTurtles(1, nil)

	turtle := m.Turtle(0)

	turtle.SetXY(0, 0)

	h := turtle.TowardsXY(1, -1)

	if h != 315 {
		t.Errorf("Expected heading to be 315 degrees, got %v", h)
	}
}

func TestTurtlesLinksDying(t *testing.T) {

	settings := model.ModelSettings{}

	m := model.NewModel(settings)

	m.CreateTurtles(6, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)
	t3 := m.Turtle(2)
	t4 := m.Turtle(3)
	t5 := m.Turtle(4)
	t6 := m.Turtle(5)

	t1.CreateLinkWithTurtle(nil, t2, nil)
	t1.Die()
	if t2.LinkedTurtles(nil).Count() != 0 {
		t.Errorf("Expected turtle to have no neighbors")
	}

	t3.CreateLinkToTurtle(nil, t4, nil)
	t3.Die()
	if t4.LinkedTurtles(nil).Count() != 0 {
		t.Errorf("Expected turtle to have no neighbors")
	}

	t5.CreateLinkFromTurtle(nil, t6, nil)
	t5.Die()
	if t6.LinkedTurtles(nil).Count() != 0 {
		t.Errorf("Expected turtle to have no neighbors")
	}

}

func TestTurtleSetBreedPatchHere(t *testing.T) {

	scouts := model.NewTurtleBreed("scouts", "", nil)
	foragers := model.NewTurtleBreed("foragers", "", nil)

	settings := model.ModelSettings{
		TurtleBreeds: []*model.TurtleBreed{scouts, foragers},
	}

	_ = model.NewModel(settings)

	// m.CreateTurtles(1, "scouts", nil)
	scouts.CreateAgents(1, nil)

	// t1 := m.Turtle("scouts", 0)
	t1 := scouts.Agent(0)

	if scouts.AgentsOnPatch(t1.PatchHere()).Count() != 1 {
		t.Errorf("Expected turtle to be on patch")
	}

	t1.SetBreed(foragers)

	if foragers.AgentsOnPatch(t1.PatchHere()).Count() != 1 {
		t.Errorf("Expected turtle to be on patch")
	}

}

func TestTurtle_Jump(t *testing.T) {
	// create a basic model
	settings := model.ModelSettings{
		WrappingX: false,
		WrappingY: false,
		MinPxCor:  -10,
		MaxPxCor:  10,
		MinPyCor:  -10,
		MaxPyCor:  10,
	}

	m := model.NewModel(settings)

	// create a turtle
	m.CreateTurtles(1, nil)

	turtle := m.Turtle(0)

	turtle.SetHeading(0)
	turtle.Jump(5)

	if turtle.XCor() != 5 || turtle.YCor() != 0 {
		t.Errorf("Expected turtle to jump to (5, 0), got (%v, %v)", turtle.XCor(), turtle.YCor())
	}

	// attempt a jump outside of the world bounds
	turtle.Jump(20)

	if turtle.XCor() != 5 || turtle.YCor() != 0 {
		t.Errorf("Expected turtle to stay in place")
	}

	// new model with wrapping
	settings = model.ModelSettings{
		WrappingX: true,
		WrappingY: true,
		MinPxCor:  -10,
		MaxPxCor:  10,
		MinPyCor:  -10,
		MaxPyCor:  10,
	}

	m = model.NewModel(settings)

	// create a turtle
	m.CreateTurtles(1, nil)

	turtle = m.Turtle(0)

	turtle.SetHeading(0)

	turtle.Jump(16) // extra 1 because of world edges

	if turtle.XCor() != -5 || turtle.YCor() != 0 {
		t.Errorf("Expected turtle to jump to (-5, 0), got (%v, %v)", turtle.XCor(), turtle.YCor())
	}
}

func TestTurtleMoveForwardDown(t *testing.T) {
	settings := model.ModelSettings{
		WrappingX:  true,
		WrappingY:  true,
		MinPxCor:   -10, // min x patch coordinate
		MaxPxCor:   10,  // max x patch coordinate
		MinPyCor:   -10, // min y patch coordinate
		MaxPyCor:   10,  // max y patch coordinate
		RandomSeed: 10,  // random seed
	}

	m := model.NewModel(settings) // create a new model with the settings

	m.ClearAll()

	m.CreateTurtles(1, func(t *model.Turtle) {
		t.SetHeading(270)
		t.SetSize(.5)
		t.SetXY(5.0, -10.4)
		t.Shape = "circle"
		t.Color = model.Green
	})

	t0 := m.Turtle(0)
	t0.Forward(1.0)

	if m.Turtle(0).XCor() != 5.0 || m.Turtle(0).YCor() != 9.6 {
		t.Errorf("Expected turtle to move to (5.0, 9.6), got (%v, %v)", m.Turtle(0).XCor(), m.Turtle(0).YCor())
	}
}

func TestTurtleMoveForwardUp(t *testing.T) {

	settings := model.ModelSettings{
		WrappingX:  true,
		WrappingY:  true,
		MinPxCor:   -10, // min x patch coordinate
		MaxPxCor:   10,  // max x patch coordinate
		MinPyCor:   -10, // min y patch coordinate
		MaxPyCor:   10,  // max y patch coordinate
		RandomSeed: 10,  // random seed
	}

	m := model.NewModel(settings) // create a new model with the settings

	m.ClearAll()

	m.CreateTurtles(1, func(t *model.Turtle) {
		t.SetHeading(90)
		t.SetSize(.5)
		t.SetXY(5.0, 10.4)
		t.Shape = "circle"
		t.Color = model.Green
	})

	t0 := m.Turtle(0)
	t0.Forward(1.0)

	if m.Turtle(0).XCor() != 5.0 || (m.Turtle(0).YCor()+9.6) > .001 {
		t.Errorf("Expected turtle to move to (5.0, -9.6), got (%v, %v)", m.Turtle(0).XCor(), m.Turtle(0).YCor())
	}
}

func TestTurtleMoveForwardLeft(t *testing.T) {

	settings := model.ModelSettings{
		WrappingX:  true,
		WrappingY:  true,
		MinPxCor:   -10, // min x patch coordinate
		MaxPxCor:   10,  // max x patch coordinate
		MinPyCor:   -10, // min y patch coordinate
		MaxPyCor:   10,  // max y patch coordinate
		RandomSeed: 10,  // random seed
	}

	m := model.NewModel(settings) // create a new model with the settings

	m.ClearAll()

	m.CreateTurtles(1, func(t *model.Turtle) {
		t.SetHeading(180)
		t.SetSize(.5)
		t.SetXY(-10.4, 5.0)
		t.Shape = "circle"
		t.Color = model.Green
	})

	t0 := m.Turtle(0)
	t0.Forward(1.0)

	if m.Turtle(0).XCor() != 9.6 || m.Turtle(0).YCor() != 5.0 {
		t.Errorf("Expected turtle to move to (-9.6, 5.0), got (%v, %v)", m.Turtle(0).XCor(), m.Turtle(0).YCor())
	}
}

func TestTurtleMoveForwardRight(t *testing.T) {

	settings := model.ModelSettings{
		WrappingX:  true,
		WrappingY:  true,
		MinPxCor:   -10, // min x patch coordinate
		MaxPxCor:   10,  // max x patch coordinate
		MinPyCor:   -10, // min y patch coordinate
		MaxPyCor:   10,  // max y patch coordinate
		RandomSeed: 10,  // random seed
	}

	m := model.NewModel(settings) // create a new model with the settings

	m.ClearAll()

	m.CreateTurtles(1, func(t *model.Turtle) {
		t.SetHeading(0)
		t.SetSize(.5)
		t.SetXY(10.4, 5.0)
		t.Shape = "circle"
		t.Color = model.Green
	})

	t0 := m.Turtle(0)
	t0.Forward(1.0)

	if m.Turtle(0).XCor()+9.6 > (.001) || m.Turtle(0).YCor() != 5.0 {
		t.Errorf("Expected turtle to move to (9.6, 5.0), got (%v, %v)", m.Turtle(0).XCor(), m.Turtle(0).YCor())
	}
}
