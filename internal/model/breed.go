//breeds are groups of turtles

package model

type LinkBreed struct {
	Links []*Link

	Directed     bool
	DefaultShape string
}

type TurtleBreed struct {
	Turtles *TurtleAgentSet

	DefaultShape string
}
