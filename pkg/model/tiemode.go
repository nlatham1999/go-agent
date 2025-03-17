package model

// descibes the behavior of a link and how the turtles tied to it should interact
type TieMode struct {
	MoveTiedTurtle   bool //if true when a parent or sibling turtle is moved the tied turtle is moved
	SwivelTiedTurtle bool //if true when a parent or sibling turtle is rotated the tied turtle is swivelled around
	RotateTiedTurtle bool //if true when a parent or sibling turtle is rotated the tied turtle is rotated
}

var (
	TieModeAllTied = TieMode{
		MoveTiedTurtle:   true,
		SwivelTiedTurtle: true,
		RotateTiedTurtle: true,
	}

	TieModeMoveAndRotate = TieMode{
		MoveTiedTurtle:   true,
		SwivelTiedTurtle: true,
		RotateTiedTurtle: false,
	}

	TieModeNone = TieMode{
		MoveTiedTurtle:   false,
		SwivelTiedTurtle: false,
		RotateTiedTurtle: false,
	}
)
