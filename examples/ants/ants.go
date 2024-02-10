package ants

import (
	"math/rand"
	"time"

	"github.com/nlatham1999/go-agent/internal/universe"
)

var (
	environment *universe.Universe

	//@TODO figure out how to handle widgets
	population int
)

//we declare the patches own variable keys as constants
const (
	chemical         = "chemical"
	food             = "food"
	nest             = "nest"
	nestScent        = "nestScent"
	foodSourceNumber = "foodSourceNumber"
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

	environment = universe.NewUniverse(patchesOwn)

}

func setup() {
	environment.ClearAll()
	environment.SetDefaultShapeTurtles("bug")
	environment.CreateTurtles(population,
		[]universe.TurtleOperation{
			universe.SetColor(environment.ColorHueMap["red"]),
			universe.SetSize(2),
		},
	)
	setupPatches()
	environment.ResetTicks()
}

func setupPatches() {
	environment.AskPatches(
		[]universe.PatchOperation{
			setupNest,
			setupFood,
			setupNest,
		},
	)
}

func setupNest(p *universe.Patch) {
	p.PatchesOwn[nest] = p.DistanceXY(0, 0) < 5
	p.PatchesOwn[nestScent] = 200 - p.DistanceXY(0, 0)
}

func setupFood(p *universe.Patch) {
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

func recolorPatch(p *universe.Patch) {

	// ;; give color to nest and food sources
	if p.PatchesOwn[nest].(bool) {
		p.Color = environment.ColorHueMap["violet"]
	} else {
		if p.PatchesOwn[food].(int) > 0 {
			if p.PatchesOwn[foodSourceNumber].(int) == 1 {
				p.Color = environment.ColorHueMap["cyan"]
			} else if p.PatchesOwn[foodSourceNumber].(int) == 2 {
				p.Color = environment.ColorHueMap["sky"]
			} else if p.PatchesOwn[foodSourceNumber].(int) == 3 {
				p.Color = environment.ColorHueMap["blue"]
			}
		} else {
			p.SetColorAndScale(p.PatchesOwn[chemical].(float64), .1, 5)
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

	environment.AskTurtles(
		[]universe.TurtleOperation{
			func(t *universe.Turtle) {
				if t.Who >= environment.Ticks {
					return
				}
				if t.Color == environment.ColorHueMap["red"] {
					lookForFood(t)
				} else {
					returnToNest(t)
				}
				wiggle(t)
			},
		},
	)
}

func returnToNest(t *universe.Turtle) {
	if t.GetPatch().PatchesOwn[nest].(bool) {
		t.Color = environment.ColorHueMap["red"]
		t.Right(180)
	} else {
		t.GetPatch().PatchesOwn[chemical] = t.GetPatch().PatchesOwn[chemical].(int) + 60
		uphillNestScent(t)
	}
}

func lookForFood(t *universe.Turtle) {
	// if food > 0
	// [ set color orange + 1     ;; pick up food
	//   set food food - 1        ;; and reduce the food source
	//   rt 180                   ;; and turn around
	//   stop ]
	// ;; go in the direction where the chemical smell is strongest
	// if (chemical >= 0.05) and (chemical < 2)
	// [ uphill-chemical ]
	p := t.GetPatch()
	if p.PatchesOwn[food].(int) > 0 {

	}

}

func uphillChemical(t *universe.Turtle) {
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

func wiggle(t *universe.Turtle) {
	t.Right(float64(environment.RandomAmount(40)))
	t.Left(float64(environment.RandomAmount(40)))

	if !t.CanMove(1) {
		t.Right(180)
	}
}

func uphillNestScent(t *universe.Turtle) {
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

func nestScentAtAngle(t *universe.Turtle, angle float64) int {
	p := t.PatchRightAndAhead(angle, 1)
	if p == nil {
		return 0
	} else {
		return p.PatchesOwn[nestScent].(int)
	}
}

func chemicalScentAtAngle(t *universe.Turtle, angle float64) float64 {
	p := t.PatchRightAndAhead(angle, 1)
	if p == nil {
		return 0
	} else {
		return p.PatchesOwn[chemical].(float64)
	}
}
