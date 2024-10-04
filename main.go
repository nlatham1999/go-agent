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
		{
			PrettyName:      "Grass Regrowth Time",
			TargetVariable:  "grass-regrowth-time",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "40",
			DefaultValue:    "20",
		},
		{
			PrettyName:      "Wolf Gain From Food",
			TargetVariable:  "wolf-gain-from-food",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "8",
			DefaultValue:    "2",
		},
		{
			PrettyName:      "Sheep Gain From Food",
			TargetVariable:  "sheep-gain-from-food",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "8",
			DefaultValue:    "2",
		},
		{
			PrettyName:      "Sheep Reproduce Rate",
			TargetVariable:  "sheep-reproduce-rate",
			WidgetType:      "slider",
			WidgetValueType: "float",
			MinValue:        "1",
			MaxValue:        "int",
			DefaultValue:    "50",
		},
		{
			PrettyName:      "Wolf Reproduce Rate",
			TargetVariable:  "wolf-reproduce-rate",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "100",
			DefaultValue:    "40",
		},
	}

	agentApi := api.NewApi(sim, widgets)

	agentApi.Serve()

	// err := converter.Convert("sample.txt")
	// if err != nil {
	// 	panic(err)
	// }

}
