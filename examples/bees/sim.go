package bees

import (
	"fmt"

	"github.com/nlatham1999/go-agent/pkg/api"
	"github.com/nlatham1999/go-agent/pkg/model"
)

// enforce that Bees implements the ModelInterface interface
var _ api.ModelInterface = (*Bees)(nil)

type Bees struct {
	model *model.Model
	step  int

	scouts int
}

func NewBees() *Bees {
	return &Bees{
		step: 0,
	}
}

func (b *Bees) Init() {

	scouts := model.NewTurtleBreed("scouts", "", nil)
	foragers := model.NewTurtleBreed("foragers", "", nil)

	modelSettings := model.ModelSettings{
		PatchProperties: map[string]interface{}{
			"nectar": 0.0,
		},
		TurtleBreeds: []*model.TurtleBreed{scouts, foragers},
		TurtleProperties: map[string]interface{}{
			"radius": 10.0,
			"group":  0,
		},
	}

	b.model = model.NewModel(modelSettings)

	b.scouts = 3

}

func (b *Bees) SetUp() error {

	if b.model == nil {
		return fmt.Errorf("model is nil")
	}

	b.model.ClearAll()

	b.step = 0

	b.model.Patches.Ask(
		func(p *model.Patch) {
			if b.model.RandomInt(100) > 95 {
				p.SetProperty("nectar", b.model.RandomFloat(100)+450)
			}
		},
	)

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

	b.model.Patches.Ask(
		func(p *model.Patch) {
			nectar := p.GetPropI("nectar")
			p.PColor.SetColorRGB(0, nectar, 0)
		},
	)

	b.model.CreateTurtles(1,
		func(t *model.Turtle) {
			t.SetXY(0, 0)
		},
	)
	b.model.CreateTurtles(1,
		func(t *model.Turtle) {
			t.SetXY(1, 0)
			t.CreateLinkWithTurtle(nil, b.model.Turtle(0), nil)
		},
	)

	b.model.Patch(0, 0).PColor.SetColor(model.Red)
	b.model.Patch(1, 0).PColor.SetColor(model.Blue)

	scouts := b.model.TurtleBreed("scouts")
	scouts.CreateAgents(b.scouts,
		func(t *model.Turtle) {
			t.SetXY(b.model.RandomXCor(), b.model.RandomYCor())
			t.Color.SetColor(model.Yellow)
			t.SetSize(.9)
			t.SetLabel(t.PatchHere().GetPropI("nectar"))
			t.SetProperty("group", t.Who())
		},
	)

	b.model.ResetTicks()

	return nil
}

func (b *Bees) Go() {
	numSteps := 2

	// first step is to create the foragers based on the scouts' findings
	scouts := b.model.TurtleBreed("scouts")
	foragers := b.model.TurtleBreed("foragers")
	if b.step%numSteps == 0 {
		scouts.Agents().Ask(
			func(scout *model.Turtle) {
				searchRadius, _ := scout.GetPropF("radius")
				if searchRadius < 1 {
					return
				}

				nectar := scout.PatchHere().GetPropI("nectar")
				numForagers := 2
				if nectar > 30 {
					numForagers = 4
				}
				foragers.CreateAgents(numForagers, func(forager *model.Turtle) {
					forager.SetSize(.8)
					forager.Color.SetColor(model.Red)
					forager.SetHeading(b.model.RandomFloat(360))
					forager.Forward(b.model.RandomFloat(searchRadius))
					scout.CreateLinkWithTurtle(nil, forager, nil)
					forager.SetLabel(forager.PatchHere().GetPropI("nectar"))
					forager.SetXY(scout.XCor(), scout.YCor())
				})

			},
		)
	}

	// determine new scout
	if b.step%numSteps == 1 {
		scouts.Agents().Ask(
			func(t *model.Turtle) {
				foragers := t.LinkedTurtles(nil)

				if foragers.Count() == 0 {
					return
				}

				foragers.SortDesc(func(f *model.Turtle) float64 {
					return f.PatchHere().GetPropF("nectar")
				})
				max, err := foragers.First()
				if err != nil {
					return
				}

				if max.PatchHere().GetPropF("nectar") > t.PatchHere().GetPropF("nectar") {
					t.Die()
					foragers.WhoAreNotTurtle(max).Ask(
						func(f *model.Turtle) {
							f.Die()
						},
					)
					max.SetBreed(scouts)
					max.Color.SetColor(model.Yellow)
				} else {
					foragers.Ask(
						func(f *model.Turtle) {
							f.Die()
						},
					)

					//shrink the search radius
					radius, _ := t.GetPropF("radius")
					t.SetProperty("radius", radius-.5)
				}
			},
		)
	}

	b.model.Links().Ask(
		func(l *model.Link) {
			l.Color.SetColor(model.Orange)
		},
	)

	b.step++

	b.model.Tick()
}

func (b *Bees) Model() *model.Model {
	return b.model
}

func (b *Bees) Stats() map[string]interface{} {

	scouts := b.model.TurtleBreed("scouts")
	stats := map[string]interface{}{}
	scouts.Agents().Ask(
		func(t *model.Turtle) {
			group := t.GetProperty("group")
			nectar := t.PatchHere().GetPropI("nectar")
			radius := t.GetProperty("radius")
			stats[fmt.Sprintf("Group %v", group)] = fmt.Sprintf("Nectar at Scout: %v, Search Radius: %v", nectar, radius)
		},
	)

	return stats
}

func (b *Bees) Stop() bool {
	scouts := b.model.TurtleBreed("scouts")
	return scouts.Agents().All(func(t *model.Turtle) bool {
		radius, _ := t.GetPropF("radius")
		return radius < 1
	})
}

func (b *Bees) Widgets() []api.Widget {
	return []api.Widget{
		{
			PrettyName:      "Scouts",
			Id:              "scouts",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "100",
			DefaultValue:    "3",
		},
	}
}
