package model

// colors
var (
	Black     = Color{Red: 0, Green: 0, Blue: 0}
	White     = Color{Red: 255, Green: 255, Blue: 255}
	Grey      = Color{Red: 128, Green: 128, Blue: 128}
	Gray      = Color{Red: 128, Green: 128, Blue: 128}
	Red       = Color{Red: 255, Green: 0, Blue: 0}
	Orange    = Color{Red: 255, Green: 165, Blue: 0}
	Brown     = Color{Red: 165, Green: 42, Blue: 42}
	Yellow    = Color{Red: 255, Green: 255, Blue: 0}
	Green     = Color{Red: 0, Green: 128, Blue: 0}
	Lime      = Color{Red: 0, Green: 255, Blue: 0}
	Turquoise = Color{Red: 64, Green: 224, Blue: 208}
	Cyan      = Color{Red: 0, Green: 255, Blue: 255}
	Sky       = Color{Red: 135, Green: 206, Blue: 235}
	Blue      = Color{Red: 0, Green: 0, Blue: 255}
	Violet    = Color{Red: 238, Green: 130, Blue: 238}
	Magenta   = Color{Red: 255, Green: 0, Blue: 255}
	Pink      = Color{Red: 255, Green: 192, Blue: 203}
)

var (
	TieModeFixed = TieMode("fixed")
	TieModeFree  = TieMode("free")
	TieModeNone  = TieMode("none")
)
