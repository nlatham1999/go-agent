package model

import "math"

type Turtle struct {
	xcor float64
	ycor float64
	who  int //the id of the turtle
	size int

	Color   Color
	heading float64 //direction the turtle is facing in radians
	Hidden  bool    //if the turtle is hidden
	breed   string
	Shape   string
	Size    float64

	parent *Model //model the turtle belongs too

	Label      interface{}
	LabelColor Color

	patch *Patch //patch the turtle is on
}

// @TODO might be faster having the patch passed in as a parameter instead of having to calculate it
func NewTurtle(m *Model, who int, breed string, x float64, y float64) *Turtle {

	if m == nil {
		return nil
	}

	//if the breed is nonexistent then return nil
	var breedSet *TurtleBreed = nil
	if breed != "" {
		found := false
		if breedSet, found = m.Breeds[breed]; !found {
			return nil
		}
	}

	t := &Turtle{
		who:    who,
		parent: m,
		xcor:   x,
		ycor:   y,
		breed:  breed,
	}

	m.Turtles.turtles[t] = nil
	m.whoToTurtles[m.turtlesWhoNumber] = t

	if breedSet != nil {
		breedSet.Turtles.turtles[t] = nil
	}

	//link the turtle to the patch
	t.patch = t.PatchHere()
	t.patch.addTurtle(t)

	return t
}

func (t *Turtle) Back(distance float64) {
	t.Forward(-distance)
}

func (t *Turtle) BreedName() string {
	return t.breed
}

func (t *Turtle) Breed() *TurtleBreed {
	return t.parent.Breeds[t.breed]
}

// Sets the breed of the turtle to the name passed in
func (t *Turtle) SetBreed(name string) {

	if t.breed == name {
		return
	}

	if t.breed != "" {
		delete(t.parent.Breeds[t.breed].Turtles.turtles, t)
	}

	t.breed = name

	if name != "" {
		t.parent.Breeds[name].Turtles.turtles[t] = nil
	}
}

// @TODO make a secondary position that just checks the new coords so that we don't have to do the move calcs twice
func (t *Turtle) CanMove(distance float64) bool {

	newX := t.xcor + distance*math.Cos(t.heading)
	newY := t.ycor + distance*math.Sin(t.heading)

	// patchX := math.Round(newX)
	// patchY := math.Round(newY)

	if t.parent.Patch(newX, newY) != nil {
		return true
	}

	return false
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
	t.heading = math.Atan2(y-t.ycor, x-t.xcor)
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
		t.Jump(1 * float64(direction))
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

func (t *Turtle) GetHeading() float64 {
	// return heading in degrees
	return t.heading * (180 / math.Pi)
}

func (t *Turtle) SetHeading(heading float64) {
	//convert heading to radians
	t.heading = heading * (math.Pi / 180)
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

// jumps ahead by the distance, if it cannot then it returns false
func (t *Turtle) Jump(distance float64) {

	if t.CanMove(distance) {
		xcor := t.xcor + distance*math.Cos(t.heading)
		ycor := t.ycor + distance*math.Sin(t.heading)
		t.SetXY(xcor, ycor)
	}
}

func (t *Turtle) Left(number float64) {

	//convert number to radians
	number = number * (math.Pi / 180)

	//add the number to the heading
	t.heading = math.Mod((t.heading + number), 2*math.Pi)
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

func (t *Turtle) Neighbors() *PatchAgentSet {
	//get the patch the turtle is on
	p := t.PatchHere()
	if p == nil {
		return nil
	}

	//get the neighbors of the patch
	neighbors := t.parent.neighbors(p)

	return neighbors
}

func (t *Turtle) Neighbors4() *PatchAgentSet {
	//get the patch the turtle is on
	p := t.PatchHere()
	if p == nil {
		return nil
	}

	//get the neighbors of the patch
	neighbors := t.parent.neighbors4(p)

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
	distX := t.xcor + distance*math.Cos(t.heading)
	distY := t.ycor + distance*math.Sin(t.heading)
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
	rightHeading := t.heading + angle
	distX := t.xcor + distance*math.Cos(rightHeading)
	distY := t.ycor + distance*math.Sin(rightHeading)
	return t.parent.getPatchAtCoords(int(distX), int(distY))
}

func (t *Turtle) PatchRightAndAhead(angle float64, distance float64) *Patch {
	rightHeading := t.heading - angle
	distX := t.xcor + distance*math.Cos(rightHeading)
	distY := t.ycor + distance*math.Sin(rightHeading)
	return t.parent.getPatchAtCoords(int(distX), int(distY))
}

// it might be faster to not use mods, the only danger is possible overflow
func (t *Turtle) Right(number float64) {
	t.Left(-number)
}

func (t *Turtle) SetXY(x float64, y float64) {
	t.xcor = x
	t.ycor = y

	oldPatch := t.patch
	t.patch = nil
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

func (t *Turtle) Who() int {
	return t.who
}

func (t *Turtle) XCor() float64 {
	return t.xcor
}

func (t *Turtle) YCor() float64 {
	return t.ycor
}
