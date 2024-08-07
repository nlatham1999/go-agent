package antpath

import (
	"fmt"

	"github.com/nlatham1999/go-agent/internal/model"
)

var (
	m *model.Model

	nestX float64
	nestY float64
	foodX float64
	foodY float64

	numAnts = 10000

	startDelay = 3

	timerVal int64
)

func SetUp() {
	if m != nil {
		m.ClearAll()
	}

	m = model.NewModel(nil, nil, nil, []string{"leader", "follower"}, nil, nil, false, false)

	m.SetDefaultShapeTurtles("bug")

	nestX = 10 + float64(m.MinPxCor)
	nestY = 0
	foodX = float64(m.MaxPxCor) - 10
	foodY = 0

	m.CreateTurtles(1, "leader", []model.TurtleOperation{
		func(t *model.Turtle) {
			t.Color.SetColor(model.Red)
		},
	})

	m.CreateTurtles(numAnts-1, "follower", []model.TurtleOperation{
		func(t *model.Turtle) {
			t.Color.SetColor(model.Yellow)
			t.SetHeading(0)
		},
	})

	m.Turtles("").MaxOneOf(func(t *model.Turtle) float64 {
		return float64(t.Who())
	}).Ask(
		[]model.TurtleOperation{
			func(t *model.Turtle) {
				t.Color.SetColor(model.Blue)
			},
		},
	)

	m.Turtles("").Ask([]model.TurtleOperation{
		func(t *model.Turtle) {
			t.SetXY(nestX, nestY)
		},
	})

	m.ResetTicks()
}

func Go() {
	for {
		if m.Turtles("").All(func(t *model.Turtle) bool {
			return t.XCor() >= foodX
		}) {
			fmt.Println("All ants have reached the food")
			fmt.Println("time in ms:", m.Timer())
			fmt.Println("num ticks:", m.Ticks)

			return
		}

		m.Turtles("leader").Ask([]model.TurtleOperation{
			func(t *model.Turtle) {
				wiggle(t, 45)
				correctPath(t)
				if t.XCor() > (foodX - 5) {
					t.FaceXY(foodX, foodY)
				}
				if t.XCor() < foodX {
					t.Forward(0.5)
				}
			},
		})

		m.Turtles("follower").Ask([]model.TurtleOperation{
			func(t *model.Turtle) {
				t.FaceTurtle(m.Turtle("", t.Who()-1))
				if timeToStart(t) && t.XCor() < foodX {
					if t.Who() == 1 {
					}
					t.Forward(0.5)
				}
			},
		})

		m.Tick()
	}
}

func wiggle(t *model.Turtle, angle float64) {
	t.Right(m.RandomFloat(angle))
	t.Left(m.RandomFloat(angle))
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

func timeToStart(t *model.Turtle) bool {
	x := m.Turtle("", t.Who()-1).XCor()
	delay := nestX + float64(startDelay) + float64(m.RandomInt(startDelay))
	if t.Who() == 1 {
		// fmt.Println(x, delay)
	}
	return x > delay
}
