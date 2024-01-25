package ants

import (
	patch "github.com/nlatham1999/go-agent/internal/patches"
	turtle "github.com/nlatham1999/go-agent/internal/turtles"
	"github.com/nlatham1999/go-agent/internal/universe"
)

var (
	environment *universe.Universe

	//@TODO figure out how to handle widgets
	population int
)

func Init() {
	patchesOwn := map[string]interface{}{
		"chemical":         0.0,
		"food":             0.0,
		"nest":             false,
		"nestScent":        0,
		"foodSourceNumber": 0,
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
	p.PatchesOwn["nest"] = p.DistanceXY(0, 0) < 5
	p.PatchesOwn["nestScent"] = 200 - p.DistanceXY(0, 0)
}

func setupFood(p *patch.Patch) {
}

func recolorPatch(p *patch.Patch) {
}
