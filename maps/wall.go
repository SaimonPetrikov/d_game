package maps

import "github.com/hajimehoshi/ebiten/v2"

type Wall struct {
	rootTile *RootTileMap
}

func NewWall(rootTile *RootTileMap) *Wall {
	return &Wall{rootTile: rootTile}
}

func (wall *Wall) Draw(screen *ebiten.Image) {
	wall.rootTile.Draw(screen)
}

func (wall *Wall) AddLink(from, to TileMapI) {
}

func (wall *Wall) GetTile() *RootTileMap {
	return wall.rootTile
}

func (wall *Wall) GetCost() int {
	return wall.rootTile.cost
}

func (wall *Wall) getImageOptions() *ebiten.DrawImageOptions {
	return wall.rootTile.getImageOptions()
}

func (wall *Wall) GetPosAndSize() (float64, float64, float64, float64) {
	return wall.rootTile.GetPosAndSize()
}

func (tile *Wall) GetPosX() float64 {
	return tile.rootTile.PosX
}

func (tile *Wall) GetPosY() float64 {
	return tile.rootTile.PosY
}

func (wall *Wall) IsResolveCollision() bool {
	return true
}

func (wall *Wall) IsDisposed() bool {
	return false
}