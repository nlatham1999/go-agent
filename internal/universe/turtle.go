package universe

import "math"

type Turtle struct {
	x       float64
	y       float64
	Who     int //the id of the turtle
	size    int
	Color   string
	Heading float64   //direction the turtle is facing in degrees
	parent  *Universe //universe the turtle belongs too
}

func NewTurtle(who int) *Turtle {
	return &Turtle{
		Who: who,
	}
}

//@TODO is this needed
//creates a new turtle from a template
//possible template attributes:
//	color
//	size
func NewTurtleFromTemplate(template *Turtle, who int) *Turtle {
	t := NewTurtle(who)
	t.Color = template.Color
	t.size = template.size

	return t
}

//just use a for loop
// func repeat(){

// }

//@TODO it might be better in the future to split the input between a whole and a decimal, so that we don't have to spend time splitting
func (t *Turtle) Forward(distance float64) {
	intPart := int(distance)
	decimalPart := distance - float64(intPart)

	direction := 1
	if intPart < 0 {
		direction = -1
		intPart *= -1
	}

	for i := 0; i < intPart; i++ {
		t.Jump(float64(intPart) * float64(direction))
	}

	if decimalPart != 0 {
		t.Jump(decimalPart * float64(direction))
	}
}

//jumps ahead by the distance, if it cannot then it returns false
func (t *Turtle) Jump(distance float64) bool {

	if t.CanMove(distance) {
		t.x = t.x + distance*math.Cos(t.Heading)
		t.y = t.y + distance*math.Sin(t.Heading)
	}

	return true
}

//@TODO make a secondary position that just checks the new coords so that we don't have to do the move calcs twice
func (t *Turtle) CanMove(distance float64) bool {

	newX := t.x + distance*math.Cos(t.Heading)
	newY := t.y + distance*math.Sin(t.Heading)

	patchX := math.Ceil(newX)
	patchY := math.Ceil(newY)

	if t.parent.getPatchAtCoords(int(patchX), int(patchY)) != nil {
		return true
	}

	return true
}
