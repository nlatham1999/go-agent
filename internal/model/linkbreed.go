package model

type LinkBreed struct {
	Links *LinkAgentSet

	name string

	Directed     bool
	DefaultShape string
}

func (l *LinkBreed) Name() string {
	return l.name
}
