package boid

import (
	"math"

	"github.com/nlatham1999/go-agent/api"
	"github.com/nlatham1999/go-agent/model"
)

// implement
// type ModelInterface interface {
// 	Init()        // runs at the very beginning
// 	SetUp() error // sets up the model
// 	Go()          // runs the model

// 	Model() *model.Model           // returns the model
// 	Stats() map[string]interface{} //returns the stats of the model
// 	Stop() bool                    // on whether to stop the model
// 	Widgets() []Widget             // returns the widgets of the model
// }

type Boid struct {
	model *model.Model
}

func NewBoid() *Boid {
	return &Boid{}
}

func (b *Boid) Init() {
	modelSettings := model.ModelSettings{
		Globals: map[string]interface{}{
			"num-birds":        100,
			"turn-factor":      .02,
			"visible-range":    .5,
			"protected-range":  .1,
			"centering-factor": 0.005,
			"avoid-factor":     .1,
			"matching-factor":  .1,
			"margin":           .2,
			"min-speed":        0.03,
			"max-speed":        0.06,
			"turtle-size":      0.03,
		},
		TurtlesOwn: map[string]interface{}{
			"vx": .01,
			"vy": .01,
		},
		MinPxCor: 0,
		MinPyCor: 0,
		MaxPxCor: 1,
		MaxPyCor: 1,
	}

	b.model = model.NewModel(modelSettings)
}

func (b *Boid) SetUp() error {
	b.model.ClearAll()

	numBirds := b.model.GetGlobal("num-birds").(int)
	b.model.CreateTurtles(numBirds, "",
		func(t *model.Turtle) {
			t.SetXY(b.model.RandomXCor(), b.model.RandomYCor())
			t.SetSize(b.model.GetGlobal("turtle-size").(float64))
		},
	)

	b.model.ResetTicks()
	return nil
}

func (b *Boid) Go() {
	b.model.Turtles("").Ask(
		func(t *model.Turtle) {
			b.seperation(t)
			b.alignment(t)
			b.cohesion(t)
			b.turnAwayFromEdges(t)
			b.limitSpeeds(t)
		},
	)

	b.model.Turtles("").Ask(
		func(t *model.Turtle) {
			b.updatePosition(t)
		},
	)
}

// seperation stage
func (b *Boid) seperation(t *model.Turtle) {
	closeDx := 0.0
	closeDy := 0.0
	b.model.Turtles("").Ask(
		func(t2 *model.Turtle) {
			if t != t2 {
				if t.DistanceTurtle(t2) < b.model.GetGlobal("protected-range").(float64) {
					closeDx += t.XCor() - t2.XCor()
					closeDy += t.YCor() - t2.YCor()
				}
			}
		},
	)
	avoidFactor := b.model.GetGlobal("avoid-factor").(float64)
	vx := t.GetOwn("vx").(float64)
	vy := t.GetOwn("vy").(float64)
	t.SetOwn("vx", vx+(closeDx*avoidFactor))
	t.SetOwn("vy", vy+(closeDy*avoidFactor))
}

func (b *Boid) alignment(t *model.Turtle) {
	xvelAvg := 0.0
	yvelAvg := 0.0
	neighboringBoids := 0
	b.model.Turtles("").Ask(
		func(t2 *model.Turtle) {
			if t != t2 {
				if t.DistanceTurtle(t2) < b.model.GetGlobal("visible-range").(float64) {
					xvelAvg += t2.GetOwn("vx").(float64)
					yvelAvg += t2.GetOwn("vy").(float64)
					neighboringBoids++
				}
			}
		},
	)
	if neighboringBoids > 0 {
		xvelAvg /= float64(neighboringBoids)
		yvelAvg /= float64(neighboringBoids)
		vx := t.GetOwn("vx").(float64)
		vy := t.GetOwn("vy").(float64)
		matchingFactor := b.model.GetGlobal("matching-factor").(float64)
		t.SetOwn("vx", vx+(xvelAvg-vx)*matchingFactor)
		t.SetOwn("vy", vy+(yvelAvg-vy)*matchingFactor)
	}
}

func (b *Boid) cohesion(t *model.Turtle) {
	xposAvg := 0.0
	yposAvg := 0.0
	neighboringBoids := 0
	b.model.Turtles("").Ask(
		func(t2 *model.Turtle) {
			if t != t2 {
				if t.DistanceTurtle(t2) < b.model.GetGlobal("visible-range").(float64) {
					xposAvg += t2.XCor()
					yposAvg += t2.YCor()
					neighboringBoids++
				}
			}
		},
	)
	if neighboringBoids > 0 {
		xposAvg /= float64(neighboringBoids)
		yposAvg /= float64(neighboringBoids)
		vx := t.GetOwn("vx").(float64)
		vy := t.GetOwn("vy").(float64)
		centeringFactor := b.model.GetGlobal("centering-factor").(float64)
		t.SetOwn("vx", vx+(xposAvg-t.XCor())*centeringFactor)
		t.SetOwn("vy", vy+(yposAvg-t.YCor())*centeringFactor)
	}
}

