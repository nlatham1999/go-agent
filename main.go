package main

import (
	"github.com/nlatham1999/go-agent/examples/boid"
	"github.com/nlatham1999/go-agent/examples/boidconcurrent"
	"github.com/nlatham1999/go-agent/examples/gol"
	"github.com/nlatham1999/go-agent/examples/prims"
	"github.com/nlatham1999/go-agent/pkg/api"
)

func main() {

	schelling := prims.NewPrims()
	boid := boid.NewBoid()
	boidConurrent := boidconcurrent.NewBoid()
	gol := gol.NewGol()

	agentApi, err := api.NewApi(
		map[string]api.ModelInterface{
			"gameoflife":     gol,
			"schelling":      schelling,
			"boid":           boid,
			"boidconcurrent": boidConurrent,
		},
		api.ApiSettings{
			StoreSteps: false,
			ButtonTitles: map[string]string{
				"gameoflife":     "🟩 Game of Life",
				"schelling":      "🏃‍♂️ Schelling's Segregation Model",
				"boid":           "🐦 Boid Movement",
				"boidconcurrent": "🐦 Boid Movement Concurrent",
			},
			ButtonDescriptions: map[string]string{
				"gameoflife":     "Conway's Game of Life",
				"schelling":      "A simple social dynamics model",
				"boid":           "Simulating flocking birds",
				"boidconcurrent": "Simulating flocking birds concurrently",
			},
		},
	)

	if err != nil {
		panic(err)
	}

	agentApi.Serve()

}
