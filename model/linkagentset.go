package model

import (
	"github.com/nlatham1999/sortedset"
)

// LinkAgentSet is an ordered set of links than can be sorted
// implements github.com/nlatham1999/sortedset
type LinkAgentSet struct {
	links *sortedset.SortedSet
}

// create a new LinkAgentSet
func NewLinkAgentSet(links []*Link) *LinkAgentSet {

	linkSet := &LinkAgentSet{
		links: sortedset.NewSortedSet(),
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
	if operation == nil {
		return false
	}

	return l.links.All(func(a interface{}) bool {
		return operation(a.(*Link))
	})
}

// returns true if any of the links in the agent set satisfy the operation
func (l *LinkAgentSet) Any(operation LinkBoolOperation) bool {
	if operation == nil {
		return false
	}

	return l.links.Any(func(a interface{}) bool {
		return operation(a.(*Link))
	})
}

// perform the operation for all links in the agent set
func (l *LinkAgentSet) Ask(operation LinkOperation) {
	if operation == nil {
		return
	}

	l.links.Ask(func(a interface{}) {
		operation(a.(*Link))
	})
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
	links := make([]*Link, 0)
	l.links.Ask(func(a interface{}) {
		links = append(links, a.(*Link))
	})
	return links
}

// returns the top n links in the agent set
func (l *LinkAgentSet) FirstNOf(n int) *LinkAgentSet {
	linkSet := sortedset.NewSortedSet()
	link := l.links.First()
	for i := 0; i < n && link != nil; i++ {
		linkSet.Add(link)
		link, _ = l.links.Next()
	}
	return &LinkAgentSet{
		links: linkSet,
	}
}

// returns the first link in the agent set
func (l *LinkAgentSet) First() (*Link, error) {
	link := l.links.First()
	if link == nil {
		return nil, ErrNoLinksInAgentSet
	}
	return link.(*Link), nil
}

// returns the last n links in the agent set
func (l *LinkAgentSet) LastNOf(n int) *LinkAgentSet {
	linkSet := sortedset.NewSortedSet()
	link := l.links.Last()
	for i := 0; i < n && link != nil; i++ {
		linkSet.Add(link)
		link, _ = l.links.Previous()
	}
	return &LinkAgentSet{
		links: linkSet,
	}
}

// returns the last link in the agent set
func (l *LinkAgentSet) Last() (*Link, error) {
	link := l.links.Last()
	if link == nil {
		return nil, ErrNoLinksInAgentSet
	}
	return link.(*Link), nil
}

// returns the next link in the set
// func (l *LinkAgentSet) Next() (*Link, error) {
// 	v, err := l.links.Next()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return v.(*Link), err
// }

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

// returns a new agent set with all the links that are not in the given agents set
func (l *LinkAgentSet) WhoAreNot(links *LinkAgentSet) *LinkAgentSet {
	return &LinkAgentSet{
		links: l.links.Difference(links.links),
	}
}

// returns a new agent set with all the links that are not the given link
func (l *LinkAgentSet) WhoAreNotLink(link *Link) *LinkAgentSet {
	return &LinkAgentSet{
		links: l.links.Difference(sortedset.NewSortedSet(link)),
	}
}

// returns a new agent set that is a subset of the agent set where all satisfy the bool operation
func (l *LinkAgentSet) With(operation LinkBoolOperation) *LinkAgentSet {
	linkSet := sortedset.NewSortedSet()

	if operation == nil {
		return nil
	}

	l.links.Ask(func(a interface{}) {
		if operation(a.(*Link)) {
			linkSet.Add(a)
		}
	})

	return &LinkAgentSet{
		links: linkSet,
	}
}
