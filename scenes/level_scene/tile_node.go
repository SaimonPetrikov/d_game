package level_scene

import (
	"d_game/core/resolve_collision"
	"d_game/maps"
	"fmt"
)

type tileNode struct {
	tile maps.TileMapI
	Object *resolve_collision.Object
	is_Disposed bool
}

func newTile(tile maps.TileMapI) *tileNode {
	x, y, h, w := tile.GetPosAndSize()
	return &tileNode{
		tile: tile,
		Object: resolve_collision.NewObject(x, y, h, w, "solid"),
		is_Disposed: true,
	}
}

func (t *tileNode) Init(s *scene) {
	if (t.tile.IsResolveCollision()) {
		s.Controller().ctx.Space.Add(t.Object)
	}
	s.AddGraphics(t.tile)
}

func (t *tileNode) IsDisposed() bool {
	return t.is_Disposed
}

func (t *tileNode) Update(delta float64) {
	t.is_Disposed = true
	if check := t.Object.Check(5, 5, "bullet"); check != nil {
		fmt.Printf("bullet")
		t.is_Disposed = true
	}
	
	// Здесь код, который раньше был в myGame Update.
}
