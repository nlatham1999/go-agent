//breeds are groups of turtles

package model

type TurtleBreed struct {
	Turtles *TurtleAgentSet

	name string

	defaultShape       string
	turtlesOwnTemplate map[string]interface{}
}

func (t *TurtleBreed) Name() string {
	return t.name
}
