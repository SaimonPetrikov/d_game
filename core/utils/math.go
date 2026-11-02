package utils

import "math"

func Atan2(v Vector) float64 {
	return math.Atan2(v.Y, v.X)
}