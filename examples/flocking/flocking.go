package flocking

import (
	"fmt"
	"math"

	"github.com/nlatham1999/go-agent/api"
	"github.com/nlatham1999/go-agent/model"
)

var _ api.ModelInterface = (*Flocking)(nil)

type Flocking struct {
	model *model.Model

	population        int
	vision            float64
	minimumSeparation float64
	maxAlignTurn      float64
	maxCohereTurn     float64
	maxSeperateTurn   float64
}

func NewFlocking() *Flocking {
	return &Flocking{}
}

func (f *Flocking) Init() {
	modelSettings := model.ModelSettings{
		TurtleProperties: map[string]interface{}{
			"flockmates":       nil,
			"nearest-neighbor": nil,
		},
		MinPxCor:  -33,
		MinPyCor:  -33,
		MaxPxCor:  33,
		MaxPyCor:  33,
		WrappingX: true,
		WrappingY: true,
	}

	f.model = model.NewModel(modelSettings)

	f.population = 300
	f.vision = 5
	f.minimumSeparation = 1
	f.maxAlignTurn = 5
	f.maxCohereTurn = 3
	f.maxSeperateTurn = 1.5
}

func (f *Flocking) SetUp() error {
	f.model.ClearAll()
	_, err := f.model.CreateTurtles(f.population, func(t *model.Turtle) {
		t.Color = f.model.RandomColor()
		t.SetSize(.5)
		t.SetXY(f.model.RandomXCor(), f.model.RandomYCor())
		t.SetProperty("flockmates", nil)
		t.Shape = "triangle"
	})
	if err != nil {
		return err
	}

	f.model.ResetTicks()

	return nil
}

func (f *Flocking) Go() {
	f.model.Turtles().Ask(
		f.flock,
	)

	f.model.Turtles().Ask(
		func(t *model.Turtle) {
			t.Forward(1)
		},
	)
	f.model.Tick()
}

func (f *Flocking) flock(t *model.Turtle) {
	f.findFlockmates(t)

	if t.GetProperty("flockmates") != nil && t.GetProperty("flockmates").(*model.TurtleAgentSet).Count() > 0 {
		if t.GetProperty("nearest-neighbor") != nil {
			neighbor := t.GetProperty("nearest-neighbor").(*model.Turtle)
			if t.DistanceTurtle(neighbor) < f.minimumSeparation {
				f.seperate(t)
			} else {
				f.align(t)
				f.cohere(t)
			}

		}
	}
}

func (f *Flocking) findFlockmates(t *model.Turtle) {

	flockMates := model.NewTurtleAgentSet(nil)

	t.SetProperty("nearest-neighbor", nil)

	minDistance := 1000000.0
	f.model.Turtles().Ask(
		func(t2 *model.Turtle) {
			distance := t.DistanceTurtle(t2)
			if t != t2 && distance < f.vision {
				flockMates.Add(t2)
				if distance < minDistance {
					minDistance = distance
					t.SetProperty("nearest-neighbor", t2)
				}
			}
		},
	)

	t.SetProperty("flockmates", flockMates)
}

func (f *Flocking) seperate(t *model.Turtle) {

	if t.GetProperty("nearest-neighbor") == nil {
		fmt.Println("nearest-neighbor is nil")
		return
	}

	nearestNeighbor := t.GetProperty("nearest-neighbor").(*model.Turtle)
	f.turnAway(t, nearestNeighbor.GetHeading(), f.maxSeperateTurn)
}

func (f *Flocking) align(t *model.Turtle) {
	f.turnTowards(t, f.averageFlockmateHeading(t), f.maxAlignTurn)
}

func (f *Flocking) averageFlockmateHeading(t *model.Turtle) float64 {
	xComponent := 0.0
	yComponent := 0.0
	flockmates := t.GetProperty("flockmates").(*model.TurtleAgentSet)
	flockmates.Ask(func(t2 *model.Turtle) {
		headingRad := t2.GetHeading() * (math.Pi / 180.0)
		xComponent += math.Cos(headingRad)
		yComponent += math.Sin(headingRad)
	})

	if xComponent == 0 && yComponent == 0 {
		return t.GetHeading()
	}

	headingRad := math.Atan2(yComponent, xComponent)
	return headingRad * (180.0 / math.Pi) // convert back to degrees
}

