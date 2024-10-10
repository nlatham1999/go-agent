package wolfsheep

import (
	"fmt"

	"github.com/nlatham1999/go-agent/internal/api"
	"github.com/nlatham1999/go-agent/internal/model"
)

type WolfSheep struct {
	m *model.Model
}

func NewWolfSheep() *WolfSheep {
	return &WolfSheep{}
}

func (ws *WolfSheep) Model() *model.Model {
	return ws.m
}

func (ws *WolfSheep) Init() {

	modelSettings := model.ModelSettings{
		TurtleBreeds: []string{"sheep", "wolves"},
		TurtlesOwn: map[string]interface{}{
			"energy": 0,
		},
		PatchesOwn: map[string]interface{}{
			"countdown": int(0),
		},
		Globals: map[string]interface{}{
			"show-energy":           false,
			"max-sheep":             300,
			"grass-regrowth-time":   30,
			"initial-number-sheep":  20,
			"initial-number-wolves": 4,
			"wolf-gain-from-food":   2,
			"sheep-gain-from-food":  2,
			"sheep-reproduce-rate":  50.0,
			"wolf-reprodue-rate":    40.0,
		},
		MinPxCor: -15,
		MaxPxCor: 15,
		MinPyCor: -15,
		MaxPyCor: 15,
	}

	ws.m = model.NewModel(modelSettings)
}

func (ws *WolfSheep) SetUp() error {
	ws.m.ClearAll()

	ws.m.Patches.Ask([]model.PatchOperation{
		func(p *model.Patch) {

			grassRegrowthTime := ws.m.GetGlobal("grass-regrowth-time")
			if ws.m.RandomFloat(1) < 0.5 {
				p.PColor.SetColor(model.Green)
				p.SetOwn("countdown", grassRegrowthTime.(int))
			} else {
				p.PColor.SetColor(model.Brown)
				p.SetOwn("countdown", ws.m.RandomInt(grassRegrowthTime.(int)))
			}
		},
	})

	initialNumberSheep, err := ws.m.GetGlobalI("initial-number-sheep")
	if err != nil {
		fmt.Println(err)
		return err
	}

	sheepGainFromFood, err := ws.m.GetGlobalI("sheep-gain-from-food")
	if err != nil {
		fmt.Println(err)
		return err
	}

	ws.m.CreateTurtles(initialNumberSheep, "sheep", []model.TurtleOperation{
		func(t *model.Turtle) {
			// t.Shape("sheep")
			t.Color.SetColor(model.White)
			// t.Size(1.5)
			t.SetLabelColor(model.Blue)
			t.SetOwn("energy", ws.m.RandomInt(2*sheepGainFromFood))
			t.SetXY(ws.m.RandomXCor(), ws.m.RandomYCor())
			t.SetSize(.5)
		},
	})

	initialNumberWolves, err := ws.m.GetGlobalI("initial-number-wolves")
	if err != nil {
		fmt.Println(err)
		return err
	}

	wolfGainFromFood, err := ws.m.GetGlobalI("wolf-gain-from-food")
	if err != nil {
		fmt.Println(err)
		return err
	}

	ws.m.CreateTurtles(initialNumberWolves, "wolves", []model.TurtleOperation{
		func(t *model.Turtle) {
			// t.Shape("wolf")
			t.Color.SetColor(model.Black)
			// t.Size(2)
			t.SetLabelColor(model.White)
			t.SetOwn("energy", ws.m.RandomInt(2*wolfGainFromFood))
			t.SetXY(float64(ws.m.RandomXCor()), ws.m.RandomYCor())
			t.SetSize(.5)
		},
	})

	showEnergy, _ := ws.m.GetGlobalB("show-energy")
	ws.m.Turtles("").Ask([]model.TurtleOperation{
		func(t *model.Turtle) {
			if showEnergy {
				t.SetLabel(fmt.Sprintf("%v", t.GetOwn("energy")))
			} else {
				t.SetLabel("")
			}
		},
	})

	ws.m.ResetTicks()
	return nil
}

