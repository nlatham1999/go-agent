package tests

import (
	"fmt"
	"testing"

	"github.com/nlatham1999/go-agent/pkg/model"
)

func TestCreateTurtles(t *testing.T) {
	ants := model.NewTurtleBreed("ants", "", nil)

	settings := model.ModelSettings{
		TurtleBreeds: []*model.TurtleBreed{ants},
	}
	environment := model.NewModel(settings)

	// creating turtles without a breed should add them to the default breed
	environment.CreateTurtles(5, nil)
	if environment.Turtles().Count() != 5 {
		t.Errorf("Expected 5 turtles, got %d", environment.Turtles().Count())
	}

	// creating turtles with a breed should add them to that breed and the default breed
	ants.CreateAgents(5, nil)
	if environment.Turtles().Count() != 10 {
		t.Errorf("Expected 10 turtles, got %d", environment.Turtles().Count())
	}
	if ants.Agents().Count() != 5 {
		t.Errorf("Expected 5 ants, got %d", ants.Agents().Count())
	}

}

func TestTurtle(t *testing.T) {

	ants := model.NewTurtleBreed("ants", "", nil)

	settings := model.ModelSettings{
		TurtleBreeds: []*model.TurtleBreed{ants},
	}
	environment := model.NewModel(settings)

	// create 5 general turtle and five ants
	environment.CreateTurtles(5, nil)
	ants.CreateAgents(5, nil)

	// get turtle from general pop
	turtle := environment.Turtle(0)
	if turtle == nil {
		t.Errorf("Expected turtle, got nil")
	}

	// get turtle from ants pop when it should not exist
	turtle = ants.Agent(0)
	if turtle != nil {
		t.Errorf("Expected nil, got turtle")
	}

	// get turtle from ants pop when it should exist
	turtle = ants.Agent(5)
	if turtle == nil {
		t.Errorf("Expected turtle, got nil")
	}

	// get turtle from general pop when it should not exist
	turtle = environment.Turtle(12)
	if turtle != nil {
		t.Errorf("Expected nil, got turtle")
	}

	// get turtle that is an ant from general pop when it should exist
	turtle = environment.Turtle(7)
	if turtle == nil {
		t.Errorf("Expected turtle, got nil")
	}
}

func TestClearTurtles(t *testing.T) {
	ants := model.NewTurtleBreed("ants", "", nil)

	settings := model.ModelSettings{
		TurtleBreeds: []*model.TurtleBreed{ants},
	}
	m := model.NewModel(settings)

	// create 5 general turtle and five ants
	m.CreateTurtles(5, nil)
	ants.CreateAgents(5, nil)

	ref := m.Turtle(0)
	ref.SetXY(1, 1)

	if m.Patch(0, 0).TurtlesHere().Count() != 9 {
		t.Errorf("Expected 10 turtles, got %d", m.Patch(0, 0).TurtlesHere().Count())
	}

	// clear general turtles
	m.ClearTurtles()
	if m.Turtles().Count() != 0 {
		t.Errorf("Expected 0 turtles, got %d", m.Turtles().Count())
	}

	if ref.XCor() != 0 {
		t.Errorf("Expected turtle to be reset")
	}

	t1 := m.Turtle(0)
	if t1 != nil {
		t.Errorf("Expected nil, got turtle")
	}

	p := m.Patch(0, 0)
	if p.TurtlesHere().Count() != 0 {
		t.Errorf("Expected 0 turtles, got %d", p.TurtlesHere().Count())
	}

	m.CreateTurtles(1, nil)

	t1 = m.Turtle(0)
	if t1 == nil {
		t.Errorf("Expected turtle, got nil")
	}

}

