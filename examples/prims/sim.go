package prims

import (
	"fmt"
	"time"

	"github.com/nlatham1999/go-agent/pkg/api"
	"github.com/nlatham1999/go-agent/pkg/model"
)

type Prims struct {
	model *model.Model

	placedTurtleBreed   *model.TurtleBreed
	unplacedTurtleBreed *model.TurtleBreed
	placedLinkBreed     *model.LinkBreed
	unplacedLinkBreed   *model.LinkBreed

	nodes int
}

func NewPrims() *Prims {
	return &Prims{}
}

func (p *Prims) Init() {

	p.placedTurtleBreed = model.NewTurtleBreed("placed", "", nil)
	p.unplacedTurtleBreed = model.NewTurtleBreed("unplaced", "", nil)

	p.placedLinkBreed = model.NewLinkBreed("unplaced")
	p.unplacedLinkBreed = model.NewLinkBreed("placed")

	modelSettings := model.ModelSettings{
		TurtleBreeds:         []*model.TurtleBreed{p.placedTurtleBreed, p.unplacedTurtleBreed},
		UndirectedLinkBreeds: []*model.LinkBreed{p.placedLinkBreed, p.unplacedLinkBreed},
		MinPxCor:             0,
		MinPyCor:             0,
		MaxPxCor:             100,
		MaxPyCor:             100,
	}

	p.model = model.NewModel(modelSettings)

	p.nodes = 5
}

func (p *Prims) SetUp() error {

	p.model.ClearAll()

	unplaced := p.model.TurtleBreed("unplaced")
	unplaced.CreateAgents(p.nodes,
		p.placeInitialNodes,
	)

	//for each turtle create a link with every other turtle
	unplaced.Agents().Ask(
		p.createInitialLinks,
	)

	p.unplacedLinkBreed.Links().SortAsc(func(l *model.Link) float64 {
		return l.Length()
	})

	t0 := p.model.Turtle(0)
	t0.Color.SetColor(model.Red)

	placed := p.model.TurtleBreed("placed")

	t0.SetBreed(placed)

	p.model.ResetTicks()
	return nil
}

func (p *Prims) placeInitialNodes(t *model.Turtle) {
	t.Color.SetColor(model.Gray)
	t.SetSize(1)
	t.SetXY(p.model.RandomXCor(), p.model.RandomYCor())
}

func (p *Prims) createInitialLinks(t *model.Turtle) {
	unplaced := p.model.TurtleBreed("unplaced")
	unplaced.Agents().Ask(
		func(t2 *model.Turtle) {
			if t != t2 && t.DistanceTurtle(t2) < 10 {
				t.CreateLinkWithTurtle(p.unplacedLinkBreed, t2,
					func(l *model.Link) {
						l.Color.SetColor(model.Gray)
						l.Hide()
					},
				)
			}
		},
	)
}

func (p *Prims) Go() {
	start := time.Now()

	// find the closest link to the cluster
	var closestLink *model.Link
	var closestTurtle *model.Turtle
	links := p.unplacedLinkBreed.Links()
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

	placed := p.model.TurtleBreed("placed")

	//add the link and turtle to the cluster
	closestLink.SetBreed(p.placedLinkBreed)
	closestLink.Color.SetColor(model.Red)
	closestLink.Hide()
	closestTurtle.SetBreed(placed)
	closestTurtle.Color.SetColor(model.Red)

	// if all nodes are placed, kill all unplaced links
	if placed.Agents().Count() == p.nodes {
		p.unplacedLinkBreed.Links().Ask(
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

	placed := p.model.TurtleBreed("placed")
	unplaced := p.model.TurtleBreed("unplaced")

	return map[string]interface{}{
		"Placed nodes":    placed.Agents().Count(),
		"Unplaced nodes":  unplaced.Agents().Count(),
		"potential links": p.unplacedLinkBreed.Links().Count(),
	}
}

func (p *Prims) Stop() bool {
	return p.placedLinkBreed.Links().Count() >= p.nodes-2
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
