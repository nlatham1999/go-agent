package model

// modelSettings holds the settings for the model
type ModelSettings struct {
	PatchesOwn           map[string]interface{}
	TurtlesOwn           map[string]interface{}
	TurtleBreedsOwn      map[string]map[string]interface{}
	TurtleBreeds         []string
	DirectedLinkBreeds   []string
	UndirectedLinkBreeds []string
	WrappingX            bool
	WrappingY            bool
	MinPxCor             int
	MaxPxCor             int
	MinPyCor             int
	MaxPyCor             int
	RandomSeed           int64
}
