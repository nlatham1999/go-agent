package main

import (
	"github.com/nlatham1999/go-agent/api"
	"github.com/nlatham1999/go-agent/examples/boid"
)

func main() {

	// sim := playgound.NewSim()
	sim := boid.NewBoid()

	agentApi := api.NewApi(sim, api.ApiSettings{
		StoreSteps: false,
	})

	agentApi.Serve()

}
