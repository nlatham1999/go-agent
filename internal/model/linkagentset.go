package model

import (
	"math"
	"sort"
)

type LinkAgentSet struct {
	links map[*Link]interface{}
}

func LinkSet(links []*Link) *LinkAgentSet {
	newLinks := make(map[*Link]interface{})
	for _, link := range links {
		newLinks[link] = nil
	}

	return &LinkAgentSet{
		links: newLinks,
	}
}

func (l *LinkAgentSet) Add(link *Link) {
	l.links[link] = nil
}

func (l *LinkAgentSet) All(operation LinkBoolOperation) bool {
	for link := range l.links {
		if !operation(link) {
			return false
		}
	}
	return true
}

func (l *LinkAgentSet) Any(operation LinkBoolOperation) bool {
	for link := range l.links {
		if operation(link) {
			return true
		}
	}
	return false
}

func (l *LinkAgentSet) Contains(link *Link) bool {
	_, ok := l.links[link]
	return ok
}

func (l *LinkAgentSet) Count() int {
	return len(l.links)
}

func (l *LinkAgentSet) List() []*Link {
	links := make([]*Link, 0)
	for link := range l.links {
		links = append(links, link)
	}
	return links
}

// gets the top n links based on the float operation
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

func (l *LinkAgentSet) MaxOneOf(operation LinkFloatOperation) *Link {
	max := math.MaxFloat64 * -1
	var maxLink *Link
	for link := range l.links {
		if operation(link) > max {
			max = operation(link)
			maxLink = link
		}
	}
	return maxLink
}

// @TODO implement
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

func (l *LinkAgentSet) MinOneOf(operation LinkFloatOperation) *Link {
	min := math.MaxFloat64
	var minLink *Link
	for link := range l.links {
		if operation(link) < min {
			min = operation(link)
			minLink = link
		}
	}
	return minLink
}

// @TODO implement
func (l *LinkAgentSet) OneOf() *Link {
	return nil
}

// @TODO implement
func (l *LinkAgentSet) UpToNOf(n int) *LinkAgentSet {
	return nil
}

// returns a new LinkAgentSet with all the links that are not in the given LinkAgentSet
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

// returns a new LinkAgentSet with all the links that are not the given link
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

func (l *LinkAgentSet) With(operation LinkBoolOperation) *LinkAgentSet {
	links := make([]*Link, 0)
	for link := range l.links {
		if operation(link) {
			links = append(links, link)
		}
	}
	return LinkSet(links)
}

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
