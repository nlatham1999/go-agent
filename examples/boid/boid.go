package boid

import (
	"math"

	"github.com/nlatham1999/go-agent/pkg/api"
	"github.com/nlatham1999/go-agent/pkg/model"
)

// enforce that Boid implements the ModelInterface interface
var _ api.ModelInterface = (*Boid)(nil)

type Boid struct {
	model *model.Model

	numBirds        int
	turnFactor      float64
	visibleRange    float64
	protectedRange  float64
	centeringFactor float64
	avoidFactor     float64
	matchingFactor  float64
	margin          float64
	minSpeed        float64
	maxSpeed        float64
	turtleSize      float64
}

func NewBoid() *Boid {
	return &Boid{}
}

func (b *Boid) Init() {
	modelSettings := model.ModelSettings{
		TurtleProperties: map[string]interface{}{
			"vx": .01,
			"vy": .01,
		},
		MinPxCor: -5,
		MinPyCor: -5,
		MaxPxCor: 5,
		MaxPyCor: 5,
	}

	b.model = model.NewModel(modelSettings)

	b.numBirds = 100
	b.turnFactor = .02
	b.visibleRange = .5
	b.protectedRange = .1
	b.centeringFactor = 0.005
	b.avoidFactor = .1
	b.matchingFactor = .1
	b.margin = .2
	b.minSpeed = 0.03
	b.maxSpeed = 0.06
	b.turtleSize = 0.16

}

func (b *Boid) SetUp() error {
	b.model.ClearAll()

	b.model.CreateTurtles(b.numBirds,
		func(t *model.Turtle) {
			t.SetXY(b.model.RandomXCor(), b.model.RandomYCor())
			t.SetSize(b.turtleSize)
			t.Shape = "triangle"
		},
	)

	b.model.ResetTicks()
	return nil
}

func (b *Boid) Go() {
	b.model.Turtles().Ask(
		func(t *model.Turtle) {
			b.seperation(t)
			b.alignment(t)
			b.cohesion(t)
			b.turnAwayFromEdges(t)
			b.limitSpeeds(t)
		},
	)

	b.model.Turtles().Ask(
		func(t *model.Turtle) {
			b.updatePosition(t)
		},
	)

	b.model.Tick()
}

// seperation stage
func (b *Boid) seperation(t *model.Turtle) {
	closeDx := 0.0
	closeDy := 0.0
	b.model.TurtlesInRadius(t.XCor(), t.YCor(), b.protectedRange).Ask(
		func(t2 *model.Turtle) {
			if t != t2 {
				closeDx += t.XCor() - t2.XCor()
				closeDy += t.YCor() - t2.YCor()
			}
		},
	)
	avoidFactor := b.avoidFactor
	vx := t.GetProperty("vx").(float64)
	vy := t.GetProperty("vy").(float64)
	t.SetProperty("vx", vx+(closeDx*avoidFactor))
	t.SetProperty("vy", vy+(closeDy*avoidFactor))
}

func (b *Boid) alignment(t *model.Turtle) {
	xvelAvg := 0.0
	yvelAvg := 0.0
	neighboringBoids := 0
	b.model.TurtlesInRadius(t.XCor(), t.YCor(), b.visibleRange).Ask(
		func(t2 *model.Turtle) {
			if t != t2 {
				xvelAvg += t2.GetProperty("vx").(float64)
				yvelAvg += t2.GetProperty("vy").(float64)
				neighboringBoids++
			}
		},
	)
	if neighboringBoids > 0 {
		xvelAvg /= float64(neighboringBoids)
		yvelAvg /= float64(neighboringBoids)
		vx := t.GetProperty("vx").(float64)
		vy := t.GetProperty("vy").(float64)
		matchingFactor := b.matchingFactor
		t.SetProperty("vx", vx+(xvelAvg-vx)*matchingFactor)
		t.SetProperty("vy", vy+(yvelAvg-vy)*matchingFactor)
	}
}

func (b *Boid) cohesion(t *model.Turtle) {
	xposAvg := 0.0
	yposAvg := 0.0
	neighboringBoids := 0
	b.model.TurtlesInRadius(t.XCor(), t.YCor(), b.visibleRange).Ask(
		func(t2 *model.Turtle) {
			if t != t2 {
				xposAvg += t2.XCor()
				yposAvg += t2.YCor()
				neighboringBoids++
			}
		},
	)
	if neighboringBoids > 0 {
		xposAvg /= float64(neighboringBoids)
		yposAvg /= float64(neighboringBoids)
		vx := t.GetProperty("vx").(float64)
		vy := t.GetProperty("vy").(float64)
		centeringFactor := b.centeringFactor
		t.SetProperty("vx", vx+(xposAvg-t.XCor())*centeringFactor)
		t.SetProperty("vy", vy+(yposAvg-t.YCor())*centeringFactor)
	}
}

