package universe

type Link struct {
	Base
	Color    Color
	End1     *Turtle
	End2     *Turtle
	Hidden   bool
	Directed bool
	Breed    string

	Label      interface{}
	LabelColor Color
}

// @TODO implement
func (t *Link) GetBreedName() string {
	return ""
}

// @TODO implement
func (t *Link) GetBreedSet() []*Link {
	return nil
}

// @TODO implement
func (t *Link) SetBreed(name string) {

}

func (t *Link) Hide() {
	t.Hidden = true
}

// @TODO implement
func (t *Link) Heading() float64 {
	return 0
}

// @TODO implement
func (t *Link) Length() float64 {
	return 0
}
