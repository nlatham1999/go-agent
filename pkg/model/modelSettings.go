package model

// modelSettings holds the settings for the model
type ModelSettings struct {
	PatchProperties      map[string]interface{}
	TurtleProperties     map[string]interface{}
	TurtleBreeds         []*TurtleBreed
	DirectedLinkBreeds   []*LinkBreed
	UndirectedLinkBreeds []*LinkBreed
	WrappingX            bool
	WrappingY            bool
	MinPxCor             int
	MaxPxCor             int
	MinPyCor             int
	MaxPyCor             int
	MinPzCor             int // in development
	MaxPzCor             int // in development
	RandomSeed           uint64
	RandomSeed2          uint64
}
