package model

func BaseColors() []Color {
	return []Color{
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

func Filter(arr []interface{}, pred func(interface{}) bool) []interface{} {
	var result []interface{}
	for _, elem := range arr {
		if pred(elem) {
			result = append(result, elem)
		}
	}
	return result
}

// @TODO implement
func SubtractHeadings(h1 float64, h2 float64) float64 {
	return 0.0
}
