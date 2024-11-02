package model

// colors
var (
	Black     = Color{Red: 0, Green: 0, Blue: 0, Alpha: 1}
	White     = Color{Red: 255, Green: 255, Blue: 255, Alpha: 1}
	Grey      = Color{Red: 128, Green: 128, Blue: 128, Alpha: 1}
	Gray      = Color{Red: 128, Green: 128, Blue: 128, Alpha: 1}
	Red       = Color{Red: 255, Green: 0, Blue: 0, Alpha: 1}
	Orange    = Color{Red: 255, Green: 165, Blue: 0, Alpha: 1}
	Brown     = Color{Red: 88, Green: 57, Blue: 39, Alpha: 1}
	Yellow    = Color{Red: 255, Green: 255, Blue: 0, Alpha: 1}
	Green     = Color{Red: 0, Green: 128, Blue: 0, Alpha: 1}
	Lime      = Color{Red: 0, Green: 255, Blue: 0, Alpha: 1}
	Turquoise = Color{Red: 64, Green: 224, Blue: 208, Alpha: 1}
	Cyan      = Color{Red: 0, Green: 255, Blue: 255, Alpha: 1}
	Sky       = Color{Red: 135, Green: 206, Blue: 235, Alpha: 1}
	Blue      = Color{Red: 0, Green: 0, Blue: 255, Alpha: 1}
	Violet    = Color{Red: 238, Green: 130, Blue: 238, Alpha: 1}
	Magenta   = Color{Red: 255, Green: 0, Blue: 255, Alpha: 1}
	Pink      = Color{Red: 255, Green: 192, Blue: 203, Alpha: 1}

	baseColorsList = []Color{
		Black,
		White,
		Grey,
		Gray,
		Red,
		Orange,
		Brown,
		Yellow,
		Green,
		Lime,
		Turquoise,
		Cyan,
		Sky,
		Blue,
		Violet,
		Magenta,
		Pink,
	}
)

// angles
var (
	RightAngle        float64 = 0
	UpAndRightAngle   float64 = 45.0
	UpAngle           float64 = 90.0
	UpAndLeftAngle    float64 = 135.0
	LeftAngle         float64 = 180.0
	DownAndLeftAngle  float64 = 225.0
	DownAngle         float64 = 270.0
	DownAndRightAngle float64 = 315.0
)
