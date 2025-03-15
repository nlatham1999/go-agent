package main

import (
	"github.com/nlatham1999/go-agent/api"
	"github.com/nlatham1999/go-agent/examples/gol"
)

func main() {

	// sim := playgound.NewSim()
	// sim := boid.NewBoid()
	sim := gol.NewGol()
	// sim := flocking.NewFlocking()

	agentApi := api.NewApi(sim, api.ApiSettings{
		StoreSteps: false,
		Title:      "Test Model",
	})

	agentApi.Serve()

}
