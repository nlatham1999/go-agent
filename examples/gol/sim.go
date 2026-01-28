package gol

import (
	"fmt"
	"time"

	"github.com/nlatham1999/go-agent/pkg/api"
	"github.com/nlatham1999/go-agent/pkg/concurrency"
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
	worldSize               int
	numAliveGraph           api.GraphWidget

	patches []*model.Patch
}

func NewGol() *Gol {
	return &Gol{}
}

func (g *Gol) Model() *model.Model {
	return g.model
}

func (g *Gol) Init() {

	g.minNeighborsToLive = 2
	g.maxNeighborsToLive = 3
	g.minNeighborsToReproduce = 3
	g.maxNeighborsToReproduce = 3
	g.initialAlive = 0.5
	g.worldSize = 12
	g.numAliveGraph = api.NewGraphWidget("Number Alive", "num-alive-graph", "ticks", "count", []string{}, []string{})

	_ = g.SetUp()
}

func (g *Gol) SetUp() error {

	settings := model.ModelSettings{
		PatchProperties: map[string]interface{}{
			"alive":      true,
			"alive-next": true,
		},
		MinPxCor: 0,
		MaxPxCor: g.worldSize,
		MinPyCor: 0,
		MaxPyCor: g.worldSize,
	}

	g.model = model.NewModel(settings)

	g.model.Patches.Ask(
		func(p *model.Patch) {
			if v := g.model.RandomFloat(1); v < g.initialAlive {
				p.SetPropertySafe("alive", true)
				p.SetPropertySafe("alive-next", true)
				p.Color.SetColor(model.Green)
			} else {
				p.SetPropertySafe("alive", false)
				p.SetPropertySafe("alive-next", false)
				p.Color.SetColor(model.Black)
			}
		},
	)

	g.numAliveGraph.XValues = []string{}
	g.numAliveGraph.YValues = []string{}

	g.patches = g.model.Patches.List()

	return nil
}

func (g *Gol) Go() {

	currentTime := time.Now()
	defer func(startTime time.Time) {
		fmt.Println("Time per tick (ms):", time.Since(currentTime).Milliseconds())
	}(currentTime)

	// g.model.Patches.Ask(
	concurrency.AskPatches(g.patches,
		func(p *model.Patch) {

			//get neighboring patches
			neighbors := p.Neighbors()

			//count the number of alive neighbors
			aliveNeighbors := neighbors.With(func(p *model.Patch) bool {
				alive := p.GetPropertySafe("alive").(bool)
				return alive
			}).Count()

			alive := p.GetPropertySafe("alive").(bool)
			if alive {
				if aliveNeighbors < g.minNeighborsToLive {
					p.SetPropertySafe("alive-next", false)
				}

				if aliveNeighbors > g.maxNeighborsToLive {
					p.SetPropertySafe("alive-next", false)
				}
			} else {
				if aliveNeighbors >= g.minNeighborsToReproduce && aliveNeighbors <= g.maxNeighborsToReproduce {
					p.SetPropertySafe("alive-next", true)
				}
			}
		},
		10,
	)

	// g.model.Patches.Ask(
	concurrency.AskPatches(g.patches,
		func(p *model.Patch) {
			p.SetPropertySafe("alive", p.GetPropertySafe("alive-next").(bool))
			if p.GetPropertySafe("alive").(bool) {
				p.Color.SetColor(model.Green)
			} else {
				p.Color.SetColor(model.Black)
			}
		},
		10,
	)

	g.numAliveGraph.XValues = append(g.numAliveGraph.XValues, fmt.Sprintf("%d", g.model.Ticks))
	g.numAliveGraph.YValues = append(g.numAliveGraph.YValues, fmt.Sprintf("%d", g.model.Patches.With(func(p *model.Patch) bool {
		return p.GetPropertySafe("alive").(bool)
	}).Count()))

	g.model.Tick()
}

func (g *Gol) Stats() map[string]interface{} {
	return map[string]interface{}{
		"num-alive": g.model.Patches.With(func(p *model.Patch) bool {
			return p.GetPropertySafe("alive").(bool)
		}).Count(),
		"num-alive-graph": g.numAliveGraph,
	}
}

func (g *Gol) Stop() bool {
	return g.model.Patches.All(func(p *model.Patch) bool {
		return !p.GetPropertySafe("alive").(bool)
	})
}

func (g *Gol) Widgets() []api.Widget {
	return []api.Widget{
		{
			PrettyName:      "Min Neighbors To Live",
			Id:              "min-neighbors-to-live",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "8",
			DefaultValue:    "2",
			ValuePointerInt: &g.minNeighborsToLive,
		},
		{
			PrettyName:      "Max Neighbors To Live",
			Id:              "max-neighbors-to-live",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "8",
			DefaultValue:    "3",
			ValuePointerInt: &g.maxNeighborsToLive,
		},
		{
			PrettyName:      "Min Neighbors To Reproduce",
			Id:              "min-neighbors-to-reproduce",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "8",
			DefaultValue:    "3",
			ValuePointerInt: &g.minNeighborsToReproduce,
		},
		{
			PrettyName:      "Max Neighbors To Reproduce",
			Id:              "max-neighbors-to-reproduce",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "8",
			DefaultValue:    "3",
			ValuePointerInt: &g.maxNeighborsToReproduce,
		},
		{
			PrettyName:        "Initial Alive",
			Id:                "initial-alive",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			MinValue:          "0",
			MaxValue:          "1",
			DefaultValue:      "0.5",
			StepAmount:        "0.01",
			ValuePointerFloat: &g.initialAlive,
		},
		api.NewIntSliderWidget("World Size", "world-size", "10", "100", "12", "1", &g.worldSize),
	}
}
