package model

// LinkBreed holds the breed information for a link
type LinkBreed struct {
	links *LinkAgentSet
	model *Model
	name  string

	directed     bool
	defaultShape string
}

// NewLinkBreed creates a new link breed
func NewLinkBreed(name string) *LinkBreed {
	return &LinkBreed{
		name:     name,
		links:    NewLinkAgentSet(nil),
		directed: false, // should get set by the model after being passed in model settings
		model:    nil,   // should get set by the model after being passed in model settings
	}
}

func (lb *LinkBreed) Link(turtle1 int, turtle2 int) *Link {
	return lb.model.linkBreeded(lb, turtle1, turtle2)
}

func (lb *LinkBreed) Links() *LinkAgentSet {
	return lb.links
}

// sets the default shape for a directed link breed
func (lb *LinkBreed) SetDefaultShape(shape string) {
	lb.defaultShape = shape
}
