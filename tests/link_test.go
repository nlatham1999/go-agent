package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/pkg/model"
)

func TestLinkCreation(t *testing.T) {

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

	// create a new link that will pass
	l, err := t1.CreateLinkToTurtle(parentChildren, t2, nil)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure it exist in the general link list
	if !m.Links().Contains(l) {
		t.Errorf("Link should have been added to the general link list")
	}

	// make sure it exist in the general directed link list
	if !m.DirectedLinks().Contains(l) {
		t.Errorf("Link should have been added to the general directed link list")
	}

	// make sure it exist in the directed link list for the breed
	if !parentChildren.Links().Contains(l) {
		t.Errorf("Link should have been added to the directed link list for the breed")
	}
}

func TestLinkBreedName(t *testing.T) {

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

	if l.BreedName() != "parent-children" {
		t.Errorf("Breed name should be parent-children")
	}
}

func TestLinkBreed(t *testing.T) {

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

	if l.BreedName() != "parent-children" {
		t.Errorf("Breed name should be parent-children")
	}
}

func TestLinkSetBreed(t *testing.T) {

	parentChildren := model.NewLinkBreed("parent-children")
	personPet := model.NewLinkBreed("person-pet")
	coworkers := model.NewLinkBreed("coworkers")

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []*model.LinkBreed{parentChildren, personPet},
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

	// set the breed to coworkers
	l.SetBreed(coworkers)

	// breed should not be changed since it is not a valid breed for direced links
	if l.BreedName() != "parent-children" {
		t.Errorf("Breed name should be parent-children")
	}

	l.SetBreed(personPet)

	// breed should be changed since it is a valid breed for directed links
	if l.BreedName() != "person-pet" {
		t.Errorf("Breed name should be person-pet")
	}

	// make sure it no longer exists for parent-children
	if parentChildren.Links().Contains(l) {
		t.Errorf("Link should have been removed from parent-children")
	}

	// make sure it exists for person-pet
	if !personPet.Links().Contains(l) {
		t.Errorf("Link should have been added to person-pet")
	}
}

func TestLinkBothEnds(t *testing.T) {

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

	if l.OtherEnd(t1) != t2 {
		t.Errorf("Other end should be turtle 2")
	}

	if l.OtherEnd(t2) != t1 {
		t.Errorf("Other end should be turtle 1")
	}
}

// make sure that when a links is created that is a duplicate, that it returns an error
func TestLinkCreationNoDuplicates(t *testing.T) {

	parentChildren := model.NewLinkBreed("parent-children")
	personPet := model.NewLinkBreed("person-pet")
	coworkers := model.NewLinkBreed("coworkers")

	// create a new model
	settings := model.ModelSettings{
		DirectedLinkBreeds:   []*model.LinkBreed{parentChildren, personPet},
		UndirectedLinkBreeds: []*model.LinkBreed{coworkers},
	}
	m := model.NewModel(settings)

	// create some turtles
	m.CreateTurtles(2, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)

	// create a new link
	_, err := t1.CreateLinkToTurtle(parentChildren, t2, nil)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// create a new link
	_, err = t1.CreateLinkToTurtle(parentChildren, t2, nil)
	if err == nil {
		t.Errorf("Error should not be nil")
	}

	// do the same for undirected links
	_, err = t1.CreateLinkWithTurtle(coworkers, t2, nil)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// create a new link
	_, err = t1.CreateLinkWithTurtle(coworkers, t2, nil)
	if err == nil {
		t.Errorf("Error should not be nil")
	}
}

