package model

//operation types

// general function for acting on a link
type LinkOperation func(t *Link)

// general function that takes in a link and returns a bool
type LinkBoolOperation func(t *Link) bool

// general function that takes in a link and returns a float
type LinkFloatOperation func(t *Link) float64

// general function for acting on a patch
type PatchOperation func(t *Patch)

// general function that takes in a patch and returns a bool
type PatchBoolOperation func(t *Patch) bool

// general function that takes in a patch and returns a float
type PatchFloatOperation func(t *Patch) float64

// general function for acting on a turtle
type TurtleOperation func(t *Turtle)

// general function that takes in a turtle and returns a bool
type TurtleBoolOperation func(t *Turtle) bool

// general function that takes in a turtle and returns a float
type TurtleFloatOperation func(t *Turtle) float64

type LinkSorter struct {
	links []*Link
	f     LinkFloatOperation
}

func (l *LinkSorter) Len() int {
	return len(l.links)
}

func (l *LinkSorter) Less(i, j int) bool {
	return l.f(l.links[i]) > l.f(l.links[j])
}

func (l *LinkSorter) Swap(i, j int) {
	l.links[i], l.links[j] = l.links[j], l.links[i]
}

type PatchSorter struct {
	patches []*Patch
	f       PatchFloatOperation
}

func (p *PatchSorter) Len() int {
	return len(p.patches)
}

func (p *PatchSorter) Less(i, j int) bool {
	return p.f(p.patches[i]) > p.f(p.patches[j])
}

func (p *PatchSorter) Swap(i, j int) {
	p.patches[i], p.patches[j] = p.patches[j], p.patches[i]
}

type TurtleSorter struct {
	turtles []*Turtle
	f       TurtleFloatOperation
}

func (t *TurtleSorter) Len() int {
	return len(t.turtles)
}

func (t *TurtleSorter) Less(i, j int) bool {
	return t.f(t.turtles[i]) > t.f(t.turtles[j])
}

func (t *TurtleSorter) Swap(i, j int) {
	t.turtles[i], t.turtles[j] = t.turtles[j], t.turtles[i]
}
