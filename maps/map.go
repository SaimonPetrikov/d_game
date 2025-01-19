package maps

import (
	"d_game/core/astar"
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
	tilesmap := make(map[int]gmap.Tile)

	for _, tile := range mapData.TileSet.Tiles {
		tilesmap[tile.Id] = tile
	}

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
				rootTile := &rootTileMap{
					TileId: tileId,
					Image: levelMap.getTileInCacheOrLoad(tileId),
					PosX: float64(mapData.TileHeight*j),
					PosY: float64(mapData.TileWidth*i),
					Height: float64(mapData.TileHeight),
					Width: float64(mapData.TileWidth),
				}
				if tilesmap[tileId].Weight == 1 {
					rootTile.cost = 1
					levelMap.TileSet = append(levelMap.TileSet, NewWall(rootTile))
					index++
					continue
				} 

				rootTile.cost = 0


						
				//it is floor
				levelMap.TileSet = append(levelMap.TileSet, rootTile)
					
				key := strconv.Itoa(i) + "." + strconv.Itoa(j)
				astar.Graph[key] = astar.Node{
					Sx: mapData.TileHeight*j + mapData.TileHeight/2,
					Sy: mapData.TileWidth*i + mapData.TileWidth/2,
					Neighbors: map[string]string{},
				}
				
				index++
			}
		}

		for i := 0; i < mapData.Height; i++ {
			for j := 0; j < mapData.Width; j++ {
				key := strconv.Itoa(i) + "." + strconv.Itoa(j)
				node, ok := astar.Graph[key]

				if !ok {
					continue
				}

				if i - 1 >= 0 {
					key := strconv.Itoa(i-1) + "." + strconv.Itoa(j)
					_, ok := astar.Graph[key]
					if (ok) {
						node.Neighbors[key] = key
					}
				}

				if i + 1 < mapData.Height {
					key := strconv.Itoa(i+1) + "." + strconv.Itoa(j)
					_, ok := astar.Graph[key]
					if (ok) {
						node.Neighbors[key] = key
					}
				}

				if j - 1 >= 0 {
					key := strconv.Itoa(i) + "." + strconv.Itoa(j-1)
					_, ok := astar.Graph[key]
					if (ok) {
						node.Neighbors[key] = key
					}
				}

				if j + 1 < mapData.Width {
					key := strconv.Itoa(i) + "." + strconv.Itoa(j+1)
					_, ok := astar.Graph[key]
					if (ok) {
						node.Neighbors[key] = key
					}
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