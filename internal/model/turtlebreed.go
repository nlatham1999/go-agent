//breeds are groups of turtles

package model

// turtleBreed holds the agentset of the turtles belonging to the breed with the name, shape, turtles own variables
type turtleBreed struct {
	turtles *TurtleAgentSet

	name string

	defaultShape       string
	turtlesOwnTemplate map[string]interface{}
}
