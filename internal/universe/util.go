package universe

// @TODO implement
func ApproximateHSB(hue int, saturation int, brightness int) float64 {
	return 0.0
}

// @TODO implement
func ApproximateRGB(red int, green int, blue int) float64 {
	return 0.0
}

func BaseColors() []float64 {
	return []float64{
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
}

func ButFirst(arr []interface{}) []interface{} {
	return arr[1:]
}

func ButLast(arr []interface{}) []interface{} {
	return arr[:len(arr)-1]
}

// @TODO implement
func ExtractHSBFromScale(scale float64) (int, int, int) {
	return 0, 0, 0
}

// @TODO implement
func ExtractHSBFromRBG(red int, green int, blue int) (int, int, int) {
	return 0, 0, 0
}

// @TODO implement
func ExtractRGBFromScale(scale float64) (int, int, int) {
	return 0, 0, 0
}

func Filter(arr []interface{}, pred func(interface{}) bool) []interface{} {
	var result []interface{}
	for _, elem := range arr {
		if pred(elem) {
			result = append(result, elem)
		}
	}
	return result
}

func HSB(hue int, saturation int, brightness int) (int, int, int) {
	return 0, 0, 0
}

// @TODO implement
func OneOf(arr []interface{}) interface{} {
	return arr[0]
}

// @TODO implement
func ScaleColor(color Color, number float64, range1 float64, range2 float64) Color {
	return color
}

// @TODO implement
func ShadeOf(color1 float64, color2 float64) bool {
	return false
}

// @TODO implement
func SubtractHeadings(h1 float64, h2 float64) float64 {
	return 0.0
}