func TestKillTurtle(t *testing.T) {

	ants := model.NewTurtleBreed("ants", "", nil)

	settings := model.ModelSettings{
		TurtleBreeds: []*model.TurtleBreed{ants},
	}
	m := model.NewModel(settings)

	// create 5 general turtle and five ants
	m.CreateTurtles(5, nil)
	// m.CreateTurtles(5, "ants", nil)
	ants.CreateAgents(5, nil)

	ref := m.Turtle(0)
	ref.SetXY(1, 1)

	if m.Patch(0, 0).TurtlesHere().Count() != 9 {
		t.Errorf("Expected 10 turtles, got %d", m.Patch(0, 0).TurtlesHere().Count())
	}

	t1 := m.Turtle(0)
	t2 := ants.Agent(5)
	t3 := ants.Agent(6)
	t4 := ants.Agent(7)

	_, err := t1.CreateLinkToTurtle(nil, t2, nil)
	if err != nil {
		t.Errorf("Error should be nil")
	}
	_, err = t1.CreateLinkToTurtle(nil, t3, nil)
	if err != nil {
		t.Errorf("Error should be nil")
	}
	_, err = t1.CreateLinkToTurtle(nil, t4, nil)
	if err != nil {
		t.Errorf("Error should be nil")
	}
	_, err = t2.CreateLinkToTurtle(nil, t3, nil)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// kill general turtle
	m.KillTurtle(m.Turtle(0))
	if m.Turtles().Count() != 9 {
		t.Errorf("Expected 9 turtles, got %d", m.Turtles().Count())
	}

	if ref.XCor() != 0 {
		t.Errorf("Expected turtle to be reset")
	}

	// make sure there's only one link left
	if m.Links().Count() != 1 {
		t.Errorf("Expected 1 link, got %d", m.Links().Count())
	}
}

// tests the model link function
func TestModelLink(t *testing.T) {

	ants := model.NewTurtleBreed("ants", "", nil)

	workers := model.NewLinkBreed("workers")

	undirectedLinkBreeds := []*model.LinkBreed{
		workers,
	}

	settings := model.ModelSettings{
		TurtleBreeds:         []*model.TurtleBreed{ants},
		UndirectedLinkBreeds: undirectedLinkBreeds,
	}
	m := model.NewModel(settings)

	// create 5 general turtle and five ants
	m.CreateTurtles(5, nil)
	ants.CreateAgents(5, nil)

	t1 := m.Turtle(0)
	t2 := ants.Agent(5)
	t3 := ants.Agent(6)
	t4 := ants.Agent(7)

	// create a directed link
	l1, err := t1.CreateLinkToTurtle(nil, t2, nil)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// create an undirected link
	l2, err := t2.CreateLinkWithTurtle(workers, t3, nil)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// create an undirected link
	l3, err := t3.CreateLinkWithTurtle(workers, t4, nil)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// create a directed link
	l4, err := t4.CreateLinkToTurtle(nil, t1, nil)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// the link between t1 and t2
	link := m.Link(0, 5)
	if link == nil {
		t.Errorf("Expected link, got nil")
	}
	if link != l1 {
		t.Errorf("Expected l1, got different link")
	}

	link = workers.Link(5, 6)
	if link == nil {
		t.Errorf("Expected link, got nil")
	}
	if link != l2 {
		t.Errorf("Expected l2, got different link")
	}

	// the link between t3 and t4
	link = workers.Link(6, 7)
	if link == nil {
		t.Errorf("Expected link, got nil")
	}
	if link != l3 {
		t.Errorf("Expected l3, got different link")
	}

	// the link between t4 and t1
	link = m.Link(7, 0)
	if link == nil {
		t.Errorf("Expected link, got nil")
	}
	if link != l4 {
		t.Errorf("Expected l4, got different link")
	}
}

