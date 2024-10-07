package maps

import "github.com/hajimehoshi/ebiten/v2"

type Wall struct {
	rootTile *rootTileMap
}

func NewWall(rootTile *rootTileMap) *Wall {
	return &Wall{rootTile: rootTile}
}

func (wall *Wall) Draw(screen *ebiten.Image) {
	wall.rootTile.Draw(screen)
}

func (wall *Wall) getImageOptions() *ebiten.DrawImageOptions {
	return wall.rootTile.getImageOptions()
}

func (wall *Wall) GetPosAndSize() (float64, float64, float64, float64) {
	return wall.rootTile.GetPosAndSize()
}

func (wall *Wall) IsResolveCollision() bool {
	return true
}

func (wall *Wall) IsDisposed() bool {
	return false
}