package model

// describes the relation that a turtle has with other linked turtles

// due to the complexity of different maps, the member variables should not be accessed directly

type turtleLinks struct {

	// all the links linked to the turtle
	allLinksDirectedOut map[*Link]interface{}
	allLinksDirectedIn  map[*Link]interface{}
	allLinksUndirected  map[*Link]interface{}

	// maps of turtles to links
	allTurtlesDirectedOut map[*Turtle]*LinkAgentSet
	allTurtlesDirectedIn  map[*Turtle]*LinkAgentSet
	allTurtlesUndirected  map[*Turtle]*LinkAgentSet

	// maps of turtles to links split by link breed
	// unbreeded is stored as empty string
	turtlesDirectedOutBreed map[*LinkBreed]map[*Turtle]*Link
	turtlesDirectedInBreed  map[*LinkBreed]map[*Turtle]*Link
	turtlesUndirectedBreed  map[*LinkBreed]map[*Turtle]*Link
}

// links to connected turtles are stored in the turtleLinks struct
// this is to allow for quick access to the links that are connected to the turtle
// the links are stored in maps to allow for quick access to the links
// allLinks* are maps of all links that are connected to the turtle
// allTurtles* are maps of all turtles that are connected to the turtle
// turtles* are maps of turtles that are connected to the turtle split by breed
func newTurtleLinks() *turtleLinks {
	return &turtleLinks{
		allLinksDirectedOut: make(map[*Link]interface{}),
		allLinksDirectedIn:  make(map[*Link]interface{}),
		allLinksUndirected:  make(map[*Link]interface{}),

		allTurtlesDirectedOut: make(map[*Turtle]*LinkAgentSet),
		allTurtlesDirectedIn:  make(map[*Turtle]*LinkAgentSet),
		allTurtlesUndirected:  make(map[*Turtle]*LinkAgentSet),

		turtlesDirectedOutBreed: make(map[*LinkBreed]map[*Turtle]*Link),
		turtlesDirectedInBreed:  make(map[*LinkBreed]map[*Turtle]*Link),
		turtlesUndirectedBreed:  make(map[*LinkBreed]map[*Turtle]*Link),
	}
}

func (t *turtleLinks) addDirectedOutBreed(breed *LinkBreed, turtle *Turtle, link *Link) {
	if _, ok := t.turtlesDirectedOutBreed[breed]; !ok {
		t.turtlesDirectedOutBreed[breed] = make(map[*Turtle]*Link)
	}
	t.turtlesDirectedOutBreed[breed][turtle] = link

	if _, ok := t.allTurtlesDirectedOut[turtle]; !ok {
		t.allTurtlesDirectedOut[turtle] = NewLinkAgentSet([]*Link{})
	}
	t.allTurtlesDirectedOut[turtle].Add(link)

	t.allLinksDirectedOut[link] = nil
}

func (t *turtleLinks) addDirectedInBreed(breed *LinkBreed, turtle *Turtle, link *Link) {
	if _, ok := t.turtlesDirectedInBreed[breed]; !ok {
		t.turtlesDirectedInBreed[breed] = make(map[*Turtle]*Link)
	}
	t.turtlesDirectedInBreed[breed][turtle] = link

	if _, ok := t.allTurtlesDirectedIn[turtle]; !ok {
		t.allTurtlesDirectedIn[turtle] = NewLinkAgentSet([]*Link{})
	}
	t.allTurtlesDirectedIn[turtle].Add(link)

	t.allLinksDirectedIn[link] = nil
}

func (t *turtleLinks) addUndirectedBreed(breed *LinkBreed, turtle *Turtle, link *Link) {
	if _, ok := t.turtlesUndirectedBreed[breed]; !ok {
		t.turtlesUndirectedBreed[breed] = make(map[*Turtle]*Link)
	}
	t.turtlesUndirectedBreed[breed][turtle] = link

	if _, ok := t.allTurtlesUndirected[turtle]; !ok {
		t.allTurtlesUndirected[turtle] = NewLinkAgentSet([]*Link{})
	}
	t.allTurtlesUndirected[turtle].Add(link)

	t.allLinksUndirected[link] = nil
}