// tests the model LinkDirected function which is like Link but only for directed links
func TestModelLinkDirected(t *testing.T) {

	ants := model.NewTurtleBreed("ants", "", nil)

	workers := model.NewLinkBreed("workers")

	directedLinkBreeds := []*model.LinkBreed{
		workers,
	}

	settings := model.ModelSettings{
		TurtleBreeds:       []*model.TurtleBreed{ants},
		DirectedLinkBreeds: directedLinkBreeds,
	}
	m := model.NewModel(settings)

	// create 5 general turtle and five ants
	m.CreateTurtles(5, nil)
	ants.CreateAgents(5, nil)

	t1 := m.Turtle(0)
	t2 := ants.Agent(5)
	t3 := ants.Agent(6)
	t4 := ants.Agent(7)

	// create a directed link
	l1, err := t1.CreateLinkToTurtle(nil, t2, nil)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// create a directed link
	l2, err := t3.CreateLinkToTurtle(workers, t4, nil)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	// the link between t1 and t2
	link := m.LinkDirected(0, 5)
	if link == nil {
		t.Errorf("Expected link, got nil")
	}
	if link != l1 {
		t.Errorf("Expected l1, got different link")
	}

	// the link between t3 and t4
	link = workers.Link(6, 7)
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
		PatchProperties: patchesOwn,
	}
	m := model.NewModel(settings)

	m.Patches.Ask(
		func(p *model.Patch) {
			p.SetProperty("heat", 0)
		},
	)

	m.Patch(0, 0).SetProperty("heat", 100)

	p1 := m.Patch(-1, 1)
	p2 := m.Patch(0, 1)
	p3 := m.Patch(1, 1)
	p4 := m.Patch(-1, 0)
	p5 := m.Patch(0, 0)
	p6 := m.Patch(1, 0)
	p7 := m.Patch(-1, -1)
	p8 := m.Patch(0, -1)
	p9 := m.Patch(1, -1)

	if p5.GetProperty("heat") != float64(100) {
		t.Errorf("Expected 100, got %d", p5.GetProperty("heat"))
	}

	m.Diffuse("heat", .25)

	if p1.GetProperty("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p1.GetProperty("heat"))
	}

	if p2.GetProperty("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p2.GetProperty("heat"))
	}

	if p3.GetProperty("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p3.GetProperty("heat"))
	}

	if p4.GetProperty("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p4.GetProperty("heat"))
	}

	if p5.GetProperty("heat") != float64(75) {
		t.Errorf("Expected 75, got %d", p5.GetProperty("heat"))
	}

	if p6.GetProperty("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p6.GetProperty("heat"))
	}

	if p7.GetProperty("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p7.GetProperty("heat"))
	}

	if p8.GetProperty("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p8.GetProperty("heat"))
	}

	if p9.GetProperty("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p9.GetProperty("heat"))
	}

	// diffuse a second round
	m.Diffuse("heat", .25)

	if p1.GetProperty("heat").(float64)-4.8828125 > 0.0001 {
		t.Errorf("Expected 25, got %d", p1.GetProperty("heat"))
	}

	if p2.GetProperty("heat").(float64)-5.078125 > 0.0001 {
		t.Errorf("Expected 25, got %d", p2.GetProperty("heat"))
	}

	if p3.GetProperty("heat").(float64)-4.8828125 > 0.0001 {
		t.Errorf("Expected 25, got %d", p3.GetProperty("heat"))
	}

	if p4.GetProperty("heat").(float64)-5.078125 > 0.0001 {
		t.Errorf("Expected 25, got %d", p4.GetProperty("heat"))
	}

	if p5.GetProperty("heat").(float64)-57.03125 > 0.0001 {
		t.Errorf("Expected 75, got %d", p5.GetProperty("heat"))
	}

	if p6.GetProperty("heat").(float64)-5.078125 > 0.0001 {
		t.Errorf("Expected 25, got %d", p6.GetProperty("heat"))
	}

	if p7.GetProperty("heat").(float64)-4.8828125 > 0.0001 {
		t.Errorf("Expected 25, got %d", p7.GetProperty("heat"))
	}

	if p8.GetProperty("heat").(float64)-5.078125 > 0.0001 {
		t.Errorf("Expected 25, got %d", p8.GetProperty("heat"))
	}

	if p9.GetProperty("heat").(float64)-4.8828125 > 0.0001 {
		t.Errorf("Expected 25, got %d", p9.GetProperty("heat"))
	}
}

