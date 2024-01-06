package ants

import (
	turtle "github.com/nlatham1999/go-agent/internal/turtles"
	"github.com/nlatham1999/go-agent/internal/universe"
)

var (
	enviroment *universe.Universe

	//@TODO figure out how to handle widgets
	population int
)

func Init() {
	enviroment = universe.NewUniverse()
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
	enviroment.ResetTickCounter()
}

func setupPatches() {

}
