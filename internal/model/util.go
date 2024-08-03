package model

func BaseColors() []Color {
	return baseColorsList
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

func LinkShapes() []string {
	return []string{
		"line", "curve", "arrow", "arrow2", "arrow3", "arrow4", "arrow5", "arrow6", "arrow7", "arrow8", "arrow9", "arrow10",
	}
}
