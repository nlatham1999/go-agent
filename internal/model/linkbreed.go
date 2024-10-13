package model

// linkBreed holds the breed information for a link
type linkBreed struct {
	links *LinkAgentSet

	name string

	Directed     bool
	DefaultShape string
}
