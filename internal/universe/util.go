package universe

//@TODO implement
func ApproximateHSB(hue int, saturation int, brightness int) float64 {
	return 0.0
}

//@TODO implement
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
