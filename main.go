package main

import (
	antpath "github.com/nlatham1999/go-agent/examples/ant-path"
	"github.com/nlatham1999/go-agent/internal/api"
)

// antpath "github.com/nlatham1999/go-agent/examples/ant-path"

func main() {

	i := antpath.Init
	s := antpath.SetUp
	g := antpath.Go
	m := antpath.Model

	agentApi := api.NewApi(m, i, s, g)

	agentApi.Serve()

	// err := converter.Convert("sample.txt")
	// if err != nil {
	// 	panic(err)
	// }

}
