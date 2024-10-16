package model

import (
	"math"

	"github.com/nlatham1999/sortedset"
)

// LinkAgentSet is a set of links
type LinkAgentSet struct {
	links sortedset.SortedSet
}

// create a new LinkAgentSet
func LinkSet(links []*Link) *LinkAgentSet {

	linkSet := &LinkAgentSet{
		links: *sortedset.NewSortedSet(),
	}

	for _, link := range links {
		linkSet.Add(link)
	}

	return linkSet
}

// add a link to the agent set
func (l *LinkAgentSet) Add(link *Link) {
	l.links.Add(link)
}

// returns true if all the links in the agent set satisfy the operation
func (l *LinkAgentSet) All(operation LinkBoolOperation) bool {
	link := l.links.First()
	for link != nil {
		if !operation(link.(*Link)) {
			return false
		}
		link, _ = l.links.Next(link)
	}
	return true
}

// returns true if any of the links in the agent set satisfy the operation
func (l *LinkAgentSet) Any(operation LinkBoolOperation) bool {
	link := l.links.First()
	for link != nil {
		if operation(link.(*Link)) {
			return true
		}
		link, _ = l.links.Next(link)
	}
	return false
}

// perform the list of operations for all links in the agent set
func (l *LinkAgentSet) Ask(operations []LinkOperation) {
	links := l.links.List()

	for _, link := range links {
		for j := 0; j < len(operations); j++ {
			operations[j](link.(*Link))
		}
	}
}

// returns true if the link is in the agent set
func (l *LinkAgentSet) Contains(link *Link) bool {
	return l.links.Contains(link)
}

// returns the length of the agent set
func (l *LinkAgentSet) Count() int {
	return l.links.Len()
}

// returns the agent set as a list
func (l *LinkAgentSet) List() []*Link {
	v := []*Link{}
	link := l.links.First()
	for link != nil {
		v = append(v, link.(*Link))
		link, _ = l.links.Next(link)
	}
	return v
}

// returns the top n links in the agent set based on the float operation
func (l *LinkAgentSet) FirstNOf(n int) *LinkAgentSet {
	links := []*Link{}
	link := l.links.First()
	for i := 0; i < n && link != nil; i++ {
		links = append(links, link.(*Link))
		link, _ = l.links.Next(link)
	}
	return LinkSet(links)
}

// returns the max link in the agent set based on the float operation
func (l *LinkAgentSet) First() (*Link, error) {
	link := l.links.First()
	if link == nil {
		return nil, ErrNoLinksInAgentSet
	}
	return link.(*Link), nil
}

// returns the min n links in the agent set based on the float operation
func (l *LinkAgentSet) LastNOf(n int) *LinkAgentSet {
	links := []*Link{}
	link := l.links.Last()
	for i := 0; i < n && link != nil; i++ {
		links = append(links, link.(*Link))
		link, _ = l.links.Previous(link)
	}
	return LinkSet(links)
}

// returns the min link in the agent set based on the float operation
func (l *LinkAgentSet) Last() (*Link, error) {
	link := l.links.Last()
	if link == nil {
		return nil, ErrNoLinksInAgentSet
	}
	return link.(*Link), nil
}

// returns one of the links
// @TODO make this actually random based on model seed
func (l *LinkAgentSet) OneOf() (*Link, error) {
	for _, link := range l.links.List() {
		return link.(*Link), nil
	}

	return nil, ErrNoLinksInAgentSet
}

// returns the next element of the agentset after the provided one
func (l *LinkAgentSet) Previous(link *Link) (*Link, error) {
	v, err := l.links.Previous(link)
	if v == nil {
		return nil, err
	}
	return v.(*Link), err
}

// remove a link from the agent set
func (l *LinkAgentSet) Remove(link *Link) {
	l.links.Remove(link)
}

// sort the agent set based on the float operation in ascending order
func (l *LinkAgentSet) SortAsc(operation LinkFloatOperation) {
	l.links.SortAsc(func(a interface{}) interface{} {
		return operation(a.(*Link))
	})
}

// sort the agent set based on the float operation in descending order
func (l *LinkAgentSet) SortDesc(operation LinkFloatOperation) {
	l.links.SortDesc(func(a interface{}) interface{} {
		return operation(a.(*Link))
	})
}

// returns n links or all the links in the agentset if the length is lower than n
// make this actually random based on model seed
func (l *LinkAgentSet) UpToNOf(n int) *LinkAgentSet {
	links := []*Link{}
	link := l.links.First()
	for i := 0; i < n && link != nil; i++ {
		links = append(links, link.(*Link))
		link, _ = l.links.Next(link)
	}
	return LinkSet(links)
}

// returns a new agent set with all the links that are not in the given agents set
func (l *LinkAgentSet) WhoAreNot(links *LinkAgentSet) *LinkAgentSet {
	newLinks := []*Link{}

	for link := l.links.First(); link != nil; link, _ = l.links.Next(link) {
		if !links.Contains(link.(*Link)) {
			newLinks = append(newLinks, link.(*Link))
		}
	}

	return LinkSet(newLinks)
}

// returns a new agent set with all the links that are not the given link
func (l *LinkAgentSet) WhoAreNotLink(link *Link) *LinkAgentSet {
	newLinks := []*Link{}

	for l1 := l.links.First(); l1 != nil; l1, _ = l.links.Next(l1) {
		if l1.(*Link) != link {
			newLinks = append(newLinks, l1.(*Link))
		}
	}

	return LinkSet(newLinks)
}

// returns a new agent set that is a subset of the agent set where all satisfy the bool operation
func (l *LinkAgentSet) With(operation LinkBoolOperation) *LinkAgentSet {
	links := make([]*Link, 0)
	for link := l.links.First(); link != nil; link, _ = l.links.Next(link) {
		if operation(link.(*Link)) {
			links = append(links, link.(*Link))
		}
	}
	return LinkSet(links)
}

// returns a subset of the agent set where all links are equal to the max value provided by the operation
func (l *LinkAgentSet) WithMax(operation LinkFloatOperation) *LinkAgentSet {
	max := math.MaxFloat64 * -1
	for link := l.links.First(); link != nil; link, _ = l.links.Next(link) {
		if operation(link.(*Link)) > max {
			max = operation(link.(*Link))
		}
	}

	//get all links where the float operation is equal to the max
	links := make([]*Link, 0)
	for link := l.links.First(); link != nil; link, _ = l.links.Next(link) {
		if operation(link.(*Link)) == max {
			links = append(links, link.(*Link))
		}
	}

	return LinkSet(links)
}

// returns a subset of the agent set where all links are equal to the min value provided by the operation
func (l *LinkAgentSet) WithMin(operation LinkFloatOperation) *LinkAgentSet {
	min := math.MaxFloat64
	for link := l.links.First(); link != nil; link, _ = l.links.Next(link) {
		if operation(link.(*Link)) < min {
			min = operation(link.(*Link))
		}
	}

	//get all links where the float operation is equal to the min
	links := make([]*Link, 0)
	for link := l.links.First(); link != nil; link, _ = l.links.Next(link) {
		if operation(link.(*Link)) == min {
			links = append(links, link.(*Link))
		}
	}

	return LinkSet(links)
}
