package model

// TurtleBreed holds the agentset of the turtles belonging to the breed with the name, shape, turtles own variables
type TurtleBreed struct {
	turtles *TurtleAgentSet

	model *Model

	name string

	defaultShape string

	turtlePropertiesTemplate map[string]interface{}
}

// creates a new turtle breed
// after creating, pass it in the model settings
func NewTurtleBreed(name string, defaultShape string, turtleProperties map[string]interface{}) *TurtleBreed {
	return &TurtleBreed{
		name:                     name,
		model:                    nil,
		turtles:                  NewTurtleAgentSet(nil),
		turtlePropertiesTemplate: turtleProperties,
	}
}

// create the specified amount of turtles with the specified operation
func (tb *TurtleBreed) CreateAgents(amount int, operation TurtleOperation) (*TurtleAgentSet, error) {
	return tb.model.createTurtlesBreeded(amount, tb, operation)
}

// sets the default shape for a turtle breed
func (tb *TurtleBreed) SetDefaultShape(shape string) {
	tb.defaultShape = shape
}

// returns a turtle given the who number if it is in the breed
func (tb *TurtleBreed) Agent(who int) *Turtle {
	t := tb.model.whoToTurtles[who]
	if t == nil {
		return nil //turtle not found
	}

	if t.breed != tb {
		return nil //turtle not in this breed
	}

	return t
}

// returns the turtles in the breed
func (tb *TurtleBreed) Agents() *TurtleAgentSet {
	return tb.turtles
}

// returns the turtle agentset that is on patch of the proviced x y coordinates that belong to the breed
// same as TurtlesOnPatch(breed, Patch(x, y))
func (tb *TurtleBreed) AgentsAtCoords(pxcor float64, pycor float64) *TurtleAgentSet {
	return tb.model.turtlesAtCoordsBreeded(tb, pxcor, pycor)
}

// returns the turtle agentset that is on the provided patch
func (tb *TurtleBreed) AgentsOnPatch(patch *Patch) *TurtleAgentSet {
	return tb.model.turtlesOnPatchBreeded(tb, patch)
}

// Returns the turtles on the provided patches
func (tb *TurtleBreed) AgentsOnPatches(patches *PatchAgentSet) *TurtleAgentSet {
	return tb.model.turtlesOnPatchesBreeded(tb, patches)
}

// Returns the turtles on the same patch as the provided turtle
func (tb *TurtleBreed) AgentsWithAgent(turtle *Turtle) *TurtleAgentSet {
	return tb.model.turtlesWithTurtleBreeded(tb, turtle)
}

// Returns the turtles on the same patch as the provided turtle
func (tb *TurtleBreed) AgentsWithAgents(turtles *TurtleAgentSet) *TurtleAgentSet {
	return tb.model.turtlesWithTurtlesBreeded(tb, turtles)
}
