package model

// LinkBreed holds the breed information for a link
type LinkBreed struct {
	links *LinkAgentSet

	name string

	Directed     bool
	DefaultShape string
}

// returns the agent set of the links that are part of the breed
func (l *LinkBreed) Links() *LinkAgentSet {
	return l.links
}

// returns the name of the links
func (l *LinkBreed) Name() string {
	return l.name
}