func (b *Boid) turnAwayFromEdges(t *model.Turtle) {
	margin := b.model.GetGlobal("margin").(float64)
	leftMargin := b.model.MinXCor() + margin
	rightMargin := b.model.MaxXCor() - margin
	topMargin := b.model.MinYCor() + margin
	bottomMargin := b.model.MaxXCor() - margin
	turnFactor := b.model.GetGlobal("turn-factor").(float64)
	vx := t.GetOwn("vx").(float64)
	vy := t.GetOwn("vy").(float64)
	if t.XCor() < leftMargin {
		t.SetOwn("vx", vx+turnFactor)
	}
	if t.XCor() > rightMargin {
		t.SetOwn("vx", vx-turnFactor)
	}
	if t.YCor() > bottomMargin {
		t.SetOwn("vy", vy-turnFactor)
	}
	if t.YCor() < topMargin {
		t.SetOwn("vy", vy+turnFactor)
	}
}

func (b *Boid) limitSpeeds(t *model.Turtle) {
	vx := t.GetOwn("vx").(float64)
	vy := t.GetOwn("vy").(float64)
	speed := math.Sqrt(vx*vx + vy*vy)
	minSpeed := b.model.GetGlobal("min-speed").(float64)
	maxSpeed := b.model.GetGlobal("max-speed").(float64)
	if speed > maxSpeed {
		t.SetOwn("vx", (vx/speed)*maxSpeed)
		t.SetOwn("vy", (vy/speed)*maxSpeed)
	}
	if speed < minSpeed {
		t.SetOwn("vx", (vx/speed)*minSpeed)
		t.SetOwn("vy", (vy/speed)*minSpeed)
	}

}

func (b *Boid) updatePosition(t *model.Turtle) {
	vx := t.GetOwn("vx").(float64)
	vy := t.GetOwn("vy").(float64)
	t.SetXY(t.XCor()+vx, t.YCor()+vy)
}

func (b *Boid) Model() *model.Model {
	return b.model
}

func (b *Boid) Stats() map[string]interface{} {
	return nil
}

func (b *Boid) Stop() bool {
	return false
}

func (b *Boid) Widgets() []api.Widget {
	return []api.Widget{
		{
			PrettyName:      "Turtle Size",
			WidgetType:      "slider",
			WidgetValueType: "float",
			TargetVariable:  "turtle-size",
			MinValue:        "0.03",
			MaxValue:        "1",
			StepAmount:      "0.01",
			DefaultValue:    "0.9",
		},
		{
			PrettyName:      "Number of Birds",
			WidgetType:      "slider",
			WidgetValueType: "int",
			TargetVariable:  "num-birds",
			MinValue:        "1",
			MaxValue:        "1000",
			StepAmount:      "1",
			DefaultValue:    "100",
		},
		{
			PrettyName:      "Turn Factor",
			WidgetType:      "slider",
			WidgetValueType: "float",
			TargetVariable:  "turn-factor",
			MinValue:        "0.01",
			MaxValue:        "10",
			StepAmount:      "0.01",
			DefaultValue:    ".02",
		},
		{
			PrettyName:      "Visible Range",
			WidgetType:      "slider",
			WidgetValueType: "float",
			TargetVariable:  "visible-range",
			MinValue:        "0.01",
			MaxValue:        "2",
			StepAmount:      "0.01",
			DefaultValue:    ".5",
		},
		{
			PrettyName:      "Protected Range",
			WidgetType:      "slider",
			WidgetValueType: "float",
			TargetVariable:  "protected-range",
			MinValue:        "0.01",
			MaxValue:        "2",
			StepAmount:      "0.01",
			DefaultValue:    ".1",
		},
		{
			PrettyName:      "Avoid Factor",
			WidgetType:      "slider",
			WidgetValueType: "float",
			TargetVariable:  "avoid-factor",
			MinValue:        "0.01",
			MaxValue:        "1",
			StepAmount:      "0.01",
			DefaultValue:    "0.05",
		},
		{
			PrettyName:      "Matching Factor",
			WidgetType:      "slider",
			WidgetValueType: "float",
			TargetVariable:  "matching-factor",
			MinValue:        "0.01",
			MaxValue:        "1",
			StepAmount:      "0.01",
			DefaultValue:    "0.05",
		},
		{
			PrettyName:      "Centering Factor",
			WidgetType:      "slider",
			WidgetValueType: "float",
			TargetVariable:  "centering-factor",
			MinValue:        "0.0001",
			MaxValue:        "0.01",
			StepAmount:      "0.0001",
			DefaultValue:    "0.0005",
		},
		{
			PrettyName:      "Margin",
			WidgetType:      "slider",
			WidgetValueType: "float",
			TargetVariable:  "margin",
			MinValue:        "0.01",
			MaxValue:        ".2",
			StepAmount:      "0.01",
			DefaultValue:    ".01",
		},
		{
			PrettyName:      "Min Speed",
			WidgetType:      "slider",
			WidgetValueType: "float",
			TargetVariable:  "min-speed",
			MinValue:        "0.01",
			MaxValue:        "2",
			StepAmount:      "0.01",
			DefaultValue:    ".03",
		},
		{
			PrettyName:      "Max Speed",
			WidgetType:      "slider",
			WidgetValueType: "float",
			TargetVariable:  "max-speed",
			MinValue:        "0.01",
			MaxValue:        "2",
			StepAmount:      "0.01",
			DefaultValue:    ".06",
		},
	}
}
