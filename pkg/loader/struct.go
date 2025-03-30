package loader

type Model struct {
	TurtleBreeds         []TurtleBreed `json:"turtleBreeds"`
	DirectedLinkBreeds   []LinkBreed   `json:"directedLinkBreeds"`
	UndirectedLinkBreeds []LinkBreed   `json:"undirectedLinkBreeds"`

	PatchProperties  map[string]interface{} `json:"patchProperties"`
	TurtleProperties map[string]interface{} `json:"turtleProperties"`

	WrappingX bool `json:"wrappingX"`
	WrappingY bool `json:"wrappingY"`

	WorldWidth  int `json:"width"`
	WorldHeight int `json:"height"`

	MinPxCor int `json:"minPxCor"`
	MaxPxCor int `json:"maxPxCor"`
	MinPyCor int `json:"minPyCor"`
	MaxPyCor int `json:"maxPyCor"`

	RandomSeed1 uint64 `json:"randomSeed1"`
	RandomSeed2 uint64 `json:"randomSeed2"`
	RandomState string `json:"randomState"`

	Patches []Patch  `json:"patches"`
	Turtles []Turtle `json:"turtles"`
	Links   []Link   `json:"links"`
	Ticks   int      `json:"ticks"`
}

type Patch struct {
	X          int                    `json:"x"`
	Y          int                    `json:"y"`
	Color      Color                  `json:"color"`
	Properties map[string]interface{} `json:"properties"`
}

type Turtle struct {
	X          float64                `json:"x"`
	Y          float64                `json:"y"`
	Color      Color                  `json:"color"`
	Size       float64                `json:"size"`
	Who        int                    `json:"who"`
	Shape      string                 `json:"shape"`
	Heading    float64                `json:"heading"`
	Label      interface{}            `json:"label"`
	LabelColor Color                  `json:"labelColor"`
	Properties map[string]interface{} `json:"properties"`
	Breed      string                 `json:"breed"`
}

type Color struct {
	Red   int `json:"r"`
	Green int `json:"g"`
	Blue  int `json:"b"`
	Alpha int `json:"a"`
}

type Link struct {
	End1       int         `json:"end1"`
	End2       int         `json:"end2"`
	End1X      float64     `json:"end1X"`
	End1Y      float64     `json:"end1Y"`
	End2X      float64     `json:"end2X"`
	End2Y      float64     `json:"end2Y"`
	Directed   bool        `json:"directed"`
	Color      Color       `json:"color"`
	Label      interface{} `json:"label"`
	LabelColor Color       `json:"labelColor"`
	Size       int         `json:"size"`
	Hidden     bool        `json:"hidden"`
	Breed      string      `json:"breed"`
}

type TurtleBreed struct {
	Name         string                 `json:"name"`
	Properties   map[string]interface{} `json:"properties"`
	DefaultShape string                 `json:"defaultShape"`
}

type LinkBreed struct {
	Name         string `json:"name"`
	DefaultShape string `json:"defaultShape"`
	Directed     bool   `json:"directed"`
}