func (t *turtleLinks) removeDirectedOutBreed(breed *LinkBreed, turtle *Turtle, link *Link) {
	if _, ok := t.turtlesDirectedOutBreed[breed]; ok {
		delete(t.turtlesDirectedOutBreed[breed], turtle)
	}

	if _, ok := t.allTurtlesDirectedOut[turtle]; ok {
		// loop through and remove the link
		t.allTurtlesDirectedOut[turtle].Remove(link)

		//if the turtle has no more links, remove it
		if t.allTurtlesDirectedOut[turtle].Count() == 0 {
			delete(t.allTurtlesDirectedOut, turtle)
		}
	}

	delete(t.allLinksDirectedOut, link)
}

func (t *turtleLinks) removeDirectedInBreed(breed *LinkBreed, turtle *Turtle, link *Link) {
	if _, ok := t.turtlesDirectedInBreed[breed]; ok {
		delete(t.turtlesDirectedInBreed[breed], turtle)
	}

	if _, ok := t.allTurtlesDirectedIn[turtle]; ok {
		// loop through and remove the link
		t.allTurtlesDirectedIn[turtle].Remove(link)

		//if the turtle has no more links, remove it
		if t.allTurtlesDirectedIn[turtle].Count() == 0 {
			delete(t.allTurtlesDirectedIn, turtle)
		}
	}

	delete(t.allLinksDirectedIn, link)
}

func (t *turtleLinks) removeUndirectedBreed(breed *LinkBreed, turtle *Turtle, link *Link) {
	if _, ok := t.turtlesUndirectedBreed[breed]; ok {
		delete(t.turtlesUndirectedBreed[breed], turtle)
	}

	if _, ok := t.allTurtlesUndirected[turtle]; ok {
		// loop through and remove the link
		t.allTurtlesUndirected[turtle].Remove(link)

		//if the turtle has no more links, remove it
		if t.allTurtlesUndirected[turtle].Count() == 0 {
			delete(t.allTurtlesUndirected, turtle)
		}
	}

	delete(t.allLinksUndirected, link)
}

func (t *turtleLinks) changeDirectedOutBreed(oldLinkBreed *LinkBreed, newLinkBreed *LinkBreed, turtle *Turtle, link *Link) {
	if _, ok := t.turtlesDirectedOutBreed[oldLinkBreed]; ok {
		delete(t.turtlesDirectedOutBreed[oldLinkBreed], turtle)
	}

	if _, ok := t.turtlesDirectedOutBreed[newLinkBreed]; !ok {
		t.turtlesDirectedOutBreed[newLinkBreed] = make(map[*Turtle]*Link)
	}
	t.turtlesDirectedOutBreed[newLinkBreed][turtle] = link
}

func (t *turtleLinks) changeDirectedInBreed(oldLinkBreed *LinkBreed, newLinkBreed *LinkBreed, turtle *Turtle, link *Link) {
	if _, ok := t.turtlesDirectedInBreed[oldLinkBreed]; ok {
		delete(t.turtlesDirectedInBreed[oldLinkBreed], turtle)
	}

	if _, ok := t.turtlesDirectedInBreed[newLinkBreed]; !ok {
		t.turtlesDirectedInBreed[newLinkBreed] = make(map[*Turtle]*Link)
	}
	t.turtlesDirectedInBreed[newLinkBreed][turtle] = link
}

func (t *turtleLinks) changeUndirectedBreed(oldLinkBreed *LinkBreed, newLinkBreed *LinkBreed, turtle *Turtle, link *Link) {
	if _, ok := t.turtlesUndirectedBreed[oldLinkBreed]; ok {
		delete(t.turtlesUndirectedBreed[oldLinkBreed], turtle)
	}

	if _, ok := t.turtlesUndirectedBreed[newLinkBreed]; !ok {
		t.turtlesUndirectedBreed[newLinkBreed] = make(map[*Turtle]*Link)
	}
	t.turtlesUndirectedBreed[newLinkBreed][turtle] = link
}

