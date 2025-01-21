package maps

import (
	"d_game/core/gmap"
	utils "d_game/core/utils"
	"fmt"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

//todo вынести map в root_map переписать стены на интерфейсы

type LevelMap struct {
	data gmap.Map
	Height, Width int
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
				}
				if tilesmap[tileId].Weight == 1 {
					rootTile.cost = 100000000
					levelMap.TileSet = append(levelMap.TileSet, NewWall(rootTile))
					index++
					continue
				} 

				rootTile.cost = 1
				levelMap.TileSet = append(levelMap.TileSet, rootTile)
				
				index++
			}
		}

		//todo: добавить диагонали
		for i := 0; i < mapData.Height; i++ {
			for j := 0; j < mapData.Width; j++ {
				current := levelMap.TileSet[j + i*mapData.Width]
				if j + (i-1)*mapData.Width >= 0 {
					fmt.Println(current.GetTile().PosX, ", ", current.GetTile().PosY)
					to := levelMap.TileSet[j + (i-1)*mapData.Width]
					current.AddLink(current, to)
				}

				if j + (i+1)*mapData.Width < 600 {
					to := levelMap.TileSet[j + (i+1)*mapData.Width]
					current.AddLink(current, to)
				}

				if j - 1 + i*mapData.Width >= 0 {
					to := levelMap.TileSet[j - 1 + i*mapData.Width]
					current.AddLink(current, to)
				}

				if j + 1 + i*mapData.Width < 600 {
					to := levelMap.TileSet[j + 1 + i*mapData.Width]
					current.AddLink(current, to)
				}
			}
		}
	}

	return levelMap
}

// func createGraph() {

// }

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