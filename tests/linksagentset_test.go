package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/internal/universe"
)

func TestAllLink(t *testing.T) {
	link1 := &universe.Link{}
	link2 := &universe.Link{}
	link3 := &universe.Link{}

	linkSet := universe.LinkSet([]*universe.Link{link1, link2, link3})

	link1.Shape = "circle"
	link2.Shape = "circle"
	link3.Shape = "circle"

	// assert that linkset has All of shape "circle"
	if !linkSet.All(func(l *universe.Link) bool {
		return l.Shape == "circle"
	}) {
		t.Errorf("Expected linkset to have all links with shape 'circle'")
	}

	link2.Shape = "square"

	if linkSet.All(func(l *universe.Link) bool {
		return l.Shape == "circle"
	}) {
		t.Errorf("Expected linkset to not have all links with shape 'circle'")
	}
}

func TestAnyLink(t *testing.T) {

	link1 := &universe.Link{}
	link2 := &universe.Link{}
	link3 := &universe.Link{}

	linkSet := universe.LinkSet([]*universe.Link{link1, link2, link3})

	link1.Shape = "circle"
	link2.Shape = "square"
	link3.Shape = "triangle"

	// assert that linkset has Any of shape "circle"
	if !linkSet.Any(func(l *universe.Link) bool {
		return l.Shape == "circle"
	}) {
		t.Errorf("Expected linkset to have a link with shape 'circle'")
	}

	link1.Shape = "square"

	if linkSet.Any(func(l *universe.Link) bool {
		return l.Shape == "circle"
	}) {
		t.Errorf("Expected linkset to not have a link with shape 'circle'")
	}

}