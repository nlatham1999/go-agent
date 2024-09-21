package tests

import (
	"testing"

	"github.com/nlatham1999/go-agent/internal/model"
)

func TestCreateTurtles(t *testing.T) {
	breeds := []string{
		"ants",
	}

	settings := model.ModelSettings{
		TurtleBreeds: breeds,
	}
	environment := model.NewModel(settings)

	// creating turtles without a breed should add them to the default breed
	environment.CreateTurtles(5, "", nil)
	if environment.Turtles("").Count() != 5 {
		t.Errorf("Expected 5 turtles, got %d", environment.Turtles("").Count())
	}

	// creating turtles with a breed should add them to that breed and the default breed
	environment.CreateTurtles(5, "ants", nil)
	if environment.Turtles("").Count() != 10 {
		t.Errorf("Expected 10 turtles, got %d", environment.Turtles("").Count())
	}
	if environment.Turtles("ants").Count() != 5 {
		t.Errorf("Expected 5 ants, got %d", environment.Turtles("ants").Count())
	}

	// creating a turtle with a nonexistent breed should return an error
	err := environment.CreateTurtles(5, "nonexistent", nil)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestTurtle(t *testing.T) {
	breeds := []string{
		"ants",
	}

	settings := model.ModelSettings{
		TurtleBreeds: breeds,
	}
	environment := model.NewModel(settings)

	// create 5 general turtle and five ants
	environment.CreateTurtles(5, "", nil)
	environment.CreateTurtles(5, "ants", nil)

	// get turtle from general pop
	turtle := environment.Turtle("", 0)
	if turtle == nil {
		t.Errorf("Expected turtle, got nil")
	}

	// get turtle from ants pop when it should not exist
	turtle = environment.Turtle("ants", 0)
	if turtle != nil {
		t.Errorf("Expected nil, got turtle")
	}

	// get turtle from ants pop when it should exist
	turtle = environment.Turtle("ants", 5)
	if turtle == nil {
		t.Errorf("Expected turtle, got nil")
	}

	// get turtle from general pop when it should not exist
	turtle = environment.Turtle("", 12)
	if turtle != nil {
		t.Errorf("Expected nil, got turtle")
	}

	// get turtle that is an ant from general pop when it should exist
	turtle = environment.Turtle("", 7)
	if turtle == nil {
		t.Errorf("Expected turtle, got nil")
	}

	// get turtle that belongs to nonexistent breed
	turtle = environment.Turtle("elephants", 0)
	if turtle != nil {
		t.Errorf("Expected nil, got turtle")
	}
}

func TestClearTurtles(t *testing.T) {
	breeds := []string{
		"ants",
	}

	settings := model.ModelSettings{
		TurtleBreeds: breeds,
	}
	m := model.NewModel(settings)

	// create 5 general turtle and five ants
	m.CreateTurtles(5, "", nil)
	m.CreateTurtles(5, "ants", nil)

	ref := m.Turtle("", 0)
	ref.SetXY(1, 1)

	if m.Patch(0, 0).TurtlesHere("").Count() != 9 {
		t.Errorf("Expected 10 turtles, got %d", m.Patch(0, 0).TurtlesHere("").Count())
	}

	// clear general turtles
	m.ClearTurtles()
	if m.Turtles("").Count() != 0 {
		t.Errorf("Expected 0 turtles, got %d", m.Turtles("").Count())
	}

	if ref.XCor() != 0 {
		t.Errorf("Expected turtle to be reset")
	}

	t1 := m.Turtle("", 0)
	if t1 != nil {
		t.Errorf("Expected nil, got turtle")
	}

	p := m.Patch(0, 0)
	if p.TurtlesHere("").Count() != 0 {
		t.Errorf("Expected 0 turtles, got %d", p.TurtlesHere("").Count())
	}

	m.CreateTurtles(1, "", nil)

	t1 = m.Turtle("", 0)
	if t1 == nil {
		t.Errorf("Expected turtle, got nil")
	}

}

func TestKillTurtle(t *testing.T) {

	breeds := []string{
		"ants",
	}

	settings := model.ModelSettings{
		TurtleBreeds: breeds,
	}
	m := model.NewModel(settings)

	// create 5 general turtle and five ants
	m.CreateTurtles(5, "", nil)
	m.CreateTurtles(5, "ants", nil)

	ref := m.Turtle("", 0)
	ref.SetXY(1, 1)

	if m.Patch(0, 0).TurtlesHere("").Count() != 9 {
		t.Errorf("Expected 10 turtles, got %d", m.Patch(0, 0).TurtlesHere("").Count())
	}

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("ants", 5)
	t3 := m.Turtle("ants", 6)
	t4 := m.Turtle("ants", 7)

	_, err := model.NewLink(m, "", t1, t2, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}
	_, err = model.NewLink(m, "", t1, t3, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}
	_, err = model.NewLink(m, "", t1, t4, false)
	if err != nil {
		t.Errorf("Error should be nil")
	}
	_, err = model.NewLink(m, "", t2, t3, false)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// kill general turtle
	m.KillTurtle(m.Turtle("", 0))
	if m.Turtles("").Count() != 9 {
		t.Errorf("Expected 9 turtles, got %d", m.Turtles("").Count())
	}

	if ref.XCor() != 0 {
		t.Errorf("Expected turtle to be reset")
	}

	// make sure there's only one link left
	if m.Links.Count() != 1 {
		t.Errorf("Expected 1 link, got %d", m.Links.Count())
	}
}

// tests the model link function
func TestModelLink(t *testing.T) {

	breeds := []string{
		"ants",
	}

	undirectedLinkBreeds := []string{
		"workers",
	}

	settings := model.ModelSettings{
		TurtleBreeds:         breeds,
		UndirectedLinkBreeds: undirectedLinkBreeds,
	}
	m := model.NewModel(settings)

	// create 5 general turtle and five ants
	m.CreateTurtles(5, "", nil)
	m.CreateTurtles(5, "ants", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("ants", 5)
	t3 := m.Turtle("ants", 6)
	t4 := m.Turtle("ants", 7)

	// create a directed link
	l1, err := model.NewLink(m, "", t1, t2, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// create an undirected link
	l2, err := model.NewLink(m, "workers", t2, t3, false)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// create an undirected link
	l3, err := model.NewLink(m, "workers", t3, t4, false)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// create a directed link
	l4, err := model.NewLink(m, "", t4, t1, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// the link between t1 and t2
	link := m.Link("", 0, 5)
	if link == nil {
		t.Errorf("Expected link, got nil")
	}
	if link != l1 {
		t.Errorf("Expected l1, got different link")
	}

	// the link between t2 and t3
	link = m.Link("ants", 5, 6)
	//link should  be nil because the breed does not exist for undirected links
	if link != nil {
		t.Errorf("Expected nil, got link")
	}
	link = m.Link("workers", 5, 6)
	if link == nil {
		t.Errorf("Expected link, got nil")
	}
	if link != l2 {
		t.Errorf("Expected l2, got different link")
	}

	// the link between t3 and t4
	link = m.Link("workers", 6, 7)
	if link == nil {
		t.Errorf("Expected link, got nil")
	}
	if link != l3 {
		t.Errorf("Expected l3, got different link")
	}

	// the link between t4 and t1
	link = m.Link("", 7, 0)
	if link == nil {
		t.Errorf("Expected link, got nil")
	}
	if link != l4 {
		t.Errorf("Expected l4, got different link")
	}
}

// tests the model LinkDirected function which is like Link but only for directed links
func TestModelLinkDirected(t *testing.T) {

	breeds := []string{
		"ants",
	}

	directedLinkBreeds := []string{
		"workers",
	}

	settings := model.ModelSettings{
		TurtleBreeds:       breeds,
		DirectedLinkBreeds: directedLinkBreeds,
	}
	m := model.NewModel(settings)

	// create 5 general turtle and five ants
	m.CreateTurtles(5, "", nil)
	m.CreateTurtles(5, "ants", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("ants", 5)
	t3 := m.Turtle("ants", 6)
	t4 := m.Turtle("ants", 7)

	// create a directed link
	l1, err := model.NewLink(m, "", t1, t2, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// create a directed link
	l2, err := model.NewLink(m, "workers", t3, t4, true)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// the link between t1 and t2
	link := m.LinkDirected("", 0, 5)
	if link == nil {
		t.Errorf("Expected link, got nil")
	}
	if link != l1 {
		t.Errorf("Expected l1, got different link")
	}

	// the link between t3 and t4
	link = m.LinkDirected("workers", 6, 7)
	if link == nil {
		t.Errorf("Expected link, got nil")
	}
	if link != l2 {
		t.Errorf("Expected l3, got different link")
	}
}

func TestDiffuse(t *testing.T) {

	patchesOwn := map[string]interface{}{
		"heat": 0.0,
	}

	// create a basic model
	settings := model.ModelSettings{
		PatchesOwn: patchesOwn,
	}
	m := model.NewModel(settings)

	m.Patches.Ask([]model.PatchOperation{
		func(p *model.Patch) {
			p.SetOwn("heat", 0)
		},
	})

	m.Patch(0, 0).Ask([]model.PatchOperation{
		func(p *model.Patch) {
			p.SetOwn("heat", 100)
		},
	})

	p1 := m.Patch(-1, 1)
	p2 := m.Patch(0, 1)
	p3 := m.Patch(1, 1)
	p4 := m.Patch(-1, 0)
	p5 := m.Patch(0, 0)
	p6 := m.Patch(1, 0)
	p7 := m.Patch(-1, -1)
	p8 := m.Patch(0, -1)
	p9 := m.Patch(1, -1)

	if p5.GetOwn("heat") != float64(100) {
		t.Errorf("Expected 100, got %d", p5.GetOwn("heat"))
	}

	m.Diffuse("heat", .25)

	if p1.GetOwn("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p1.GetOwn("heat"))
	}

	if p2.GetOwn("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p2.GetOwn("heat"))
	}

	if p3.GetOwn("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p3.GetOwn("heat"))
	}

	if p4.GetOwn("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p4.GetOwn("heat"))
	}

	if p5.GetOwn("heat") != float64(75) {
		t.Errorf("Expected 75, got %d", p5.GetOwn("heat"))
	}

	if p6.GetOwn("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p6.GetOwn("heat"))
	}

	if p7.GetOwn("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p7.GetOwn("heat"))
	}

	if p8.GetOwn("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p8.GetOwn("heat"))
	}

	if p9.GetOwn("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p9.GetOwn("heat"))
	}

	// diffuse a second round
	m.Diffuse("heat", .25)

	if p1.GetOwn("heat").(float64)-4.8828125 > 0.0001 {
		t.Errorf("Expected 25, got %d", p1.GetOwn("heat"))
	}

	if p2.GetOwn("heat").(float64)-5.078125 > 0.0001 {
		t.Errorf("Expected 25, got %d", p2.GetOwn("heat"))
	}

	if p3.GetOwn("heat").(float64)-4.8828125 > 0.0001 {
		t.Errorf("Expected 25, got %d", p3.GetOwn("heat"))
	}

	if p4.GetOwn("heat").(float64)-5.078125 > 0.0001 {
		t.Errorf("Expected 25, got %d", p4.GetOwn("heat"))
	}

	if p5.GetOwn("heat").(float64)-57.03125 > 0.0001 {
		t.Errorf("Expected 75, got %d", p5.GetOwn("heat"))
	}

	if p6.GetOwn("heat").(float64)-5.078125 > 0.0001 {
		t.Errorf("Expected 25, got %d", p6.GetOwn("heat"))
	}

	if p7.GetOwn("heat").(float64)-4.8828125 > 0.0001 {
		t.Errorf("Expected 25, got %d", p7.GetOwn("heat"))
	}

	if p8.GetOwn("heat").(float64)-5.078125 > 0.0001 {
		t.Errorf("Expected 25, got %d", p8.GetOwn("heat"))
	}

	if p9.GetOwn("heat").(float64)-4.8828125 > 0.0001 {
		t.Errorf("Expected 25, got %d", p9.GetOwn("heat"))
	}
}

func TestDiffuseCorner(t *testing.T) {

	patchesOwn := map[string]interface{}{
		"heat": 0.0,
	}

	// create a basic model
	settings := model.ModelSettings{
		PatchesOwn: patchesOwn,
	}
	m := model.NewModel(settings)

	m.Patches.Ask([]model.PatchOperation{
		func(p *model.Patch) {
			p.SetOwn("heat", 0)
		},
	})

	m.Patch(-15, -15).Ask([]model.PatchOperation{
		func(p *model.Patch) {
			p.SetOwn("heat", 100)
		},
	})

	p1 := m.Patch(-15, -14)
	p2 := m.Patch(-14, -14)
	p3 := m.Patch(-15, -15)
	p4 := m.Patch(-15, -14)

	m.Diffuse("heat", .25)

	if p1.GetOwn("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p1.GetOwn("heat"))
	}

	if p2.GetOwn("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p2.GetOwn("heat"))
	}

	if p3.GetOwn("heat") != 90.625 {
		t.Errorf("Expected 25, got %d", p3.GetOwn("heat"))
	}

	if p4.GetOwn("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p4.GetOwn("heat"))
	}

}

// tests the Diffuse4 function
func TestDiffuse4(t *testing.T) {

	patchesOwn := map[string]interface{}{
		"heat": 0.0,
	}

	// create a basic model
	settings := model.ModelSettings{
		PatchesOwn: patchesOwn,
	}
	m := model.NewModel(settings)

	m.Patches.Ask([]model.PatchOperation{
		func(p *model.Patch) {
			p.SetOwn("heat", 0)
		},
	})

	m.Patch(0, 0).Ask([]model.PatchOperation{
		func(p *model.Patch) {
			p.SetOwn("heat", 100)
		},
	})

	p1 := m.Patch(-1, 1)
	p2 := m.Patch(0, 1)
	p3 := m.Patch(1, 1)
	p4 := m.Patch(-1, 0)
	p5 := m.Patch(0, 0)
	p6 := m.Patch(1, 0)
	p7 := m.Patch(-1, -1)
	p8 := m.Patch(0, -1)
	p9 := m.Patch(1, -1)

	if p5.GetOwn("heat") != float64(100) {
		t.Errorf("Expected 100, got %d", p5.GetOwn("heat"))
	}

	m.Diffuse4("heat", .25)

	if p1.GetOwn("heat") != 0.0 {
		t.Errorf("Expected 0, got %d", p1.GetOwn("heat"))
	}

	if p2.GetOwn("heat") != 6.25 {
		t.Errorf("Expected 6.25, got %d", p2.GetOwn("heat"))
	}

	if p3.GetOwn("heat") != 0.0 {
		t.Errorf("Expected 0, got %d", p3.GetOwn("heat"))
	}

	if p4.GetOwn("heat") != 6.25 {
		t.Errorf("Expected 6.25, got %d", p4.GetOwn("heat"))
	}

	if p5.GetOwn("heat") != float64(75) {
		t.Errorf("Expected 75, got %d", p5.GetOwn("heat"))
	}

	if p6.GetOwn("heat") != 6.25 {
		t.Errorf("Expected 6.25, got %d", p6.GetOwn("heat"))
	}

	if p7.GetOwn("heat") != 0.0 {
		t.Errorf("Expected 0, got %d", p7.GetOwn("heat"))
	}

	if p8.GetOwn("heat") != 6.25 {
		t.Errorf("Expected 6.25, got %d", p8.GetOwn("heat"))
	}

	if p9.GetOwn("heat") != 0.0 {
		t.Errorf("Expected 0, got %d", p9.GetOwn("heat"))
	}

	// diffuse a second round
	m.Diffuse4("heat", .25)

	if p1.GetOwn("heat") != 0.78125 {
		t.Errorf("Expected 0, got %d", p1.GetOwn("heat"))
	}

	if p2.GetOwn("heat") != 9.375 {
		t.Errorf("Expected 6.25, got %d", p2.GetOwn("heat"))
	}

	if p3.GetOwn("heat") != 0.78125 {
		t.Errorf("Expected 0, got %d", p3.GetOwn("heat"))
	}

	if p4.GetOwn("heat") != 9.375 {
		t.Errorf("Expected 6.25, got %d", p4.GetOwn("heat"))
	}

	if p5.GetOwn("heat") != float64(57.8125) {
		t.Errorf("Expected 57.8125, got %d", p5.GetOwn("heat"))
	}

	if p6.GetOwn("heat") != 9.375 {
		t.Errorf("Expected 6.25, got %d", p6.GetOwn("heat"))
	}

	if p7.GetOwn("heat") != 0.78125 {
		t.Errorf("Expected 0, got %d", p7.GetOwn("heat"))
	}

	if p8.GetOwn("heat") != 9.375 {
		t.Errorf("Expected 6.25, got %d", p8.GetOwn("heat"))
	}

	if p9.GetOwn("heat") != 0.78125 {
		t.Errorf("Expected 0, got %d", p9.GetOwn("heat"))
	}
}

func TestTurtlesOnPatch(t *testing.T) {

	// create a basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	// create 5 turtles on a patch
	m.CreateTurtles(5, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)
	t3 := m.Turtle("", 2)

	t1.SetXY(1, 1)
	t2.SetXY(1, 1)
	t3.SetXY(1, 1)

	turtles := m.TurtlesOnPatch(m.Patch(1, 1))

	if turtles.Count() != 3 {
		t.Errorf("Expected 3 turtles, got %d", turtles.Count())
	}

	if !turtles.Contains(t1) {
		t.Errorf("Expected turtles to contain t1")
	}

	if !turtles.Contains(t2) {
		t.Errorf("Expected turtles to contain t2")
	}

	if !turtles.Contains(t3) {
		t.Errorf("Expected turtles to contain t3")
	}
}

func TestTurtlesOnPatches(t *testing.T) {

	// create a basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	// create 5 turtles on a patch
	m.CreateTurtles(5, "", nil)

	t1 := m.Turtle("", 0)
	t2 := m.Turtle("", 1)
	t3 := m.Turtle("", 2)
	t4 := m.Turtle("", 3)

	t1.SetXY(1, 1)
	t2.SetXY(1, 1)
	t3.SetXY(1, 1)
	t4.SetXY(1, 2)

	patchSet := model.PatchSet([]*model.Patch{m.Patch(1, 1), m.Patch(1, 2)})

	turtles := m.TurtlesOnPatches(patchSet)

	if turtles.Count() != 4 {
		t.Errorf("Expected 4 turtles, got %d", turtles.Count())
	}

	if !turtles.Contains(t1) {
		t.Errorf("Expected turtles to contain t1")
	}

	if !turtles.Contains(t2) {
		t.Errorf("Expected turtles to contain t2")
	}

	if !turtles.Contains(t3) {
		t.Errorf("Expected turtles to contain t3")
	}

	if !turtles.Contains(t4) {
		t.Errorf("Expected turtles to contain t4")
	}
}
