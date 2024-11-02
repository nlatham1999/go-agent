package model

import (
	"fmt"
	"math"
)

// Turtle is an agent that can move around the world
// it can have links to other turtles
type Turtle struct {
	xcor float64
	ycor float64
	who  int //the id of the turtle
	size float64

	Color   Color
	heading float64 //direction the turtle is facing in radians
	Hidden  bool    //if the turtle is hidden
	breed   string
	Shape   string

	parent *Model //model the turtle belongs too

	label      interface{}
	labelColor Color

	turtlesOwnGeneral map[string]interface{} // turtles own variables
	turtlesOwnBreed   map[string]interface{} // turtles own variables

	// turtles the current turtle is linked to/by/with
	linkedTurtles *turtleLinks

	patch *Patch //patch the turtle is on
}

func newTurtle(m *Model, who int, breed string, x float64, y float64) *Turtle {

	if m == nil {
		return nil
	}

	//if the breed is nonexistent then return nil
	var breedSet *turtleBreed = nil
	if breed != "" {
		found := false
		if breedSet, found = m.breeds[breed]; !found {
			return nil
		}
	}

	t := &Turtle{
		who:           who,
		parent:        m,
		xcor:          x,
		ycor:          y,
		breed:         breed,
		linkedTurtles: newTurtleLinks(),
		size:          .8,
		label:         "",
		labelColor:    Black,
	}

	m.turtles.Add(t)
	m.whoToTurtles[m.turtlesWhoNumber] = t

	if breedSet != nil {
		breedSet.turtles.Add(t)
	}

	//link the turtle to the patch
	t.patch = t.PatchHere()
	t.patch.addTurtle(t)

	//set the turtle own variables
	//breed specific variables can override general variables
	t.turtlesOwnGeneral = make(map[string]interface{})
	t.turtlesOwnBreed = make(map[string]interface{})
	generalTemplate := m.breeds[""].turtlesOwnTemplate
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

// moves the turtle backwards by the distance passed in and in relation to its heading
func (t *Turtle) Back(distance float64) {
	t.Forward(-distance)
}

// returns the breed of the turtle
func (t *Turtle) BreedName() string {
	return t.breed
}

// Sets the breed of the turtle to the name passed in
func (t *Turtle) SetBreed(name string) {

	if t.breed == name {
		return
	}

	// remove the turtle from the patch and add it back in at the end
	t.patch.removeTurtle(t)
	defer t.patch.addTurtle(t)

	if t.breed != "" {
		t.parent.breeds[t.breed].turtles.Remove(t)
	}

	t.breed = name

	if name != "" {
		t.parent.breeds[name].turtles.Add(t)
	}

	// switch the turtles own variables to the new breed
	t.turtlesOwnBreed = make(map[string]interface{})
	if name != "" {
		breedTemplate := t.parent.breeds[name].turtlesOwnTemplate
		for key, value := range breedTemplate {
			t.turtlesOwnBreed[key] = value
		}
	}

	// add the turtle to the patch

}

// returns if the turtle can move foward by the distance passed in
func (t *Turtle) CanMove(distance float64) bool {
	if t.parent.wrappingY && t.parent.wrappingX {
		return true
	}

	newX := t.xcor + distance*math.Cos(t.heading)
	newY := t.ycor + distance*math.Sin(t.heading)

	if newX < float64(t.parent.minPxCor) || newX >= float64(t.parent.maxPxCor) {
		if !t.parent.wrappingX {
			return false
		} else {
			return true
		}
	}

	if newY < float64(t.parent.minPyCor) || newY >= float64(t.parent.maxPyCor) {
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

func (t *Turtle) descendents(checkForRotated bool, checkForMoving bool, checkForSwivelling bool) *TurtleAgentSet {
	d := NewTurtleAgentSet([]*Turtle{})
	outgoing := t.linkedTurtles.getLinksOutgoing("")
	for outgoing.Count() > 0 {
		l, _ := outgoing.First()

		if checkForRotated && !l.TieMode.RotateTiedTurtle {
			outgoing.Remove(l)
			continue
		}

		if checkForMoving && !l.TieMode.MoveTiedTurtle {
			outgoing.Remove(l)
			continue
		}

		if checkForSwivelling && !l.TieMode.SwivelTiedTurtle {
			outgoing.Remove(l)
			continue
		}

		t1 := l.end1
		t2 := l.end2

		if t1 != t && !d.Contains(t1) {
			d.Add(t1)
			nextLinks := t1.linkedTurtles.getLinksOutgoing("")
			nextLinks.Ask(func(l2 *Link) {
				if d.Contains(l2.end2) && d.Contains(l2.end1) {
					return
				}
				if !outgoing.Contains(l2) {
					outgoing.Add(l2)
				}
			})
		}

		if t2 != t && !d.Contains(t2) {
			d.Add(t2)
			nextLinks := t2.linkedTurtles.getLinksOutgoing("")
			nextLinks.Ask(func(l2 *Link) {
				if d.Contains(l2.end2) && d.Contains(l2.end1) {
					return
				}
				if !outgoing.Contains(l2) {
					outgoing.Add(l2)
				}
			})
		}

		outgoing.Remove(l)
	}
	return d
}

// kill the turtle
func (t *Turtle) Die() {
	t.parent.KillTurtle(t)
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
		t.setHeadingRadians(math.Atan2(dy, dx))
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

	deltaXInverse := float64(t.parent.worldWidth) - math.Abs(dx)
	deltaYInverse := float64(t.parent.worldHeight) - math.Abs(dy)

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

	a := math.Atan2(newDy, newDx)
	t.setHeadingRadians(a)
}

// moves the turtle forward by the distance passed in and in relation to its heading
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
func (t *Turtle) Hatch(breed string, amount int, operation TurtleOperation) {

	turtles := make([]*Turtle, amount)
	for i := 0; i < amount; i++ {
		newBreed := t.breed
		if breed != "" {
			newBreed = breed
		}
		turtles[i] = newTurtle(t.parent, t.parent.turtlesWhoNumber, newBreed, t.xcor, t.ycor)
		t.parent.turtlesWhoNumber++

		// copy the variables
		turtles[i].Color = t.Color
		turtles[i].heading = t.heading
		turtles[i].Hidden = t.Hidden
		turtles[i].Shape = t.Shape
		turtles[i].size = t.size
		turtles[i].label = t.label
		turtles[i].labelColor = t.labelColor

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
		if operation != nil {
			operation(turtle)
		}
	}
}

// returns the heading of the turtle in degrees
func (t *Turtle) GetHeading() float64 {
	return radiansToDegrees(t.heading)
}

// takes in the heading in degrees and sets the heading of the turtle
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

	if t.linkedTurtles.count() == 0 {
		return
	}

	swivellDescendents := t.descendents(false, false, true)
	rotateDescendents := t.descendents(true, false, false)

	// rotate the heading for all descendents where the tiemode is at least free
	swivellDescendents.Ask(func(turtle *Turtle) {
		t.rotateTiedTurtle(turtle, headingDifference)
	})

	// rotate the heading for all descendents where the tiemode is fixed
	rotateDescendents.Ask(func(turtle *Turtle) {
		turtle.heading += headingDifference
	})
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

// Hide the turtle
func (t *Turtle) Hide() {
	t.Hidden = true
}

// sets the turtle to be in the middle of the world
func (t *Turtle) Home() {
	t.SetXY(0, 0)
}

// jumps ahead by the distance, if it cannot then it returns false
func (t *Turtle) Jump(distance float64) {
	// if t.CanMove(distance) {
	xcor := t.xcor + distance*math.Cos(t.heading)
	ycor := t.ycor + distance*math.Sin(t.heading)
	t.SetXY(xcor, ycor)
	// }
}

func (t *Turtle) SetLabel(label interface{}) {
	if label == nil {
		return
	}
	t.label = label
}

func (t *Turtle) GetLabel() interface{} {
	return t.label
}

func (t *Turtle) SetLabelColor(color Color) {
	t.labelColor = color
}

func (t *Turtle) GetLabelColor() Color {
	return t.labelColor
}

func (t *Turtle) Left(number float64) {
	// convert number to radians
	number = -number * (math.Pi / 180)

	// add the number to the heading
	heading := math.Mod((t.heading + number), 2*math.Pi)

	t.setHeadingRadians(heading)

}

// moves the turtle to the patch passed in
func (t *Turtle) MoveToPatch(patch *Patch) {
	if patch == nil {
		return
	}

	t.SetXY(patch.xFloat64, patch.yFloat64)
}

// moves the turtle to the x y coordinates passed in
func (t *Turtle) MoveToTurtle(turtle *Turtle) {
	if turtle == nil {
		return
	}

	t.SetXY(turtle.xcor, turtle.ycor)
}

// returns the neighbors of the patch that the turtle is on
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

// returns the 4 neighbors of the patch that the turtle is on in the cardinal directions
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

// returns the end of the given link that is not the current turtle
func (t *Turtle) OtherEnd(link *Link) *Turtle {
	if link.end1 == t {
		return link.end2
	}
	return link.end1
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

// returns the turtle own variable as an int
func (t *Turtle) GetOwnI(key string) (int, error) {
	v := t.GetOwn(key)
	if v == nil {
		return 0, fmt.Errorf(fmt.Sprintf("key not found: %s", key))
	}
	switch v := v.(type) {
	case int:
		return v, nil
	case float64:
		return int(v), nil
	default:
		return 0, fmt.Errorf("not a number")
	}
}

// returns the turtle own variable as a float
func (t *Turtle) GetOwnF(key string) (float64, error) {
	v := t.GetOwn(key)
	if v == nil {
		return 0, fmt.Errorf(fmt.Sprintf("key not found: %s", key))
	}
	switch v := v.(type) {
	case int:
		return float64(v), nil
	case float64:
		return v, nil
	default:
		return 0, fmt.Errorf("not a number")
	}
}

// returns the turtle own variable as a string
func (t *Turtle) GetOwnS(key string) (string, error) {
	v := t.GetOwn(key)
	if v == nil {
		return "", fmt.Errorf(fmt.Sprintf("key not found: %s", key))
	}
	switch v := v.(type) {
	case string:
		return v, nil
	default:
		return "", fmt.Errorf("not a string")
	}
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

// returns the patch that is ahead of the turtle by the distance passed in in relation to its heading
func (t *Turtle) PatchAhead(distance float64) *Patch {
	distX := t.xcor + distance*math.Cos(t.heading)
	distY := t.ycor + distance*math.Sin(t.heading)
	return t.parent.Patch(distX, distY)
}

// returns that patch that is the the left and right of the turtle by the coordinates passed in
func (t *Turtle) PatchAt(dx float64, dy float64) *Patch {

	//round the coords
	px := t.xcor + dx
	py := t.ycor + dy

	return t.parent.Patch(px, py)
}

// returns the patch that is the left of the turtle by the distance passed in and in relation to the heading passed in
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
	return t.parent.Patch(distX, distY)
}

func (t *Turtle) PatchRightAndAhead(angle float64, distance float64) *Patch {
	rightHeading := t.heading - angle
	distX := t.xcor + distance*math.Cos(rightHeading)
	distY := t.ycor + distance*math.Sin(rightHeading)
	return t.parent.Patch(distX, distY)
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

	if t.linkedTurtles.count() == 0 {
		return
	}

	moveDescendents := t.descendents(false, true, false)

	// move the linked turtles
	moveDescendents.Ask(func(turtle *Turtle) {
		t.moveTiedTurtle(turtle, dx, dy)
	})
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

func (t *Turtle) SetSize(size float64) {
	t.size = size
}

func (t *Turtle) GetSize() float64 {
	return t.size
}

func (t *Turtle) TowardsPatch(patch *Patch) float64 {
	//returns heading that faces the patch
	return t.TowardsXY(patch.xFloat64, patch.yFloat64)
}

func (t *Turtle) TowardsTurtle(turtle *Turtle) float64 {
	//returns heading that faces the turtle
	return t.TowardsXY(turtle.xcor, turtle.ycor)
}

func (t *Turtle) TowardsXY(x float64, y float64) float64 {
	//returns heading that faces the x y coordinates
	return radiansToDegrees(math.Atan2(y-t.ycor, x-t.xcor))
}

// returns the turtles that are on the patch
func (t *Turtle) TurtlesHere(breed string) *TurtleAgentSet {
	return t.PatchHere().TurtlesHere(breed)
}

func (t *Turtle) Uphill(patchVariable string) {
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
		if patch.patchesOwn[patchVariable].(float64) > minValue {
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

func (t *Turtle) Uphill4(patchVariable string) {

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
		if patch.patchesOwn[patchVariable].(float64) > minValue {
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

func (t *Turtle) Who() int {
	return t.who
}

func (t *Turtle) XCor() float64 {
	return t.xcor
}

func (t *Turtle) YCor() float64 {
	return t.ycor
}