func (f *Flocking) cohere(t *model.Turtle) {
	f.turnTowards(t, f.averageHeadingTowardsFlockmates(t), f.maxCohereTurn)
}

func (f *Flocking) averageHeadingTowardsFlockmates(t *model.Turtle) float64 {
	xComponent := 0.0
	yComponent := 0.0
	flockmates := t.GetProperty("flockmates").(*model.TurtleAgentSet)

	count := float64(flockmates.Count())
	if count == 0 {
		return t.GetHeading()
	}

	flockmates.Ask(func(t2 *model.Turtle) {
		towards := t.TowardsTurtle(t2)
		headingRad := (towards + 180) * (math.Pi / 180.0) // convert to radians

		xComponent += math.Sin(headingRad) // netlogo uses sin for x
		yComponent += math.Cos(headingRad) // and cos for y
	})

	xComponent /= count
	yComponent /= count

	if xComponent == 0 && yComponent == 0 {
		return t.GetHeading()
	}

	headingRad := math.Atan2(yComponent, xComponent) // atan2(y, x) is correct
	return headingRad * (180.0 / math.Pi)            // convert back to degrees
}

// ;;; HELPER PROCEDURES

func (f *Flocking) turnTowards(t *model.Turtle, newHeading float64, maxTurn float64) {

	turn := newHeading - t.GetHeading()
	if turn > 180 {
		turn -= 360
	} else if turn < -180 {
		turn += 360
	}

	f.turnAtMost(t, turn, maxTurn)
}

func (f *Flocking) turnAway(t *model.Turtle, newHeading float64, maxTurn float64) {
	turn := t.GetHeading() - newHeading
	if turn > 180 {
		turn -= 360
	} else if turn < -180 {
		turn += 360
	}
	f.turnAtMost(t, turn, maxTurn)
}

func (f *Flocking) turnAtMost(t *model.Turtle, turn float64, maxTurn float64) {
	if turn > maxTurn {
		if turn > 0 {
			t.Right(maxTurn)
		} else {
			t.Left(maxTurn)
		}
	} else {
		t.Right(turn)
	}
}

func (f *Flocking) Model() *model.Model {
	return f.model
}

func (f *Flocking) Stats() map[string]interface{} {
	return map[string]interface{}{}
}

func (f *Flocking) Stop() bool {
	return false
}

func (f *Flocking) Widgets() []api.Widget {
	return []api.Widget{
		{
			PrettyName:      "Population",
			WidgetType:      "slider",
			WidgetValueType: "int",
			TargetVariable:  "population",
			MinValue:        "1",
			MaxValue:        "500",
			DefaultValue:    "300",
			StepAmount:      "1",
			ValuePointerInt: &f.population,
		},
		{
			PrettyName:        "Vision",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			TargetVariable:    "vision",
			MinValue:          "1",
			MaxValue:          "10",
			DefaultValue:      "5",
			StepAmount:        "0.1",
			ValuePointerFloat: &f.vision,
		},
		{
			PrettyName:        "Minimum Separation",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			TargetVariable:    "minimum-separation",
			MinValue:          "0.1",
			MaxValue:          "5",
			DefaultValue:      "1",
			StepAmount:        "0.1",
			ValuePointerFloat: &f.minimumSeparation,
		},
		{
			PrettyName:        "Max Align Turn",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			TargetVariable:    "max-align-turn",
			MinValue:          "0.1",
			MaxValue:          "10",
			DefaultValue:      "5",
			StepAmount:        "0.1",
			ValuePointerFloat: &f.maxAlignTurn,
		},
		{
			PrettyName:        "Max Cohere Turn",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			TargetVariable:    "max-cohere-turn",
			MinValue:          "0.1",
			MaxValue:          "10",
			DefaultValue:      "5",
			StepAmount:        "0.1",
			ValuePointerFloat: &f.maxCohereTurn,
		},
		{
			PrettyName:        "Max Seperate Turn",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			TargetVariable:    "max-seperate-turn",
			MinValue:          "0.1",
			MaxValue:          "10",
			DefaultValue:      "5",
			StepAmount:        "0.1",
			ValuePointerFloat: &f.maxSeperateTurn,
		},
	}
}
