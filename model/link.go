package model

import (
	"fmt"
	"math"
)

// A Link represents a connection between two turtles
// A Link can be be a one way directed link or a two way undirected link
type Link struct {
	Color      Color       // Color of the link
	end1       *Turtle     // the two ends of the link
	end2       *Turtle     // the two ends of the link
	Hidden     bool        // whether the link is hidden or not
	directed   bool        // whether the link is directed or not
	breed      *LinkBreed  // Breed of the link
	Shape      string      // Shape of the link
	Thickness  float64     // Thickness of the link
	TieMode    TieMode     // Tie mode of the link
	parent     *Model      // the parent model
	Size       int         // Size of the link
	Label      interface{} // Label of the link
	LabelColor Color       // Color of the label
}

// newLink creates a new link between two turtles
func newLink(model *Model, breed *LinkBreed, end1 *Turtle, end2 *Turtle, directed bool) (*Link, error) {

	// make sure the link doesn't already exist
	if directed && model.linkedTurtles[end1].existsOutgoing(breed, end2) {
		return nil, fmt.Errorf("Link already exists")
	}

	if !directed && model.linkedTurtles[end1].existsUndirected(breed, end2) {
		return nil, fmt.Errorf("Link already exists")
	}

	l := &Link{
		breed:    breed,
		end1:     end1,
		end2:     end2,
		directed: directed,
		parent:   model,
		Size:     1,
		Hidden:   false,
		Color:    White,
	}

	model.links.Add(l)

	if directed {
		model.directedLinkBreeds[breed.name].links.Add(l)
		model.directedLinkBreeds[BreedNone].links.Add(l)
	} else {
		model.undirectedLinkBreeds[breed.name].links.Add(l)
		model.undirectedLinkBreeds[BreedNone].links.Add(l)
	}

	// add the link to the turtle's link map
	if directed {
		model.linkedTurtles[end1].addDirectedOutBreed(breed, end2, l)
		model.linkedTurtles[end2].addDirectedInBreed(breed, end1, l)
	} else {
		model.linkedTurtles[end1].addUndirectedBreed(breed, end2, l)
		model.linkedTurtles[end2].addUndirectedBreed(breed, end1, l)
	}

	return l, nil
}

// Returns the name of the breed of the link
func (l *Link) BreedName() string {
	return l.breed.name
}

// returns an agentset of the turtles at the ends of the link
func (l *Link) BothEnds() *TurtleAgentSet {
	return NewTurtleAgentSet([]*Turtle{l.end1, l.end2})
}

func (l *Link) Die() {
	l.parent.KillLink(l)
}

func (l *Link) Directed() bool {
	return l.directed
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
func (l *Link) SetBreed(breed *LinkBreed) {

	if l.breed == breed {
		return
	}

	oldBreed := l.breed

	// make sure the breed exists
	// we can only change to a breed that is the same directedness
	if l.directed {
		if _, ok := l.parent.directedLinkBreeds[breed.name]; !ok {
			return
		}
	} else {
		if _, ok := l.parent.undirectedLinkBreeds[breed.name]; !ok {
			return
		}
	}

	// remove the link from the old breed if it exists
	if l.breed.name != "" {
		var breed *LinkBreed
		if l.directed {
			breed = l.parent.directedLinkBreeds[l.breed.name]
		} else {
			breed = l.parent.undirectedLinkBreeds[l.breed.name]
		}

		breed.links.links.Remove(l)
	}

	l.breed = breed

	if l.breed.name != BreedNone {
		if l.directed {
			l.parent.directedLinkBreeds[l.breed.name].links.Add(l)
		} else {
			l.parent.undirectedLinkBreeds[l.breed.name].links.Add(l)
		}
	}

	//change the breed on the turtles
	if l.directed {
		l.parent.linkedTurtles[l.end1].changeDirectedOutBreed(oldBreed, breed, l.end2, l)
		l.parent.linkedTurtles[l.end2].changeDirectedInBreed(oldBreed, breed, l.end1, l)
	} else {
		l.parent.linkedTurtles[l.end1].changeUndirectedBreed(oldBreed, breed, l.end2, l)
		l.parent.linkedTurtles[l.end2].changeUndirectedBreed(oldBreed, breed, l.end1, l)
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

// sets the tie mode to be closely tied
func (l *Link) Tie() {
	l.TieMode.MoveTiedTurtle = true
	l.TieMode.SwivelTiedTurtle = true
	l.TieMode.RotateTiedTurtle = true
}

// sets the tie mode to completely free
func (l *Link) Untie() {
	l.TieMode.MoveTiedTurtle = false
	l.TieMode.SwivelTiedTurtle = false
	l.TieMode.RotateTiedTurtle = false
}
