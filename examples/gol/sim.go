package gol

import (
	"github.com/nlatham1999/go-agent/api"
	"github.com/nlatham1999/go-agent/model"
)

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
		PatchesOwn: map[string]interface{}{
			"alive":      true,
			"alive-next": true,
		},
		MinPxCor: 0,
		MaxPxCor: 200,
		MinPyCor: 0,
		MaxPyCor: 200,
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
				p.SetOwn("alive", true)
				p.SetOwn("alive-next", true)
				p.PColor.SetColor(model.Green)
			} else {
				p.SetOwn("alive", false)
				p.SetOwn("alive-next", false)
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
				alive := p.GetOwnB("alive")
				return alive
			}).Count()

			alive := p.GetOwnB("alive")
			if alive {
				if aliveNeighbors < g.minNeighborsToLive {
					p.SetOwn("alive-next", false)
				}

				if aliveNeighbors > g.maxNeighborsToLive {
					p.SetOwn("alive-next", false)
				}
			} else {
				if aliveNeighbors >= g.minNeighborsToReproduce && aliveNeighbors <= g.maxNeighborsToReproduce {
					p.SetOwn("alive-next", true)
				}
			}
		},
	)

	g.model.Patches.Ask(
		func(p *model.Patch) {
			p.SetOwn("alive", p.GetOwnB("alive-next"))
			if p.GetOwnB("alive") {
				p.PColor.SetColor(model.Green)
			} else {
				p.PColor.SetColor(model.Black)
			}
		},
	)
}

func (g *Gol) Stats() map[string]interface{} {
	return map[string]interface{}{
		"num-alive": g.model.Patches.With(func(p *model.Patch) bool {
			return p.GetOwnB("alive")
		}).Count(),
	}
}

func (g *Gol) Stop() bool {
	return g.model.Patches.All(func(p *model.Patch) bool {
		return !p.GetOwnB("alive")
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
		},
		{
			PrettyName:      "Max Neighbors To Live",
			TargetVariable:  "max-neighbors-to-live",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "8",
			DefaultValue:    "3",
		},
		{
			PrettyName:      "Min Neighbors To Reproduce",
			TargetVariable:  "min-neighbors-to-reproduce",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "8",
			DefaultValue:    "3",
		},
		{
			PrettyName:      "Max Neighbors To Reproduce",
			TargetVariable:  "max-neighbors-to-reproduce",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "8",
			DefaultValue:    "3",
		},
		{
			PrettyName:      "Initial Alive",
			TargetVariable:  "initial-alive",
			WidgetType:      "slider",
			WidgetValueType: "float",
			MinValue:        "0",
			MaxValue:        "1",
			DefaultValue:    "0.5",
			StepAmount:      "0.01",
		},
	}
}
