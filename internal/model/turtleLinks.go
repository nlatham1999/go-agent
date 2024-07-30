package model

// describes the relation that a turtle has with other linked turtles

// due to the different maps, the member variables should not be accessed directly

type turtleLinks struct {

	// all the links linked to the turtle
	allLinksDirectedOut map[*Link]interface{}
	allLinksDirectedIn  map[*Link]interface{}
	allLinksUndirected  map[*Link]interface{}

	// maps of turtles to links
	allTurtlesDirectedOut map[*Turtle][]*Link
	allTurtlesDirectedIn  map[*Turtle][]*Link
	allTurtlesUndirected  map[*Turtle][]*Link

	// maps of turtles to links split by breed
	// unbreeded is stored as empty string
	turtlesDirectedOutBreed map[string]map[*Turtle]*Link
	turtlesDirectedInBreed  map[string]map[*Turtle]*Link
	turtlesUndirectedBreed  map[string]map[*Turtle]*Link
}

func newTurtleLinks() *turtleLinks {
	return &turtleLinks{
		allLinksDirectedOut: make(map[*Link]interface{}),
		allLinksDirectedIn:  make(map[*Link]interface{}),
		allLinksUndirected:  make(map[*Link]interface{}),

		allTurtlesDirectedOut: make(map[*Turtle][]*Link),
		allTurtlesDirectedIn:  make(map[*Turtle][]*Link),
		allTurtlesUndirected:  make(map[*Turtle][]*Link),

		turtlesDirectedOutBreed: make(map[string]map[*Turtle]*Link),
		turtlesDirectedInBreed:  make(map[string]map[*Turtle]*Link),
		turtlesUndirectedBreed:  make(map[string]map[*Turtle]*Link),
	}
}

func (t *turtleLinks) addDirectedOutBreed(breed string, turtle *Turtle, link *Link) {
	if _, ok := t.turtlesDirectedOutBreed[breed]; !ok {
		t.turtlesDirectedOutBreed[breed] = make(map[*Turtle]*Link)
	}
	t.turtlesDirectedOutBreed[breed][turtle] = link

	if _, ok := t.allTurtlesDirectedOut[turtle]; !ok {
		t.allTurtlesDirectedOut[turtle] = []*Link{}
	}
	t.allTurtlesDirectedOut[turtle] = append(t.allTurtlesDirectedOut[turtle], link)

	t.allLinksDirectedOut[link] = nil
}

func (t *turtleLinks) addDirectedInBreed(breed string, turtle *Turtle, link *Link) {
	if _, ok := t.turtlesDirectedInBreed[breed]; !ok {
		t.turtlesDirectedInBreed[breed] = make(map[*Turtle]*Link)
	}
	t.turtlesDirectedInBreed[breed][turtle] = link

	if _, ok := t.allTurtlesDirectedIn[turtle]; !ok {
		t.allTurtlesDirectedIn[turtle] = []*Link{}
	}
	t.allTurtlesDirectedIn[turtle] = append(t.allTurtlesDirectedIn[turtle], link)

	t.allLinksDirectedIn[link] = nil
}

func (t *turtleLinks) addUndirectedBreed(breed string, turtle *Turtle, link *Link) {
	if _, ok := t.turtlesUndirectedBreed[breed]; !ok {
		t.turtlesUndirectedBreed[breed] = make(map[*Turtle]*Link)
	}
	t.turtlesUndirectedBreed[breed][turtle] = link

	if _, ok := t.allTurtlesUndirected[turtle]; !ok {
		t.allTurtlesUndirected[turtle] = []*Link{}
	}
	t.allTurtlesUndirected[turtle] = append(t.allTurtlesUndirected[turtle], link)

	t.allLinksUndirected[link] = nil
}

func (t *turtleLinks) removeDirectedOutBreed(breed string, turtle *Turtle, link *Link) {
	if _, ok := t.turtlesDirectedOutBreed[breed]; ok {
		delete(t.turtlesDirectedOutBreed[breed], turtle)
	}

	if _, ok := t.allTurtlesDirectedOut[turtle]; ok {
		// loop through and remove the link
		for i, l := range t.allTurtlesDirectedOut[turtle] {
			if l == link {
				t.allTurtlesDirectedOut[turtle] = append(t.allTurtlesDirectedOut[turtle][:i], t.allTurtlesDirectedOut[turtle][i+1:]...)
				break
			}
		}
	}

	delete(t.allLinksDirectedOut, link)
}

func (t *turtleLinks) removeDirectedInBreed(breed string, turtle *Turtle, link *Link) {
	if _, ok := t.turtlesDirectedInBreed[breed]; ok {
		delete(t.turtlesDirectedInBreed[breed], turtle)
	}

	if _, ok := t.allTurtlesDirectedIn[turtle]; ok {
		// loop through and remove the link
		for i, l := range t.allTurtlesDirectedIn[turtle] {
			if l == link {
				t.allTurtlesDirectedIn[turtle] = append(t.allTurtlesDirectedIn[turtle][:i], t.allTurtlesDirectedIn[turtle][i+1:]...)
				break
			}
		}
	}

	delete(t.allLinksDirectedIn, link)
}

func (t *turtleLinks) removeUndirectedBreed(breed string, turtle *Turtle, link *Link) {
	if _, ok := t.turtlesUndirectedBreed[breed]; ok {
		delete(t.turtlesUndirectedBreed[breed], turtle)
	}

	if _, ok := t.allTurtlesUndirected[turtle]; ok {
		// loop through and remove the link
		for i, l := range t.allTurtlesUndirected[turtle] {
			if l == link {
				t.allTurtlesUndirected[turtle] = append(t.allTurtlesUndirected[turtle][:i], t.allTurtlesUndirected[turtle][i+1:]...)
				break
			}
		}
	}

	delete(t.allLinksUndirected, link)
}

// get a turtle that is connected from the current turtle
// provides a link that has a path from the current turtle to the target turtle
func (t turtleLinks) getLink(breed string, turtle *Turtle) *Link {
	if breed == "" {
		// look in all directed
		for _, link := range t.allTurtlesDirectedOut[turtle] {
			return link
		}

		// look in all undirected
		for _, link := range t.allTurtlesUndirected[turtle] {
			return link
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

func (t *turtleLinks) exists(breed string, directed bool, turtle *Turtle) bool {
	if breed == "" {
		if directed {
			if v, ok := t.allTurtlesDirectedOut[turtle]; ok {
				return len(v) > 0
			} else {
				return false
			}
		} else {
			if v, ok := t.allTurtlesUndirected[turtle]; ok {
				return len(v) > 0
			} else {
				return false
			}
		}
	} else {
		if directed {
			_, ok := t.turtlesDirectedOutBreed[breed][turtle]
			return ok
		}

		_, ok := t.turtlesUndirectedBreed[breed][turtle]
		return ok
	}

}