func TestLinkBreedSetting(t *testing.T) {
	//create breeded link betwen turtles switch breed and make sure link under the old breed does not exist any more for either turtle

	a := model.NewLinkBreed("a")
	b := model.NewLinkBreed("b")
	c := model.NewLinkBreed("c")
	d := model.NewLinkBreed("d")

	modelSettings := model.ModelSettings{
		UndirectedLinkBreeds: []*model.LinkBreed{a, b},
		DirectedLinkBreeds:   []*model.LinkBreed{c, d},
	}

	m := model.NewModel(modelSettings)

	m.CreateTurtles(4, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)

	l, err := t1.CreateLinkWithTurtle(a, t2, nil)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	if t1.Links(a).Count() != 1 {
		t.Errorf("Turtle 1 should have 1 link of breed a")
	}

	if t2.Links(a).Count() != 1 {
		t.Errorf("Turtle 2 should have 1 link of breed a")
	}

	l.SetBreed(b)

	if a.Links().Count() != 0 {
		t.Errorf("Link should not exist under breed a")
	}

	if b.Links().Count() != 1 {
		t.Errorf("Link should exist under breed b")
	}

	if t1.Links(a).Count() != 0 {
		t.Errorf("Turtle 1 should have 0 links of breed a")
	}

	if t2.Links(a).Count() != 0 {
		t.Errorf("Turtle 2 should have 0 links of breed a")
	}

	if t1.Links(b).Count() != 1 {
		t.Errorf("Turtle 1 should have 1 link of breed b")
	}

	if t2.Links(b).Count() != 1 {
		t.Errorf("Turtle 2 should have 1 link of breed b")
	}

	t3 := m.Turtle(2)
	t4 := m.Turtle(3)

	l, err = t3.CreateLinkToTurtle(c, t4, nil)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	if t3.OutLinks(c).Count() != 1 {
		t.Errorf("Turtle 3 should have 1 out link of breed c")
	}

	if t4.InLinks(c).Count() != 1 {
		t.Errorf("Turtle 4 should have 1 in link of breed c")
	}

	l.SetBreed(d)

	if c.Links().Count() != 0 {
		t.Errorf("Link should not exist under breed c")
	}

	if d.Links().Count() != 1 {
		t.Errorf("Link should exist under breed d")
	}

	if t3.OutLinks(c).Count() != 0 {
		t.Errorf("Turtle 3 should have 0 out links of breed c")
	}

	if t4.InLinks(c).Count() != 0 {
		t.Errorf("Turtle 4 should have 0 in links of breed c")
	}

	if t3.OutLinks(d).Count() != 1 {
		t.Errorf("Turtle 3 should have 1 out link of breed d")
	}

	if t4.InLinks(d).Count() != 1 {
		t.Errorf("Turtle 4 should have 1 in link of breed d")
	}
}

func TestLinkDying(t *testing.T) {
	//create a link and kill it, make sure it is removed from the model and the turtles

	a := model.NewLinkBreed("a")
	b := model.NewLinkBreed("b")
	c := model.NewLinkBreed("c")
	d := model.NewLinkBreed("d")

	modelSettings := model.ModelSettings{
		UndirectedLinkBreeds: []*model.LinkBreed{a, b},
		DirectedLinkBreeds:   []*model.LinkBreed{c, d},
	}

	m := model.NewModel(modelSettings)

	m.CreateTurtles(5, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)
	t3 := m.Turtle(2)
	t4 := m.Turtle(3)
	t5 := m.Turtle(4)

	t1.CreateLinkWithTurtle(a, t2, nil)
	t1.CreateLinkWithTurtle(a, t3, nil)
	t1.CreateLinkWithTurtle(a, t4, nil)
	t1.CreateLinkWithTurtle(a, t5, nil)
	t2.CreateLinkWithTurtle(a, t3, nil)
	t2.CreateLinkWithTurtle(a, t4, nil)
	t2.CreateLinkWithTurtle(a, t5, nil)
	t3.CreateLinkWithTurtle(a, t4, nil)
	t3.CreateLinkWithTurtle(a, t5, nil)
	t4.CreateLinkWithTurtle(a, t5, nil)

	if m.Links().Count() != 10 {
		t.Errorf("Model should have 4 links")
	}

	if a.Links().Count() != 10 {
		t.Errorf("Model should have 4 links of breed a")
	}

	for i, link := range a.Links().List() {
		if i == 0 {
			link.Die()
		}
	}

	if m.Links().Count() != 9 {
		t.Errorf("Model should have 9 links, got %d", m.Links().Count())
	}

	if len(m.Links().List()) != 9 {
		t.Errorf("Model should have 9 links, got %d", len(m.Links().List()))
	}
}
