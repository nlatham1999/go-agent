package main

import (
	wolfsheep "github.com/nlatham1999/go-agent/examples/wolf-sheep"
	"github.com/nlatham1999/go-agent/internal/api"
)

// antpath "github.com/nlatham1999/go-agent/examples/ant-path"

func main() {

	// sim := antpath.NewAntPath()
	// sim := simplesim.NewSimpleSim()
	sim := wolfsheep.NewWolfSheep()
	// sim := gol.NewGol()

	agentApi := api.NewApi(sim, api.ApiSettings{
		StoreSteps: true,
	})

	agentApi.Serve()

	// err := converter.Convert("sample.txt")
	// if err != nil {
	// 	panic(err)
	// }

}
