package level_scene

import "d_game/maps"

type SceneState struct {
	Policeman *Policeman
	Tiles     []maps.TileMapI
}