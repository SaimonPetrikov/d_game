package utils

import "github.com/hajimehoshi/ebiten/v2"

type DrawImageOptions struct {
	*ebiten.DrawImageOptions
}

func (d DrawImageOptions) GeomTranslate(v Vector) {
	d.GeoM.Translate(v.X, v.Y)
}