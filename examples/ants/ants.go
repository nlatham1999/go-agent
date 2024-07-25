package ants

import (
	"math/rand"
	"time"

	"github.com/nlatham1999/go-agent/internal/model"
	"github.com/nlatham1999/go-agent/internal/slider"
)

var (
	environment *model.Model

	sliders map[string]*slider.Slider
)

const (
	//patches own
	chemical         = "chemical"
	food             = "food"
	nest             = "nest"
	nestScent        = "nestScent"
	foodSourceNumber = "foodSourceNumber"

	//widgets
	population      = "population"
	diffusionRate   = "diffusionRate"
	evaporationRate = "evaporationRate"
)

func Init() {

	rand.Seed(time.Now().UnixNano())

	patchesOwn := map[string]interface{}{
		chemical:         0.0,
		food:             0.0,
		nest:             false,
		nestScent:        0,
		foodSourceNumber: 0,
	}

	environment = model.NewModel(patchesOwn, nil, nil, nil, nil, nil, false, false)

	sliders = map[string]*slider.Slider{
		population:      slider.NewSlider(0, 1, 200, 125),
		diffusionRate:   slider.NewSlider(0, 1, 99, 50),
		evaporationRate: slider.NewSlider(0, 1, 99, 10),
	}

}

func setup() {
	environment.ClearAll()
	environment.SetDefaultShapeTurtles("bug")
	environment.CreateTurtles(int(sliders[population].GetValue()),
		"",
		[]model.TurtleOperation{
			// model.SetColor(environment.ColorHueMap["red"]),
			// model.SetSize(2),
		},
	)
	setupPatches()
	environment.ResetTicks()
}

func setupPatches() {
	model.AskPatches(
		environment.Patches,
		[]model.PatchOperation{
			setupNest,
			setupFood,
			setupNest,
		},
	)
}

func setupNest(p *model.Patch) {
	p.PatchesOwn[nest] = p.DistanceXY(0, 0) < 5
	p.PatchesOwn[nestScent] = 200 - p.DistanceXY(0, 0)
}

func setupFood(p *model.Patch) {
	// ;; setup food source one on the right
	if p.DistanceXY(.6*float64(environment.MaxPxCor), 0) < 5 {
		p.PatchesOwn[foodSourceNumber] = 1
	}

	// ;; setup food source two on the lower-left
	if p.DistanceXY(-.6*float64(environment.MaxPxCor), -.6*float64(environment.MaxPxCor)) < 5 {
		p.PatchesOwn[foodSourceNumber] = 2
	}

	// ;; setup food source three on the upper-left
	if p.DistanceXY(-.8*float64(environment.MaxPxCor), .8*float64(environment.MaxPxCor)) < 5 {
		p.PatchesOwn[foodSourceNumber] = 3
	}

	// ;; set "food" at sources to either 1 or 2, randomly
	if p.PatchesOwn[foodSourceNumber].(int) > 0 {
		p.PatchesOwn[food] = environment.OneOfInt([]int{1, 2})
	}
}

func recolorPatch(p *model.Patch) {

	// ;; give color to nest and food sources
	if p.PatchesOwn[nest].(bool) {
		p.PColor.SetColor(model.Violet)
	} else {
		if p.PatchesOwn[food].(int) > 0 {
			if p.PatchesOwn[foodSourceNumber].(int) == 1 {
				p.PColor.SetColor(model.Cyan)
			} else if p.PatchesOwn[foodSourceNumber].(int) == 2 {
				p.PColor.SetColor(model.Sky)
			} else if p.PatchesOwn[foodSourceNumber].(int) == 3 {
				p.PColor.SetColor(model.Blue)
			}
		} else {
			// p.SetColorAndScale(p.PatchesOwn[chemical].(float64), .1, 5)
		}
	}
}

func run() {

	// ask turtles
	// [ if who >= ticks [ stop ] ;; delay initial departure
	//   ifelse color = red
	//   [ look-for-food  ]       ;; not carrying food? look for it
	//   [ return-to-nest ]       ;; carrying food? take it back to nest
	//   wiggle
	//   fd 1 ]
	// diffuse chemical (diffusion-rate / 100)
	// ask patches
	// [ set chemical chemical * (100 - evaporation-rate) / 100  ;; slowly evaporate chemical
	//   recolor-patch ]
	// tick

	model.AskTurtles(
		environment.Turtles,
		[]model.TurtleOperation{
			func(t *model.Turtle) {
				if t.Who() >= environment.Ticks {
					return
				}
				if t.Color == model.Red {
					lookForFood(t)
				} else {
					returnToNest(t)
				}
				wiggle(t)
				t.Forward(1)
			},
		},
	)
	environment.Diffuse(chemical, sliders[diffusionRate].GetValue()/100)
	model.AskPatches(
		environment.Patches,
		[]model.PatchOperation{
			func(p *model.Patch) {
				p.PatchesOwn[chemical] = p.PatchesOwn[chemical].(float64) * (100 - sliders[evaporationRate].GetValue()) / 100
				recolorPatch(p)
			},
		},
	)
	environment.Tick()

}

func returnToNest(t *model.Turtle) {
	if t.PatchHere().PatchesOwn[nest].(bool) {
		t.Color.SetColor(model.Red)
		t.Right(180)
	} else {
		t.PatchHere().PatchesOwn[chemical] = t.PatchHere().PatchesOwn[chemical].(int) + 60
		uphillNestScent(t)
	}
}

func lookForFood(t *model.Turtle) {
	p := t.PatchHere()
	if p.PatchesOwn[food].(int) > 0 {
		t.Color.SetColor(model.Orange)
		p.PatchesOwn[food] = p.PatchesOwn[food].(int) - 1
		t.Right(180)
		return
	}

	if p.PatchesOwn[chemical].(float64) >= .05 && p.PatchesOwn[chemical].(float64) < 2 {
		uphillChemical(t)
	}

}

func uphillChemical(t *model.Turtle) {
	scentAhead := chemicalScentAtAngle(t, 0)
	scentRight := chemicalScentAtAngle(t, 45)
	scentLeft := chemicalScentAtAngle(t, -45)
	if scentRight > scentAhead || scentLeft > scentAhead {
		if scentRight > scentLeft {
			t.Right(45)
		} else {
			t.Left(45)
		}
	}
}

func wiggle(t *model.Turtle) {
	t.Right(float64(environment.RandomAmount(40)))
	t.Left(float64(environment.RandomAmount(40)))

	if !t.CanMove(1) {
		t.Right(180)
	}
}

func uphillNestScent(t *model.Turtle) {
	scentAhead := nestScentAtAngle(t, 0)
	scentRight := nestScentAtAngle(t, 45)
	scentLeft := nestScentAtAngle(t, -45)
	if scentRight > scentAhead || scentLeft > scentAhead {
		if scentRight > scentLeft {
			t.Right(45)
		} else {
			t.Left(45)
		}
	}
}

func nestScentAtAngle(t *model.Turtle, angle float64) int {
	p := t.PatchRightAndAhead(angle, 1)
	if p == nil {
		return 0
	} else {
		return p.PatchesOwn[nestScent].(int)
	}
}

func chemicalScentAtAngle(t *model.Turtle, angle float64) float64 {
	p := t.PatchRightAndAhead(angle, 1)
	if p == nil {
		return 0
	} else {
		return p.PatchesOwn[chemical].(float64)
	}
}
