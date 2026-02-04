package golgrowth

import (
	"fmt"
	"time"

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
	worldSize               int
	numGenerationsToGrow    int
	numAliveGraph           api.GraphWidget
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
	g.numGenerationsToGrow = 12
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
		MinPzCor: 0,
		MaxPzCor: g.numGenerationsToGrow - 1,
	}

	g.model = model.NewModel(settings)

	firstLayerPatches := g.model.PatchAtZLayer(0)

	firstLayerPatches.Ask(
		func(p *model.Patch) {
			if v := g.model.RandomFloat(1); v < g.initialAlive {
				p.SetProperty("alive", true)
				p.SetProperty("alive-next", true)
				p.Color.SetColor(model.Green)
			} else {
				p.SetProperty("alive", false)
				p.SetProperty("alive-next", false)
				p.Color.SetColor(model.Black)
			}
		},
	)

	g.numAliveGraph.XValues = []string{}
	g.numAliveGraph.YValues = []string{}

	return nil
}

func (g *Gol) Go() {

	currentTime := time.Now()
	defer func(startTime time.Time) {
		fmt.Println("Time per tick (ms):", time.Since(currentTime).Milliseconds())
	}(currentTime)

	patches := g.model.PatchAtZLayer(g.model.Ticks)

	// g.model.Patches.Ask(
	patches.Ask(
		func(p *model.Patch) {

			//get neighboring patches
			neighbors := p.NeighborsAtZOffset(0)

			//count the number of alive neighbors
			aliveNeighbors := neighbors.With(func(p *model.Patch) bool {
				alive := p.GetProperty("alive").(bool)
				return alive
			}).Count()

			alive := p.GetProperty("alive").(bool)
			if alive {
				// Alive cell: survives with 2-3 neighbors, dies otherwise
				if aliveNeighbors < g.minNeighborsToLive || aliveNeighbors > g.maxNeighborsToLive {
					p.SetProperty("alive-next", false)
				} else {
					p.SetProperty("alive-next", true)
				}
			} else {
				// Dead cell: becomes alive with exactly 3 neighbors
				if aliveNeighbors >= g.minNeighborsToReproduce && aliveNeighbors <= g.maxNeighborsToReproduce {
					p.SetProperty("alive-next", true)
				} else {
					p.SetProperty("alive-next", false)
				}
			}
		},
	)

	patches.Ask(
		func(p *model.Patch) {

			patchAbove := g.model.Patch3D(float64(p.XCor()), float64(p.YCor()), float64(p.ZCor()+1))

			if patchAbove == nil {
				return
			}

			patchAbove.SetProperty("alive", p.GetProperty("alive-next").(bool))
			if patchAbove.GetProperty("alive").(bool) {
				patchAbove.Color.SetColor(model.White)
			} else {
				patchAbove.Color.SetColor(model.Black)
			}

			// p.Color.SetColor(model.Black)
		},
	)

	g.numAliveGraph.XValues = append(g.numAliveGraph.XValues, fmt.Sprintf("%d", g.model.Ticks))
	g.numAliveGraph.YValues = append(g.numAliveGraph.YValues, fmt.Sprintf("%d", g.model.Patches.With(func(p *model.Patch) bool {
		return p.GetProperty("alive").(bool)
	}).Count()))

	g.model.Tick()
}

func (g *Gol) Stats() map[string]interface{} {
	return map[string]interface{}{
		"num-alive": g.model.Patches.With(func(p *model.Patch) bool {
			return p.GetProperty("alive").(bool)
		}).Count(),
		"num-alive-graph": g.numAliveGraph,
	}
}

func (g *Gol) Stop() bool {
	return g.model.Patches.All(func(p *model.Patch) bool {
		return !p.GetProperty("alive").(bool) || g.model.Ticks >= g.numGenerationsToGrow
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
		api.NewIntSliderWidget("Generations To Grow", "generations-to-grow", "1", "100", "12", "1", &g.numGenerationsToGrow),
	}
}
