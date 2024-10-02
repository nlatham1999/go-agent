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

	// "max-sheep":             300,
	// "grass-regrowth-time":   30,
	// "initial-number-sheep":  20,
	// "initial-number-wolves": 4,
	// "wolf-gain-from-food":   2,
	// "sheep-gain-from-food":  2,
	// "sheep-reproduce-rate":  50.0,
	// "wolf-reprodue-rate":    40.0,

	widgets := []api.Widget{
		{
			PrettyName:      "Max Sheep",
			TargetVariable:  "max-sheep",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "200",
			MaxValue:        "1000",
			DefaultValue:    "300",
		},
		{
			PrettyName:      "Initial Number Of Sheep",
			TargetVariable:  "initial-number-sheep",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "100",
			DefaultValue:    "20",
		},
		{
			PrettyName:      "Initial Number Of Wolves",
			TargetVariable:  "initial-number-wolves",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "100",
			DefaultValue:    "4",
		},
	}

	agentApi := api.NewApi(sim, widgets)

	agentApi.Serve()

	// err := converter.Convert("sample.txt")
	// if err != nil {
	// 	panic(err)
	// }

}
