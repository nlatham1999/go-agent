package model

import (
	"fmt"
	"math"
)

type TieMode int

// Link represents a link between two turtles
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
	Size      int

	Label      interface{}
	LabelColor Color
}

// NewLink creates a new link between two turtles
func NewLink(model *Model, breed string, end1 *Turtle, end2 *Turtle, directed bool) (*Link, error) {

	// make sure the breed exists
	if directed {
		if _, ok := model.directedLinkBreeds[breed]; !ok {
			return nil, fmt.Errorf("directed link breed %s does not exist", breed)
		}
	} else {
		if _, ok := model.undirectedLinkBreeds[breed]; !ok {
			return nil, fmt.Errorf("undirected link breed %s does not exist", breed)
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
		Size:     1,
	}

	model.links.Add(l)

	if directed {
		model.directedLinkBreeds[breed].links.Add(l)
		model.directedLinkBreeds[""].links.Add(l)
	} else {
		model.undirectedLinkBreeds[breed].links.Add(l)
		model.undirectedLinkBreeds[""].links.Add(l)
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

// Returns the name of the breed of the link
func (l *Link) BreedName() string {
	return l.breed
}

// Returns the breed of this link
func (l *Link) Breed() *linkBreed {
	if l.breed == "" {
		return nil
	}

	if l.Directed {
		return l.parent.directedLinkBreeds[l.breed]
	} else {
		return l.parent.undirectedLinkBreeds[l.breed]
	}
}

// returns an agentset of the turtles at the ends of the link
func (l *Link) BothEnds() *TurtleAgentSet {
	return TurtleSet([]*Turtle{l.end1, l.end2})
}

// takes in a list of link operations and runs them for the link
func (l *Link) Ask(operations []LinkOperation) {
	for j := 0; j < len(operations); j++ {
		operations[j](l)
	}
}

func (l *Link) Die() {
	l.parent.KillLink(l)
}

// returns the first end of the link
func (l *Link) End1() *Turtle {
	return l.end1
}

// returns the second end of the link
func (l *Link) End2() *Turtle {
	return l.end2
}

// sets the link to be a breed
func (l *Link) SetBreed(name string) {

	oldBreedName := l.breed

	// make sure the breed exists
	if l.Directed {
		if _, ok := l.parent.directedLinkBreeds[name]; !ok {
			return
		}
	} else {
		if _, ok := l.parent.undirectedLinkBreeds[name]; !ok {
			return
		}
	}

	// remove the link from the old breed if it exists
	if l.breed != "" {
		var breed *linkBreed
		if l.Directed {
			breed = l.parent.directedLinkBreeds[l.breed]
		} else {
			breed = l.parent.undirectedLinkBreeds[l.breed]
		}

		delete(breed.links.links, l)
	}

	l.breed = name

	if l.breed != "" {
		if l.Directed {
			l.parent.directedLinkBreeds[name].links.Add(l)
		} else {
			l.parent.undirectedLinkBreeds[name].links.Add(l)
		}
	}

	//change the breed on the turtles
	if l.Directed {
		l.end1.linkedTurtles.changeDirectedOutBreed(oldBreedName, name, l.end2, l)
		l.end2.linkedTurtles.changeDirectedInBreed(oldBreedName, name, l.end1, l)
	} else {
		l.end1.linkedTurtles.changeUndirectedBreed(oldBreedName, name, l.end2, l)
		l.end2.linkedTurtles.changeUndirectedBreed(oldBreedName, name, l.end1, l)
	}
}

// sets the link to be hidden
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

// sets the link to be visible
func (l *Link) Show() {
	l.Hidden = false
}

// sets the tie mode to be fixed
func (l *Link) Tie() {
	l.TieMode = TieModeFixed
}

// sets the tie mode to be none
func (l *Link) Untie() {
	l.TieMode = TieModeNone
}
