package api

type Model struct {
	Patches          []Patch  `json:"patches"`
	Turtles          []Turtle `json:"turtles"`
	Links            []Link   `json:"links"`
	DynamicVariables map[string]interface{}
	Ticks            int `json:"ticks"`
	WorldWidth       int `json:"width"`
	WorldHeight      int `json:"height"`
	MinPxCor         int `json:"minPxCor"`
	MaxPxCor         int `json:"maxPxCor"`
	MinPyCor         int `json:"minPyCor"`
	MaxPyCor         int `json:"maxPyCor"`
}

type Patch struct {
	X     int   `json:"x"`
	Y     int   `json:"y"`
	Color Color `json:"color"`
}

type Turtle struct {
	X          float64     `json:"x"`
	Y          float64     `json:"y"`
	Color      Color       `json:"color"`
	Size       float64     `json:"size"`
	Who        int         `json:"who"`
	Shape      string      `json:"shape"`
	Heading    float64     `json:"heading"`
	Label      interface{} `json:"label"`
	LabelColor Color       `json:"labelColor"`
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
	End1Size   float64     `json:"end1Size"`
	End2Size   float64     `json:"end2Size"`
	Directed   bool        `json:"directed"`
	Color      Color       `json:"color"`
	Label      interface{} `json:"label"`
	LabelColor Color       `json:"labelColor"`
	Size       int         `json:"size"`
	Hidden     bool        `json:"hidden"`
}
