package prims

import (
	"fmt"
	"time"

	"github.com/nlatham1999/go-agent/api"
	"github.com/nlatham1999/go-agent/model"
)

type Prims struct {
	model *model.Model
}

func NewPrims() *Prims {
	return &Prims{}
}

func (p *Prims) Init() {
	modelSettings := model.ModelSettings{
		TurtleBreeds:         []string{"unplaced", "placed"},
		UndirectedLinkBreeds: []string{"unplaced", "placed"},
		Globals: map[string]interface{}{
			"nodes": 5,
		},
		MinPxCor: 0,
		MinPyCor: 0,
		MaxPxCor: 100,
		MaxPyCor: 100,
	}

	p.model = model.NewModel(modelSettings)
}

func (p *Prims) SetUp() error {

	p.model.ClearAll()

	numNodes := p.model.GetGlobal("nodes").(int)

	p.model.CreateTurtles(numNodes, "unplaced", []model.TurtleOperation{
		p.placeInitialNodes,
	})

	//for each turtle create a link with every other turtle
	p.model.Turtles("unplaced").Ask(
		p.createInitialLinks,
	)

	p.model.UndirectedLinks("unplaced").SortAsc(func(l *model.Link) float64 {
		return l.Length()
	})

	t0 := p.model.Turtle("", 0)
	t0.Color.SetColor(model.Red)
	t0.SetBreed("placed")

	p.model.ResetTicks()
	return nil
}

func (p *Prims) placeInitialNodes(t *model.Turtle) {
	t.Color.SetColor(model.Gray)
	t.SetSize(1)
	t.SetXY(p.model.RandomXCor(), p.model.RandomYCor())
}

func (p *Prims) createInitialLinks(t *model.Turtle) {
	p.model.Turtles("unplaced").Ask(
		func(t2 *model.Turtle) {
			if t != t2 && t.DistanceTurtle(t2) < 10 {
				t.CreateLinkWithTurtle("unplaced", t2, []model.LinkOperation{
					func(l *model.Link) {
						l.Color.SetColor(model.Gray)
						l.Hide()
					},
				})
			}
		},
	)
}

func (p *Prims) Go() {
	start := time.Now()

	// find the closest link to the cluster
	var closestLink *model.Link
	var closestTurtle *model.Turtle
	links := p.model.UndirectedLinks("unplaced")
	done := false
	links.Ask(func(l *model.Link) {
		if done {
			return
		}
		breedName1 := l.End1().BreedName()
		breedName2 := l.End2().BreedName()
		if breedName1 == "placed" && breedName2 == "placed" {
			l.Die()
			return
		}
		if breedName1 == "unplaced" && breedName2 == "unplaced" {
			return
		}

		closestLink = l
		if breedName1 == "unplaced" {
			closestTurtle = l.End1()
		} else {
			closestTurtle = l.End2()
		}
		done = true
	})

	if closestLink == nil {
		return
	}

	//add the link and turtle to the cluster
	closestLink.SetBreed("placed")
	closestLink.Color.SetColor(model.Red)
	closestLink.Hide()
	closestTurtle.SetBreed("placed")
	closestTurtle.Color.SetColor(model.Red)

	// if all nodes are placed, kill all unplaced links
	if p.model.Turtles("placed").Count() == p.model.GetGlobal("nodes").(int) {
		p.model.UndirectedLinks("unplaced").Ask(
			func(l *model.Link) {
				l.Die()
			},
		)
	}

	p.model.Tick()

	fmt.Println("Time taken: ", time.Since(start))

}

func (p *Prims) Model() *model.Model {
	return p.model
}

func (p *Prims) Stats() map[string]interface{} {
	return map[string]interface{}{
		"Placed nodes":    p.model.Turtles("placed").Count(),
		"Unplaced nodes":  p.model.Turtles("unplaced").Count(),
		"potential links": p.model.UndirectedLinks("unplaced").Count(),
	}
}

func (p *Prims) Stop() bool {
	return p.model.UndirectedLinks("placed").Count() >= p.model.GetGlobal("nodes").(int)-2
}

func (p *Prims) Widgets() []api.Widget {
	return []api.Widget{
		{
			PrettyName:      "Nodes",
			TargetVariable:  "nodes",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "2",
			MaxValue:        "4000",
			DefaultValue:    "5",
		},
	}
}
