package model

import "math"

func BaseColors() []Color {
	return baseColorsList
}

func radiansToDegrees(radians float64) float64 {
	degrees := radians * (180 / math.Pi)
	if degrees < 0 {
		degrees += 360
	}
	return degrees
}