func TestDiffuseCorner(t *testing.T) {

	patchesOwn := map[string]interface{}{
		"heat": 0.0,
	}

	// create a basic model
	settings := model.ModelSettings{
		PatchProperties: patchesOwn,
	}
	m := model.NewModel(settings)

	m.Patches.Ask(
		func(p *model.Patch) {
			p.SetProperty("heat", 0)
		},
	)

	m.Patch(-15, -15).SetProperty("heat", 100)

	p1 := m.Patch(-15, -14)
	p2 := m.Patch(-14, -14)
	p3 := m.Patch(-15, -15)
	p4 := m.Patch(-15, -14)

	m.Diffuse("heat", .25)

	if p1.GetProperty("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p1.GetProperty("heat"))
	}

	if p2.GetProperty("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p2.GetProperty("heat"))
	}

	if p3.GetProperty("heat") != 90.625 {
		t.Errorf("Expected 25, got %d", p3.GetProperty("heat"))
	}

	if p4.GetProperty("heat") != 3.125 {
		t.Errorf("Expected 25, got %d", p4.GetProperty("heat"))
	}

}

// tests the Diffuse4 function
func TestDiffuse4(t *testing.T) {

	patchesOwn := map[string]interface{}{
		"heat": 0.0,
	}

	// create a basic model
	settings := model.ModelSettings{
		PatchProperties: patchesOwn,
	}
	m := model.NewModel(settings)

	m.Patches.Ask(
		func(p *model.Patch) {
			p.SetProperty("heat", 0)
		},
	)

	m.Patch(0, 0).SetProperty("heat", 100)

	p1 := m.Patch(-1, 1)
	p2 := m.Patch(0, 1)
	p3 := m.Patch(1, 1)
	p4 := m.Patch(-1, 0)
	p5 := m.Patch(0, 0)
	p6 := m.Patch(1, 0)
	p7 := m.Patch(-1, -1)
	p8 := m.Patch(0, -1)
	p9 := m.Patch(1, -1)

	if p5.GetProperty("heat") != float64(100) {
		t.Errorf("Expected 100, got %d", p5.GetProperty("heat"))
	}

	m.Diffuse4("heat", .25)

	if p1.GetProperty("heat") != 0.0 {
		t.Errorf("Expected 0, got %d", p1.GetProperty("heat"))
	}

	if p2.GetProperty("heat") != 6.25 {
		t.Errorf("Expected 6.25, got %d", p2.GetProperty("heat"))
	}

	if p3.GetProperty("heat") != 0.0 {
		t.Errorf("Expected 0, got %d", p3.GetProperty("heat"))
	}

	if p4.GetProperty("heat") != 6.25 {
		t.Errorf("Expected 6.25, got %d", p4.GetProperty("heat"))
	}

	if p5.GetProperty("heat") != float64(75) {
		t.Errorf("Expected 75, got %d", p5.GetProperty("heat"))
	}

	if p6.GetProperty("heat") != 6.25 {
		t.Errorf("Expected 6.25, got %d", p6.GetProperty("heat"))
	}

	if p7.GetProperty("heat") != 0.0 {
		t.Errorf("Expected 0, got %d", p7.GetProperty("heat"))
	}

	if p8.GetProperty("heat") != 6.25 {
		t.Errorf("Expected 6.25, got %d", p8.GetProperty("heat"))
	}

	if p9.GetProperty("heat") != 0.0 {
		t.Errorf("Expected 0, got %d", p9.GetProperty("heat"))
	}

	// diffuse a second round
	m.Diffuse4("heat", .25)

	if p1.GetProperty("heat") != 0.78125 {
		t.Errorf("Expected 0, got %d", p1.GetProperty("heat"))
	}

	if p2.GetProperty("heat") != 9.375 {
		t.Errorf("Expected 6.25, got %d", p2.GetProperty("heat"))
	}

	if p3.GetProperty("heat") != 0.78125 {
		t.Errorf("Expected 0, got %d", p3.GetProperty("heat"))
	}

	if p4.GetProperty("heat") != 9.375 {
		t.Errorf("Expected 6.25, got %d", p4.GetProperty("heat"))
	}

	if p5.GetProperty("heat") != float64(57.8125) {
		t.Errorf("Expected 57.8125, got %d", p5.GetProperty("heat"))
	}

	if p6.GetProperty("heat") != 9.375 {
		t.Errorf("Expected 6.25, got %d", p6.GetProperty("heat"))
	}

	if p7.GetProperty("heat") != 0.78125 {
		t.Errorf("Expected 0, got %d", p7.GetProperty("heat"))
	}

	if p8.GetProperty("heat") != 9.375 {
		t.Errorf("Expected 6.25, got %d", p8.GetProperty("heat"))
	}

	if p9.GetProperty("heat") != 0.78125 {
		t.Errorf("Expected 0, got %d", p9.GetProperty("heat"))
	}
}

