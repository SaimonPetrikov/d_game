package maps

import (
	"d_game/core/gmap"
	utils "d_game/core/utils"
	"strconv"
	"github.com/hajimehoshi/ebiten/v2"
)

//todo вынести map в root_map переписать стены на интерфейсы

type LevelMap struct {
	data gmap.Map
	TileImages map[int]*ebiten.Image
	tileOptions map[int]*ebiten.DrawImageOptions
	TileSet []TileMapI
}

func InitMap (nameMap string) *LevelMap {
	mapData := gmap.LoadLevelMapData(nameMap)

	index := 0

	levelMap := &LevelMap{
		data: mapData,
		TileImages: map[int]*ebiten.Image{},
		tileOptions: map[int]*ebiten.DrawImageOptions{},
	}

	for indexLayer := 0; indexLayer < len(mapData.Layers); indexLayer++ {
		layer := mapData.Layers[indexLayer]
		for i := 0; i < mapData.Height; i++ {
			for j := 0; j < mapData.Width; j++ {
				tileId := layer.Data[index]
				for _, tile := range mapData.TileSet.Tiles {
					rootTile := &rootTileMap{
						TileId: tileId,
						Image: levelMap.getTileInCacheOrLoad(tileId),
						PosX: float64(mapData.TileHeight*j),
						PosY: float64(mapData.TileWidth*i),
						Height: float64(mapData.TileHeight),
						Width: float64(mapData.TileWidth),
					}

					if tile.Id == tileId && tile.Weight == 1 {
						levelMap.TileSet = append(levelMap.TileSet, NewWall(rootTile))
					} else {
						//it is floor
						levelMap.TileSet = append(levelMap.TileSet, rootTile)
					}
				}
				index++
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