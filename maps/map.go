package maps

import (
	"d_game/core/gmap"
	"d_game/core/resolve_collision"
	utils "d_game/core/utils"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

//todo вынести map в root_map переписать стены на интерфейсы

type LevelMap struct {
	data gmap.Map
	Height, Width int
	HeightCount, WidthCount int
	TileImages map[int]*ebiten.Image
	tileOptions map[int]*ebiten.DrawImageOptions
	TileSet []TileMapI
}

func InitMap (nameMap string) *LevelMap {
	mapData := gmap.LoadLevelMapData(nameMap)
	tilesmap := make(map[int]gmap.Tile)

	levelMap := &LevelMap{
		data: mapData,
		TileImages: map[int]*ebiten.Image{},
		tileOptions: map[int]*ebiten.DrawImageOptions{},
	}

	for _, tile := range mapData.TileSet.Tiles {
		tilesmap[tile.Id] = tile
	}

	index := 0

	height := mapData.TileHeight
	width := mapData.TileWidth

	levelMap.Height, levelMap.Width = height, width	
	levelMap.HeightCount, levelMap.WidthCount = mapData.Height, mapData.Width

	for indexLayer := 0; indexLayer < len(mapData.Layers); indexLayer++ {
		layer := mapData.Layers[indexLayer]
		for i := 0; i < mapData.Height; i++ {
			for j := 0; j < mapData.Width; j++ {
				tileId := layer.Data[index]
				rootTile := &RootTileMap{
					TileId: tileId,
					Image: levelMap.getTileInCacheOrLoad(tileId),
					PosX: float64(height*j),
					PosY: float64(width*i),
					Height: float64(height),
					Width: float64(width),
					Object: resolve_collision.NewObject(float64(height*j), float64(width*i), float64(height), float64(width), "floor"),
				}
				rootTile.Object.SetShape(resolve_collision.NewRectangle(float64(height*j), float64(width*i), float64(height), float64(width)))
				if tilesmap[tileId].Weight == 1 {
					rootTile.Cost = 100000000
					levelMap.TileSet = append(levelMap.TileSet, NewWall(rootTile))
					index++
					continue
				} 

				rootTile.Cost = 1
				levelMap.TileSet = append(levelMap.TileSet, rootTile)
				
				index++
			}
		}

		//todo: добавить диагонали
		for i := 0; i < mapData.Height; i++ {
			for j := 0; j < mapData.Width; j++ {
				current := levelMap.TileSet[j + i*mapData.Width]
				if j + (i-1)*mapData.Width >= 0 {
					to := levelMap.TileSet[j + (i-1)*mapData.Width]
					if !to.IsResolveCollision() {
						current.AddLink(current, to)
					}
				}

				if j + (i+1)*mapData.Width < 600 {
					to := levelMap.TileSet[j + (i+1)*mapData.Width]
					if !to.IsResolveCollision() {
						current.AddLink(current, to)
					}
				}

				if j - 1 + i*mapData.Width >= 0 {
					to := levelMap.TileSet[j - 1 + i*mapData.Width]
					if !to.IsResolveCollision() {
						current.AddLink(current, to)
					}
				}

				if j + 1 + i*mapData.Width < 600 {
					to := levelMap.TileSet[j + 1 + i*mapData.Width]
					if !to.IsResolveCollision() {
						current.AddLink(current, to)
					}
				}
			}
		}
	}

	return levelMap
}

func (levelMap *LevelMap) getTileInCacheOrLoad(id int) *ebiten.Image {
	if tile, ok := levelMap.TileImages[id]; ok {
		return tile
	}

	tile := utils.LoadImage("assets/map/Tile" + strconv.Itoa(id) + ".png")
	levelMap.TileImages[id] = tile

	return tile
}

func (levelMap *LevelMap) IsDisposed() bool {
	return false
}