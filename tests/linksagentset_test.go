package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/internal/model"
)

func TestAllLink(t *testing.T) {
	link1 := &model.Link{}
	link2 := &model.Link{}
	link3 := &model.Link{}

	linkSet := model.NewLinkAgentSet([]*model.Link{link1, link2, link3})

	link1.Shape = "circle"
	link2.Shape = "circle"
	link3.Shape = "circle"

	// assert that linkset has All of shape "circle"
	if !linkSet.All(func(l *model.Link) bool {
		return l.Shape == "circle"
	}) {
		t.Errorf("Expected linkset to have all links with shape 'circle'")
	}

	link2.Shape = "square"

	if linkSet.All(func(l *model.Link) bool {
		return l.Shape == "circle"
	}) {
		t.Errorf("Expected linkset to not have all links with shape 'circle'")
	}
}

func TestAnyLink(t *testing.T) {

	link1 := &model.Link{}
	link2 := &model.Link{}
	link3 := &model.Link{}

	linkSet := model.NewLinkAgentSet([]*model.Link{link1, link2, link3})

	link1.Shape = "circle"
	link2.Shape = "square"
	link3.Shape = "triangle"

	// assert that linkset has Any of shape "circle"
	if !linkSet.Any(func(l *model.Link) bool {
		return l.Shape == "circle"
	}) {
		t.Errorf("Expected linkset to have a link with shape 'circle'")
	}

	link1.Shape = "square"

	if linkSet.Any(func(l *model.Link) bool {
		return l.Shape == "circle"
	}) {
		t.Errorf("Expected linkset to not have a link with shape 'circle'")
	}

}

func TestLinksWhoAreNotLinks(t *testing.T) {

	// create links
	link1 := &model.Link{}
	link2 := &model.Link{}
	link3 := &model.Link{}

	// create linkset
	linkSet := model.NewLinkAgentSet([]*model.Link{link1, link2, link3})

	// create a second linkset
	linkSet2 := model.NewLinkAgentSet([]*model.Link{link1, link2})

	// create a third linkset that is the difference between the first and second linkset
	linkSet3 := linkSet.WhoAreNot(linkSet2)

	// assert that the third linkset has only one link
	if linkSet3.Count() != 1 {
		t.Errorf("Expected linkset3 to have 1 link")
	}

	// assert that the third linkset has link3
	if !linkSet3.Contains(link3) {
		t.Errorf("Expected linkset3 to have link3")
	}
}

func TestLinksWhoAreNotLink(t *testing.T) {

	// create links
	link1 := &model.Link{}
	link2 := &model.Link{}
	link3 := &model.Link{}

	// create linkset
	linkSet := model.NewLinkAgentSet([]*model.Link{link1, link2, link3})

	// create a second linkset
	linkSet2 := linkSet.WhoAreNotLink(link1)

	// assert that the second linkset has only two links
	if linkSet2.Count() != 2 {
		t.Errorf("Expected linkset2 to have 2 links")
	}

	// assert that the second linkset does not have link1
	if linkSet2.Contains(link1) {
		t.Errorf("Expected linkset2 to not have link1")
	}
}
