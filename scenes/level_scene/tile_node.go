package level_scene

import (
	"d_game/core/gobjects"
	"d_game/core/resolve_collision"
	"d_game/maps"
)

type tileNode struct {
	tile maps.TileMapI
	gameObject
}

func newTile(tile maps.TileMapI) *tileNode {
	x, y, h, w := tile.GetPosAndSize()
	return &tileNode{
		tile: tile,
		gameObject: gobjects.Object{
			Object: resolve_collision.NewObject(x, y, h, w, "solid"),
			IsDeleted: true,
		},
	}
}

func (t *tileNode) Init(s *scene) {
	if (t.tile.IsResolveCollision()) {
		s.Controller().ctx.Space.Add(t.Object)
	}
	s.AddGraphics(t.tile)
}

func (t *tileNode) IsDisposed() bool {
	return t.IsDeleted
}

func (t *tileNode) Update(delta float64) {
	t.IsDeleted = true
	if check := t.Object.Check(5, 5, "bullet"); check != nil {
		t.IsDeleted = true
	}
}
