package model

// all api functions for turtles that deal with links

// creates a directed link from the current turtle to the turtle passed in
func (t *Turtle) CreateLinkToTurtle(breed string, turtle *Turtle, operations []LinkOperation) error {
	l, err := NewLink(t.parent, breed, t, turtle, true)
	if err != nil {
		return err
	}

	for _, operation := range operations {
		operation(l)
	}

	return nil
}

// creates a directed link from the current turtle to the turtles passed in
// if a link creation errors, than it is skipped
func (t *Turtle) CreateLinksToSet(breed string, turtles *TurtleAgentSet, operations []LinkOperation) {
	if breed == "" {
		return
	}

	linksAdded := LinkSet([]*Link{})
	for turtle := range turtles.turtles {
		l, err := NewLink(t.parent, breed, t, turtle, true)
		if err != nil {
			continue
		}
		linksAdded.Add(l)
	}

	linksAdded.Ask(operations)
}

// creates an undirected breed link from the current turtle with the turtle passed in
func (t *Turtle) CreateLinkWithTurtle(breed string, turtle *Turtle, operations []LinkOperation) error {
	l, err := NewLink(t.parent, breed, t, turtle, false)
	if err != nil {
		return err
	}

	for _, operation := range operations {
		operation(l)
	}

	return nil

}

// creates an undirected breed link from the current turtle with the turtles passed in
// if a link creation errors, than it is skipped
func (t *Turtle) CreateLinksWithSet(breed string, turtles *TurtleAgentSet, operations []LinkOperation) {
	linksAdded := LinkSet([]*Link{})
	for turtle := range turtles.turtles {
		l, err := NewLink(t.parent, breed, t, turtle, false)
		if err != nil {
			continue
		}
		linksAdded.Add(l)
	}

	linksAdded.Ask(operations)
}

// creates a directed breed link from the current turtle with the turtle passed in
func (t *Turtle) CreateLinkFromTurtle(breed string, turtle *Turtle, operations []LinkOperation) error {
	l, err := NewLink(t.parent, breed, turtle, t, true)
	if err != nil {
		return err
	}

	for _, operation := range operations {
		operation(l)
	}

	return nil
}

// creates a directed breed link from the turtles passed in to the current turtle
// if a link creation errors, than it is skipped
func (t *Turtle) CreateLinksFromSet(breed string, turtles *TurtleAgentSet, operations []LinkOperation) {
	linksAdded := LinkSet([]*Link{})
	for turtle := range turtles.turtles {
		l, err := NewLink(t.parent, breed, turtle, t, true)
		if err != nil {
			continue
		}
		linksAdded.Add(l)
	}

	linksAdded.Ask(operations)
}

// returns if there is any sort of link between the current turtle and the turtle passed in
func (t *Turtle) LinkNeighbor(breed string, turtle *Turtle) bool {
	return t.linkedTurtles.existsIncoming(breed, turtle) || t.linkedTurtles.existsOutgoing(breed, turtle) || t.linkedTurtles.existsUndirected(breed, turtle)
}

// returns all turtles that are linked to the current turtle
//
//	incoming, outgoing, or undirected
func (t *Turtle) LinkNeighbors(breed string) *TurtleAgentSet {
	return t.linkedTurtles.getTurtlesAll(breed)
}

// returns if there is a directed link from turtle to t or an undirected link connecting the two
func (t *Turtle) InLinkNeighbor(breed string, turtle *Turtle) bool {

	return t.linkedTurtles.existsIncoming(breed, turtle) || t.linkedTurtles.existsUndirected(breed, turtle)
}

// returns all turtles that have a directed link to the current turtle
//
//	or an undirected link connecting the two
//
// basically all turtles where there is a path from the turtle to the current turtle
func (t *Turtle) InLinkNeighbors(breed string) *TurtleAgentSet {
	return t.linkedTurtles.getTurtlesIncoming(breed)
}

// returns whether there is a directed link connecting t to turtle or an undirected link connecting the two
func (t *Turtle) OutLinkNeighbor(breed string, turtle *Turtle) bool {
	return t.linkedTurtles.existsOutgoing(breed, turtle) || t.linkedTurtles.existsUndirected(breed, turtle)
}

// returns all turtles that have a directed link from the current turtle to them
func (t *Turtle) OutLinkNeighbors(breed string) *TurtleAgentSet {
	return t.linkedTurtles.getTurtlesOutgoing(breed)
}

// finds a link from the turtle passed int to the current turtle
func (t *Turtle) LinkFrom(breed string, turtle *Turtle) *Link {

	if turtle.linkedTurtles == nil {
		return nil
	}

	return turtle.linkedTurtles.getLink(breed, t)
}

// finds a link from the current turtle to the turtle passed in
func (t *Turtle) LinkTo(breed string, turtle *Turtle) *Link {
	if t.linkedTurtles == nil {
		return nil
	}

	return t.linkedTurtles.getLink(breed, turtle)
}

// finds a link between the current turtle and the turtle passed in
func (t *Turtle) LinkWith(breed string, turtle *Turtle) *Link {
	if t.linkedTurtles != nil {
		if link := t.linkedTurtles.getLink(breed, turtle); link != nil {
			return link
		}
	}

	if turtle.linkedTurtles != nil {
		if link := turtle.linkedTurtles.getLink(breed, t); link != nil {
			return link
		}
	}

	return nil
}

// returns all links that are connected to a turtle, undirected or directed, incoming or outgoing
func (t *Turtle) MyLinks(breed string) *LinkAgentSet {
	return t.linkedTurtles.getLinksAll(breed)
}

// returns all incoming links that are connected to the turtle
func (t *Turtle) MyInLinks(breed string) *LinkAgentSet {
	return t.linkedTurtles.getLinksIncoming(breed)
}

// returns all outgoing links that are connected to the turtle
func (t *Turtle) MyOutLinks(breed string) *LinkAgentSet {
	return t.linkedTurtles.getLinksOutgoing(breed)
}
