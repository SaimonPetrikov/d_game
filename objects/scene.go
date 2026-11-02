package objects

import (
	"d_game/camera"
	"d_game/core/astar"
	"d_game/core/resolve_collision"
	"d_game/maps"
)

type State[ControllerAccessor any] struct {
	NpcLst 	  []*Npc[ControllerAccessor]
	Tiles     []maps.TileMapI
	Cam 	  *camera.Camera
	PathMap   map[string][]astar.Pather
	PlayerPosition *Vector
	PlayerTile *int
	Space *resolve_collision.Space
}

func (s *State[ControllerAccessor]) AddToSpace(object *resolve_collision.Object) {
	s.Space.Add(object)
}