package universe

import "math"

type Turtle struct {
	x       float64
	y       float64
	Who     int //the id of the turtle
	size    int
	Color   float64
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

	patchX := math.Round(newX)
	patchY := math.Round(newY)

	if t.parent.getPatchAtCoords(int(patchX), int(patchY)) != nil {
		return true
	}

	return true
}

func (t *Turtle) GetPatch() *Patch {
	px := int(math.Round(t.x))
	py := int(math.Round(t.y))
	p := t.parent.getPatchAtCoords(px, py)
	return p
}

func (t *Turtle) PatchRightAndAhead(angle float64, distance float64) *Patch {
	rightHeading := t.Heading - angle
	distX := t.x + distance*math.Cos(rightHeading)
	distY := t.y + distance*math.Sin(rightHeading)
	return t.parent.getPatchAtCoords(int(distX), int(distY))
}

func (t *Turtle) PatchLeftAndAhead(angle float64, distance float64) *Patch {
	rightHeading := t.Heading + angle
	distX := t.x + distance*math.Cos(rightHeading)
	distY := t.y + distance*math.Sin(rightHeading)
	return t.parent.getPatchAtCoords(int(distX), int(distY))
}

// it might be faster to not use mods, the only danger is possible overflow
func (t *Turtle) Right(number float64) {
	t.Heading = math.Mod((t.Heading - number), 360)
	if t.Heading < 0 {
		t.Heading += 360
	}
}

func (t *Turtle) Left(number float64) {
	t.Heading = math.Mod((t.Heading + number), 360)
}
