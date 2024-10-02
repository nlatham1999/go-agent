package main

import (
	antpath "github.com/nlatham1999/go-agent/examples/ant-path"
	"github.com/nlatham1999/go-agent/internal/api"
)

// antpath "github.com/nlatham1999/go-agent/examples/ant-path"

func main() {

	sim := antpath.NewAntPath()
	// sim := simplesim.NewSimpleSim()

	widgets := []api.Widget{
		{
			PrettyName:      "Number of Turtles",
			TargetVariable:  "num-turtles",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "2",
			MaxValue:        "1000",
			DefaultValue:    "50",
		},
	}

	agentApi := api.NewApi(sim, widgets)

	agentApi.Serve()

	// err := converter.Convert("sample.txt")
	// if err != nil {
	// 	panic(err)
	// }

}
