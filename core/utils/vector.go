package utils

type Vector struct {
	X, Y float64
}

func GetDistance(v1, v2 Vector) Vector {
	return Vector{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}