//breeds are groups of turtles

package universe

type LinkBreed struct {
	Links []*Link

	Directed     bool
	DefaultShape string
}

type TurtleBreed struct {
	Turtles []*Turtle

	DefaultShape string
}
