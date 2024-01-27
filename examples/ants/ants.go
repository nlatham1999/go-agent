package ants

import (
	"math/rand"
	"time"

	patch "github.com/nlatham1999/go-agent/internal/patches"
	turtle "github.com/nlatham1999/go-agent/internal/turtles"
	"github.com/nlatham1999/go-agent/internal/universe"
	"github.com/nlatham1999/go-agent/internal/util"
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
		[]turtle.TurtleOperation{
			turtle.SetColor("red"),
			turtle.SetSize(2),
		},
	)
	setupPatches()
	environment.ResetTicks()
}

func setupPatches() {
	environment.AskPatches(
		[]patch.PatchOperation{
			setupNest,
			setupFood,
			setupNest,
		},
	)
}

func setupNest(p *patch.Patch) {
	p.PatchesOwn[nest] = p.DistanceXY(0, 0) < 5
	p.PatchesOwn[nestScent] = 200 - p.DistanceXY(0, 0)
}

func setupFood(p *patch.Patch) {
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
		p.PatchesOwn[food] = util.OneOfInt([]int{1, 2})
	}
}

func recolorPatch(p *patch.Patch) {

	// ;; give color to nest and food sources
	if p.PatchesOwn[nest].(bool) {
		p.Color = "violet"
	} else {
		if p.PatchesOwn[food].(int) > 0 {
			if p.PatchesOwn[foodSourceNumber].(int) == 1 {
				p.Color = "cyan"
			} else if p.PatchesOwn[foodSourceNumber].(int) == 2 {
				p.Color = "sky"
			} else if p.PatchesOwn[foodSourceNumber].(int) == 3 {
				p.Color = "blue"
			}
		} else {
			p.SetColorAndScale(p.PatchesOwn[chemical].(float64), .1, 5)
		}
	}
}