func (b *Boid) turnAwayFromEdges(t *model.Turtle) {
	margin := b.margin
	leftMargin := b.model.MinXCor() + margin
	rightMargin := b.model.MaxXCor() - margin
	topMargin := b.model.MinYCor() + margin
	bottomMargin := b.model.MaxXCor() - margin
	turnFactor := b.turnFactor
	vx := t.GetProperty("vx").(float64)
	vy := t.GetProperty("vy").(float64)
	if t.XCor() < leftMargin {
		t.SetProperty("vx", vx+turnFactor)
	}
	if t.XCor() > rightMargin {
		t.SetProperty("vx", vx-turnFactor)
	}
	if t.YCor() > bottomMargin {
		t.SetProperty("vy", vy-turnFactor)
	}
	if t.YCor() < topMargin {
		t.SetProperty("vy", vy+turnFactor)
	}
}

func (b *Boid) limitSpeeds(t *model.Turtle) {
	vx := t.GetProperty("vx").(float64)
	vy := t.GetProperty("vy").(float64)
	speed := math.Sqrt(vx*vx + vy*vy)
	minSpeed := b.minSpeed
	maxSpeed := b.maxSpeed
	if speed > maxSpeed {
		t.SetProperty("vx", (vx/speed)*maxSpeed)
		t.SetProperty("vy", (vy/speed)*maxSpeed)
	}
	if speed < minSpeed {
		t.SetProperty("vx", (vx/speed)*minSpeed)
		t.SetProperty("vy", (vy/speed)*minSpeed)
	}

}

func (b *Boid) updatePosition(t *model.Turtle) {
	vx := t.GetProperty("vx").(float64)
	vy := t.GetProperty("vy").(float64)
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
			PrettyName:        "Turtle Size",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			TargetVariable:    "turtle-size",
			MinValue:          "0.01",
			MaxValue:          "1",
			StepAmount:        "0.01",
			DefaultValue:      "0.16",
			ValuePointerFloat: &b.turtleSize,
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
			ValuePointerInt: &b.numBirds,
		},
		{
			PrettyName:        "Turn Factor",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			TargetVariable:    "turn-factor",
			MinValue:          "0.01",
			MaxValue:          "10",
			StepAmount:        "0.01",
			DefaultValue:      ".02",
			ValuePointerFloat: &b.turnFactor,
		},
		{
			PrettyName:        "Visible Range",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			TargetVariable:    "visible-range",
			MinValue:          "0.01",
			MaxValue:          "2",
			StepAmount:        "0.01",
			DefaultValue:      ".5",
			ValuePointerFloat: &b.visibleRange,
		},
		{
			PrettyName:        "Protected Range",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			TargetVariable:    "protected-range",
			MinValue:          "0.01",
			MaxValue:          "2",
			StepAmount:        "0.01",
			DefaultValue:      ".1",
			ValuePointerFloat: &b.protectedRange,
		},
		{
			PrettyName:        "Avoid Factor",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			TargetVariable:    "avoid-factor",
			MinValue:          "0.01",
			MaxValue:          "1",
			StepAmount:        "0.01",
			DefaultValue:      "0.05",
			ValuePointerFloat: &b.avoidFactor,
		},
		{
			PrettyName:        "Matching Factor",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			TargetVariable:    "matching-factor",
			MinValue:          "0.01",
			MaxValue:          "1",
			StepAmount:        "0.01",
			DefaultValue:      "0.05",
			ValuePointerFloat: &b.matchingFactor,
		},
		{
			PrettyName:        "Centering Factor",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			TargetVariable:    "centering-factor",
			MinValue:          "0.0001",
			MaxValue:          "0.01",
			StepAmount:        "0.0001",
			DefaultValue:      "0.0005",
			ValuePointerFloat: &b.centeringFactor,
		},
		{
			PrettyName:        "Margin",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			TargetVariable:    "margin",
			MinValue:          "0.01",
			MaxValue:          ".2",
			StepAmount:        "0.01",
			DefaultValue:      ".01",
			ValuePointerFloat: &b.margin,
		},
		{
			PrettyName:        "Min Speed",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			TargetVariable:    "min-speed",
			MinValue:          "0.01",
			MaxValue:          "2",
			StepAmount:        "0.01",
			DefaultValue:      ".03",
			ValuePointerFloat: &b.minSpeed,
		},
		{
			PrettyName:        "Max Speed",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			TargetVariable:    "max-speed",
			MinValue:          "0.01",
			MaxValue:          "2",
			StepAmount:        "0.01",
			DefaultValue:      ".06",
			ValuePointerFloat: &b.maxSpeed,
		},
	}
}
