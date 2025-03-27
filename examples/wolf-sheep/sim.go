package wolfsheep

import (
	"fmt"

	"github.com/nlatham1999/go-agent/pkg/api"
	"github.com/nlatham1999/go-agent/pkg/model"
)

type WolfSheep struct {
	m *model.Model

	showEnergy          bool
	maxSheep            int
	grassRegrowthTime   int
	initialNumberSheep  int
	initialNumberWolves int
	wolfGainFromFood    int
	sheepGainFromFood   int
	sheepReproduceRate  float64
	wolfReproduceRate   float64
}

func NewWolfSheep() *WolfSheep {
	return &WolfSheep{}
}

func (ws *WolfSheep) Model() *model.Model {
	return ws.m
}

func (ws *WolfSheep) Init() {

	sheep := model.NewTurtleBreed("sheep", "", nil)
	wolves := model.NewTurtleBreed("wolves", "", nil)

	modelSettings := model.ModelSettings{
		TurtleBreeds: []*model.TurtleBreed{sheep, wolves},
		TurtleProperties: map[string]interface{}{
			"energy": 0,
		},
		PatchProperties: map[string]interface{}{
			"countdown": int(0),
		},
		MinPxCor: -15,
		MaxPxCor: 15,
		MinPyCor: -15,
		MaxPyCor: 15,
	}

	ws.m = model.NewModel(modelSettings)

	ws.showEnergy = false
	ws.maxSheep = 300
	ws.grassRegrowthTime = 30
	ws.initialNumberSheep = 20
	ws.initialNumberWolves = 4
	ws.wolfGainFromFood = 2
	ws.sheepGainFromFood = 2
	ws.sheepReproduceRate = 50.0
	ws.wolfReproduceRate = 40.0
}

func (ws *WolfSheep) SetUp() error {
	ws.m.ClearAll()

	ws.m.Patches.Ask(
		func(p *model.Patch) {

			if ws.m.RandomFloat(1) < 0.5 {
				p.PColor.SetColor(model.Green)
				p.SetProperty("countdown", ws.grassRegrowthTime)
			} else {
				p.PColor.SetColor(model.Brown)
				p.SetProperty("countdown", ws.grassRegrowthTime)
			}
		},
	)

	sheep := ws.m.TurtleBreed("sheep")
	wolves := ws.m.TurtleBreed("wolves")

	sheep.CreateAgents(ws.initialNumberSheep,
		func(t *model.Turtle) {
			// t.Shape("sheep")
			t.Color.SetColor(model.White)
			// t.Size(1.5)
			t.SetLabelColor(model.Blue)
			t.SetProperty("energy", ws.m.RandomInt(2*ws.sheepGainFromFood))
			t.SetXY(ws.m.RandomXCor(), ws.m.RandomYCor())
			t.SetSize(.5)
		},
	)

	wolves.CreateAgents(ws.initialNumberWolves,
		func(t *model.Turtle) {
			// t.Shape("wolf")
			t.Color.SetColor(model.Black)
			// t.Size(2)
			t.SetLabelColor(model.White)
			t.SetProperty("energy", ws.m.RandomInt(2*ws.wolfGainFromFood))
			t.SetXY(float64(ws.m.RandomXCor()), ws.m.RandomYCor())
			t.SetSize(.5)
		},
	)

	ws.m.Turtles().Ask(
		func(t *model.Turtle) {
			if ws.showEnergy {
				t.SetLabel(fmt.Sprintf("%v", t.GetProperty("energy")))
			} else {
				t.SetLabel("")
			}
		},
	)

	ws.m.ResetTicks()
	return nil
}

func (ws *WolfSheep) Go() {
	if ws.m.Turtles().Count() == 0 {
		return
	}

	wolves := ws.m.TurtleBreed("wolves")
	sheep := ws.m.TurtleBreed("sheep")

	if wolves.Agents().Count() == 0 && sheep.Agents().Count() > ws.maxSheep {
		fmt.Println("The sheep have inherited the earth")
		return
	}

	sheep.Agents().Ask(
		func(t *model.Turtle) {
			ws.move(t)
			energy, err := t.GetPropI("energy")
			if err != nil {
				fmt.Println(err)
				return
			}
			t.SetProperty("energy", energy-1)
			ws.EatGrass(t)
			ws.Death(t)
			ws.reproduceSheep(t)
		},
	)

	wolves.Agents().Ask(
		ws.move,
	)
	wolves.Agents().Ask(
		func(t *model.Turtle) {
			ws.move(t)
			t.SetProperty("energy", t.GetProperty("energy").(int)-1)
			ws.EatSheep(t)
			ws.Death(t)
			ws.reproduceWolves(t)
		},
	)

	ws.m.Patches.Ask(
		ws.growGrass,
	)

	wolves.Agents().Ask(
		func(t *model.Turtle) {
			if ws.showEnergy {
				t.SetLabel(fmt.Sprintf("%v", t.GetProperty("energy")))
			} else {
				t.SetLabel("")
			}
		},
	)

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
		t.SetProperty("energy", t.GetProperty("energy").(int)+ws.sheepGainFromFood)
	}
}

