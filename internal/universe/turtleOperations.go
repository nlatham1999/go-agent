package universe

//general function for acting on a turtle
type TurtleOperation func(t *Turtle)

func SetColor(color string) TurtleOperation {
	return func(t *Turtle) {
		t.Color = color
	}
}

func SetSize(size int) TurtleOperation {
	return func(t *Turtle) {
		t.size = size
	}
}
