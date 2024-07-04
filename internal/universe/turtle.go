package universe

import "math"

type Turtle struct {
	x    float64
	y    float64
	Who  int //the id of the turtle
	size int

	Color   Color
	Heading float64 //direction the turtle is facing in degrees
	Hidden  bool    //if the turtle is hidden
	Breed   string

	parent *Universe //universe the turtle belongs too

	Base

	Label      interface{}
	LabelColor Color
}

func NewTurtle(who int) *Turtle {
	return &Turtle{
		Who: who,
	}
}

// @TODO is this needed
// creates a new turtle from a template
// possible template attributes:
//
//	color
//	size
func NewTurtleFromTemplate(template *Turtle, who int) *Turtle {
	t := NewTurtle(who)
	t.Color = template.Color
	t.size = template.size

	return t
}

// @TODO Implement
func (t *Turtle) Back(distance float64) {

}

// @TODO implement
func (t *Turtle) GetBreedName() string {
	return ""
}

// @TODO implement
func (t *Turtle) GetBreedSet() []*Turtle {
	return nil
}

// @TODO implement
func (t *Turtle) SetBreed(name string) {

}

// @TODO make a secondary position that just checks the new coords so that we don't have to do the move calcs twice
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

// @TODO implement
// creates a directed breed link from the current turtle to the turtle passed in
func (t *Turtle) CreateBreedTo(breed string, turtle *Turtle, operations []TurtleOperation) {

}

// @TODO implement
// creates a directed breed link from the current turtle to the turtles passed in
func (t *Turtle) CreateBreedsTo(breed string, turtles []*Turtle, operations []TurtleOperation) {

}

// @TODO implement
// creates an undirected breed link from the current turtle with the turtle passed in
func (t *Turtle) CreateBreedWith(breed string, turtle Turtle, operations []TurtleOperation) {

}

// @TODO implement
// creates an undirected breed link from the current turtle with the turtles passed in
func (t *Turtle) CreateBreedsWith(breed string, turtles []*Turtle, operations []TurtleOperation) {

}

// @TODO implement
// creates a directed breed link from the current turtle with the turtle passed in
func (t *Turtle) CreateBreedFrom(breed string, turtle *Turtle, operations []TurtleOperation) {

}

// @TODO implement
func (t *Turtle) CreateBreedsFrom(breed string, turtles []*Turtle, operations []TurtleOperation) {

}

// @TODO implement
func (t *Turtle) CreateLinkTo(turtle *Turtle, operations []TurtleOperation) {

}

// @TODO implement
func (t *Turtle) CreateLinksTo(turtles []*Turtle, operations []TurtleOperation) {

}

// @TODO implement
func (t *Turtle) CreateLinkWith(turtle *Turtle, operations []TurtleOperation) {

}

// @TODO implement
func (t *Turtle) CreateLinksWith(turtles []*Turtle, operations []TurtleOperation) {

}

// @TODO implement
func (t *Turtle) CreateLinkFrom(turtle *Turtle, operations []TurtleOperation) {

}

// @TODO implement
func (t *Turtle) CreateLinksFrom(turtles []*Turtle, operations []TurtleOperation) {

}

// @TODO implement
func (t *Turtle) DistanceTurtle(turtle *Turtle) float64 {
	return 0
}

// @TODO implement
func (t *Turtle) DistancePatch(patch *Patch) float64 {
	return 0
}

// @TODO implement
func (t *Turtle) DistanceXY(x float64, y float64) float64 {
	return 0
}

// @TODO implement
func (t *Turtle) Downhill(patchVariable string) {

}

// @TODO implement
func (t *Turtle) Downhill4(patchVariable string) {

}

// @TODO implement
func (t *Turtle) DX() float64 {
	return 0
}

// @TODO implement
func (t *Turtle) DY() float64 {
	return 0
}

// @TODO implement
func (t *Turtle) FaceTurtle(turtle *Turtle) {

}

// @TODO implement
func (t *Turtle) FacePatch(patch *Patch) {

}

// @TODO implement
func (t *Turtle) FaceXY(x float64, y float64) {
	t.Heading = math.Atan2(y-t.y, x-t.x)
}

// @TODO it might be better in the future to split the input between a whole and a decimal, so that we don't have to spend time splitting
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

// @TODO implement
func (t *Turtle) Hatch(amount int, operations []TurtleOperation) {

}

// @TODO implement
func (t *Turtle) HatchBreed(breed string, amount int, operations []TurtleOperation) {

}

func (t *Turtle) Hide() {
	t.Hidden = true
}

func (t *Turtle) Home() {
	t.SetXY(0, 0)
}

// @TODO implement
func (t *Turtle) InConePatches(distance float64, angle float64) []*Patch {
	return nil
}

// @TODO implement
func (t *Turtle) InConeTurtles(distance float64, angle float64) []*Turtle {
	return nil
}

// @TODO implement
func (t *Turtle) InLinkBreedNeighbor(breed string, turtle *Turtle) bool {
	return false
}

// @TODO implement
func (t *Turtle) InLinkNeighbor(turtle *Turtle) bool {
	return false
}

// @TODO implement
func (t *Turtle) InLinkBreedNeighbors(breed string, turtle *Turtle) []*Turtle {
	return nil
}

// @TODO implement
func (t *Turtle) InLinkNeighbors(turtle *Turtle) []*Turtle {
	return nil
}

// @TODO implement
func (t *Turtle) InLinkBreedFrom(breed string, turtle *Turtle) *Link {
	return nil
}

// @TODO implement
func (t *Turtle) InLinkFrom(turtle *Turtle) *Link {
	return nil
}

// @TODO implement
func (t *Turtle) InRadiusPatches(distance float64) []*Patch {
	return nil
}

// @TODO implement
func (t *Turtle) InRadiusTurtles(distance float64) []*Turtle {
	return nil
}

// jumps ahead by the distance, if it cannot then it returns false
func (t *Turtle) Jump(distance float64) {

	if t.CanMove(distance) {
		t.x = t.x + distance*math.Cos(t.Heading)
		t.y = t.y + distance*math.Sin(t.Heading)
	}
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
