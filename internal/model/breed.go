//breeds are groups of turtles

package model

type LinkBreed struct {
	Links []*Link

	Name string

	Directed     bool
	DefaultShape string
}

type TurtleBreed struct {
	Turtles *TurtleAgentSet

	Name string

	DefaultShape string
}
