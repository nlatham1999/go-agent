package universe

import "math"

type Turtle struct {
	xcor float64
	ycor float64
	Who  int //the id of the turtle
	size int

	Color   Color
	Heading float64 //direction the turtle is facing in degrees
	Hidden  bool    //if the turtle is hidden
	breed   string
	Shape   string
	Size    float64

	parent *Universe //universe the turtle belongs too

	Label      interface{}
	LabelColor Color

	patch *Patch //patch the turtle is on
}

func NewTurtle(u *Universe, who int, breed string) *Turtle {

	if u == nil {
		return nil
	}

	//if the breed is nonexistent then return nil
	if breed != "" {
		if _, found := u.Breeds[breed]; !found {
			return nil
		}
	}

	t := &Turtle{
		Who:    who,
		parent: u,
		xcor:   0,
		ycor:   0,
		breed:  breed,
	}

	//link the turtle to the patch
	t.patch = t.PatchHere()
	t.patch.addTurtle(t)

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

	newX := t.xcor + distance*math.Cos(t.Heading)
	newY := t.ycor + distance*math.Sin(t.Heading)

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
	t.Heading = math.Atan2(y-t.ycor, x-t.xcor)
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
		xcor := t.xcor + distance*math.Cos(t.Heading)
		ycor := t.ycor + distance*math.Sin(t.Heading)
		t.SetXY(xcor, ycor)
	}
}

func (t *Turtle) Left(number float64) {
	t.Heading = math.Mod((t.Heading + number), 360)
}

// @TODO implement
func (t *Turtle) LinkNeighbors(breed string) []*Turtle {
	return nil
}

// @TODO implement
func (t *Turtle) LinkNeighbor(turtle *Turtle) bool {
	return false
}

func (t *Turtle) MoveToPatch(patch *Patch) {
	t.xcor = patch.xFloat64
	t.ycor = patch.yFloat64
}

func (t *Turtle) MoveToTurtle(turtle *Turtle) {
	t.xcor = turtle.xcor
	t.ycor = turtle.ycor
}

// @TODO implement
func (t *Turtle) MyLinks(breed string) []*Link {
	return nil
}

// @TODO implement
func (t *Turtle) MyInLinks(breed string) []*Link {
	return nil
}

// @TODO implement
func (t *Turtle) MyOutLinks(breed string) []*Link {
	return nil
}

func (t *Turtle) Neighbors() []*Patch {
	//get the patch the turtle is on
	p := t.PatchHere()
	if p == nil {
		return nil
	}

	//get the index of the patch
	index := p.index

	//get the neighbors of the patch
	neighbors := t.parent.Neighbors(index)

	return neighbors
}

func (t *Turtle) Neighbors4() []*Patch {
	//get the patch the turtle is on
	p := t.PatchHere()
	if p == nil {
		return nil
	}

	//get the index of the patch
	index := p.index

	//get the neighbors of the patch
	neighbors := t.parent.Neighbors4(index)

	return neighbors
}

// TODO implement
func (t *Turtle) Other(turtles TurtleAgentSet) *TurtleAgentSet {
	return nil
}

// @TODO implement
func (t *Turtle) OtherEnd(link *Link) *Turtle {
	return nil
}

// @TODO implement
func (t *Turtle) OutLinkNeighbor(breed string, turtle *Turtle) bool {
	return false
}

// @TODO implement
func (t *Turtle) OutLinkNeighbors(breed string, turtle *Turtle) *TurtleAgentSet {
	return nil
}

// @TODO implement
func (t *Turtle) OutLinkTo(breed string, turtle *Turtle) *Link {
	return nil
}

func (t *Turtle) PatchAhead(distance float64) *Patch {
	distX := t.xcor + distance*math.Cos(t.Heading)
	distY := t.ycor + distance*math.Sin(t.Heading)
	return t.parent.getPatchAtCoords(int(distX), int(distY))
}

func (t *Turtle) PatchAt(dx float64, dy float64) *Patch {

	//round the coords
	px := int(math.Round(t.xcor + dx))
	py := int(math.Round(t.ycor + dy))

	return t.parent.getPatchAtCoords(px, py)
}

func (t *Turtle) PatchAtHeadingAndDistance(heading float64, distance float64) *Patch {
	distX := t.xcor + distance*math.Cos(heading)
	distY := t.ycor + distance*math.Sin(heading)
	return t.parent.Patch(distX, distY)
}

func (t *Turtle) PatchHere() *Patch {

	if t.patch != nil {
		return t.patch
	}

	p := t.parent.Patch(t.xcor, t.ycor)

	t.patch = p

	return p
}

func (t *Turtle) PatchLeftAndAhead(angle float64, distance float64) *Patch {
	rightHeading := t.Heading + angle
	distX := t.xcor + distance*math.Cos(rightHeading)
	distY := t.ycor + distance*math.Sin(rightHeading)
	return t.parent.getPatchAtCoords(int(distX), int(distY))
}

func (t *Turtle) PatchRightAndAhead(angle float64, distance float64) *Patch {
	rightHeading := t.Heading - angle
	distX := t.xcor + distance*math.Cos(rightHeading)
	distY := t.ycor + distance*math.Sin(rightHeading)
	return t.parent.getPatchAtCoords(int(distX), int(distY))
}

// it might be faster to not use mods, the only danger is possible overflow
func (t *Turtle) Right(number float64) {
	t.Heading = math.Mod((t.Heading - number), 360)
	if t.Heading < 0 {
		t.Heading += 360
	}
}

func (t *Turtle) SetXY(x float64, y float64) {
	t.xcor = x
	t.ycor = y

	oldPatch := t.patch
	t.patch = t.PatchHere()
	if t.patch != oldPatch {
		oldPatch.removeTurtle(t)
		t.patch.addTurtle(t)
	}
}

func (t *Turtle) Show() {
	t.Hidden = false
}

func (t *Turtle) TowardsPatch(patch *Patch) float64 {
	//returns heading that faces the patch
	return math.Atan2(patch.yFloat64-t.ycor, patch.xFloat64-t.xcor)
}

func (t *Turtle) TowardsTurtle(turtle *Turtle) float64 {
	//returns heading that faces the turtle
	return math.Atan2(turtle.ycor-t.ycor, turtle.xcor-t.xcor)
}

func (t *Turtle) TowardsXY(x float64, y float64) float64 {
	//returns heading that faces the x y coordinates
	return math.Atan2(y-t.ycor, x-t.xcor)
}

// @TODO implement
func (t *Turtle) TurtlesHere(breed string) *TurtleAgentSet {
	return nil
}

// @TODO implement
func (t *Turtle) Uphill(patchVariable string) {

}

// @TODO implement
func (t *Turtle) Uphill4(patchVariable string) {

}

// @TODO implement
func (t *Turtle) LinkWith(turtle *Turtle) *Link {
	return nil
}

func (t *Turtle) XCor() float64 {
	return t.xcor
}

func (t *Turtle) YCor() float64 {
	return t.ycor
}
