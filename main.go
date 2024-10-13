package main

import (
	"github.com/nlatham1999/go-agent/examples/prims"
	"github.com/nlatham1999/go-agent/internal/api"
)

// antpath "github.com/nlatham1999/go-agent/examples/ant-path"

func main() {

	// sim := antpath.NewAntPath()
	// sim := simplesim.NewSimpleSim()
	// sim := wolfsheep.NewWolfSheep()
	// sim := gol.NewGol()
	// sim := bees.NewBees()
	sim := prims.NewPrims()

	agentApi := api.NewApi(sim, api.ApiSettings{
		StoreSteps: true,
	})

	agentApi.Serve()

	// err := converter.Convert("sample.txt")
	// if err != nil {
	// 	panic(err)
	// }

}
