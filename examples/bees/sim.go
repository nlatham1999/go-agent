package bees

import (
	"fmt"

	"github.com/nlatham1999/go-agent/internal/api"
	"github.com/nlatham1999/go-agent/internal/model"
)

type Bees struct {
	model *model.Model
	step  int
}

func NewBees() *Bees {
	return &Bees{
		step: 0,
	}
}

func (b *Bees) Init() {
	modelSettings := model.ModelSettings{
		PatchesOwn: map[string]interface{}{
			"nectar": 0.0,
		},
		TurtleBreeds: []string{"scouts", "foragers"},
		TurtlesOwn: map[string]interface{}{
			"radius": 10.0,
			"group":  0,
		},
		Globals: map[string]interface{}{
			"scouts": 3,
		},
	}

	b.model = model.NewModel(modelSettings)

}

func (b *Bees) SetUp() error {

	b.model.ClearAll()

	b.step = 0

	b.model.Patches.Ask([]model.PatchOperation{
		func(p *model.Patch) {
			if b.model.RandomInt(100) > 95 {
				p.SetOwn("nectar", b.model.RandomFloat(100)+450)
			}
		},
	})

	b.model.Diffuse("nectar", .8)
	b.model.Diffuse("nectar", .8)
	b.model.Diffuse("nectar", .8)
	b.model.Diffuse("nectar", .8)
	// b.model.Diffuse("nectar", .1)
	// b.model.Diffuse("nectar", .1)
	// b.model.Diffuse("nectar", .1)
	// b.model.Diffuse("nectar", .1)
	// b.model.Diffuse("nectar", .1)
	// b.model.Diffuse("nectar", .1)
	// b.model.Diffuse("nectar", .1)

	b.model.Patches.Ask([]model.PatchOperation{
		func(p *model.Patch) {
			nectar := p.GetOwnI("nectar")
			p.PColor.SetColorRGB(0, nectar, 0)
		},
	})

	b.model.CreateTurtles(1, "", []model.TurtleOperation{
		func(t *model.Turtle) {
			t.SetXY(0, 0)
		},
	})
	b.model.CreateTurtles(1, "", []model.TurtleOperation{
		func(t *model.Turtle) {
			t.SetXY(1, 0)
			t.CreateLinkWithTurtle("", b.model.Turtle("", 0), nil)
		},
	})

	b.model.Patch(0, 0).PColor.SetColor(model.Red)
	b.model.Patch(1, 0).PColor.SetColor(model.Blue)

	numScouts, _ := b.model.GetGlobalI("scouts")
	b.model.CreateTurtles(numScouts, "scouts", []model.TurtleOperation{
		func(t *model.Turtle) {
			t.SetXY(b.model.RandomXCor(), b.model.RandomYCor())
			t.Color.SetColor(model.Yellow)
			t.SetSize(.9)
			t.SetLabel(t.PatchHere().GetOwnI("nectar"))
			t.SetOwn("group", t.Who())
		},
	})

	b.model.ResetTicks()

	return nil
}

func (b *Bees) Go() {
	numSteps := 2

	// first step is to create the foragers based on the scouts' findings
	if b.step%numSteps == 0 {
		b.model.Turtles("scouts").Ask([]model.TurtleOperation{
			func(scout *model.Turtle) {
				searchRadius, _ := scout.GetOwnF("radius")
				if searchRadius < 1 {
					return
				}

				nectar := scout.PatchHere().GetOwnI("nectar")
				numForagers := 2
				if nectar > 30 {
					numForagers = 4
				}
				scout.Hatch("foragers", numForagers, []model.TurtleOperation{
					func(forager *model.Turtle) {
						forager.SetSize(.8)
						forager.Color.SetColor(model.Red)
						forager.SetHeading(b.model.RandomFloat(360))
						forager.Forward(b.model.RandomFloat(searchRadius))
						scout.CreateLinkWithTurtle("", forager, []model.LinkOperation{
							func(l *model.Link) {
								l.Label = l.Length()
								l.LabelColor = model.Red
							},
						})
						forager.SetLabel(forager.PatchHere().GetOwnI("nectar"))
					},
				})
			},
		})
	}

	// determine new scout
	if b.step%numSteps == 1 {
		b.model.Turtles("scouts").Ask([]model.TurtleOperation{
			func(t *model.Turtle) {
				foragers := t.LinkNeighbors("")

				if foragers.Count() == 0 {
					return
				}

				max := foragers.MaxOneOf(func(f *model.Turtle) float64 {
					return f.PatchHere().GetOwnF("nectar")
				})

				if max.PatchHere().GetOwnF("nectar") > t.PatchHere().GetOwnF("nectar") {
					t.Die()
					foragers.WhoAreNotTurtle(max).Ask([]model.TurtleOperation{
						func(f *model.Turtle) {
							f.Die()
						},
					})
					max.SetBreed("scouts")
					max.Color.SetColor(model.Yellow)
				} else {
					foragers.Ask([]model.TurtleOperation{
						func(f *model.Turtle) {
							f.Die()
						},
					})

					//shrink the search radius
					radius, _ := t.GetOwnF("radius")
					t.SetOwn("radius", radius-.5)
				}
			},
		})
	}

	b.model.Links().Ask([]model.LinkOperation{
		func(l *model.Link) {
			l.Color.SetColor(model.Orange)
		},
	})

	b.step++

	b.model.Tick()
}

func (b *Bees) Model() *model.Model {
	return b.model
}

func (b *Bees) Stats() map[string]interface{} {

	stats := map[string]interface{}{}
	b.model.Turtles("scouts").Ask([]model.TurtleOperation{
		func(t *model.Turtle) {
			group := t.GetOwn("group")
			nectar := t.PatchHere().GetOwnI("nectar")
			radius := t.GetOwn("radius")
			stats[fmt.Sprintf("Group %v", group)] = fmt.Sprintf("Nectar at Scout: %v, Search Radius: %v", nectar, radius)
		},
	})

	return stats
}

func (b *Bees) Stop() bool {
	return b.model.Turtles("scouts").All(func(t *model.Turtle) bool {
		radius, _ := t.GetOwnF("radius")
		return radius < 1
	})
}

func (b *Bees) Widgets() []api.Widget {
	return []api.Widget{
		{
			PrettyName:      "Scouts",
			TargetVariable:  "scouts",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "100",
			DefaultValue:    "3",
		},
	}
}
