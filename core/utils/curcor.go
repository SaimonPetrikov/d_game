package utils

import "github.com/hajimehoshi/ebiten/v2"

func CursorPosition() Vector {
	x, y := ebiten.CursorPosition()

	return Vector{
		X: float64(x),
		Y: float64(y),
	}
}