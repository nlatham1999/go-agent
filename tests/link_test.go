package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/internal/model"
)

func TestLinkCreation(t *testing.T) {

	// create a new model
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false)

	// create some turtles
	m.CreateTurtles(2, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	// create a new link that will fail
	l := model.NewLink(m, "parent-children3", t1, t2, true)

	if l != nil {
		t.Errorf("Link should not have been created")
	}

	// create a new link that will pass
	l = model.NewLink(m, "parent-children", t1, t2, true)

	if l == nil {
		t.Errorf("Link should have been created")
	}

	// make sure it exist in the general link list
	if !m.Links.Contains(l) {
		t.Errorf("Link should have been added to the general link list")
	}

	// make sure it exist in the general directed link list
	if !m.DirectedLinkBreeds[""].Links.Contains(l) {
		t.Errorf("Link should have been added to the general directed link list")
	}

	// make sure it exist in the directed link list for the breed
	if !m.DirectedLinkBreeds["parent-children"].Links.Contains(l) {
		t.Errorf("Link should have been added to the directed link list for the breed")
	}
}

func TestLinkBreedName(t *testing.T) {

	// create a new model
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false)

	// create some turtles
	m.CreateTurtles(2, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	// create a new link
	l := model.NewLink(m, "parent-children", t1, t2, true)

	if l.BreedName() != "parent-children" {
		t.Errorf("Breed name should be parent-children")
	}
}

func TestLinkBreed(t *testing.T) {

	// create a new model
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false)

	// create some turtles
	m.CreateTurtles(2, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	// create a new link
	l := model.NewLink(m, "parent-children", t1, t2, true)

	if l.Breed().Name() != "parent-children" {
		t.Errorf("Breed name should be parent-children")
	}
}

func TestLinkSetBreed(t *testing.T) {

	// create a new model
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children", "person-pet"}, []string{"coworkers"}, false)

	// create some turtles
	m.CreateTurtles(2, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	// create a new link
	l := model.NewLink(m, "parent-children", t1, t2, true)

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
	if m.DirectedLinkBreeds["parent-children"].Links.Contains(l) {
		t.Errorf("Link should have been removed from parent-children")
	}

	// make sure it exists for person-pet
	if !m.DirectedLinkBreeds["person-pet"].Links.Contains(l) {
		t.Errorf("Link should have been added to person-pet")
	}
}

func TestLinkCreateLinkToTurtle(t *testing.T) {

	// create a new model
	m := model.NewModel(nil, nil, nil, nil, []string{"parent-children"}, []string{"coworkers"}, false)

	// create some turtles
	m.CreateTurtles(2, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)

	// create a new link
	t1.CreateLinkTo("parent-children", t2, nil)

	// make sure the link exists from t1 to t2
	l := t2.InLinkFrom("parent-children", t1)

	if l == nil {
		t.Errorf("Link should have been created")
	}
}
