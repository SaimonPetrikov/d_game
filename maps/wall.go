package maps

import (
	"d_game/core/resolve_collision"
	"d_game/core/utils"
	_ "embed"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed shaders/wall_damaged.go
	wall_damaged_go []byte
)


type Wall struct {
	rootTile *RootTileMap
	Object *resolve_collision.Object
}

func NewWall(rootTile *RootTileMap) *Wall {
	x, y, w, h := rootTile.GetPosAndSize()
	object := resolve_collision.NewObject(x, y, w, h, "solid")
	object.SetSprite(1.0, rootTile.Image)
	object.SetAction(func(hp float64, damage float64) (float32, float64) {
		hp -= damage
		if hp <= 0 {
			return float32(0), float64(0)
		}

		return float32(hp / 1), hp
	})
	shader, _ := ebiten.NewShader(wall_damaged_go)
	object.SetShader(shader, utils.LoadImage("assets/masks/wall/" + strconv.Itoa(1+rand.Intn(3)) + ".png"), map[string]any{"HP": 1.0})
	return &Wall{rootTile: rootTile, Object: object}
}

func (wall *Wall) Draw(screen *ebiten.Image) {
	if wall.Object.Sprite.Shader == nil {
        var options ebiten.DrawImageOptions
        options.GeoM.Translate(wall.rootTile.PosX, wall.rootTile.PosY)
        screen.DrawImage(wall.rootTile.Image, &options)
        return
    }

    var options ebiten.DrawRectShaderOptions
    options.GeoM.Translate(wall.rootTile.PosX, wall.rootTile.PosY)
    options.Images[0] = wall.rootTile.Image // Src0
    options.Images[1] = wall.Object.Sprite.ShaderTexture // Src1
    options.Uniforms = wall.Object.Sprite.ShaderParams
    b := wall.rootTile.Image.Bounds()
    screen.DrawRectShader(b.Dx(), b.Dy(), wall.Object.Sprite.Shader, &options)
}

func (wall *Wall) GetObject() *resolve_collision.Object {
	return wall.Object
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