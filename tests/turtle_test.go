package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/internal/model"
)

func TestTurtleBack(t *testing.T) {

	//create a basic model
	m := model.NewModel(nil, nil, nil, nil, nil, nil, false)

	//create a turtle
	turtle := model.NewTurtle(m, 0, "")

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