func (ws *WolfSheep) reproduceSheep(t *model.Turtle) {
	energy, err := t.GetPropI("energy")
	if err != nil {
		return
	}

	if energy <= 0 {
		return
	}
	if ws.m.RandomFloat(100) < ws.sheepReproduceRate {

		t.SetProperty("energy", energy/2)
		t.Hatch(1,
			func(t *model.Turtle) {
				t.Right(ws.m.RandomFloat(360))
				t.Forward(1)
			},
		)
	}
}

func (ws *WolfSheep) reproduceWolves(t *model.Turtle) {
	energy, err := t.GetPropI("energy")
	if err != nil {
		return
	}

	if energy <= 0 {
		return
	}
	if ws.m.RandomFloat(100) < ws.sheepReproduceRate {
		t.SetProperty("energy", energy/2)
		t.Hatch(1,
			func(t *model.Turtle) {
				t.Right(ws.m.RandomFloat(360))
				t.Forward(1)
			},
		)
	}
}

func (ws *WolfSheep) EatSheep(t *model.Turtle) {

	sheep := ws.m.TurtleBreed("sheep")

	prey, err := sheep.Agents().First()
	if err != nil {
		return
	}
	if prey != nil {
		prey.Die()
		t.SetProperty("energy", t.GetProperty("energy").(int)+ws.wolfGainFromFood)
	}
}

func (ws *WolfSheep) Death(t *model.Turtle) {
	if t.GetProperty("energy").(int) <= 0 {
		t.Die()
	}
}

func (ws *WolfSheep) growGrass(p *model.Patch) {
	if p.PColor == model.Brown {
		if p.GetPropI("countdown") <= 0 {
			p.PColor.SetColor(model.Green)
			p.SetProperty("countdown", ws.grassRegrowthTime)
		} else {
			p.SetProperty("countdown", p.GetPropI("countdown")-1)
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
	if ws.m.Turtles().Count() == 0 {
		return true
	}

	wolves := ws.m.TurtleBreed("wolves")
	sheep := ws.m.TurtleBreed("sheep")

	if wolves.Agents().Count() == 0 && sheep.Agents().Count() > ws.maxSheep {
		fmt.Println("The sheep have inherited the earth")
		return true
	}

	return false
}

func (ws *WolfSheep) Widgets() []api.Widget {
	return []api.Widget{
		{
			PrettyName:      "Max Sheep",
			Id:              "max-sheep",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "200",
			MaxValue:        "1000",
			DefaultValue:    "300",
		},
		{
			PrettyName:      "Initial Number Of Sheep",
			Id:              "initial-number-sheep",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "100",
			DefaultValue:    "20",
		},
		{
			PrettyName:      "Initial Number Of Wolves",
			Id:              "initial-number-wolves",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "100",
			DefaultValue:    "4",
		},
		{
			PrettyName:      "Grass Regrowth Time",
			Id:              "grass-regrowth-time",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "40",
			DefaultValue:    "20",
		},
		{
			PrettyName:      "Wolf Gain From Food",
			Id:              "wolf-gain-from-food",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "8",
			DefaultValue:    "2",
		},
		{
			PrettyName:      "Sheep Gain From Food",
			Id:              "sheep-gain-from-food",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "8",
			DefaultValue:    "2",
		},
		{
			PrettyName:      "Sheep Reproduce Rate",
			Id:              "sheep-reproduce-rate",
			WidgetType:      "slider",
			WidgetValueType: "float",
			MinValue:        "1",
			MaxValue:        "int",
			DefaultValue:    "50",
		},
		{
			PrettyName:      "Wolf Reproduce Rate",
			Id:              "wolf-reproduce-rate",
			WidgetType:      "slider",
			WidgetValueType: "int",
			MinValue:        "1",
			MaxValue:        "100",
			DefaultValue:    "40",
		},
	}
}
