package antpath

import (
	"fmt"

	"github.com/nlatham1999/go-agent/internal/api"
	"github.com/nlatham1999/go-agent/internal/model"
)

var ()

type AntPath struct {
	m *model.Model

	nestX float64
	nestY float64
	foodX float64
	foodY float64

	startDelay int

	timerVal int64
}

func NewAntPath() *AntPath {
	return &AntPath{
		startDelay: 3,
	}
}

func (a *AntPath) Model() *model.Model {
	return a.m
}

func (a *AntPath) Init() {

	fmt.Println("Initializing model")

	settings := model.ModelSettings{
		TurtleBreeds: []string{"leader", "follower"},
	}
	a.m = model.NewModel(settings)

	a.m.SetGlobal("max-ticks", 5000)
	a.m.SetGlobal("num-turtles", 100)

	if a.m != nil {
		fmt.Println("Model initialized")
	}
}

func (a *AntPath) SetUp() error {
	if a.m != nil {
		a.m.ClearAll()
	}

	a.m.SetDefaultShapeTurtles("bug")

	a.nestX = 10 + float64(a.m.MinPxCor())
	a.nestY = 0
	a.foodX = float64(a.m.MaxPxCor()) - 10
	a.foodY = 0

	a.m.CreateTurtles(1, "leader", []model.TurtleOperation{
		func(t *model.Turtle) {
			t.Color.SetColor(model.Red)
		},
	})

	numAnts := a.m.GetGlobal("num-turtles").(int)
	a.m.CreateTurtles(numAnts-1, "follower", []model.TurtleOperation{
		func(t *model.Turtle) {
			t.Color.SetColor(model.Yellow)
			t.SetHeading(0)
		},
	})

	a.m.Turtles("").MaxOneOf(func(t *model.Turtle) float64 {
		return float64(t.Who())
	}).Ask(
		[]model.TurtleOperation{
			func(t *model.Turtle) {
				t.Color.SetColor(model.Blue)
			},
		},
	)

	a.m.Turtles("").Ask([]model.TurtleOperation{
		func(t *model.Turtle) {
			t.SetXY(a.nestX, a.nestY)
		},
	})

	a.m.ResetTicks()

	return nil
}

func (a *AntPath) Go() {
	if a.m.Turtles("").All(func(t *model.Turtle) bool {
		return t.XCor() >= a.foodX
	}) {
		fmt.Println("All ants have reached the food")
		fmt.Println("time in ms:", a.m.Timer())
		fmt.Println("num ticks:", a.m.Ticks)

		return
	}

	a.m.Turtles("leader").Ask([]model.TurtleOperation{
		func(t *model.Turtle) {
			a.wiggle(t, 45)
			correctPath(t)
			if t.XCor() > (a.foodX - 5) {
				t.FaceXY(a.foodX, a.foodY)
			}
			if t.XCor() < a.foodX {
				t.Forward(0.5)
			}
		},
	})

	a.m.Turtles("follower").Ask([]model.TurtleOperation{
		func(t *model.Turtle) {
			t.FaceTurtle(a.m.Turtle("", t.Who()-1))
			if a.timeToStart(t) && t.XCor() < a.foodX {
				if t.Who() == 1 {
				}
				t.Forward(0.5)
			}
		},
	})

	a.m.Tick()
}

func (a *AntPath) wiggle(t *model.Turtle, angle float64) {
	t.Right(a.m.RandomFloat(angle))
	t.Left(a.m.RandomFloat(angle))
}

func correctPath(t *model.Turtle) {
	if t.GetHeading() > 90 && t.GetHeading() < 270 {
		t.Right(180)
	} else {
		if t.PatchAt(0, -5) == nil {
			t.Right(100)
		}
		if t.PatchAt(0, 5) == nil {
			t.Left(100)
		}
	}
}

func (a *AntPath) timeToStart(t *model.Turtle) bool {
	x := a.m.Turtle("", t.Who()-1).XCor()
	delay := a.nestX + float64(a.startDelay) + float64(a.m.RandomInt(a.startDelay))
	if t.Who() == 1 {
		// fmt.Println(x, delay)
	}
	return x > delay
}

func (a *AntPath) Stats() map[string]interface{} {
	return map[string]interface{}{
		"num-turtles": a.m.Turtles("").Count(),
		"num-at-food": a.m.Turtles("").With(func(t *model.Turtle) bool {
			return t.XCor() >= a.foodX
		}).Count(),
	}
}

// stop the model when all the ants have reached the food
func (a *AntPath) Stop() bool {
	return a.m.Turtles("").All(func(t *model.Turtle) bool {
		return t.XCor() >= a.foodX
	})
}

func (a *AntPath) Widgets() []api.Widget {
	return []api.Widget{}
}