func (ws *WolfSheep) Go() {
	if ws.m.Turtles("").Count() == 0 {
		return
	}

	if ws.m.Turtles("wolves").Count() == 0 && ws.m.Turtles("sheep").Count() > ws.m.GetGlobal("max-sheep").(int) {
		fmt.Println("The sheep have inherited the earth")
		return
	}

	ws.m.Turtles("sheep").Ask([]model.TurtleOperation{
		ws.move,
		func(t *model.Turtle) {

			energy, err := t.GetOwnI("energy")
			if err != nil {
				fmt.Println(err)
				return
			}

			t.SetOwn("energy", energy-1)
		},
		ws.EatGrass,
		ws.Death,
		ws.reproduceSheep,
	})

	ws.m.Turtles("wolves").Ask([]model.TurtleOperation{
		ws.move,
		func(t *model.Turtle) {
			t.SetOwn("energy", t.GetOwn("energy").(int)-1)
		},
		ws.EatSheep,
		ws.Death,
		ws.reproduceWolves,
	})

	ws.m.Patches.Ask([]model.PatchOperation{
		ws.growGrass,
	})

	showEnergy, _ := ws.m.GetGlobalB("show-energy")
	ws.m.Turtles("").Ask([]model.TurtleOperation{
		func(t *model.Turtle) {
			if showEnergy {
				t.SetLabel(fmt.Sprintf("%v", t.GetOwn("energy")))
			} else {
				t.SetLabel("")
			}
		},
	})

	ws.m.Tick()
}

func (ws *WolfSheep) move(t *model.Turtle) {
	t.Right(ws.m.RandomFloat(50))
	t.Left(ws.m.RandomFloat(50))
	t.Forward(1)
}

func (ws *WolfSheep) EatGrass(t *model.Turtle) {
	if t.PatchHere().PColor == model.Green {
		t.PatchHere().PColor.SetColor(model.Brown)
		sheepGainFromFood := ws.m.GetGlobal("sheep-gain-from-food").(int)
		t.SetOwn("energy", t.GetOwn("energy").(int)+sheepGainFromFood)
	}
}

func (ws *WolfSheep) reproduceSheep(t *model.Turtle) {
	energy, err := t.GetOwnI("energy")
	if err != nil {
		return
	}

	if energy <= 0 {
		return
	}
	if ws.m.RandomFloat(100) < ws.m.GetGlobal("sheep-reproduce-rate").(float64) {

		t.SetOwn("energy", energy/2)
		t.Hatch("", 1, []model.TurtleOperation{
			func(t *model.Turtle) {
				t.Right(ws.m.RandomFloat(360))
				t.Forward(1)
			},
		})
	}
}

func (ws *WolfSheep) reproduceWolves(t *model.Turtle) {
	energy, err := t.GetOwnI("energy")
	if err != nil {
		return
	}

	if energy <= 0 {
		return
	}
	if ws.m.RandomFloat(100) < ws.m.GetGlobal("sheep-reproduce-rate").(float64) {
		t.SetOwn("energy", energy/2)
		t.Hatch("", 1, []model.TurtleOperation{
			func(t *model.Turtle) {
				t.Right(ws.m.RandomFloat(360))
				t.Forward(1)
			},
		})
	}
}

func (ws *WolfSheep) EatSheep(t *model.Turtle) {
	prey := t.PatchHere().TurtlesHere("sheep").OneOf()
	if prey != nil {
		prey.Die()
		t.SetOwn("energy", t.GetOwn("energy").(int)+ws.m.GetGlobal("wolf-gain-from-food").(int))
	}
}

func (ws *WolfSheep) Death(t *model.Turtle) {
	if t.GetOwn("energy").(int) <= 0 {
		t.Die()
	}
}

func (ws *WolfSheep) growGrass(p *model.Patch) {
	if p.PColor == model.Brown {
		if p.GetOwnI("countdown") <= 0 {
			p.PColor.SetColor(model.Green)
			p.SetOwn("countdown", ws.m.GetGlobal("grass-regrowth-time").(int))
		} else {
			p.SetOwn("countdown", p.GetOwnI("countdown")-1)
		}
	}
}

func (ws *WolfSheep) grass() *model.PatchAgentSet {
	return ws.m.Patches.With(func(p *model.Patch) bool {
		return p.PColor == model.Green
	})
}

func (ws *WolfSheep) Stats() map[string]interface{} {
	return map[string]interface{}{}
}

// stop the model when all the ants have reached the food
func (ws *WolfSheep) Stop() bool {
	if ws.m.Turtles("").Count() == 0 {
		return true
	}

	if ws.m.Turtles("wolves").Count() == 0 && ws.m.Turtles("sheep").Count() > ws.m.GetGlobal("max-sheep").(int) {
		fmt.Println("The sheep have inherited the earth")
		return true
	}

	return false
}

func (ws *WolfSheep) Widgets() []api.Widget {
	return []api.Widget{
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
}
