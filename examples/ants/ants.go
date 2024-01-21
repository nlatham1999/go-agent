package ants

import (
	patch "github.com/nlatham1999/go-agent/internal/patches"
	turtle "github.com/nlatham1999/go-agent/internal/turtles"
	"github.com/nlatham1999/go-agent/internal/universe"
)

var (
	enviroment *universe.Universe

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

	enviroment = universe.NewUniverse(patchesOwn)

}

func setup() {
	enviroment.ClearAll()
	enviroment.SetDefaultShapeTurtles("bug")
	enviroment.CreateTurtles(population,
		[]turtle.TurtleOperation{
			turtle.SetColor("red"),
			turtle.SetSize(2),
		},
	)
	setupPatches()
	enviroment.ResetTicks()
}

func setupPatches() {
	enviroment.AskPatches(
		[]patch.PatchOperation{
			setupNest,
			setupFood,
			setupNest,
		},
	)
}

func setupNest(p *patch.Patch) {
}

func setupFood(p *patch.Patch) {
}

func recolorPatch(p *patch.Patch) {
}
