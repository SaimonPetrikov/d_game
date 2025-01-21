package maps

import (
	"d_game/core/astar"

	"github.com/hajimehoshi/ebiten/v2"
)

type TileMapI interface {
	Draw(screen *ebiten.Image)
	getImageOptions() *ebiten.DrawImageOptions
	GetPosX() float64
	GetPosY() float64
	GetPosAndSize() (float64, float64, float64, float64)
	IsResolveCollision() bool
	IsDisposed() bool
	AddLink(from, to TileMapI)
	GetCost() int
	GetTile() *RootTileMap
}

type Link struct {
	From, To TileMapI
	Cost float64
}


type RootTileMap struct {
	TileId int
	Image  *ebiten.Image
	PosX   float64
	PosY   float64
	Height float64
	Width  float64
	Links []Link
	cost int
}

func (tile *RootTileMap) AddLink(from, to TileMapI) {
	tile.Links = append(tile.Links, Link{From: from, To: to, Cost: float64(to.GetCost())})
}

func (tile *RootTileMap) GetCost() int {
	return tile.cost
}

func (tile *RootTileMap) GetTile() *RootTileMap {
	return tile
}

func (tile *RootTileMap) PathNeighbors() []astar.Pather {
	neighbors := []astar.Pather{}

	for _, link := range tile.Links {
		neighbors = append(neighbors, astar.Pather(link.To.GetTile()))
	}
	return neighbors
}

func (tile *RootTileMap) PathNeighborCost(to astar.Pather) float64 {
	for _, link := range (tile).Links {
		if astar.Pather((link.To.GetTile())) == to {
			return link.Cost
		}
	}
	return 10000000
}

func (tile *RootTileMap) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*RootTileMap)
	absX := toT.PosX - tile.PosY
	if absX < 0 {
		absX = -absX
	}
	absY := toT.PosY - tile.PosX
	if absY < 0 {
		absY = -absY
	}
	r := float64(absX + absY)

	return r
}

func (tile *RootTileMap) Draw(screen *ebiten.Image) {
	screen.DrawImage(tile.Image, tile.getImageOptions())
}

func (tile *RootTileMap) getImageOptions() *ebiten.DrawImageOptions {
	tileOption := &ebiten.DrawImageOptions{}
	tileOption.GeoM.Translate(tile.PosX, tile.PosY)
	return tileOption
}

func (tile *RootTileMap) GetPosX() float64 {
	return tile.PosX
}

func (tile *RootTileMap) GetPosY() float64 {
	return tile.PosY
}

func (tile *RootTileMap) GetPosAndSize() (float64, float64, float64, float64) {
	return tile.PosX, tile.PosY, tile.Height, tile.Width
}

func (tile *RootTileMap) IsResolveCollision() bool {
	return false
}

func (tile *RootTileMap) IsDisposed() bool {
	return false
}