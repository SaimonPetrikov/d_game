package level_scene

import (
	"d_game/maps"

	"github.com/hajimehoshi/ebiten/v2"
)

type tileNode struct {
	tile maps.TileMapI
	gameObject
}

func newTile(tile maps.TileMapI) *tileNode {
	return &tileNode{
		tile: tile,
	}
}

func (t *tileNode) Init(s *scene) {
	object := t.tile.GetObject()
	if (object != nil) {
		s.Controller().ctx.Space.Add(object)
	}
	s.AddGraphics(t)
}

func (t *tileNode) Draw(screen *ebiten.Image) {
    t.tile.Draw(screen)
}

func (t *tileNode) IsDisposed() bool {
	return t.IsDeleted
}

func (t *tileNode) Update(delta float64) {
}
