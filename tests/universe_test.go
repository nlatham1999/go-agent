package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/internal/model"
)

func TestCreateTurtles(t *testing.T) {
	breeds := []string{
		"ants",
	}

	environment := model.NewModel(nil, nil, nil, breeds, nil, nil, false, false)

	// creating turtles without a breed should add them to the default breed
	environment.CreateTurtles(5, "", nil)
	if environment.Turtles.Count() != 5 {
		t.Errorf("Expected 5 turtles, got %d", environment.Turtles.Count())
	}

	// creating turtles with a breed should add them to that breed and the default breed
	environment.CreateTurtles(5, "ants", nil)
	if environment.Turtles.Count() != 10 {
		t.Errorf("Expected 10 turtles, got %d", environment.Turtles.Count())
	}
	if environment.Breeds["ants"].Turtles.Count() != 5 {
		t.Errorf("Expected 5 ants, got %d", environment.Breeds["ants"].Turtles.Count())
	}

	// creating a turtle with a nonexistent breed should return an error
	err := environment.CreateTurtles(5, "nonexistent", nil)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestTurtle(t *testing.T) {
	breeds := []string{
		"ants",
	}

	environment := model.NewModel(nil, nil, nil, breeds, nil, nil, false, false)

	// create 5 general turtle and five ants
	environment.CreateTurtles(5, "", nil)
	environment.CreateTurtles(5, "ants", nil)

	// get turtle from general pop
	turtle := environment.Turtle("", 0)
	if turtle == nil {
		t.Errorf("Expected turtle, got nil")
	}

	// get turtle from ants pop when it should not exist
	turtle = environment.Turtle("ants", 0)
	if turtle != nil {
		t.Errorf("Expected nil, got turtle")
	}

	// get turtle from ants pop when it should exist
	turtle = environment.Turtle("ants", 5)
	if turtle == nil {
		t.Errorf("Expected turtle, got nil")
	}

	// get turtle from general pop when it should not exist
	turtle = environment.Turtle("", 12)
	if turtle != nil {
		t.Errorf("Expected nil, got turtle")
	}

	// get turtle that is an ant from general pop when it should exist
	turtle = environment.Turtle("", 7)
	if turtle == nil {
		t.Errorf("Expected turtle, got nil")
	}

	// get turtle that belongs to nonexistent breed
	turtle = environment.Turtle("elephants", 0)
	if turtle != nil {
		t.Errorf("Expected nil, got turtle")
	}
}

func TestClearTurtles(t *testing.T) {
	breeds := []string{
		"ants",
	}

	m := model.NewModel(nil, nil, nil, breeds, nil, nil, false, false)

	// create 5 general turtle and five ants
	m.CreateTurtles(5, "", nil)
	m.CreateTurtles(5, "ants", nil)

	ref := m.Turtle("", 0)
	ref.SetXY(1, 1)

	if m.Patch(0, 0).TurtlesHere("").Count() != 9 {
		t.Errorf("Expected 10 turtles, got %d", m.Patch(0, 0).TurtlesHere("").Count())
	}

	// clear general turtles
	m.ClearTurtles()
	if m.Turtles.Count() != 0 {
		t.Errorf("Expected 0 turtles, got %d", m.Turtles.Count())
	}

	if ref.XCor() != 0 {
		t.Errorf("Expected turtle to be reset")
	}

	t1 := m.Turtle("", 0)
	if t1 != nil {
		t.Errorf("Expected nil, got turtle")
	}

	p := m.Patch(0, 0)
	if p.TurtlesHere("").Count() != 0 {
		t.Errorf("Expected 0 turtles, got %d", p.TurtlesHere("").Count())
	}

	m.CreateTurtles(1, "", nil)

	t1 = m.Turtle("", 0)
	if t1 == nil {
		t.Errorf("Expected turtle, got nil")
	}

}

func TestKillTurtle(t *testing.T) {

	breeds := []string{
		"ants",
	}

	m := model.NewModel(nil, nil, nil, breeds, nil, nil, false, false)

	// create 5 general turtle and five ants
	m.CreateTurtles(5, "", nil)
	m.CreateTurtles(5, "ants", nil)

	ref := m.Turtle("", 0)
	ref.SetXY(1, 1)

	if m.Patch(0, 0).TurtlesHere("").Count() != 9 {
		t.Errorf("Expected 10 turtles, got %d", m.Patch(0, 0).TurtlesHere("").Count())
	}

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("ants", 5)
	t3 := m.Turtle("ants", 6)
	t4 := m.Turtle("ants", 7)

	_, err := model.NewLink(m, "", t1, t2, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}
	_, err = model.NewLink(m, "", t1, t3, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}
	_, err = model.NewLink(m, "", t1, t4, false)
	if err != nil {
		t.Errorf("Error should be nil")
	}
	_, err = model.NewLink(m, "", t2, t3, false)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// kill general turtle
	m.KillTurtle(m.Turtle("", 0))
	if m.Turtles.Count() != 9 {
		t.Errorf("Expected 9 turtles, got %d", m.Turtles.Count())
	}

	if ref.XCor() != 0 {
		t.Errorf("Expected turtle to be reset")
	}

	// make sure there's only one link left
	if m.Links.Count() != 1 {
		t.Errorf("Expected 1 link, got %d", m.Links.Count())
	}
}
