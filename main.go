package main

import (
	antpath "github.com/nlatham1999/go-agent/examples/ant-path"
	"github.com/nlatham1999/go-agent/internal/api"
)

// antpath "github.com/nlatham1999/go-agent/examples/ant-path"

func main() {

	sim := antpath.NewAntPath()
	// sim := simplesim.NewSimpleSim()

	agentApi := api.NewApi(sim)

	agentApi.Serve()

	// err := converter.Convert("sample.txt")
	// if err != nil {
	// 	panic(err)
	// }

}
