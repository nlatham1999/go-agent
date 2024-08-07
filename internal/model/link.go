package model

import (
	"fmt"
	"math"
)

type TieMode int

type Link struct {
	Color     Color
	end1      *Turtle
	end2      *Turtle
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

// @TODO return an error if the link already exists
func NewLink(model *Model, breed string, end1 *Turtle, end2 *Turtle, directed bool) (*Link, error) {

	// make sure the breed exists
	if directed {
		if _, ok := model.DirectedLinkBreeds[breed]; !ok {
			return nil, fmt.Errorf("Directed link breed %s does not exist", breed)
		}
	} else {
		if _, ok := model.UndirectedLinkBreeds[breed]; !ok {
			return nil, fmt.Errorf("Undirected link breed %s does not exist", breed)
		}
	}

	// make sure the link doesn't already exist
	if directed && end1.linkedTurtles.existsOutgoing(breed, end2) {
		return nil, fmt.Errorf("Link already exists")
	}

	if !directed && end1.linkedTurtles.existsUndirected(breed, end2) {
		return nil, fmt.Errorf("Link already exists")
	}

	l := &Link{
		breed:    breed,
		end1:     end1,
		end2:     end2,
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
		end1.linkedTurtles.addDirectedOutBreed(breed, end2, l)
		end2.linkedTurtles.addDirectedInBreed(breed, end1, l)
	} else {
		end1.linkedTurtles.addUndirectedBreed(breed, end2, l)
		end2.linkedTurtles.addUndirectedBreed(breed, end1, l)
	}

	return l, nil
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
	return TurtleSet([]*Turtle{l.end1, l.end2})
}

func (l *Link) Ask(operations []LinkOperation) {
	for j := 0; j < len(operations); j++ {
		operations[j](l)
	}
}

func (l *Link) End1() *Turtle {
	return l.end1
}

func (l *Link) End2() *Turtle {
	return l.end2
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
	if l.end1.xcor == l.end2.xcor && l.end1.ycor == l.end2.ycor {
		return 0, fmt.Errorf("Link has zero length")
	}

	// get the heading which is in radians
	heading := l.end1.heading - l.end2.heading

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
	return l.parent.DistanceBetweenPoints(l.end1.xcor, l.end1.ycor, l.end2.xcor, l.end2.ycor)
}

// returns the other end of the link that is not the given turtle
func (l *Link) OtherEnd(t *Turtle) *Turtle {
	if t == l.end1 {
		return l.end2
	} else {
		return l.end1
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