// get a turtle that is connected from the current turtle
// provides a link that has a path from the current turtle to the target turtle
func (t turtleLinks) getLink(breed *LinkBreed, turtle *Turtle) *Link {
	if breed.name == BreedNone {
		// look in all directed
		if _, ok := t.allTurtlesDirectedOut[turtle]; ok {
			if t.allTurtlesDirectedOut[turtle].Count() > 0 {
				link, _ := t.allTurtlesDirectedOut[turtle].First()
				return link
			}
		}

		// look in all undirected
		if _, ok := t.allTurtlesUndirected[turtle]; ok {
			if t.allTurtlesUndirected[turtle].Count() > 0 {
				link, _ := t.allTurtlesUndirected[turtle].First()
				return link
			}
		}
	} else {
		// look in directed
		if link, ok := t.turtlesDirectedOutBreed[breed][turtle]; ok {
			return link
		}

		// look in undirected
		if link, ok := t.turtlesUndirectedBreed[breed][turtle]; ok {
			return link
		}
	}

	return nil
}

func (t *turtleLinks) getLinkDirected(breed *LinkBreed, turtle *Turtle) *Link {
	if breed.name == BreedNone {
		// look in all directed
		if _, ok := t.allTurtlesDirectedOut[turtle]; ok {
			if t.allTurtlesDirectedOut[turtle].Count() > 0 {
				link, _ := t.allTurtlesDirectedOut[turtle].First()
				return link
			}
		}
	} else {
		// look in directed
		if link, ok := t.turtlesDirectedOutBreed[breed][turtle]; ok {
			return link
		}
	}

	return nil
}

// get all the turtles that are connected to the current turtle
func (t *turtleLinks) getTurtlesIncoming(breed *LinkBreed) *TurtleAgentSet {
	agentSet := NewTurtleAgentSet(nil)
	if breed.name == BreedNone {
		for turtle := range t.allTurtlesDirectedIn {
			agentSet.Add(turtle)
		}
		for turtle := range t.allTurtlesUndirected {
			agentSet.Add(turtle)
		}
	} else {
		for turtle := range t.turtlesDirectedInBreed[breed] {
			agentSet.Add(turtle)
		}
		for turtle := range t.turtlesUndirectedBreed[breed] {
			agentSet.Add(turtle)
		}
	}

	return agentSet
}

// gets all the turtles that are connected from the current turtle with a path from the current turtle to the target turtle
// so undirected links are included
func (t *turtleLinks) getTurtlesOutgoing(breed *LinkBreed) *TurtleAgentSet {
	agentSet := NewTurtleAgentSet(nil)
	if breed.name == BreedNone {
		for turtle := range t.allTurtlesDirectedOut {
			agentSet.Add(turtle)
		}
		for turtle := range t.allTurtlesUndirected {
			agentSet.Add(turtle)
		}
	} else {
		for turtle := range t.turtlesDirectedOutBreed[breed] {
			agentSet.Add(turtle)
		}
		for turtle := range t.turtlesUndirectedBreed[breed] {
			agentSet.Add(turtle)
		}
	}

	return agentSet
}

func (t *turtleLinks) getTurtlesAll(breed *LinkBreed) *TurtleAgentSet {
	agentSet := NewTurtleAgentSet(nil)
	if breed.name == BreedNone {
		for turtle := range t.allTurtlesDirectedOut {
			agentSet.Add(turtle)
		}
		for turtle := range t.allTurtlesDirectedIn {
			agentSet.Add(turtle)
		}
		for turtle := range t.allTurtlesUndirected {
			agentSet.Add(turtle)
		}
	} else {
		for turtle := range t.turtlesDirectedOutBreed[breed] {
			agentSet.Add(turtle)
		}
		for turtle := range t.turtlesDirectedInBreed[breed] {
			agentSet.Add(turtle)
		}
		for turtle := range t.turtlesUndirectedBreed[breed] {
			agentSet.Add(turtle)
		}
	}

	return agentSet
}

