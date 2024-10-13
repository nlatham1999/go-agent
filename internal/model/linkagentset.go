package model

import (
	"math"
	"sort"
)

// LinkAgentSet is a set of links
type LinkAgentSet struct {
	links map[*Link]interface{}
}

// create a new LinkAgentSet
func LinkSet(links []*Link) *LinkAgentSet {
	newLinks := make(map[*Link]interface{})
	for _, link := range links {
		newLinks[link] = nil
	}

	return &LinkAgentSet{
		links: newLinks,
	}
}

// add a link to the agent set
func (l *LinkAgentSet) Add(link *Link) {
	l.links[link] = nil
}

// returns true if all the links in the agent set satisfy the operation
func (l *LinkAgentSet) All(operation LinkBoolOperation) bool {
	for link := range l.links {
		if !operation(link) {
			return false
		}
	}
	return true
}

// returns true if any of the links in the agent set satisfy the operation
func (l *LinkAgentSet) Any(operation LinkBoolOperation) bool {
	for link := range l.links {
		if operation(link) {
			return true
		}
	}
	return false
}

// perform the list of operations for all links in the agent set
func (l *LinkAgentSet) Ask(operations []LinkOperation) {
	for link := range l.links {
		for j := 0; j < len(operations); j++ {
			operations[j](link)
		}
	}
}

// returns true if the link is in the agent set
func (l *LinkAgentSet) Contains(link *Link) bool {
	_, ok := l.links[link]
	return ok
}

// returns the length of the agent set
func (l *LinkAgentSet) Count() int {
	return len(l.links)
}

// returns the agent set as a list
func (l *LinkAgentSet) List() []*Link {
	links := make([]*Link, 0)
	for link := range l.links {
		links = append(links, link)
	}
	return links
}

// returns the top n links in the agent set based on the float operation
func (l *LinkAgentSet) MaxNOf(n int, operation LinkFloatOperation) *LinkAgentSet {
	if n < 0 {
		return nil
	}

	links := l.List()
	sorter := &LinkSorter{links, operation, false}
	sort.Sort(sorter)

	if n > len(links) {
		n = len(links)
	}

	return LinkSet(links[:n])
}

// returns the max link in the agent set based on the float operation
func (l *LinkAgentSet) MaxOneOf(operation LinkFloatOperation) (*Link, error) {

	if len(l.links) == 0 {
		return nil, ErrNoLinksInAgentSet
	}

	max := math.MaxFloat64 * -1
	var maxLink *Link
	for link := range l.links {
		if operation(link) > max {
			max = operation(link)
			maxLink = link
		}
	}
	return maxLink, nil
}

// returns the min n links in the agent set based on the float operation
func (l *LinkAgentSet) MinNOf(n int, operation LinkFloatOperation) *LinkAgentSet {
	if n < 0 {
		return nil
	}

	links := l.List()
	sorter := &LinkSorter{links, operation, true}
	sort.Sort(sorter)

	if n > len(links) {
		n = len(links)
	}

	return LinkSet(links[:n])
}

// returns the min link in the agent set based on the float operation
func (l *LinkAgentSet) MinOneOf(operation LinkFloatOperation) (*Link, error) {

	if len(l.links) == 0 {
		return nil, ErrNoLinksInAgentSet
	}

	min := math.MaxFloat64
	var minLink *Link
	for link := range l.links {
		if operation(link) < min {
			min = operation(link)
			minLink = link
		}
	}
	return minLink, nil
}

// returns one of the links
// @TODO make this actually random based on model seed
func (l *LinkAgentSet) OneOf() (*Link, error) {
	for link := range l.links {
		return link, nil
	}

	return nil, ErrNoLinksInAgentSet
}

// returns n links or all the links in the agentset if the length is lower than n
func (l *LinkAgentSet) UpToNOf(n int) *LinkAgentSet {
	links := []*Link{}

	for link := range l.links {
		links = append(links, link)
		if len(links) == n {
			break
		}
	}

	return LinkSet(links)
}

// returns a new agent set with all the links that are not in the given agents set
func (l *LinkAgentSet) WhoAreNot(links *LinkAgentSet) *LinkAgentSet {
	linkMap := make(map[*Link]interface{})

	for link := range l.links {
		if _, ok := links.links[link]; !ok {
			linkMap[link] = nil
		}
	}

	return &LinkAgentSet{
		links: linkMap,
	}
}

// returns a new agent set with all the links that are not the given link
func (l *LinkAgentSet) WhoAreNotLink(link *Link) *LinkAgentSet {
	linkMap := make(map[*Link]interface{})

	for l1 := range l.links {
		if l1 != link {
			linkMap[l1] = nil
		}
	}

	return &LinkAgentSet{
		links: linkMap,
	}
}

// returns a new agent set that is a subset of the agent set where all satisfy the bool operation
func (l *LinkAgentSet) With(operation LinkBoolOperation) *LinkAgentSet {
	links := make([]*Link, 0)
	for link := range l.links {
		if operation(link) {
			links = append(links, link)
		}
	}
	return LinkSet(links)
}

// returns a subset of the agent set where all links are equal to the max value provided by the operation
func (l *LinkAgentSet) WithMax(operation LinkFloatOperation) *LinkAgentSet {
	max := math.MaxFloat64 * -1
	for link := range l.links {
		if operation(link) > max {
			max = operation(link)
		}
	}

	//get all links where the float operation is equal to the max
	links := make([]*Link, 0)
	for link := range l.links {
		if operation(link) == max {
			links = append(links, link)
		}
	}

	return LinkSet(links)
}

// returns a subset of the agent set where all links are equal to the min value provided by the operation
func (l *LinkAgentSet) WithMin(operation LinkFloatOperation) *LinkAgentSet {
	min := math.MaxFloat64
	for link := range l.links {
		if operation(link) < min {
			min = operation(link)
		}
	}

	//get all links where the float operation is equal to the min
	links := make([]*Link, 0)
	for link := range l.links {
		if operation(link) == min {
			links = append(links, link)
		}
	}

	return LinkSet(links)
}
