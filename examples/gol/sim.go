package gol

import (
	"github.com/nlatham1999/go-agent/pkg/api"
	"github.com/nlatham1999/go-agent/pkg/model"
)

var _ api.ModelInterface = &Gol{}

type Gol struct {
	model *model.Model

	minNeighborsToLive      int
	maxNeighborsToLive      int
	minNeighborsToReproduce int
	maxNeighborsToReproduce int
	initialAlive            float64
}

func NewGol() *Gol {
	return &Gol{}
}

func (g *Gol) Model() *model.Model {
	return g.model
}

func (g *Gol) Init() {
	settings := model.ModelSettings{
		PatchProperties: map[string]interface{}{
			"alive":      true,
			"alive-next": true,
		},
		MinPxCor: 0,
		MaxPxCor: 20,
		MinPyCor: 0,
		MaxPyCor: 20,
	}

	g.model = model.NewModel(settings)

	g.minNeighborsToLive = 2
	g.maxNeighborsToLive = 3
	g.minNeighborsToReproduce = 3
	g.maxNeighborsToReproduce = 3
	g.initialAlive = 0.5
}

func (g *Gol) SetUp() error {
	g.model.ClearAll()

	g.model.Patches.Ask(
		func(p *model.Patch) {
			if v := g.model.RandomFloat(1); v < g.initialAlive {
				p.SetProperty("alive", true)
				p.SetProperty("alive-next", true)
				p.PColor.SetColor(model.Green)
			} else {
				p.SetProperty("alive", false)
				p.SetProperty("alive-next", false)
				p.PColor.SetColor(model.Black)
			}
		},
	)

	return nil
}

func (g *Gol) Go() {

	g.model.Patches.Ask(
		func(p *model.Patch) {

			//get neighboring patches
			neighbors := p.Neighbors()

			//count the number of alive neighbors
			aliveNeighbors := neighbors.With(func(p *model.Patch) bool {
				alive := p.GetPropB("alive")
				return alive
			}).Count()

			alive := p.GetPropB("alive")
			if alive {
				if aliveNeighbors < g.minNeighborsToLive {
					p.SetProperty("alive-next", false)
				}

				if aliveNeighbors > g.maxNeighborsToLive {
					p.SetProperty("alive-next", false)
				}
			} else {
				if aliveNeighbors >= g.minNeighborsToReproduce && aliveNeighbors <= g.maxNeighborsToReproduce {
					p.SetProperty("alive-next", true)
				}
			}
		},
	)

	g.model.Patches.Ask(
		func(p *model.Patch) {
			p.SetProperty("alive", p.GetPropB("alive-next"))
			if p.GetPropB("alive") {
				p.PColor.SetColor(model.Green)
			} else {
				p.PColor.SetColor(model.Black)
			}
		},
	)

	g.model.Tick()
}

func (g *Gol) Stats() map[string]interface{} {
	return map[string]interface{}{
		"num-alive": g.model.Patches.With(func(p *model.Patch) bool {
			return p.GetPropB("alive")
		}).Count(),
	}
}

func (g *Gol) Stop() bool {
	return g.model.Patches.All(func(p *model.Patch) bool {
		return !p.GetPropB("alive")
	})
}

func (g *Gol) Widgets() []api.Widget {
	return []api.Widget{
		{
			PrettyName:      "Min Neighbors To Live",
			TargetVariable:  "min-neighbors-to-live",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "8",
			DefaultValue:    "2",
			ValuePointerInt: &g.minNeighborsToLive,
		},
		{
			PrettyName:      "Max Neighbors To Live",
			TargetVariable:  "max-neighbors-to-live",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "8",
			DefaultValue:    "3",
			ValuePointerInt: &g.maxNeighborsToLive,
		},
		{
			PrettyName:      "Min Neighbors To Reproduce",
			TargetVariable:  "min-neighbors-to-reproduce",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "8",
			DefaultValue:    "3",
			ValuePointerInt: &g.minNeighborsToReproduce,
		},
		{
			PrettyName:      "Max Neighbors To Reproduce",
			TargetVariable:  "max-neighbors-to-reproduce",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "8",
			DefaultValue:    "3",
			ValuePointerInt: &g.maxNeighborsToReproduce,
		},
		{
			PrettyName:        "Initial Alive",
			TargetVariable:    "initial-alive",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			MinValue:          "0",
			MaxValue:          "1",
			DefaultValue:      "0.5",
			StepAmount:        "0.01",
			ValuePointerFloat: &g.initialAlive,
		},
	}
}
