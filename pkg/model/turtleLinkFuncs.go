package model

// all api functions for turtles that deal with links

// creates a directed link from the current turtle to the turtle passed in
func (t *Turtle) CreateLinkToTurtle(breed *LinkBreed, turtle *Turtle, operation LinkOperation) (*Link, error) {

	// if the breed is nil, then use the default breed
	if breed == nil {
		breed = t.parent.directedLinkBreeds[BreedNone]
	}

	l, err := newLink(t.parent, breed, t, turtle, true)
	if err != nil {
		return l, err
	}

	if operation != nil {
		operation(l)
	}

	return l, nil
}

// creates a directed link from the current turtle to the turtles passed in
// if a link creation errors, than it is skipped
func (t *Turtle) CreateLinksToSet(breed *LinkBreed, turtles *TurtleAgentSet, operation LinkOperation) {

	// if the breed is nil, then use the default breed
	if breed == nil {
		breed = t.parent.directedLinkBreeds[BreedNone]
	}

	linksAdded := NewLinkAgentSet([]*Link{})
	turtles.Ask(func(turtle *Turtle) {
		l, err := newLink(t.parent, breed, t, turtle, true)
		if err != nil {
			return
		}
		linksAdded.Add(l)
	})

	linksAdded.Ask(operation)
}

// creates an undirected breed link from the current turtle with the turtle passed in
func (t *Turtle) CreateLinkWithTurtle(breed *LinkBreed, turtle *Turtle, operation LinkOperation) (*Link, error) {

	// if the breed is nil, then use the default breed
	if breed == nil {
		breed = t.parent.undirectedLinkBreeds[BreedNone]
	}

	l, err := newLink(t.parent, breed, t, turtle, false)
	if err != nil {
		return l, err
	}

	if operation != nil {
		operation(l)
	}

	return l, nil

}

// creates an undirected breed link from the current turtle with the turtles passed in
// if a link creation errors, than it is skipped
func (t *Turtle) CreateLinksWithSet(breed *LinkBreed, turtles *TurtleAgentSet, operation LinkOperation) {

	// if the breed is nil, then use the default breed
	if breed == nil {
		breed = t.parent.undirectedLinkBreeds[BreedNone]
	}

	linksAdded := NewLinkAgentSet([]*Link{})
	turtles.Ask(func(turtle *Turtle) {
		l, err := newLink(t.parent, breed, t, turtle, false)
		if err != nil {
			return
		}
		linksAdded.Add(l)
	})

	linksAdded.Ask(operation)
}

// creates a directed breed link from the current turtle with the turtle passed in
func (t *Turtle) CreateLinkFromTurtle(breed *LinkBreed, turtle *Turtle, operation LinkOperation) (*Link, error) {

	// if the breed is nil, then use the default breed
	if breed == nil {
		breed = t.parent.directedLinkBreeds[BreedNone]
	}

	l, err := newLink(t.parent, breed, turtle, t, true)
	if err != nil {
		return l, err
	}

	if operation != nil {
		operation(l)
	}

	return l, nil
}

// creates a directed breed link from the turtles passed in to the current turtle
// if a link creation errors, than it is skipped
func (t *Turtle) CreateLinksFromSet(breed *LinkBreed, turtles *TurtleAgentSet, operation LinkOperation) {

	// if the breed is nil, then use the default breed
	if breed == nil {
		breed = t.parent.directedLinkBreeds[BreedNone]
	}

	linksAdded := NewLinkAgentSet([]*Link{})
	turtles.Ask(func(turtle *Turtle) {
		l, err := newLink(t.parent, breed, turtle, t, true)
		if err != nil {
			return
		}
		linksAdded.Add(l)
	})

	linksAdded.Ask(operation)
}

// returns if there is any sort of link between the current turtle and the turtle passed in
func (t *Turtle) LinkExists(breed *LinkBreed, turtle *Turtle) bool {

	// if the breed is nil, then use the default breed
	if breed == nil {
		breed = t.parent.directedLinkBreeds[BreedNone]
	}

	return t.parent.linkedTurtles[t].existsIncoming(breed, turtle) || t.parent.linkedTurtles[t].existsOutgoing(breed, turtle) || t.parent.linkedTurtles[t].existsUndirected(breed, turtle)
}

// returns all turtles that are linked to the current turtle
//
//	incoming, outgoing, or undirected
func (t *Turtle) LinkedTurtles(breed *LinkBreed) *TurtleAgentSet {

	// if the breed is nil, then use the default breed
	if breed == nil {
		breed = t.parent.directedLinkBreeds[BreedNone]
	}

	return t.parent.linkedTurtles[t].getTurtlesAll(breed)
}

