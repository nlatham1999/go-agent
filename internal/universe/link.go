package universe

type Link struct {
	Base
	Color     Color
	End1      *Turtle
	End2      *Turtle
	Hidden    bool
	Directed  bool
	Breed     string
	Shape     string
	Thickness float64
	TieMode   string

	Label      interface{}
	LabelColor Color
}

// @TODO implement
func (l *Link) GetBreedName() string {
	return ""
}

// @TODO implement
func (l *Link) GetBreedSet() []*Link {
	return nil
}

// @TODO implement
func (l *Link) SetBreed(name string) {

}

func (l *Link) Hide() {
	l.Hidden = true
}

// @TODO implement
func (l *Link) Heading() float64 {
	return 0
}

// @TODO implement
func (l *Link) Length() float64 {
	return 0
}

// @TODO implement
func (l *Link) OtherEnd(t *Turtle) *Turtle {
	return nil
}

func (l *Link) Show() {
	l.Hidden = false
}

// @TODO implement
func (l *Link) Tie() {

}
