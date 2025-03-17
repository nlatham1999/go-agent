package model

import (
	"github.com/nlatham1999/sortedset"
)

type TurtleAgentSet struct {
	turtles *sortedset.SortedSet //map of turtles so we can quickly check if a turtle is in the set
}

func NewTurtleAgentSet(turtles []*Turtle) *TurtleAgentSet {
	turtleSet := &TurtleAgentSet{
		turtles: sortedset.NewSortedSet(),
	}
	for _, turtle := range turtles {
		turtleSet.turtles.Add(turtle)
	}
	return turtleSet
}

func (t *TurtleAgentSet) Add(turtle *Turtle) {
	t.turtles.Add(turtle)
}

func (t *TurtleAgentSet) All(operation TurtleBoolOperation) bool {
	if operation == nil {
		return false
	}

	return t.turtles.All(func(a interface{}) bool {
		return operation(a.(*Turtle))
	})
}

func (t *TurtleAgentSet) Any(operation TurtleBoolOperation) bool {
	if operation == nil {
		return false
	}

	return t.turtles.Any(func(a interface{}) bool {
		return operation(a.(*Turtle))
	})
}

func (t *TurtleAgentSet) Ask(operation TurtleOperation) {
	if operation == nil {
		return
	}

	t.turtles.Ask(func(a interface{}) {
		operation(a.(*Turtle))
	})
}

func (t *TurtleAgentSet) AtPoints(m *Model, points []Coordinate) *TurtleAgentSet {

	// convert the points to patches
	patchesAtPoints := sortedset.NewSortedSet()
	for _, point := range points {
		patch := m.Patch(point.X, point.Y)
		patchesAtPoints.Add(patch)
	}

	return t.With(func(turtle *Turtle) bool {
		return patchesAtPoints.Contains(turtle.PatchHere())
	})

}

func (t *TurtleAgentSet) Contains(turtle *Turtle) bool {
	return t.turtles.Contains(turtle)
}

func (t *TurtleAgentSet) Count() int {
	return t.turtles.Len()
}

func (t *TurtleAgentSet) InRadiusPatch(radius float64, patch *Patch) *TurtleAgentSet {
	return t.With(func(t *Turtle) bool {
		return t.DistancePatch(patch) <= radius
	})
}

func (t *TurtleAgentSet) InRadiusTurtle(radius float64, turtle *Turtle) *TurtleAgentSet {
	return t.With(func(t *Turtle) bool {
		return t.DistanceTurtle(turtle) <= radius
	})
}

func (t *TurtleAgentSet) List() []*Turtle {
	turtles := make([]*Turtle, 0)
	t.turtles.Ask(func(turtle interface{}) {
		turtles = append(turtles, turtle.(*Turtle))
	})
	return turtles
}

func (t *TurtleAgentSet) FirstNOf(n int) *TurtleAgentSet {
	turtleSet := sortedset.NewSortedSet()
	turtle := t.turtles.First()
	for i := 0; i < n && turtle != nil; i++ {
		turtleSet.Add(turtle)
		turtle, _ = t.turtles.Next()
	}
	return &TurtleAgentSet{
		turtles: turtleSet,
	}
}

func (t *TurtleAgentSet) First() (*Turtle, error) {
	turtle := t.turtles.First()
	if turtle == nil {
		return nil, ErrNoTurtlesInAgentSet
	}
	return turtle.(*Turtle), nil
}

func (t *TurtleAgentSet) LastNOf(n int) *TurtleAgentSet {
	turtleSet := sortedset.NewSortedSet()
	turtle := t.turtles.Last()
	for i := 0; i < n && turtle != nil; i++ {
		turtleSet.Add(turtle)
		turtle, _ = t.turtles.Previous()
	}
	return &TurtleAgentSet{
		turtles: turtleSet,
	}
}

func (t *TurtleAgentSet) Last() (*Turtle, error) {
	turtle := t.turtles.Last()
	if turtle == nil {
		return nil, ErrNoTurtlesInAgentSet
	}
	return turtle.(*Turtle), nil
}

// func (t *TurtleAgentSet) Next() (*Turtle, error) {
// 	turtle, _ := t.turtles.Next()
// 	if turtle == nil {
// 		return nil, ErrNoTurtlesInAgentSet
// 	}
// 	return turtle.(*Turtle), nil
// }

func (t *TurtleAgentSet) Remove(turtle *Turtle) {
	t.turtles.Remove(turtle)
}

func (t *TurtleAgentSet) SortAsc(operation TurtleFloatOperation) {
	t.turtles.SortAsc(func(a interface{}) interface{} {
		return operation(a.(*Turtle))
	})
}

func (t *TurtleAgentSet) SortDesc(operation TurtleFloatOperation) {
	t.turtles.SortDesc(func(a interface{}) interface{} {
		return operation(a.(*Turtle))
	})
}

// returns a new TurtleAgentSet with all the turtles that are not in the given TurtleAgentSet
func (t *TurtleAgentSet) WhoAreNot(turtles *TurtleAgentSet) *TurtleAgentSet {
	return &TurtleAgentSet{
		turtles: t.turtles.Difference(turtles.turtles),
	}
}

// returns a new TurtleAgentSet with all the turtles that are not the given turtle
func (t *TurtleAgentSet) WhoAreNotTurtle(turtle *Turtle) *TurtleAgentSet {
	return &TurtleAgentSet{
		turtles: t.turtles.Difference(sortedset.NewSortedSet(turtle)),
	}
}

func (t *TurtleAgentSet) With(operation TurtleBoolOperation) *TurtleAgentSet {
	turtleSet := sortedset.NewSortedSet()

	if operation == nil {
		return nil
	}

	t.turtles.Ask(func(a interface{}) {
		if operation(a.(*Turtle)) {
			turtleSet.Add(a)
		}
	})

	return &TurtleAgentSet{
		turtles: turtleSet,
	}
}
