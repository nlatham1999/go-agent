package main

import (
	"github.com/nlatham1999/go-agent/api"
	"github.com/nlatham1999/go-agent/playgound"
)

// antpath "github.com/nlatham1999/go-agent/examples/ant-path"

func main() {

	// sim := antpath.NewAntPath()
	// sim := simplesim.NewSimpleSim()
	// sim := wolfsheep.NewWolfSheep()
	// sim := gol.NewGol()
	// sim := bees.NewBees()
	// sim := prims.NewPrims()
	sim := playgound.NewSim()

	agentApi := api.NewApi(sim, api.ApiSettings{
		StoreSteps: false,
	})

	agentApi.Serve()

	// fmt.Println("Setting up model")
	// sim.Init()
	// sim.Model().SetGlobal("nodes", 4000)
	// sim.SetUp()
	// fmt.Println("running model")
	// i := 0
	// for !sim.Stop() {
	// 	sim.Go()
	// 	if i%100 == 0 {
	// 		fmt.Println("Ticks: ", sim.Model().Ticks)
	// 		fmt.Println(sim.Stats())
	// 	}
	// 	i++
	// }
	// fmt.Println("model finished")

	// err := converter.Convert("sample.txt")
	// if err != nil {
	// 	panic(err)
	// }

}