// returns all links that have an incoming path to the turtle. This includes undirected links
func (t *turtleLinks) getLinksIncoming(breed *LinkBreed) *LinkAgentSet {
	if breed.name == BreedNone {
		links := make([]*Link, 0)
		for link := range t.allLinksDirectedIn {
			links = append(links, link)
		}
		for link := range t.allLinksUndirected {
			links = append(links, link)
		}
		return NewLinkAgentSet(links)
	} else {
		links := make([]*Link, 0)
		for _, link := range t.turtlesDirectedInBreed[breed] {
			links = append(links, link)
		}
		for _, link := range t.turtlesUndirectedBreed[breed] {
			links = append(links, link)
		}
		return NewLinkAgentSet(links)
	}
}

func (t turtleLinks) getLinksOutgoing(breed *LinkBreed) *LinkAgentSet {
	if breed.name == BreedNone {
		links := make([]*Link, 0)
		for link := range t.allLinksDirectedOut {
			links = append(links, link)
		}
		for link := range t.allLinksUndirected {
			links = append(links, link)
		}
		return NewLinkAgentSet(links)
	} else {
		links := make([]*Link, 0)
		for _, link := range t.turtlesDirectedOutBreed[breed] {
			links = append(links, link)
		}
		for _, link := range t.turtlesUndirectedBreed[breed] {
			links = append(links, link)
		}
		return NewLinkAgentSet(links)
	}
}

func (t *turtleLinks) getLinksAll(breed *LinkBreed) *LinkAgentSet {
	if breed.name == BreedNone {
		links := make([]*Link, 0)
		for link := range t.allLinksDirectedOut {
			links = append(links, link)
		}
		for link := range t.allLinksDirectedIn {
			links = append(links, link)
		}
		for link := range t.allLinksUndirected {
			links = append(links, link)
		}
		return NewLinkAgentSet(links)
	} else {
		links := make([]*Link, 0)
		for _, link := range t.turtlesDirectedOutBreed[breed] {
			links = append(links, link)
		}
		for _, link := range t.turtlesDirectedInBreed[breed] {
			links = append(links, link)
		}
		for _, link := range t.turtlesUndirectedBreed[breed] {
			links = append(links, link)
		}
		return NewLinkAgentSet(links)
	}
}

func (t *turtleLinks) count() int {
	return len(t.allLinksDirectedOut) + len(t.allLinksUndirected)
}

func (t *turtleLinks) getAllDirectedOutLinks() map[*Link]interface{} {
	return t.allLinksDirectedOut
}

func (t *turtleLinks) getAllDirectedInLinks() map[*Link]interface{} {
	return t.allLinksDirectedIn
}

func (t *turtleLinks) getAllUndirectedLinks() map[*Link]interface{} {
	return t.allLinksUndirected
}

func (t *turtleLinks) existsOutgoing(breed *LinkBreed, turtle *Turtle) bool {
	if breed.name == BreedNone {
		if v, ok := t.allTurtlesDirectedOut[turtle]; ok {
			return v.Count() > 0
		} else {
			return false
		}
	} else {
		// make sure the breed exists
		if _, ok := t.turtlesDirectedOutBreed[breed]; !ok {
			return false
		}
		_, ok := t.turtlesDirectedOutBreed[breed][turtle]
		return ok
	}
}

func (t *turtleLinks) existsIncoming(breed *LinkBreed, turtle *Turtle) bool {
	if breed.name == BreedNone {
		if v, ok := t.allTurtlesDirectedIn[turtle]; ok {
			return v.Count() > 0
		} else {
			return false
		}
	} else {
		// make sure the breed exists
		if _, ok := t.turtlesDirectedInBreed[breed]; !ok {
			return false
		}
		_, ok := t.turtlesDirectedInBreed[breed][turtle]
		return ok
	}
}

func (t *turtleLinks) existsUndirected(breed *LinkBreed, turtle *Turtle) bool {
	if breed.name == BreedNone {
		if v, ok := t.allTurtlesUndirected[turtle]; ok {
			return v.Count() > 0
		} else {
			return false
		}
	} else {
		// make sure the breed exists
		if _, ok := t.turtlesUndirectedBreed[breed]; !ok {
			return false
		}
		_, ok := t.turtlesUndirectedBreed[breed][turtle]
		return ok
	}
}
