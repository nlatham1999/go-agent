package universe

import "math"

type LinkAgentSet struct {
	links []*Link
}

func LinkSet(links []*Link) *LinkAgentSet {
	newLinks := make([]*Link, len(links))
	copy(newLinks, links)

	return &LinkAgentSet{
		links: newLinks,
	}
}

func (l *LinkAgentSet) All(operation LinkBoolOperation) bool {
	for _, link := range l.links {
		if !operation(link) {
			return false
		}
	}
	return true
}

func (l *LinkAgentSet) Any(operation LinkBoolOperation) bool {
	for _, link := range l.links {
		if operation(link) {
			return true
		}
	}
	return false
}

func (l *LinkAgentSet) Count() int {
	return len(l.links)
}

// @TODO implement
func (l *LinkAgentSet) MaxNOf(n int, operation LinkFloatOperation) *LinkAgentSet {
	return nil
}

func (l *LinkAgentSet) MaxOneOf(operation LinkFloatOperation) *Link {
	max := math.MaxFloat64 * -1
	var maxLink *Link
	for _, link := range l.links {
		if operation(link) > max {
			max = operation(link)
			maxLink = link
		}
	}
	return maxLink
}

// @TODO implement
func (l *LinkAgentSet) MinNOf(n int, operation LinkFloatOperation) *LinkAgentSet {
	return nil
}

func (l *LinkAgentSet) MinOneOf(operation LinkFloatOperation) *Link {
	min := math.MaxFloat64
	var minLink *Link
	for _, link := range l.links {
		if operation(link) < min {
			min = operation(link)
			minLink = link
		}
	}
	return minLink
}

// @TODO implement
func (l *LinkAgentSet) UpToNOf(n int) *LinkAgentSet {
	return nil
}

// @TODO implement
func (l *LinkAgentSet) WhoAreNot(links *LinkAgentSet) *LinkAgentSet {
	return nil
}

// @TODO implement
func (l *LinkAgentSet) WhoAreNotLink(link *Link) *LinkAgentSet {
	return nil
}

func (l *LinkAgentSet) With(operation LinkBoolOperation) *LinkAgentSet {
	links := make([]*Link, 0)
	for _, link := range l.links {
		if operation(link) {
			links = append(links, link)
		}
	}
	return LinkSet(links)
}

func (l *LinkAgentSet) WithMax(operation LinkFloatOperation) *LinkAgentSet {
	max := math.MaxFloat64 * -1
	for _, link := range l.links {
		if operation(link) > max {
			max = operation(link)
		}
	}

	//get all links where the float operation is equal to the max
	links := make([]*Link, 0)
	for _, link := range l.links {
		if operation(link) == max {
			links = append(links, link)
		}
	}

	return LinkSet(links)
}

func (l *LinkAgentSet) WithMin(operation LinkFloatOperation) *LinkAgentSet {
	min := math.MaxFloat64
	for _, link := range l.links {
		if operation(link) < min {
			min = operation(link)
		}
	}

	//get all links where the float operation is equal to the min
	links := make([]*Link, 0)
	for _, link := range l.links {
		if operation(link) == min {
			links = append(links, link)
		}
	}

	return LinkSet(links)
}