// returns if there is a directed link from turtle passed in to the current turtle or an undirected link connecting the two
func (t *Turtle) LinkToTurtleExists(breed *LinkBreed, turtle *Turtle) bool {

	// if the breed is nil, then use the default breed
	if breed == nil {
		breed = t.parent.directedLinkBreeds[BreedNone]
	}

	return t.parent.linkedTurtles[t].existsIncoming(breed, turtle) || t.parent.linkedTurtles[t].existsUndirected(breed, turtle)
}

// returns all turtles that have a directed link to the current turtle
//
//	or an undirected link connecting the two
//
// basically all turtles where there is a path from the turtle to the current turtle
func (t *Turtle) LinkedTurtlesToThis(breed *LinkBreed) *TurtleAgentSet {

	// if the breed is nil, then use the default breed
	if breed == nil {
		breed = t.parent.directedLinkBreeds[BreedNone]
	}

	return t.parent.linkedTurtles[t].getTurtlesIncoming(breed)
}

// returns whether there is a directed link connecting the current turtle to the turtle passed in or an undirected link connecting the two
// (this -> turtle) or (this <-> turtle)
func (t *Turtle) LinkFromTurtleExists(breed *LinkBreed, turtle *Turtle) bool {

	// if the breed is nil, then use the default breed
	if breed == nil {
		breed = t.parent.directedLinkBreeds[BreedNone]
	}

	return t.parent.linkedTurtles[t].existsOutgoing(breed, turtle) || t.parent.linkedTurtles[t].existsUndirected(breed, turtle)
}

// returns all turtles that have a directed link from the current turtle to them
func (t *Turtle) LinkedTurtlesFromThis(breed *LinkBreed) *TurtleAgentSet {

	// if the breed is nil, then use the default breed
	if breed == nil {
		breed = t.parent.directedLinkBreeds[BreedNone]
	}

	return t.parent.linkedTurtles[t].getTurtlesOutgoing(breed)
}

// finds a link from the turtle passed int to the current turtle (turtle -> this)
// to get all the links use InLinks
func (t *Turtle) LinkFrom(breed *LinkBreed, turtle *Turtle) *Link {

	// if the breed is nil, then use the default breed
	if breed == nil {
		breed = t.parent.directedLinkBreeds[BreedNone]
	}

	if turtle.parent.linkedTurtles[turtle] == nil {
		return nil
	}

	return turtle.parent.linkedTurtles[turtle].getLink(breed, t)
}

// finds a link from the current turtle to the turtle passed in (this -> turtle)
// To get all the links use OutLinks
func (t *Turtle) LinkTo(breed *LinkBreed, turtle *Turtle) *Link {

	// if the breed is nil, then use the default breed
	if breed == nil {
		breed = t.parent.directedLinkBreeds[BreedNone]
	}

	if t.parent.linkedTurtles[t] == nil {
		return nil
	}

	return t.parent.linkedTurtles[t].getLink(breed, turtle)
}

// finds a link between the current turtle and the turtle passed in (this <-> turtle)
// to get all the links use Links
func (t *Turtle) LinkWith(breed *LinkBreed, turtle *Turtle) *Link {

	// if the breed is nil, then use the default breed
	if breed == nil {
		breed = t.parent.undirectedLinkBreeds[BreedNone]
	}

	if t.parent.linkedTurtles[t] != nil {
		if link := t.parent.linkedTurtles[t].getLink(breed, turtle); link != nil {
			return link
		}
	}

	if turtle.parent.linkedTurtles[t] != nil {
		if link := turtle.parent.linkedTurtles[t].getLink(breed, t); link != nil {
			return link
		}
	}

	return nil
}

// returns all links that are connected to a turtle, undirected or directed, incoming or outgoing
func (t *Turtle) Links(breed *LinkBreed) *LinkAgentSet {

	// if the breed is nil, then use the default breed
	if breed == nil {
		breed = t.parent.directedLinkBreeds[BreedNone]
	}

	return t.parent.linkedTurtles[t].getLinksAll(breed)
}

// returns all incoming links that are connected to the turtle
// this includes directed links going in and undirected links
func (t *Turtle) InLinks(breed *LinkBreed) *LinkAgentSet {

	// if the breed is nil, then use the default breed
	if breed == nil {
		breed = t.parent.directedLinkBreeds[BreedNone]
	}

	return t.parent.linkedTurtles[t].getLinksIncoming(breed)
}

// returns all outgoing links that are connected from the turtle
// this includes directed links going out and undirected links
func (t *Turtle) OutLinks(breed *LinkBreed) *LinkAgentSet {

	// if the breed is nil, then use the default breed
	if breed == nil {
		breed = t.parent.directedLinkBreeds[BreedNone]
	}

	return t.parent.linkedTurtles[t].getLinksOutgoing(breed)
}

// returns the end of the given link that is not the current turtle
func (t *Turtle) OtherEnd(link *Link) *Turtle {
	if link.end1 == t {
		return link.end2
	}
	if link.end2 == t {
		return link.end1
	}
	return nil
}
