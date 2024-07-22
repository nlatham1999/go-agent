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

	turtlesOwnGeneral map[string]interface{} // turtles own variables
	turtlesOwnBreed   map[string]interface{} // turtles own variables

	// 3D map of the links the turtle is involved in
	// the first bool is if the link is directed
	// the string is the breed of the link - the empty string key is all the links
	// the third map is the turtles the link is connected to
	linkedTurtles map[bool]map[string]map[*Turtle]*Link

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

	//set the turtle own variables
	//breed specific variables can override general variables
	t.turtlesOwnGeneral = make(map[string]interface{})
	t.turtlesOwnBreed = make(map[string]interface{})
	generalTemplate := m.Breeds[""].turtlesOwnTemplate
	for key, value := range generalTemplate {
		t.turtlesOwnGeneral[key] = value
	}
	if breedSet != nil {
		breedTemplate := breedSet.turtlesOwnTemplate
		for key, value := range breedTemplate {
			t.turtlesOwnBreed[key] = value
		}
	}

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

	// switch the turtles own variables to the new breed
	t.turtlesOwnBreed = make(map[string]interface{})
	if name != "" {
		breedTemplate := t.parent.Breeds[name].turtlesOwnTemplate
		for key, value := range breedTemplate {
			t.turtlesOwnBreed[key] = value
		}
	}
}

func (t *Turtle) CanMove(distance float64) bool {

	if t.parent.wrapping {
		return true
	}

	newX := t.xcor + distance*math.Cos(t.heading)
	newY := t.ycor + distance*math.Sin(t.heading)

	// patchX := math.Round(newX)
	// patchY := math.Round(newY)

	if t.parent.Patch(newX, newY) != nil {
		return true
	}

	return false
}

// creates a directed link from the current turtle to the turtle passed in
func (t *Turtle) CreateLinkTo(breed string, turtle *Turtle, operations []LinkOperation) {
	l := NewLink(t.parent, breed, t, turtle, true)

	for _, operation := range operations {
		operation(l)
	}
}

// creates a directed link from the current turtle to the turtles passed in
func (t *Turtle) CreateLinkToSet(breed string, turtles *TurtleAgentSet, operations []LinkOperation) {
	if breed == "" {
		return
	}

	linksAdded := LinkSet([]*Link{})
	for turtle := range turtles.turtles {
		l := NewLink(t.parent, breed, t, turtle, true)
		linksAdded.Add(l)
	}

	AskLinks(linksAdded, operations)
}

// creates an undirected breed link from the current turtle with the turtle passed in
func (t *Turtle) CreateLinkWith(breed string, turtle *Turtle, operations []LinkOperation) {
	l := NewLink(t.parent, breed, t, turtle, false)

	for _, operation := range operations {
		operation(l)
	}

}

// creates an undirected breed link from the current turtle with the turtles passed in
func (t *Turtle) CreateLinkWithSet(breed string, turtles *TurtleAgentSet, operations []LinkOperation) {
	linksAdded := LinkSet([]*Link{})
	for turtle := range turtles.turtles {
		l := NewLink(t.parent, breed, t, turtle, false)
		linksAdded.Add(l)
	}

	AskLinks(linksAdded, operations)
}

// creates a directed breed link from the current turtle with the turtle passed in
func (t *Turtle) CreateBreedFrom(breed string, turtle *Turtle, operations []LinkOperation) {
	l := NewLink(t.parent, breed, turtle, t, true)

	for _, operation := range operations {
		operation(l)
	}
}

// creates a directed breed link from the turtles passed in to the current turtle
func (t *Turtle) CreateBreedsFrom(breed string, turtles *TurtleAgentSet, operations []LinkOperation) {
	linksAdded := LinkSet([]*Link{})
	for turtle := range turtles.turtles {
		l := NewLink(t.parent, breed, turtle, t, true)
		linksAdded.Add(l)
	}

	AskLinks(linksAdded, operations)
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

// finds a
func (t *Turtle) InLinkFrom(breed string, turtle *Turtle) *Link {

	if turtle.linkedTurtles != nil {
		// look in the directed links
		if turtle.linkedTurtles[true] != nil {
			if turtle.linkedTurtles[true][breed] != nil {
				if link, found := turtle.linkedTurtles[true][breed][t]; found {
					return link
				}
			}
		}
		// look in the undirected links
		if turtle.linkedTurtles[false] != nil {
			if turtle.linkedTurtles[false][breed] != nil {
				if link, found := turtle.linkedTurtles[false][breed][t]; found {
					return link
				}
			}
		}
	}

	return nil
}

// jumps ahead by the distance, if it cannot then it returns false
// @TODO implement - wrapping and don't call Can Move since that is expensive
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

// returns the turtle own variable
func (t *Turtle) GetOwn(key string) interface{} {
	if val, found := t.turtlesOwnBreed[key]; found {
		return val
	}
	if val, found := t.turtlesOwnGeneral[key]; found {
		return val
	}
	return nil
}

// sets the turtle own variable
func (t *Turtle) SetOwn(key string, value interface{}) {
	if _, found := t.turtlesOwnBreed[key]; found {
		t.turtlesOwnBreed[key] = value
		return
	} else {
		if _, found := t.turtlesOwnGeneral[key]; found {
			t.turtlesOwnGeneral[key] = value
		}
	}
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
