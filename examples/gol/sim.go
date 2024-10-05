package gol

import (
	"github.com/nlatham1999/go-agent/internal/api"
	"github.com/nlatham1999/go-agent/internal/model"
)

type Gol struct {
	model *model.Model
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
			"alive": true,
		},
		MinPxCor: 0,
		MaxPxCor: 50,
		MinPyCor: 0,
		MaxPyCor: 50,
		Globals: map[string]interface{}{
			"min-neighbors-to-live":      2,
			"max-neighbors-to-live":      3,
			"min-neighbors-to-reproduce": 3,
			"max-neighbors-to-reproduce": 3,
			"initial-alive":              0.5,
		},
	}

	g.model = model.NewModel(settings)
}

func (g *Gol) SetUp() error {
	g.model.ClearAll()

	g.model.Patches.Ask([]model.PatchOperation{
		func(p *model.Patch) {
			if v := g.model.RandomFloat(1); v < g.model.GetGlobal("initial-alive").(float64) {
				p.SetOwn("alive", true)
				p.PColor.SetColor(model.Green)
			} else {
				p.SetOwn("alive", false)
				p.PColor.SetColor(model.Black)
			}
		},
	})

	return nil
}

func (g *Gol) Go() {

	minNeighborsToLive := g.model.GetGlobal("min-neighbors-to-live").(int)
	maxNeighborsToLive := g.model.GetGlobal("max-neighbors-to-live").(int)
	minNeighborsToReproduce := g.model.GetGlobal("min-neighbors-to-reproduce").(int)
	maxNeighborsToReproduce := g.model.GetGlobal("max-neighbors-to-reproduce").(int)

	g.model.Patches.Ask([]model.PatchOperation{
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
				if aliveNeighbors < minNeighborsToLive {
					p.SetOwn("alive", false)
					p.PColor.SetColor(model.Black)
				}

				if aliveNeighbors > maxNeighborsToLive {
					p.SetOwn("alive", false)
					p.PColor.SetColor(model.Black)
				}
			} else {

				if aliveNeighbors >= minNeighborsToReproduce && aliveNeighbors <= maxNeighborsToReproduce {
					p.SetOwn("alive", true)
					p.PColor.SetColor(model.Green)
				}
			}
		},
	})

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