func TestTurtlesOnPatch(t *testing.T) {

	// create a basic model
	settings := model.ModelSettings{}
	m := model.NewModel(settings)

	// create 5 turtles on a patch
	m.CreateTurtles(5, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)
	t3 := m.Turtle(2)

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
	m.CreateTurtles(5, nil)

	t1 := m.Turtle(0)
	t2 := m.Turtle(1)
	t3 := m.Turtle(2)
	t4 := m.Turtle(3)

	t1.SetXY(1, 1)
	t2.SetXY(1, 1)
	t3.SetXY(1, 1)
	t4.SetXY(1, 2)

	patchSet := model.NewPatchAgentSet([]*model.Patch{m.Patch(1, 1), m.Patch(1, 2)})

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

func TestModelSize(t *testing.T) {

	sheep := model.NewTurtleBreed("sheep", "", nil)
	wolves := model.NewTurtleBreed("wolves", "", nil)

	modelSettings := model.ModelSettings{
		TurtleBreeds: []*model.TurtleBreed{sheep, wolves},
		TurtleProperties: map[string]interface{}{
			"energy": 0,
		},
		PatchProperties: map[string]interface{}{
			"countdown": int(0),
		},
		MinPxCor: 0,
		MaxPxCor: 50,
		MinPyCor: 0,
		MaxPyCor: 50,
	}

	m := model.NewModel(modelSettings)

	m.ClearAll()

	m.Patches.Ask(
		func(p *model.Patch) {

			grassRegrowthTime := 30
			if m.RandomFloat(1) < 0.5 {
				p.PColor.SetColor(model.Green)
				p.SetProperty("countdown", grassRegrowthTime)
			} else {
				p.PColor.SetColor(model.Brown)
				p.SetProperty("countdown", m.RandomInt(grassRegrowthTime))
			}
		},
	)

	initialNumberSheep := 20

	sheepGainFromFood := 2

	sheep.CreateAgents(initialNumberSheep,
		func(t *model.Turtle) {
			// t.Shape("sheep")
			t.Color.SetColor(model.White)
			// t.Size(1.5)
			t.SetLabelColor(model.Blue)
			t.SetProperty("energy", m.RandomInt(2*sheepGainFromFood))
			t.SetXY(m.RandomXCor(), m.RandomYCor())
			t.SetSize(.5)
		},
	)

	initialNumberWolves := 4

	wolfGainFromFood := 2

	wolves.CreateAgents(initialNumberWolves,
		func(t *model.Turtle) {
			// t.Shape("wolf")
			t.Color.SetColor(model.Black)
			// t.Size(2)
			t.SetLabelColor(model.White)
			t.SetProperty("energy", m.RandomInt(2*wolfGainFromFood))
			t.SetXY(float64(m.RandomXCor()), m.RandomYCor())
			t.SetSize(.5)
		},
	)

	showEnergy := true
	m.Turtles().Ask(
		func(t *model.Turtle) {
			if showEnergy {
				t.SetLabel(fmt.Sprintf("%v", t.GetProperty("energy")))
			} else {
				t.SetLabel("")
			}
		},
	)

	m.ResetTicks()
}
