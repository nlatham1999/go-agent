package main

import (
	"github.com/nlatham1999/go-agent/examples/boid"
	"github.com/nlatham1999/go-agent/examples/boidconcurrent"
	"github.com/nlatham1999/go-agent/examples/flocking"
	"github.com/nlatham1999/go-agent/examples/gol"
	"github.com/nlatham1999/go-agent/examples/prims"
	"github.com/nlatham1999/go-agent/pkg/api"
	"github.com/nlatham1999/go-agent/playground"
)

func main() {

	schelling := prims.NewPrims()
	boid := boid.NewBoid()
	boidConurrent := boidconcurrent.NewBoid()
	gol := gol.NewGol()
	playground := playground.NewSim()
	flocking := flocking.NewFlocking()

	agentApi, err := api.NewApi(
		map[string]api.ModelInterface{
			"gameoflife":     gol,
			"schelling":      schelling,
			"boid":           boid,
			"boidconcurrent": boidConurrent,
			"playground":     playground,
			"flocking":       flocking,
		},
		api.ApiSettings{
			StoreSteps: false,
			ButtonTitles: map[string]string{
				"gameoflife":     "🟩 Game of Life",
				"schelling":      "🏃‍♂️ Schelling's Segregation Model",
				"boid":           "🐦 Boid Movement",
				"boidconcurrent": "🐦 Boid Movement Concurrent",
				"playground":     "🎮 Playground",
				"flocking":       "🐦 Flocking",
			},
			ButtonDescriptions: map[string]string{
				"gameoflife":     "Conway's Game of Life",
				"schelling":      "A simple social dynamics model",
				"boid":           "Simulating flocking birds",
				"boidconcurrent": "Simulating flocking birds concurrently",
				"playground":     "Playground for quick testing",
				"flocking":       "Flocking simulation",
			},
		},
	)

	if err != nil {
		panic(err)
	}

	agentApi.Serve()

}
