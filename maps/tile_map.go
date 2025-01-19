package maps

import "github.com/hajimehoshi/ebiten/v2"

type TileMapI interface {
	Draw(screen *ebiten.Image)
	getImageOptions() *ebiten.DrawImageOptions
	GetPosAndSize() (float64, float64, float64, float64)
	IsResolveCollision() bool
	IsDisposed() bool
}


type rootTileMap struct {
	TileId int
	Image  *ebiten.Image
	PosX   float64
	PosY   float64
	Height float64
	Width  float64
	cost int
}

func AStar(start, goal *rootTileMap, grid []*rootTileMap) []*rootTileMap {
    var closedSet []*rootTileMap
    var openSet = []*rootTileMap{start}
    start.cost = 0

    for len(openSet) > 0 {
        var current = openSet[0]
        if current == goal {
            var path []*rootTileMap
            for current != nil {
                path = append([]*rootTileMap{current}, path...)
            }
            return path
        }
        openSet = openSet[1:]
        closedSet = append(closedSet, current)
        // Code for updating cost and checking for better paths
        // Code for adding neighbors to the openSet
    }
    return nil
}

func (tile *rootTileMap) Draw(screen *ebiten.Image) {
	screen.DrawImage(tile.Image, tile.getImageOptions())
}

func (tile *rootTileMap) getImageOptions() *ebiten.DrawImageOptions {
	tileOption := &ebiten.DrawImageOptions{}
	tileOption.GeoM.Translate(tile.PosX, tile.PosY)
	return tileOption
}

func (tile *rootTileMap) GetPosAndSize() (float64, float64, float64, float64) {
	return tile.PosX, tile.PosY, tile.Height, tile.Width
}

func (tile *rootTileMap) IsResolveCollision() bool {
	return false
}

func (tile *rootTileMap) IsDisposed() bool {
	return false
}