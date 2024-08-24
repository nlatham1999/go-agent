package api

type Model struct {
	Patches []Patch  `json:"patches"`
	Turtles []Turtle `json:"turtles"`
	Links   []Link   `json:"links"`
}

type Patch struct {
	X     int   `json:"x"`
	Y     int   `json:"y"`
	Color Color `json:"color"`
}

type Turtle struct {
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Color   Color   `json:"color"`
	Size    float64 `json:"size"`
	Who     int     `json:"who"`
	Shape   string  `json:"shape"`
	Heading float64 `json:"heading"`
}

type Color struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
}

type Link struct {
	End1     int  `json:"end1"`
	End2     int  `json:"end2"`
	Directed bool `json:"directed"`
}