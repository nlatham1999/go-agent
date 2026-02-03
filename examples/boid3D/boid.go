package boid3D

import (
	"fmt"
	"math"
	"time"

	"github.com/nlatham1999/go-agent/pkg/api"
	"github.com/nlatham1999/go-agent/pkg/model"
)

// enforce that Boid implements the ModelInterface interface
var _ api.ModelInterface = (*Boid3D)(nil)

type Boid3D struct {
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
	avgSpeedGraph   api.GraphWidget
}

func NewBoid3D() *Boid3D {
	return &Boid3D{}
}

func (b *Boid3D) Init() {
	modelSettings := model.ModelSettings{
		TurtleProperties: map[string]interface{}{
			"vx":     .01,
			"vy":     .01,
			"vz":     .01,
			"vx-new": .01,
			"vy-new": .01,
			"vz-new": .01,
		},
		MinPxCor: -5,
		MinPyCor: -5,
		MaxPxCor: 5,
		MaxPyCor: 5,
		MinPzCor: -5,
		MaxPzCor: 5,
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
	b.avgSpeedGraph = api.NewGraphWidget("Average Speed", "avg-speed-graph", "ticks", "speed", []string{}, []string{})

}

func (b *Boid3D) SetUp() error {
	b.model.ClearAll()

	b.model.CreateTurtles(b.numBirds,
		func(t *model.Turtle) {
			t.SetXYZ(b.model.RandomXCor(), b.model.RandomYCor(), b.model.RandomZCor())
			t.SetSize(b.turtleSize)
			t.Shape = "triangle"
		},
	)

	b.avgSpeedGraph.XValues = []string{}
	b.avgSpeedGraph.YValues = []string{}

	b.model.ResetTicks()
	return nil
}

func (b *Boid3D) Go() {

	timeNow := time.Now()

	// remove all links - collect them first to avoid modifying during iteration
	links := b.model.Links().List()
	for _, l := range links {
		b.model.KillLink(l)
	}

	// Initialize vx-new and vy-new with current values
	b.model.Turtles().Ask(
		func(t *model.Turtle) {
			t.SetProperty("vx-new", t.GetProperty("vx"))
			t.SetProperty("vy-new", t.GetProperty("vy"))
			t.SetProperty("vz-new", t.GetProperty("vz"))
		},
	)

	// Compute all forces (reads vx/vy, writes vx-new/vy-new)
	b.model.Turtles().Ask(
		func(t *model.Turtle) {
			b.computeSeperation(t)
			b.computeAlignment(t)
			b.computeCohesion(t)
			b.computeTurnAwayFromEdges(t)
			b.computeLimitSpeeds(t)
		},
	)

	// Apply velocities (copies vx-new/vy-new to vx/vy)
	b.model.Turtles().Ask(
		func(t *model.Turtle) {
			b.applyVelocities(t)
		},
	)

	// Update positions
	b.model.Turtles().Ask(
		func(t *model.Turtle) {
			b.updatePosition(t)
		},
	)

	// Calculate average speed for graph
	totalSpeed := 0.0
	b.model.Turtles().Ask(
		func(t *model.Turtle) {
			vx := t.GetProperty("vx").(float64)
			vy := t.GetProperty("vy").(float64)
			vz := t.GetProperty("vz").(float64)
			speed := math.Sqrt(vx*vx + vy*vy + vz*vz)
			totalSpeed += speed
		},
	)
	avgSpeed := totalSpeed / float64(b.model.Turtles().Count())

	b.avgSpeedGraph.XValues = append(b.avgSpeedGraph.XValues, fmt.Sprintf("%d", b.model.Ticks))
	b.avgSpeedGraph.YValues = append(b.avgSpeedGraph.YValues, fmt.Sprintf("%.4f", avgSpeed))

	b.model.Tick()

	fmt.Println("Time taken: ", time.Since(timeNow))
	fmt.Println("Number of turtles: ", b.model.Turtles().Count())
}

// seperation stage - reads from vx/vy, writes to vx-new/vy-new
func (b *Boid3D) computeSeperation(t *model.Turtle) {
	closeDx := 0.0
	closeDy := 0.0
	closeDz := 0.0
	b.model.TurtlesInRadiusXYZ(t.XCor(), t.YCor(), t.ZCor(), b.protectedRange).Ask(
		func(t2 *model.Turtle) {
			if t != t2 {
				closeDx += t.XCor() - t2.XCor()
				closeDy += t.YCor() - t2.YCor()
				closeDz += t.ZCor() - t2.ZCor()
			}
		},
	)
	avoidFactor := b.avoidFactor
	vxNew := t.GetProperty("vx-new").(float64)
	vyNew := t.GetProperty("vy-new").(float64)
	vzNew := t.GetProperty("vz-new").(float64)
	t.SetProperty("vx-new", vxNew+(closeDx*avoidFactor))
	t.SetProperty("vy-new", vyNew+(closeDy*avoidFactor))
	t.SetProperty("vz-new", vzNew+(closeDz*avoidFactor))
}

func (b *Boid3D) computeAlignment(t *model.Turtle) {
	xvelAvg := 0.0
	yvelAvg := 0.0
	zvelAvg := 0.0
	neighboringBoids := 0
	b.model.TurtlesInRadiusXYZ(t.XCor(), t.YCor(), t.ZCor(), b.visibleRange).Ask(
		func(t2 *model.Turtle) {
			if t != t2 {
				xvelAvg += t2.GetProperty("vx").(float64)
				yvelAvg += t2.GetProperty("vy").(float64)
				zvelAvg += t2.GetProperty("vz").(float64)
				neighboringBoids++

				// Create link to visualize connections
				t.CreateLinkWithTurtle(nil, t2, func(l *model.Link) {
					l.Color = b.model.RandomColor()
					l.Show()
				})
			}
		},
	)
	if neighboringBoids > 0 {
		xvelAvg /= float64(neighboringBoids)
		yvelAvg /= float64(neighboringBoids)
		zvelAvg /= float64(neighboringBoids)
		vx := t.GetProperty("vx").(float64)
		vy := t.GetProperty("vy").(float64)
		vz := t.GetProperty("vz").(float64)
		matchingFactor := b.matchingFactor
		vxNew := t.GetProperty("vx-new").(float64)
		vyNew := t.GetProperty("vy-new").(float64)
		vzNew := t.GetProperty("vz-new").(float64)
		t.SetProperty("vx-new", vxNew+(xvelAvg-vx)*matchingFactor)
		t.SetProperty("vy-new", vyNew+(yvelAvg-vy)*matchingFactor)
		t.SetProperty("vz-new", vzNew+(zvelAvg-vz)*matchingFactor)
	}
}

func (b *Boid3D) computeCohesion(t *model.Turtle) {
	xposAvg := 0.0
	yposAvg := 0.0
	zposAvg := 0.0
	neighboringBoids := 0
	b.model.TurtlesInRadiusXYZ(t.XCor(), t.YCor(), t.ZCor(), b.visibleRange).Ask(
		func(t2 *model.Turtle) {
			if t != t2 {
				xposAvg += t2.XCor()
				yposAvg += t2.YCor()
				zposAvg += t2.ZCor()
				neighboringBoids++
			}
		},
	)
	if neighboringBoids > 0 {
		xposAvg /= float64(neighboringBoids)
		yposAvg /= float64(neighboringBoids)
		zposAvg /= float64(neighboringBoids)
		centeringFactor := b.centeringFactor
		vxNew := t.GetProperty("vx-new").(float64)
		vyNew := t.GetProperty("vy-new").(float64)
		vzNew := t.GetProperty("vz-new").(float64)
		t.SetProperty("vx-new", vxNew+(xposAvg-t.XCor())*centeringFactor)
		t.SetProperty("vy-new", vyNew+(yposAvg-t.YCor())*centeringFactor)
		t.SetProperty("vz-new", vzNew+(zposAvg-t.ZCor())*centeringFactor)
	}
}

func (b *Boid3D) computeTurnAwayFromEdges(t *model.Turtle) {
	margin := b.margin
	leftMargin := b.model.MinXCor() + margin
	rightMargin := b.model.MaxXCor() - margin
	topMargin := b.model.MinYCor() + margin
	bottomMargin := b.model.MaxXCor() - margin
	turnFactor := b.turnFactor
	vxNew := t.GetProperty("vx-new").(float64)
	vyNew := t.GetProperty("vy-new").(float64)
	vzNew := t.GetProperty("vz-new").(float64)
	if t.XCor() < leftMargin {
		vxNew += turnFactor
	}
	if t.XCor() > rightMargin {
		vxNew -= turnFactor
	}
	if t.YCor() > bottomMargin {
		vyNew -= turnFactor
	}
	if t.YCor() < topMargin {
		vyNew += turnFactor
	}
	if t.ZCor() < leftMargin {
		vzNew += turnFactor
	}
	if t.ZCor() > rightMargin {
		vzNew -= turnFactor
	}
	t.SetProperty("vx-new", vxNew)
	t.SetProperty("vy-new", vyNew)
	t.SetProperty("vz-new", vzNew)
}

func (b *Boid3D) computeLimitSpeeds(t *model.Turtle) {
	vxNew := t.GetProperty("vx-new").(float64)
	vyNew := t.GetProperty("vy-new").(float64)
	vzNew := t.GetProperty("vz-new").(float64)
	speed := math.Sqrt(vxNew*vxNew + vyNew*vyNew + vzNew*vzNew)
	minSpeed := b.minSpeed
	maxSpeed := b.maxSpeed
	if speed > maxSpeed {
		t.SetProperty("vx-new", (vxNew/speed)*maxSpeed)
		t.SetProperty("vy-new", (vyNew/speed)*maxSpeed)
		t.SetProperty("vz-new", (vzNew/speed)*maxSpeed)
	}
	if speed < minSpeed {
		t.SetProperty("vx-new", (vxNew/speed)*minSpeed)
		t.SetProperty("vy-new", (vyNew/speed)*minSpeed)
		t.SetProperty("vz-new", (vzNew/speed)*minSpeed)
	}
}

// apply velocities - copies vx-new/vy-new to vx/vy
func (b *Boid3D) applyVelocities(t *model.Turtle) {
	t.SetProperty("vx", t.GetProperty("vx-new"))
	t.SetProperty("vy", t.GetProperty("vy-new"))
	t.SetProperty("vz", t.GetProperty("vz-new"))
}

func (b *Boid3D) updatePosition(t *model.Turtle) {
	vx := t.GetProperty("vx").(float64)
	vy := t.GetProperty("vy").(float64)
	vz := t.GetProperty("vz").(float64)
	t.SetXYZ(t.XCor()+vx, t.YCor()+vy, t.ZCor()+vz)

	// Update heading to point in direction of movement
	// atan2 returns radians, convert to degrees
	heading := math.Atan2(vy, vx) * (180.0 / math.Pi)
	t.SetHeading(heading)
}

func (b *Boid3D) Model() *model.Model {
	return b.model
}

func (b *Boid3D) Stats() map[string]interface{} {
	return map[string]interface{}{
		"avg-speed-graph": b.avgSpeedGraph,
	}
}

func (b *Boid3D) Stop() bool {
	return false
}

func (b *Boid3D) Widgets() []api.Widget {
	return []api.Widget{
		{
			PrettyName:        "Turtle Size",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			Id:                "turtle-size",
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
			Id:              "num-birds",
			MinValue:        "1",
			MaxValue:        "5000",
			StepAmount:      "1",
			DefaultValue:    "100",
			ValuePointerInt: &b.numBirds,
		},
		{
			PrettyName:        "Turn Factor",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			Id:                "turn-factor",
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
			Id:                "visible-range",
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
			Id:                "protected-range",
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
			Id:                "avoid-factor",
			MinValue:          "0.01",
			MaxValue:          "1",
			StepAmount:        "0.01",
			DefaultValue:      "0.1",
			ValuePointerFloat: &b.avoidFactor,
		},
		{
			PrettyName:        "Matching Factor",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			Id:                "matching-factor",
			MinValue:          "0.01",
			MaxValue:          "1",
			StepAmount:        "0.01",
			DefaultValue:      "0.1",
			ValuePointerFloat: &b.matchingFactor,
		},
		{
			PrettyName:        "Centering Factor",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			Id:                "centering-factor",
			MinValue:          "0.0001",
			MaxValue:          "0.01",
			StepAmount:        "0.0001",
			DefaultValue:      "0.005",
			ValuePointerFloat: &b.centeringFactor,
		},
		{
			PrettyName:        "Margin",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			Id:                "margin",
			MinValue:          "0.01",
			MaxValue:          ".2",
			StepAmount:        "0.01",
			DefaultValue:      "0.2",
			ValuePointerFloat: &b.margin,
		},
		{
			PrettyName:        "Min Speed",
			WidgetType:        "slider",
			WidgetValueType:   "float",
			Id:                "min-speed",
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
			Id:                "max-speed",
			MinValue:          "0.01",
			MaxValue:          "2",
			StepAmount:        "0.01",
			DefaultValue:      ".06",
			ValuePointerFloat: &b.maxSpeed,
		},
	}
}
