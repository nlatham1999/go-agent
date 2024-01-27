package util

import "math/rand"

func OneOfInt(arr []int) interface{} {

	return arr[rand.Intn(len(arr))-1]
}
