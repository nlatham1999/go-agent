package model

type ModelSettings struct {
	PatchesOwn           map[string]interface{}
	TurtlesOwn           map[string]interface{}
	TurtleBreedsOwn      map[string]map[string]interface{}
	TurtleBreeds         []string
	DirectedLinkBreeds   []string
	UndirectedLinkBreeds []string
	WrappingX            bool
	WrappingY            bool
	Globals              map[string]interface{}
}
