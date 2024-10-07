package gmap

type Map struct {
	Width      int         `json:"width"`
	Height     int         `json:"height"`
	TileHeight int         `json:"tileheight"`
	TileWidth  int         `json:"tilewidth"`
	Layers     []Layer     `json:"layers"`
	TileSet    TileSetType `json:"tileset"`
}

type Layer struct {
	Data []int `json:"data"`
}

type TileSetType struct {
	ImageHeight int    `json:"imageheight"`
	ImageWidth  int    `json:"imagewidth"`
	Tiles       []Tile `json:"tiles"`
}

type Tile struct {
	Id         int              `json:"id"`
	Weight     int              `json:"weight"`
	Properties []TileProperties `json:"properties"`
}

type TileProperties struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value int    `json:"value"`
}
