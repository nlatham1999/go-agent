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
	linkedTurtles map[linkedTurtle]*Link

	// map of the turtles that are connected to the current turtle
	linkedTurtlesConnectedFrom map[*Turtle]*Link

	patch *Patch //patch the turtle is on
}

type linkedTurtle struct {
	directed bool
	breed    string
	turtle   *Turtle
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
		who:                        who,
		parent:                     m,
		xcor:                       x,
		ycor:                       y,
		breed:                      breed,
		linkedTurtles:              make(map[linkedTurtle]*Link),
		linkedTurtlesConnectedFrom: make(map[*Turtle]*Link),
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

	if t.parent.wrappingY && t.parent.wrappingX {
		return true
	}

	newX := t.xcor + distance*math.Cos(t.heading)
	newY := t.ycor + distance*math.Sin(t.heading)

	if newX < float64(t.parent.MinPxCor) || newX >= float64(t.parent.MaxPxCor) {
		if !t.parent.wrappingX {
			return false
		} else {
			return true
		}
	}

	if newY < float64(t.parent.MinPyCor) || newY >= float64(t.parent.MaxPyCor) {
		if !t.parent.wrappingY {
			return false
		} else {
			return true
		}
	}

	// patchX := math.Round(newX)
	// patchY := math.Round(newY)

	return true
}

// creates a directed link from the current turtle to the turtle passed in
func (t *Turtle) CreateLinkToTurtle(breed string, turtle *Turtle, operations []LinkOperation) {
	l := NewLink(t.parent, breed, t, turtle, true)

	for _, operation := range operations {
		operation(l)
	}
}

