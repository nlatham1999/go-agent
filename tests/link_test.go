package tests

import (
	"math"
	"testing"

	"github.com/nlatham1999/go-agent/internal/model"
)

func TestLinkCreation(t *testing.T) {

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []string{"parent-children"},
		UndirectedLinkBreeds: []string{"coworkers"},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(2, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	// create a new link that will fail
	l, err := model.NewLink(m, "parent-children3", t1, t2, true)
	if err == nil {
		t.Errorf("Error should not be nil")
	}

	if l != nil {
		t.Errorf("Link should not have been created")
	}

	// create a new link that will pass
	l, err = model.NewLink(m, "parent-children", t1, t2, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure it exist in the general link list
	if !m.Links.Contains(l) {
		t.Errorf("Link should have been added to the general link list")
	}

	// make sure it exist in the general directed link list
	if !m.DirectedLinkBreeds[""].Links().Contains(l) {
		t.Errorf("Link should have been added to the general directed link list")
	}

	// make sure it exist in the directed link list for the breed
	if !m.DirectedLinkBreeds["parent-children"].Links().Contains(l) {
		t.Errorf("Link should have been added to the directed link list for the breed")
	}
}

func TestLinkBreedName(t *testing.T) {

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []string{"parent-children"},
		UndirectedLinkBreeds: []string{"coworkers"},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(2, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	// create a new link
	l, err := model.NewLink(m, "parent-children", t1, t2, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	if l.BreedName() != "parent-children" {
		t.Errorf("Breed name should be parent-children")
	}
}

func TestLinkBreed(t *testing.T) {

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []string{"parent-children"},
		UndirectedLinkBreeds: []string{"coworkers"},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(2, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	// create a new link
	l, err := model.NewLink(m, "parent-children", t1, t2, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	if l.Breed().Name() != "parent-children" {
		t.Errorf("Breed name should be parent-children")
	}
}

func TestLinkSetBreed(t *testing.T) {

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []string{"parent-children", "person-pet"},
		UndirectedLinkBreeds: []string{"coworkers"},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(2, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	// create a new link
	l, err := model.NewLink(m, "parent-children", t1, t2, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// set the breed to coworkers
	l.SetBreed("coworkers")

	// breed should not be changed since it is not a valid breed for direced links
	if l.BreedName() != "parent-children" {
		t.Errorf("Breed name should be parent-children")
	}

	l.SetBreed("person-pet")

	// breed should be changed since it is a valid breed for directed links
	if l.BreedName() != "person-pet" {
		t.Errorf("Breed name should be person-pet")
	}

	// make sure it no longer exists for parent-children
	if m.DirectedLinkBreeds["parent-children"].Links().Contains(l) {
		t.Errorf("Link should have been removed from parent-children")
	}

	// make sure it exists for person-pet
	if !m.DirectedLinkBreeds["person-pet"].Links().Contains(l) {
		t.Errorf("Link should have been added to person-pet")
	}
}

func TestLinkBothEnds(t *testing.T) {

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []string{"parent-children"},
		UndirectedLinkBreeds: []string{"coworkers"},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(2, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	// create a new link
	l, err := model.NewLink(m, "parent-children", t1, t2, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// get the both ends
	ends := l.BothEnds()

	// make sure both ends are in the set
	if !ends.Contains(t1) {
		t.Errorf("Link should have turtle 1")
	}

	if !ends.Contains(t2) {
		t.Errorf("Link should have turtle 2")
	}
}

func TestLinkHeading(t *testing.T) {

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []string{"parent-children"},
		UndirectedLinkBreeds: []string{"coworkers"},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(2, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	// create a new link
	l, err := model.NewLink(m, "parent-children", t1, t2, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	_, err = l.Heading()

	//err should not be nil since the turtles are at the same location
	if err == nil {
		t.Errorf("Error should not be nil, got %v", err)
	}

	t1.SetXY(0, 0)
	t2.SetXY(1, 1)

	t1.SetHeading(0)
	t2.SetHeading(90)

	heading, _ := l.Heading()

	if heading != 270 {
		t.Errorf("Heading should be 180, got %f", heading)
	}

	t1.SetHeading(90)
	t2.SetHeading(0)

	heading, _ = l.Heading()

	if heading != 90 {
		t.Errorf("Heading should be 90, got %f", heading)
	}

	t1.SetHeading(0)
	t2.SetHeading(450)

	heading, _ = l.Heading()

	if heading != 270 {
		t.Errorf("Heading should be 270, got %f", heading)
	}
}

func TestLinkOtherEnd(t *testing.T) {

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []string{"parent-children"},
		UndirectedLinkBreeds: []string{"coworkers"},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(2, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	// create a new link
	l, err := model.NewLink(m, "parent-children", t1, t2, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	if l.OtherEnd(t1) != t2 {
		t.Errorf("Other end should be turtle 2")
	}

	if l.OtherEnd(t2) != t1 {
		t.Errorf("Other end should be turtle 1")
	}
}

func TestRotatingTiedTurtles(t *testing.T) {

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []string{"parent-children", "person-pet"},
		UndirectedLinkBreeds: []string{"coworkers"},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(3, "", []model.TurtleOperation{
		func(t *model.Turtle) {
			t.SetHeading(0)
		},
	})

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)
	t3 := m.Turtle("", 2)

	t1.SetHeading(30)
	t1.SetXY(0, 0)

	t2.SetHeading(60)
	t2.SetXY(3, 4)

	t3.SetHeading(90)
	t3.SetXY(5, 1)

	// create a new link
	l1, err := model.NewLink(m, "parent-children", t1, t2, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}
	l1.TieMode = model.TieModeFixed

	l2, err := model.NewLink(m, "parent-children", t2, t3, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}
	l2.TieMode = model.TieModeFixed

	l3, err := model.NewLink(m, "parent-children", t3, t1, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}
	l3.TieMode = model.TieModeFixed

	// rotate the turtles
	t1.Right(10)

	// make sure the heading of each has increased by 10
	if t1.GetHeading() != 40 {
		t.Errorf("Turtle 1 heading should be 40, got %f", t1.GetHeading())
	}

	if t2.GetHeading() != 70 {
		t.Errorf("Turtle 2 heading should be 70, got %f", t2.GetHeading())
	}

	if t3.GetHeading() != 100 {
		t.Errorf("Turtle 3 heading should be 100, got %f", t3.GetHeading())
	}

	// make sure that t1 has not moved
	if t1.XCor() != 0 {
		t.Errorf("Turtle 1 x should be 0, got %f", t1.YCor())
	}

	if t1.YCor() != 0 {
		t.Errorf("Turtle 1 y should be 0, got %f", t1.XCor())
	}

	if t2.XCor() != 3.6490159697043456 || t2.YCor() != 3.4182864790480414 {
		t.Errorf("Turtle 2 x should be 3.6490159697043456 and y should be 3.4182864790480414, got %f and %f", t2.XCor(), t2.YCor())
	}

	if t3.XCor() != 5.097686942727971 || t3.YCor() != 0.11656686467755639 {
		t.Errorf("Turtle 3 x should be 5.097686942727971 and y should be 0.11656686467755639, got %f and %f", t3.XCor(), t3.YCor())
	}

	// rotate the second turtle
	t2.Left(20)

	if t1.GetHeading()-20 > .00001 {
		t.Errorf("Turtle 1 heading should be 20, got %f", t1.GetHeading())
	}

	if t2.GetHeading()-50 > .00001 {
		t.Errorf("Turtle 2 heading should be 50, got %f", t2.GetHeading())
	}

	if t3.GetHeading() != 80 {
		t.Errorf("Turtle 3 heading should be 80, got %f", t3.GetHeading())
	}

	// make sure that t2 has not moved
	if t2.XCor() != 3.6490159697043456 || t2.YCor() != 3.4182864790480414 {
		t.Errorf("Turtle 2 x should be 3.6490159697043456, got %f", t2.XCor())
	}

	if t1.XCor()-1.3891854213354429 > .0001 || t1.YCor()+1.0418890660015827 > .0001 {
		t.Errorf("Turtle 1 x should be 1.3891854213354429 and y should be 0=-1.0418890660015827, got %f and %f", t1.XCor(), t1.YCor())
	}

	if t3.XCor()-6.139576008729553 > .0001 || t3.YCor()-0.8111595753452776 > .0001 {
		t.Errorf("Turtle 3 x should be 6.139576008729553 and y should be 0.8111595753452776, got %f and %f", t3.XCor(), t3.YCor())
	}

	// create another turtle at 14, 14 that is fixed to t1
	m.CreateTurtles(2, "", nil)
	t4 := m.Turtle("", 3)
	t5 := m.Turtle("", 4)
	t4.SetXY(0, 0)
	t5.SetXY(14, 14)
	l4, err := model.NewLink(m, "parent-children", t4, t5, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}
	l4.TieMode = model.TieModeFixed

	// rotate t1
	t4.Right(20)

	// t5 should not move because it would be off the world
	if t5.XCor() != 14 || t5.YCor() != 14 {
		t.Errorf("Turtle 5 should not have moved, got %f and %f", t5.XCor(), t5.YCor())
	}

	// revert the rotation
	l4.TieMode = model.TieModeNone
	t4.Left(10)
	l4.TieMode = model.TieModeFixed

	m.WrappingXOn()

	t4.Right(20)

	if math.Abs(t5.XCor()+13.05602130243792) > .0001 || math.Abs(t5.YCor()-8.367414684443355) > .0001 {
		t.Errorf("Turtle 5 should not have moved, got %f and %f", t5.XCor(), t5.YCor())
	}

	// create another 2 turtles, t6 at 0,0 and t7 at -14, 14 that is fixed to t6
	m.CreateTurtles(2, "", nil)
	t6 := m.Turtle("", 5)
	t7 := m.Turtle("", 6)

	t6.SetXY(0, 0)
	t7.SetXY(-14, 14)

	l5, err := model.NewLink(m, "parent-children", t6, t7, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}
	l5.TieMode = model.TieModeFixed

	// rotate t6
	t6.Left(20)

	if t7.XCor()-14.05602130243792 > .0001 || t7.YCor()-8.367414684443355 > .0001 {
		t.Errorf("Turtle 7 should be at 14.05602130243792 and 8.367414684443355, got %f and %f", t7.XCor(), t7.YCor())
	}
}

// test that when a turtle is tied to another turtle, it moves with it, when the turtle moves forward, back or setxy
func TestMovingTiedTurtles(t *testing.T) {

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []string{"parent-children", "person-pet"},
		UndirectedLinkBreeds: []string{"coworkers"},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(2, "", []model.TurtleOperation{
		func(t *model.Turtle) {
			t.SetHeading(0)
		},
	})

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	t1.SetXY(0, 0)
	t2.SetXY(3, 4)

	// create a new link
	l, err := model.NewLink(m, "parent-children", t1, t2, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}
	l.TieMode = model.TieModeFixed

	// move t1 forward
	t1.Forward(5)

	if t2.XCor() != 8 || t2.YCor() != 4 {
		t.Errorf("Turtle 2 should have moved to 8, 4, got %f, %f", t2.XCor(), t2.YCor())
	}

	// move t1 back
	t1.Back(5)

	if t2.XCor() != 3 || t2.YCor() != 4 {
		t.Errorf("Turtle 2 should have moved to 3, 4, got %f, %f", t2.XCor(), t2.YCor())
	}

	// set t1 to 10, 10
	t1.SetXY(10, 10)

	if t2.XCor() != 13 || t2.YCor() != 14 {
		t.Errorf("Turtle 2 should have moved to 15, 10, got %f, %f", t2.XCor(), t2.YCor())
	}

	t1.SetXY(15, 15)

	// should not move since it would be off the world
	if t2.XCor() != 13 || t2.YCor() != 14 {
		t.Errorf("Turtle 2 should not have moved, got %f, %f", t2.XCor(), t2.YCor())
	}

	m.WrappingXOn()
	m.WrappingYOn()

	t1.SetXY(16, 16)

	if t2.XCor() != 14 || t2.YCor() != 15 {
		t.Errorf("Turtle 2 should have moved to 14, 15, got %f, %f", t2.XCor(), t2.YCor())
	}
}

// make sure that when a links is created that is a duplicate, that it returns an error
func TestLinkCreationNoDuplicates(t *testing.T) {

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []string{"parent-children", "person-pet"},
		UndirectedLinkBreeds: []string{"coworkers"},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(2, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	// create a new link
	_, err := model.NewLink(m, "parent-children", t1, t2, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// create a new link
	_, err = model.NewLink(m, "parent-children", t1, t2, true)
	if err == nil {
		t.Errorf("Error should not be nil")
	}

	// do the same for undirected links
	_, err = model.NewLink(m, "coworkers", t1, t2, false)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// create a new link
	_, err = model.NewLink(m, "coworkers", t1, t2, false)
	if err == nil {
		t.Errorf("Error should not be nil")
	}
}
