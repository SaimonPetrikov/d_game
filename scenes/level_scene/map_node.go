package level_scene

import (
	"d_game/maps"
)

type mapNode struct {
	levelMap *maps.LevelMap
}

func newMap() *mapNode {
	levelMap := maps.InitMap("first_map")
	return &mapNode{levelMap: levelMap}
}

func (mapNode *mapNode) Init(s *scene) {
	for _, tile := range mapNode.levelMap.TileSet {
		tileNode := newTile(tile)
		s.AddObject(tileNode)
	}
}

func (mapNode *mapNode) IsDisposed() bool {
	return false
}

func (mapNode *mapNode) Update(delta float64) {
	// Здесь код, который раньше был в myGame Update.
}
