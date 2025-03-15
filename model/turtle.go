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
	breed   *TurtleBreed
	Shape   string

	parent *Model //model the turtle belongs too

	label      interface{}
	labelColor Color

	turtlePropertiesGeneral map[string]interface{} // turtles own variables
	turtlePropertiesBreed   map[string]interface{} // turtle properties variables

	patch *Patch //patch the turtle is on
}

func newTurtle(m *Model, who int, breed *TurtleBreed, x float64, y float64) *Turtle {

	if m == nil {
		return nil
	}

	t := &Turtle{
		who:        who,
		parent:     m,
		xcor:       x,
		ycor:       y,
		breed:      breed,
		size:       .8,
		label:      "",
		labelColor: Black,
		Shape:      "circle",
	}

	// add in the linked turtles
	t.parent.linkedTurtles[t] = newTurtleLinks()

	m.turtles.Add(t)
	m.whoToTurtles[m.turtlesWhoNumber] = t

	if breed != nil {
		breed.turtles.Add(t)
	}

	//link the turtle to the patch
	t.patch = t.PatchHere()
	t.patch.addTurtle(t)

	//set the turtle properties variables
	//breed specific variables can override general variables
	t.turtlePropertiesGeneral = make(map[string]interface{})
	t.turtlePropertiesBreed = make(map[string]interface{})
	generalTemplate := m.breeds[""].turtlePropertiesTemplate
	for key, value := range generalTemplate {
		t.turtlePropertiesGeneral[key] = value
	}
	if breed != nil {
		breedTemplate := breed.turtlePropertiesTemplate
		for key, value := range breedTemplate {
			t.turtlePropertiesBreed[key] = value
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
	if t.breed == nil {
		return ""
	}

	return t.breed.name
}

// Sets the breed of the turtle to the name passed in
func (t *Turtle) SetBreed(breed *TurtleBreed) {

	if t.breed == breed {
		return
	}

	// remove the turtle from the patch and add it back in at the end
	t.patch.removeTurtle(t)
	defer t.patch.addTurtle(t)

	if t.breed != nil {
		t.parent.breeds[t.breed.name].turtles.Remove(t)
	}

	t.breed = breed

	if t.breed != nil {
		t.parent.breeds[t.breed.name].turtles.Add(t)
	}

	// switch the turtles properties variables to the new breed
	t.turtlePropertiesBreed = make(map[string]interface{})
	if t.breed != nil {
		breedTemplate := t.parent.breeds[t.breed.name].turtlePropertiesTemplate
		for key, value := range breedTemplate {
			t.turtlePropertiesBreed[key] = value
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
	if _, ok := t.parent.patchPropertiesTemplate[patchVariable].(float64); !ok {
		return
	}

	minPatch := p
	minValue := minPatch.patchProperties[patchVariable].(float64)

	for patch := range p.patchNeighborsMap {
		if patch == nil {
			continue
		}
		if patch.patchProperties[patchVariable].(float64) < minValue {
			minPatch = patch
			minValue = patch.patchProperties[patchVariable].(float64)
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
	if _, ok := t.parent.patchPropertiesTemplate[patchVariable].(float64); !ok {
		return
	}

	minPatch := t.PatchHere()
	minValue := minPatch.patchProperties[patchVariable].(float64)

	for patch := range neighborsMap {
		if patch.patchProperties[patchVariable].(float64) < minValue {
			minPatch = patch
			minValue = patch.patchProperties[patchVariable].(float64)
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
func (t *Turtle) Hatch(amount int, operation TurtleOperation) {

	turtles := make([]*Turtle, amount)
	for i := 0; i < amount; i++ {
		turtles[i] = newTurtle(t.parent, t.parent.turtlesWhoNumber, t.breed, t.xcor, t.ycor)
		t.parent.turtlesWhoNumber++

		// copy the variables
		turtles[i].Color = t.Color
		turtles[i].heading = t.heading
		turtles[i].Hidden = t.Hidden
		turtles[i].Shape = t.Shape
		turtles[i].size = t.size
		turtles[i].label = t.label
		turtles[i].labelColor = t.labelColor

		// copy the property variables
		for key, value := range t.turtlePropertiesGeneral {
			turtles[i].turtlePropertiesGeneral[key] = value
		}

		// copy the breed variables if the breed is the same
		for key, value := range t.turtlePropertiesBreed {
			turtles[i].turtlePropertiesBreed[key] = value
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

	if t.parent.linkedTurtles[t].count() == 0 {
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

// returns the turtle property variable
func (t *Turtle) GetProperty(key string) interface{} {
	if val, found := t.turtlePropertiesBreed[key]; found {
		return val
	}
	if val, found := t.turtlePropertiesGeneral[key]; found {
		return val
	}
	return nil
}

// returns the turtle property variable as an int
func (t *Turtle) GetPropI(key string) (int, error) {
	v := t.GetProperty(key)
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

// returns the turtle property variable as a float
func (t *Turtle) GetPropF(key string) (float64, error) {
	v := t.GetProperty(key)
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

// returns the turtle property variable as a string
func (t *Turtle) GetPropS(key string) (string, error) {
	v := t.GetProperty(key)
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

// sets the turtle property variable
func (t *Turtle) SetProperty(key string, value interface{}) {
	if _, found := t.turtlePropertiesBreed[key]; found {
		t.turtlePropertiesBreed[key] = value
		return
	} else {
		if _, found := t.turtlePropertiesGeneral[key]; found {
			t.turtlePropertiesGeneral[key] = value
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

	if t.parent.linkedTurtles[t].count() == 0 {
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

// returns the turtles that are on the patch regardless of breed
// if you want to get the turtles of a specific breed, use turtleBreed.TurtlesOnPatch(patch)
// @TODO this function, is it needed?
func (t *Turtle) TurtlesHere() *TurtleAgentSet {
	return t.PatchHere().turtlesHereBreeded("")
}

func (t *Turtle) Uphill(patchVariable string) {
	p := t.PatchHere()

	// if the patch variable is not a number then return
	if _, ok := t.parent.patchPropertiesTemplate[patchVariable].(float64); !ok {
		return
	}

	minPatch := p
	minValue := minPatch.patchProperties[patchVariable].(float64)

	for patch := range p.patchNeighborsMap {
		if patch == nil {
			continue
		}
		if patch.patchProperties[patchVariable].(float64) > minValue {
			minPatch = patch
			minValue = patch.patchProperties[patchVariable].(float64)
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
	if _, ok := t.parent.patchPropertiesTemplate[patchVariable].(float64); !ok {
		return
	}

	minPatch := t.PatchHere()
	minValue := minPatch.patchProperties[patchVariable].(float64)

	for patch := range neighborsMap {
		if patch.patchProperties[patchVariable].(float64) > minValue {
			minPatch = patch
			minValue = patch.patchProperties[patchVariable].(float64)
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
