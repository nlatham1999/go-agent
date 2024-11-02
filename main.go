package main

import (
	"github.com/nlatham1999/go-agent/api"
	"github.com/nlatham1999/go-agent/playgound"
)

func main() {

	sim := playgound.NewSim()

	agentApi := api.NewApi(sim, api.ApiSettings{
		StoreSteps: false,
	})

	agentApi.Serve()

}
