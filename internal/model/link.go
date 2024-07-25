package model

import (
	"fmt"
	"math"
)

type Link struct {
	Color     Color
	End1      *Turtle
	End2      *Turtle
	Hidden    bool
	Directed  bool
	breed     string
	Shape     string
	Thickness float64
	TieMode   string
	parent    *Model

	Label      interface{}
	LabelColor Color
}

func NewLink(model *Model, breed string, end1 *Turtle, end2 *Turtle, directed bool) *Link {

	// make sure the breed exists
	if directed {
		if _, ok := model.DirectedLinkBreeds[breed]; !ok {
			return nil
		}
	} else {
		if _, ok := model.UndirectedLinkBreeds[breed]; !ok {
			return nil
		}
	}

	l := &Link{
		breed:    breed,
		End1:     end1,
		End2:     end2,
		Directed: directed,
		parent:   model,
	}

	model.Links.Add(l)

	if directed {
		model.DirectedLinkBreeds[breed].Links.Add(l)
		model.DirectedLinkBreeds[""].Links.Add(l)
	} else {
		model.UndirectedLinkBreeds[breed].Links.Add(l)
		model.UndirectedLinkBreeds[""].Links.Add(l)
	}

	// add the link to the turtle's link map
	if directed {
		if end1.linkedTurtles[true] == nil {
			end1.linkedTurtles[true] = make(map[string]map[*Turtle]*Link)
		}
		if end1.linkedTurtles[true][""] == nil {
			end1.linkedTurtles[true][""] = make(map[*Turtle]*Link)
		}
		end1.linkedTurtles[true][""][end2] = l
		if breed != "" {
			if end1.linkedTurtles[true][breed] == nil {
				end1.linkedTurtles[true][breed] = make(map[*Turtle]*Link)
			}
			end1.linkedTurtles[true][breed][end2] = l
		}
	} else {
		if end1.linkedTurtles[false] == nil {
			end1.linkedTurtles[false] = make(map[string]map[*Turtle]*Link)
		}
		if end1.linkedTurtles[false][""] == nil {
			end1.linkedTurtles[false][""] = make(map[*Turtle]*Link)
		}
		if end2.linkedTurtles[false] == nil {
			end2.linkedTurtles[false] = make(map[string]map[*Turtle]*Link)
		}
		if end2.linkedTurtles[false][""] == nil {
			end2.linkedTurtles[false][""] = make(map[*Turtle]*Link)
		}
		end1.linkedTurtles[false][""][end2] = l
		end2.linkedTurtles[false][""][end1] = l
		if breed != "" {
			if end1.linkedTurtles[false][breed] == nil {
				end1.linkedTurtles[false][breed] = make(map[*Turtle]*Link)
			}
			if end2.linkedTurtles[false][breed] == nil {
				end2.linkedTurtles[false][breed] = make(map[*Turtle]*Link)
			}
			end2.linkedTurtles[false][breed][end1] = l
			end1.linkedTurtles[false][breed][end2] = l
		}
	}

	return l
}

func (l *Link) BreedName() string {
	return l.breed
}

func (l *Link) Breed() *LinkBreed {
	if l.breed == "" {
		return nil
	}

	if l.Directed {
		return l.parent.DirectedLinkBreeds[l.breed]
	} else {
		return l.parent.UndirectedLinkBreeds[l.breed]
	}
}

// returns an agentset of the turtles at the ends of the link
func (l *Link) BothEnds() *TurtleAgentSet {
	return TurtleSet([]*Turtle{l.End1, l.End2})
}

func (l *Link) SetBreed(name string) {

	// make sure the breed exists
	if l.Directed {
		if _, ok := l.parent.DirectedLinkBreeds[name]; !ok {
			return
		}
	} else {
		if _, ok := l.parent.UndirectedLinkBreeds[name]; !ok {
			return
		}
	}

	// remove the link from the old breed if it exists
	if l.breed != "" {
		var breed *LinkBreed
		if l.Directed {
			breed = l.parent.DirectedLinkBreeds[l.breed]
		} else {
			breed = l.parent.UndirectedLinkBreeds[l.breed]
		}

		delete(breed.Links.links, l)
	}

	l.breed = name

	if l.breed != "" {
		if l.Directed {
			l.parent.DirectedLinkBreeds[name].Links.Add(l)
		} else {
			l.parent.UndirectedLinkBreeds[name].Links.Add(l)
		}
	}
}

func (l *Link) Hide() {
	l.Hidden = true
}

// returns the heading in degrees from end1 to end2. Returns an error if the link has zero length
func (l *Link) Heading() (float64, error) {
	if l.End1.xcor == l.End2.xcor && l.End1.ycor == l.End2.ycor {
		return 0, fmt.Errorf("Link has zero length")
	}

	// get the heading which is in radians
	heading := l.End1.heading - l.End2.heading

	// convert to degrees
	heading = heading * 180 / math.Pi

	// if negative, add 360 to make it positive
	if heading < 0 {
		heading += 360
	}

	return heading, nil
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

// @TODO implement
func (l *Link) Untie() {

}
