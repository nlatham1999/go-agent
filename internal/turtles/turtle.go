package turtle

type Turtle struct {
	x     float64
	y     float64
	id    int
	size  int
	color string
}

func NewTurtle(id int) *Turtle {
	return &Turtle{
		id: id,
	}
}

//@TODO is this needed
//creates a new turtle from a template
//possible template attributes:
//	color
//	size
func NewTurtleFromTemplate(template *Turtle, id int) *Turtle {
	t := NewTurtle(id)
	t.color = template.color
	t.size = template.size

	return t
}