// creates a directed link from the current turtle to the turtles passed in
func (t *Turtle) CreateLinksToSet(breed string, turtles *TurtleAgentSet, operations []LinkOperation) {
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
func (t *Turtle) CreateLinkWithTurtle(breed string, turtle *Turtle, operations []LinkOperation) {
	l := NewLink(t.parent, breed, t, turtle, false)

	for _, operation := range operations {
		operation(l)
	}

}

// creates an undirected breed link from the current turtle with the turtles passed in
func (t *Turtle) CreateLinksWithSet(breed string, turtles *TurtleAgentSet, operations []LinkOperation) {
	linksAdded := LinkSet([]*Link{})
	for turtle := range turtles.turtles {
		l := NewLink(t.parent, breed, t, turtle, false)
		linksAdded.Add(l)
	}

	AskLinks(linksAdded, operations)
}

// creates a directed breed link from the current turtle with the turtle passed in
func (t *Turtle) CreateLinkFromTurtle(breed string, turtle *Turtle, operations []LinkOperation) {
	l := NewLink(t.parent, breed, turtle, t, true)

	for _, operation := range operations {
		operation(l)
	}
}

// creates a directed breed link from the turtles passed in to the current turtle
func (t *Turtle) CreateLinksFromSet(breed string, turtles *TurtleAgentSet, operations []LinkOperation) {
	linksAdded := LinkSet([]*Link{})
	for turtle := range turtles.turtles {
		l := NewLink(t.parent, breed, turtle, t, true)
		linksAdded.Add(l)
	}

	AskLinks(linksAdded, operations)
}

// returns an agentset of the node turtles that are tied to the current turtle
func (t *Turtle) descendents(minTieMode TieMode) *TurtleAgentSet {
	tiedTurtlesMap := make(map[*Turtle]interface{})
	tiedTurtles := []*Turtle{}
	for _, l := range t.linkedTurtles {
		if l.TieMode >= minTieMode {
			tiedTurtles = append(tiedTurtles, l.OtherEnd(t))
		}
	}
	for len(tiedTurtles) > 0 {
		for _, turtle := range tiedTurtles {
			tiedTurtlesMap[turtle] = nil
		}
		newTiedTurtles := []*Turtle{}
		for _, turtle := range tiedTurtles {
			for _, l := range turtle.linkedTurtles {
				if l.TieMode != TieModeNone {
					otherEnd := l.OtherEnd(turtle)
					if _, found := tiedTurtlesMap[otherEnd]; !found && otherEnd != t {
						newTiedTurtles = append(newTiedTurtles, otherEnd)
					}
				}
			}
		}
		tiedTurtles = newTiedTurtles
	}

	return &TurtleAgentSet{tiedTurtlesMap}
}

// returns the distance between the two turtles
func (t *Turtle) DistanceTurtle(turtle *Turtle) float64 {
	return t.parent.DistanceBetweenPoints(t.xcor, t.ycor, turtle.xcor, turtle.ycor)
}

// returns the distance between the turtle and the middle of the patch
func (t *Turtle) DistancePatch(patch *Patch) float64 {
	return t.parent.DistanceBetweenPoints(t.xcor, t.ycor, patch.xFloat64, patch.yFloat64)
}

// returns the distance between the turtle and the x y coordinates
func (t *Turtle) DistanceXY(x float64, y float64) float64 {
	return t.parent.DistanceBetweenPoints(t.xcor, t.ycor, x, y)
}

// moves the turtle to the neighboring patch that has the lowest value of the patch variable
// if the current patch variable is the lowest then the turtle stays in place
func (t *Turtle) Downhill(patchVariable string) {

	p := t.PatchHere()

	// if the patch variable is not a number then return
	if _, ok := t.parent.patchesOwnTemplate[patchVariable].(float64); !ok {
		return
	}

	minPatch := p
	minValue := minPatch.patchesOwn[patchVariable].(float64)

	for patch := range p.patchNeighborsMap {
		if patch == nil {
			continue
		}
		if patch.patchesOwn[patchVariable].(float64) < minValue {
			minPatch = patch
			minValue = patch.patchesOwn[patchVariable].(float64)
		}
	}

	if minPatch != p {

		pos := p.patchNeighborsMap[minPatch]

		switch pos {
		case "left":
			t.SetHeading(LeftAngle)
		case "topLeft":
			t.SetHeading(UpAndLeftAngle)
		case "top":
			t.SetHeading(UpAngle)
		case "topRight":
			t.SetHeading(UpAndRightAngle)
		case "right":
			t.SetHeading(RightAngle)
		case "bottomRight":
			t.SetHeading(DownAndRightAngle)
		case "bottom":
			t.SetHeading(DownAngle)
		case "bottomLeft":
			t.SetHeading(DownAndLeftAngle)
		}

		t.MoveToPatch(minPatch)

	}
}

// moves the turtle to the neighboring4 patch that has the lowest value of the patch variable
// if the current patch variable is the lowest then the turtle stays in place
func (t *Turtle) Downhill4(patchVariable string) {

	neighborsMap := make(map[*Patch]string)

	p := t.PatchHere()

	topNeighbor := p.neighborsPatchMap["top"]
	if topNeighbor != nil {
		neighborsMap[topNeighbor] = "top"
	}

	bottomNeighbor := p.neighborsPatchMap["bottom"]
	if bottomNeighbor != nil {
		neighborsMap[bottomNeighbor] = "bottom"
	}

	leftNeighbor := p.neighborsPatchMap["left"]
	if leftNeighbor != nil {
		neighborsMap[leftNeighbor] = "left"
	}

	rightNeighbor := p.neighborsPatchMap["right"]
	if rightNeighbor != nil {
		neighborsMap[rightNeighbor] = "right"
	}

	// if the patch variable is not a number then return
	if _, ok := t.parent.patchesOwnTemplate[patchVariable].(float64); !ok {
		return
	}

	minPatch := t.PatchHere()
	minValue := minPatch.patchesOwn[patchVariable].(float64)

	for patch := range neighborsMap {
		if patch.patchesOwn[patchVariable].(float64) < minValue {
			minPatch = patch
			minValue = patch.patchesOwn[patchVariable].(float64)
		}
	}

	if minPatch != t.PatchHere() {

		pos := neighborsMap[minPatch]
		switch pos {
		case "top":
			t.SetHeading(UpAngle)
		case "bottom":
			t.SetHeading(DownAngle)
		case "left":
			t.SetHeading(LeftAngle)
		case "right":
			t.SetHeading(RightAngle)
		}

		t.MoveToPatch(minPatch)

	}
}

func (t *Turtle) DX() float64 {
	return math.Cos(t.heading)
}

func (t *Turtle) DY() float64 {
	return math.Sin(t.heading)
}

// faces the turtle passed in
func (t *Turtle) FaceTurtle(turtle *Turtle) {
	t.FaceXY(turtle.xcor, turtle.ycor)
}

// faces the patch passed in
func (t *Turtle) FacePatch(patch *Patch) {
	t.FaceXY(patch.xFloat64, patch.yFloat64)
}

// faces the x y coordinates passed in
func (t *Turtle) FaceXY(x float64, y float64) {
	if x == t.xcor && y == t.ycor {
		return
	}

	dx := x - t.xcor
	dy := y - t.ycor

	if !t.parent.wrappingX && !t.parent.wrappingY {
		t.setHeadingRadians(math.Atan2(dx, dy))
		return
	}

	newDx := dx
	newDy := dy

	dxSign := 1.0
	if dx < 0 {
		dxSign = -1.0
	}

	dySign := 1.0
	if dy < 0 {
		dySign = -1.0
	}

	distance := math.Abs(math.Sqrt(dx*dx + dy*dy))

	deltaXInverse := float64(t.parent.WorldWidth) - math.Abs(dx)
	deltaYInverse := float64(t.parent.WorldHeight) - math.Abs(dy)

	if t.parent.wrappingX {
		d := math.Min(distance, math.Abs(math.Sqrt(deltaXInverse*deltaXInverse+dy*dy)))
		if d < distance {
			distance = d
			newDx = deltaXInverse * dxSign * -1
		}
	}

	if t.parent.wrappingY {
		d := math.Min(distance, math.Abs(math.Sqrt(dx*dx+deltaYInverse*deltaYInverse)))
		if d < distance {
			distance = d
			newDy = deltaYInverse * dySign * -1
		}
	}

	if t.parent.wrappingX && t.parent.wrappingY {
		d := math.Min(distance, math.Abs(math.Sqrt(deltaXInverse*deltaXInverse+deltaYInverse*deltaYInverse)))
		if d < distance {
			newDx = deltaXInverse * dxSign * -1
			newDy = deltaYInverse * dySign * -1
		}
	}

	a := math.Atan2(newDx, newDy)
	t.setHeadingRadians(a)
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

// creates new turtles that are a copy of the current turtle
// if a breed is passed in then the new turtles will be of that breed
func (t *Turtle) Hatch(breed string, amount int, operations []TurtleOperation) {

	turtles := make([]*Turtle, amount)
	for i := 0; i < amount; i++ {
		newBreed := t.breed
		if breed != "" {
			newBreed = breed
		}
		turtles[i] = NewTurtle(t.parent, t.parent.turtlesWhoNumber, newBreed, t.xcor, t.ycor)
		t.parent.turtlesWhoNumber++

		// copy the variables
		turtles[i].Color = t.Color
		turtles[i].heading = t.heading
		turtles[i].Hidden = t.Hidden
		turtles[i].Shape = t.Shape
		turtles[i].Size = t.Size
		turtles[i].Label = t.Label
		turtles[i].LabelColor = t.LabelColor

		// copy the own variables
		for key, value := range t.turtlesOwnGeneral {
			turtles[i].turtlesOwnGeneral[key] = value
		}

		// copy the breed variables if the breed is the same
		if t.breed == newBreed {
			for key, value := range t.turtlesOwnBreed {
				turtles[i].turtlesOwnBreed[key] = value
			}
		}
	}

	for _, turtle := range turtles {
		for _, operation := range operations {
			operation(turtle)
		}
	}
}

func (t *Turtle) GetHeading() float64 {
	// return heading in degrees
	return t.heading * (180 / math.Pi)
}

// takes in a heading in degrees and sets the heading in radians
func (t *Turtle) SetHeading(heading float64) {

	//make sure the heading is between -360 and 360
	if heading > 360 || heading < -360 {
		heading = math.Mod(heading, 360)
	}

	//make sure the heading is between 0 and 360
	if heading < 0 {
		heading += 360
	}

	//convert heading to radians
	t.setHeadingRadians(heading * (math.Pi / 180))
}

func (t *Turtle) setHeadingRadians(heading float64) {

	headingDifference := heading - t.heading

	t.heading = heading

	if len(t.linkedTurtles) == 0 {
		return
	}

	// rotate the heading for all descendents where the tiemode is at least free
	for turtle := range t.descendents(TieModeFree).turtles {
		t.rotateTiedTurtle(turtle, headingDifference)
	}

	// rotate the heading for all descendents where the tiemode is fixed
	for turtle := range t.descendents(TieModeFixed).turtles {
		turtle.heading += headingDifference
	}
}

// swivvels the turtle to the left or right by the amount passed in
// the turtle passed in will maintain the same distance from the current turtle
func (t *Turtle) rotateTiedTurtle(turtle *Turtle, amount float64) {
	if t == turtle {
		return
	}

	// get the distance between the two turtles
	distanceX := turtle.xcor - t.xcor
	distanceY := turtle.ycor - t.ycor

	// get the new x and y coordinates
	newX := t.xcor + distanceX*math.Cos(-amount) - distanceY*math.Sin(-amount)
	newY := t.ycor + distanceX*math.Sin(-amount) + distanceY*math.Cos(-amount)

	newX, newY, allowed := t.parent.convertXYToInBounds(newX, newY)
	if !allowed {
		return
	}

	turtle.xcor = newX
	turtle.ycor = newY
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

// finds a link from the current turtle to the turtle passed in
func (t *Turtle) InLinkFrom(breed string, turtle *Turtle) *Link {

	if turtle.linkedTurtles != nil {
		// look in the directed links
		s := linkedTurtle{true, breed, t}
		if link, found := turtle.linkedTurtles[s]; found {
			return link
		}

		// look in the undirected links
		s = linkedTurtle{false, breed, t}
		if link, found := turtle.linkedTurtles[s]; found {
			return link
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
	// convert number to radians
	number = -number * (math.Pi / 180)

	// add the number to the heading
	heading := math.Mod((t.heading + number), 2*math.Pi)

	t.setHeadingRadians(heading)

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
	t.SetXY(patch.xFloat64, patch.yFloat64)
}

func (t *Turtle) MoveToTurtle(turtle *Turtle) {
	t.SetXY(turtle.xcor, turtle.ycor)
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

// returns the end of the given link that is not the current turtle
func (t *Turtle) OtherEnd(link *Link) *Turtle {
	if link.End1 == t {
		return link.End2
	}
	return link.End1
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

	dx := x - t.xcor
	dy := y - t.ycor

	x, y, allowed := t.parent.convertXYToInBounds(x, y)
	if !allowed {
		return
	}

	t.xcor = x
	t.ycor = y

	t.transferPatchOwnership()

	if len(t.linkedTurtles) == 0 {
		return
	}

	// move the linked turtles
	for turtle := range t.descendents(TieModeFree).turtles {
		t.moveTiedTurtle(turtle, dx, dy)
	}
}

func (t *Turtle) transferPatchOwnership() {
	oldPatch := t.patch
	t.patch = nil
	t.patch = t.PatchHere()
	if t.patch != oldPatch {
		oldPatch.removeTurtle(t)
		t.patch.addTurtle(t)
	}
}

func (t *Turtle) moveTiedTurtle(turtle *Turtle, dx float64, dy float64) {
	if t == turtle {
		return
	}

	newX := turtle.xcor + dx
	newY := turtle.ycor + dy

	newX, newY, allowed := t.parent.convertXYToInBounds(newX, newY)
	if !allowed {
		return
	}

	turtle.xcor = newX
	turtle.ycor = newY

	t.transferPatchOwnership()

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
