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

	if turtle.Breed().Name != "ants" {
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
