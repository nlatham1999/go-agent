package model

import (
	"fmt"
	"math"
)

type TieMode int

type Link struct {
	Color     Color
	End1      *Turtle
	End2      *Turtle
	Hidden    bool
	Directed  bool
	breed     string
	Shape     string
	Thickness float64
	TieMode   TieMode
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
		s := linkedTurtle{true, "", end2}
		end1.linkedTurtles[s] = l

		end2.linkedTurtlesConnectedFrom[end1] = l

		if breed != "" {
			s = linkedTurtle{true, breed, end2}
			end1.linkedTurtles[s] = l
		}
	} else {
		s1 := linkedTurtle{false, "", end2}
		s2 := linkedTurtle{false, "", end1}

		end1.linkedTurtles[s1] = l
		end2.linkedTurtles[s2] = l
		if breed != "" {
			s1 = linkedTurtle{false, breed, end2}
			s2 = linkedTurtle{false, breed, end1}

			end1.linkedTurtles[s1] = l
			end2.linkedTurtles[s2] = l
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

// returns the distance between the two ends of the link
func (l *Link) Length() float64 {
	return l.parent.DistanceBetweenPoints(l.End1.xcor, l.End1.ycor, l.End2.xcor, l.End2.ycor)
}

// returns the other end of the link that is not the given turtle
func (l *Link) OtherEnd(t *Turtle) *Turtle {
	if t == l.End1 {
		return l.End2
	} else {
		return l.End1
	}
}

func (l *Link) Show() {
	l.Hidden = false
}

func (l *Link) Tie() {
	l.TieMode = TieModeFixed
}

func (l *Link) Untie() {
	l.TieMode = TieModeNone
}
