package universe

import "math"

type TurtleAgentSet struct {
	turtles []*Turtle
}

// @TODO implement
func TurtleSet(turtles []*Turtle) *TurtleAgentSet {
	return nil
}

func (t *TurtleAgentSet) All(operation TurtleBoolOperation) bool {
	for _, turtle := range t.turtles {
		if !operation(turtle) {
			return false
		}
	}
	return true
}

func (t *TurtleAgentSet) Any(operation TurtleBoolOperation) bool {
	for _, turtle := range t.turtles {
		if operation(turtle) {
			return true
		}
	}
	return false
}

// @TODO implement
func (t *TurtleAgentSet) MaxNOf(n int, operation TurtleFloatOperation) *TurtleAgentSet {
	return nil
}

func (t *TurtleAgentSet) MaxOneOf(operation TurtleFloatOperation) *Turtle {
	max := math.MaxFloat64 * -1
	var maxTurtle *Turtle
	for _, turtle := range t.turtles {
		if operation(turtle) > max {
			max = operation(turtle)
			maxTurtle = turtle
		}
	}
	return maxTurtle
}

func (t *TurtleAgentSet) MinNOf(n int, operation TurtleFloatOperation) *TurtleAgentSet {
	return nil
}

func (t *TurtleAgentSet) MinOneOf(operation TurtleFloatOperation) *Turtle {
	min := math.MaxFloat64
	var minTurtle *Turtle
	for _, turtle := range t.turtles {
		if operation(turtle) < min {
			min = operation(turtle)
			minTurtle = turtle
		}
	}
	return minTurtle
}

func (t *TurtleAgentSet) TurtlesOn(breed string) *TurtleAgentSet {
	return nil
}

// @TODO implement
func (t *TurtleAgentSet) UpToNOf(n int) *TurtleAgentSet {
	return nil
}

// @TODO implement
func (t *TurtleAgentSet) WhoAreNot(turtles *TurtleAgentSet) *TurtleAgentSet {
	return nil
}

// @TODO implement
func (t *TurtleAgentSet) WhoAreNotTurtle(turtle *Turtle) *TurtleAgentSet {
	return nil
}

func (t *TurtleAgentSet) With(operation TurtleBoolOperation) *TurtleAgentSet {
	turtles := make([]*Turtle, 0)
	for _, turtle := range t.turtles {
		if operation(turtle) {
			turtles = append(turtles, turtle)
		}
	}
	return TurtleSet(turtles)
}

func (t *TurtleAgentSet) WithMax(operation TurtleFloatOperation) *TurtleAgentSet {
	max := math.MaxFloat64 * -1
	for _, turtle := range t.turtles {
		if operation(turtle) > max {
			max = operation(turtle)
		}
	}

	//get all turtles where the float operation is equal to the max
	turtles := make([]*Turtle, 0)
	for _, turtle := range t.turtles {
		if operation(turtle) == max {
			turtles = append(turtles, turtle)
		}
	}

	return TurtleSet(turtles)
}

func (t *TurtleAgentSet) WithMin(operation TurtleFloatOperation) *TurtleAgentSet {
	min := math.MaxFloat64
	for _, turtle := range t.turtles {
		if operation(turtle) < min {
			min = operation(turtle)
		}
	}

	//get all turtles where the float operation is equal to the min
	turtles := make([]*Turtle, 0)
	for _, turtle := range t.turtles {
		if operation(turtle) == min {
			turtles = append(turtles, turtle)
		}
	}

	return TurtleSet(turtles)
}
