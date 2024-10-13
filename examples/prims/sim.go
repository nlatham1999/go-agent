package prims

import (
	"github.com/nlatham1999/go-agent/internal/api"
	"github.com/nlatham1999/go-agent/internal/model"
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
		func(t *model.Turtle) {
			t.Color.SetColor(model.Gray)
			t.SetSize(1)
			t.SetXY(p.model.RandomXCor(), p.model.RandomYCor())
		},
	})

	//for each turtle create a link with every other turtle
	p.model.Turtles("unplaced").Ask([]model.TurtleOperation{
		func(t *model.Turtle) {
			p.model.Turtles("unplaced").Ask([]model.TurtleOperation{
				func(t2 *model.Turtle) {
					if t != t2 {
						t.CreateLinkWithTurtle("unplaced", t2, []model.LinkOperation{
							func(l *model.Link) {
								l.Color.SetColor(model.Gray)
							},
						})
					}
				},
			})
		},
	})

	t0 := p.model.Turtle("", 0)
	t0.Color.SetColor(model.Red)
	t0.SetBreed("placed")

	p.model.ResetTicks()
	return nil
}

func (p *Prims) Go() {

	// find the closest link to the cluster
	var closestLink *model.Link
	var closestTurtle *model.Turtle
	var closestDistance float64 = 100000 //max float
	p.model.Turtles("placed").Ask([]model.TurtleOperation{
		func(t *model.Turtle) {

			t.Links("unplaced").Ask([]model.LinkOperation{
				func(l *model.Link) {
					if l.OtherEnd(t).BreedName() == "placed" {
						l.Die()
						return
					}

					d := l.Length()
					if d < closestDistance {
						closestDistance = d
						closestLink = l
						closestTurtle = l.OtherEnd(t)
					}
				},
			})

		},
	})

	if closestLink == nil {
		return
	}

	//if the closest link is linked to two placed nodes already then delete it
	if closestLink.End1().BreedName() == "placed" && closestLink.End2().BreedName() == "placed" {
		closestLink.Die()
		return
	}

	//add the link and turtle to the cluster
	closestLink.SetBreed("placed")
	closestLink.Color.SetColor(model.Red)
	closestTurtle.SetBreed("placed")
	closestTurtle.Color.SetColor(model.Red)

	p.model.Tick()
}

func (p *Prims) Model() *model.Model {
	return p.model
}

func (p *Prims) Stats() map[string]interface{} {
	return nil
}

func (p *Prims) Stop() bool {
	return p.model.UndirectedLinks("unplaced").Count() == 0
}

func (p *Prims) Widgets() []api.Widget {
	return []api.Widget{
		{
			PrettyName:      "Nodes",
			TargetVariable:  "nodes",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "2",
			MaxValue:        "200",
			DefaultValue:    "5",
		},
	}
}
