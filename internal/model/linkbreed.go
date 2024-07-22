package model

type LinkBreed struct {
	Links []*Link

	Name string

	Directed     bool
	DefaultShape string
}
