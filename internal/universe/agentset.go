package universe

type TurtleAgentSet struct {
	turtles []*Turtle
}

type PatchAgentSet struct {
	patches []*Patch
}

type LinkAgentSet struct {
	links []*Link
}

// @TODO implement
func LinkSet(links []*Link) *LinkAgentSet {
	return nil
}

// @TODO implement
func PatchSet(patches []*Patch) *PatchAgentSet {
	return nil
}

// @TODO implement
func TurtleSet(turtles []*Turtle) *TurtleAgentSet {
	return nil
}
