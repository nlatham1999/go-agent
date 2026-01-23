package main

import (
	"fmt"

	"github.com/nlatham1999/go-agent/examples/boid"
	"github.com/nlatham1999/go-agent/examples/boidconcurrent"
	"github.com/nlatham1999/go-agent/examples/flocking"
	"github.com/nlatham1999/go-agent/examples/gol"
	mouseinteractions "github.com/nlatham1999/go-agent/examples/mouse-interactions"
	"github.com/nlatham1999/go-agent/examples/prims"
	"github.com/nlatham1999/go-agent/pkg/api"
	"github.com/nlatham1999/go-agent/playground"
)

func main() {
	RunServer()

	// RunSingleModel(geneticalgorithm.NewGeneticAlgorithm())
}

func RunServer() {

	schelling := prims.NewPrims()
	boid := boid.NewBoid()
	boidConurrent := boidconcurrent.NewBoid()
	gol := gol.NewGol()
	playground := playground.NewSim()
	flocking := flocking.NewFlocking()
	mouseInteractions := mouseinteractions.NewSim()

	agentApi, err := api.NewApi(
		map[string]api.ModelInterface{
			"gameoflife":        gol,
			"schelling":         schelling,
			"boid":              boid,
			"boidconcurrent":    boidConurrent,
			"playground":        playground,
			"flocking":          flocking,
			"mouseinteractions": mouseInteractions,
		},
		api.ApiSettings{
			StoreSteps: false,
			ButtonTitles: map[string]string{
				"gameoflife":        "🟩 Game of Life",
				"schelling":         "🏃‍♂️ Schelling's Segregation Model",
				"boid":              "🐦 Boid Movement",
				"boidconcurrent":    "🐦 Boid Movement Concurrent",
				"playground":        "🎮 Playground",
				"flocking":          "🐦 Flocking",
				"mouseinteractions": "🖱️ Mouse Interactions",
			},
			ButtonDescriptions: map[string]string{
				"gameoflife":        "Conway's Game of Life",
				"schelling":         "A simple social dynamics model",
				"boid":              "Simulating flocking birds",
				"boidconcurrent":    "Simulating flocking birds concurrently",
				"playground":        "Playground for quick testing",
				"flocking":          "Flocking simulation",
				"mouseinteractions": "Mouse interactions example",
			},
		},
	)

	if err != nil {
		panic(err)
	}

	agentApi.Serve()
}

func RunSingleModel(model api.ModelInterface) {

	model.Init()

	model.SetUp()

	fmt.Println("generation, high score, average_score, highest amount picked")

	for !model.Stop() {
		model.Go()

		// PrintStats(model)
		stats := model.Stats()

		fmt.Println(fmt.Sprintf("%d,%d,%d,%d", stats["generation"], stats["high score"], stats["average score"], stats["highest amount picked"]))
	}
}

func PrintStats(model api.ModelInterface) {
	stats := model.Stats()
	for key, val := range stats {
		fmt.Println(key, val)
	}
}
