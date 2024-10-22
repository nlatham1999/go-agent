package model

//operation types

// general function for acting on a link
// func(l *Link)
type LinkOperation func(l *Link)

// general function that takes in a link and returns a bool
// func(l *Link) bool
type LinkBoolOperation func(l *Link) bool

// general function that takes in a link and returns a float
// func(l *Link) float64
type LinkFloatOperation func(l *Link) float64

// general function for acting on a patch
// func(p *Patch)
type PatchOperation func(p *Patch)

// general function that takes in a patch and returns a bool
// func(p *Patch) bool
type PatchBoolOperation func(p *Patch) bool

// general function that takes in a patch and returns a float
// func(p *Patch) float64
type PatchFloatOperation func(p *Patch) float64

// general function for acting on a turtle
// func(t *Turtle)
type TurtleOperation func(t *Turtle)

// general function that takes in a turtle and returns a bool
// func(t *Turtle) bool
type TurtleBoolOperation func(t *Turtle) bool

// general function that takes in a turtle and returns a float
// func(t *Turtle) float64
type TurtleFloatOperation func(t *Turtle) float64

type LinkSorter struct {
	links   []*Link
	f       LinkFloatOperation
	reverse bool
}

func (l *LinkSorter) Len() int {
	return len(l.links)
}

func (l *LinkSorter) Less(i, j int) bool {
	if l.reverse {
		return l.f(l.links[i]) < l.f(l.links[j])
	}
	return l.f(l.links[i]) > l.f(l.links[j])
}

func (l *LinkSorter) Swap(i, j int) {
	l.links[i], l.links[j] = l.links[j], l.links[i]
}

type PatchSorter struct {
	patches []*Patch
	f       PatchFloatOperation
	reverse bool
}

func (p *PatchSorter) Len() int {
	return len(p.patches)
}

func (p *PatchSorter) Less(i, j int) bool {
	if p.reverse {
		return p.f(p.patches[i]) < p.f(p.patches[j])
	}
	return p.f(p.patches[i]) > p.f(p.patches[j])
}

func (p *PatchSorter) Swap(i, j int) {
	p.patches[i], p.patches[j] = p.patches[j], p.patches[i]
}
