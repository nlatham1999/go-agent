package model

import "math"

type TurtleAgentSet struct {
	turtles map[*Turtle]interface{} //map of turtles so we can quickly check if a turtle is in the set
}

// @TODO implement
func TurtleSet(turtles []*Turtle) *TurtleAgentSet {
	turtleSet := make(map[*Turtle]interface{})
	for _, turtle := range turtles {
		turtleSet[turtle] = nil
	}

	return &TurtleAgentSet{
		turtles: turtleSet,
	}
}

func (t *TurtleAgentSet) Add(turtle *Turtle) {
	t.turtles[turtle] = len(t.turtles) - 1
}

func (t *TurtleAgentSet) Remove(turtle *Turtle) {
	if _, ok := t.turtles[turtle]; ok {
		delete(t.turtles, turtle)
	}
}

func (t *TurtleAgentSet) All(operation TurtleBoolOperation) bool {
	for turtle := range t.turtles {
		if !operation(turtle) {
			return false
		}
	}
	return true
}

func (t *TurtleAgentSet) Any(operation TurtleBoolOperation) bool {
	for turtle := range t.turtles {
		if operation(turtle) {
			return true
		}
	}
	return false
}

func (t *TurtleAgentSet) AtPoints(m *Model, points []Coordinate) *TurtleAgentSet {
	//get the turtles at the patches
	turtlesAtPatches := []*Turtle{}
	for _, point := range points {
		patch := m.Patch(point.X, point.Y)
		if patch != nil && patch.turtles[""] != nil {
			for turtle := range patch.turtles[""].turtles {
				if _, ok := t.turtles[turtle]; ok {
					turtlesAtPatches = append(turtlesAtPatches, turtle)
				}
			}
		}
	}

	return TurtleSet(turtlesAtPatches)
}

func (t *TurtleAgentSet) Count() int {
	return len(t.turtles)
}

// @TODO implement
func (t *TurtleAgentSet) InRadiusPatch(radius float64, patch *Patch) *TurtleAgentSet {
	return nil
}

// @TODO implement
func (t *TurtleAgentSet) InRadiusTurtle(radius float64, turtle *Turtle) *TurtleAgentSet {
	return nil
}

// @TODO implement
func (t *TurtleAgentSet) MaxNOf(n int, operation TurtleFloatOperation) *TurtleAgentSet {
	return nil
}

func (t *TurtleAgentSet) MaxOneOf(operation TurtleFloatOperation) *Turtle {
	max := math.MaxFloat64 * -1
	var maxTurtle *Turtle
	for turtle := range t.turtles {
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
	for turtle := range t.turtles {
		if operation(turtle) < min {
			min = operation(turtle)
			minTurtle = turtle
		}
	}
	return minTurtle
}

// @TODO implement
func (t *TurtleAgentSet) OneOf() *Turtle {
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
	for turtle := range t.turtles {
		if operation(turtle) {
			turtles = append(turtles, turtle)
		}
	}
	return TurtleSet(turtles)
}

func (t *TurtleAgentSet) WithMax(operation TurtleFloatOperation) *TurtleAgentSet {
	max := math.MaxFloat64 * -1
	for turtle := range t.turtles {
		if operation(turtle) > max {
			max = operation(turtle)
		}
	}

	//get all turtles where the float operation is equal to the max
	turtles := make([]*Turtle, 0)
	for turtle := range t.turtles {
		if operation(turtle) == max {
			turtles = append(turtles, turtle)
		}
	}

	return TurtleSet(turtles)
}

func (t *TurtleAgentSet) WithMin(operation TurtleFloatOperation) *TurtleAgentSet {
	min := math.MaxFloat64
	for turtle := range t.turtles {
		if operation(turtle) < min {
			min = operation(turtle)
		}
	}

	//get all turtles where the float operation is equal to the min
	turtles := make([]*Turtle, 0)
	for turtle := range t.turtles {
		if operation(turtle) == min {
			turtles = append(turtles, turtle)
		}
	}

	return TurtleSet(turtles)
}
