//breeds are groups of turtles

package model

// TurtleBreed holds the agentset of the turtles belonging to the breed with the name, shape, turtles own variables
type TurtleBreed struct {
	turtles *TurtleAgentSet

	name string

	defaultShape       string
	turtlesOwnTemplate map[string]interface{}
}

// Get the name of the breed
func (t *TurtleBreed) Name() string {
	return t.name
}

// get the turtles of the breed
func (t *TurtleBreed) Turtles() *TurtleAgentSet {
	return t.turtles
}
