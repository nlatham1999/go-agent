package main

import (
	"fmt"

	"github.com/nlatham1999/go-agent/examples/boid"
	"github.com/nlatham1999/go-agent/examples/boid3D"
	"github.com/nlatham1999/go-agent/examples/boidconcurrent"
	"github.com/nlatham1999/go-agent/examples/flocking"
	"github.com/nlatham1999/go-agent/examples/gol"
	"github.com/nlatham1999/go-agent/examples/golgrowth"
	mouseinteractions "github.com/nlatham1999/go-agent/examples/mouse-interactions"
	"github.com/nlatham1999/go-agent/examples/prims"
	"github.com/nlatham1999/go-agent/examples/prims3d"
	"github.com/nlatham1999/go-agent/pkg/api"
	"github.com/nlatham1999/go-agent/playground"
)

func main() {
	RunServer()

	// RunSingleModel(geneticalgorithm.NewGeneticAlgorithm())
}

func RunServer() {

	prims := prims.NewPrims()
	boid := boid.NewBoid()
	boidConurrent := boidconcurrent.NewBoid()
	gol := gol.NewGol()
	playground := playground.NewSim()
	flocking := flocking.NewFlocking()
	mouseInteractions := mouseinteractions.NewSim()
	boid3d := boid3D.NewBoid3D()
	golgrowth := golgrowth.NewGol()
	prims3d := prims3d.NewPrims()

	agentApi, err := api.NewApi(
		map[string]api.ModelInterface{
			"gameoflife":        gol,
			"prims":             prims,
			"boid":              boid,
			"boidconcurrent":    boidConurrent,
			"playground":        playground,
			"flocking":          flocking,
			"mouseinteractions": mouseInteractions,
			"boid3D":            boid3d,
			"golgrowth":         golgrowth,
			"prims3d":           prims3d,
		},
		api.ApiSettings{
			StoreSteps: false,
			ButtonTitles: map[string]string{
				"gameoflife":        "🟩 Game of Life",
				"prims":             "🧱 Prim's Maze Generation",
				"boid":              "🐦 Boid Movement",
				"boidconcurrent":    "🐦 Boid Movement Concurrent",
				"playground":        "🎮 Playground",
				"flocking":          "🐦 Flocking",
				"mouseinteractions": "🖱️ Mouse Interactions",
				"boid3D":            "🐦 3D Boid Movement",
				"golgrowth":         "🟩 3D visualization of Game of Life growth",
				"prims3d":           "🧱 3D Prim's Maze Generation",
			},
			ButtonDescriptions: map[string]string{
				"gameoflife":        "Conway's Game of Life",
				"prims":             "Generating mazes using Prim's algorithm",
				"boid":              "Simulating flocking birds",
				"boidconcurrent":    "Simulating flocking birds concurrently",
				"playground":        "Playground for quick testing",
				"flocking":          "Flocking simulation",
				"mouseinteractions": "Mouse interactions example",
				"boid3D":            "Simulating flocking birds in 3D",
				"golgrowth":         "3D visualization of Game of Life growth",
				"prims3d":           "Generating mazes using Prim's algorithm in 3D",
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
