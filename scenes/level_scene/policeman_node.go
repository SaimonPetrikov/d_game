package level_scene

import (
	"d_game/core/astar"
	"d_game/core/gobjects"
	"d_game/core/gscene"
	"d_game/core/resolve_collision"
	utils "d_game/core/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Npc interface {
	pathFind()
	setTarget()
}

// type PatrolledArea struct {
// 	tiles []maps.TileMapI
// 	startArea maps.TileMapI
// 	nextArea maps.TileMapI
// }

type Policeman struct {
	gameObject
	health int
	angle float64
	path []astar.Pather
	targetTile astar.Pather
	indexTargetTile int
}

func newPoliceman() *Policeman {
	from := sceneState.Tiles[31]
	to := sceneState.Tiles[41]
	path, _, _ := astar.Path(from.GetTile(), to.GetTile())

	//todo: position to center of tiles
	object := resolve_collision.NewObject(float64(path[0].GetPosX()), path[0].GetPosY(), 5, 5, "enemy")
	object.SetShape(resolve_collision.NewRectangle(0, 0, 5, 5))
	p := &Policeman{
		health: 1,
		path: path,
		gameObject: gobjects.Object{
			Sprite: utils.LoadImage("assets/policeman/policeman.png"),
			Object: object,
		},
	}
	p.targetTile = p.path[0]
	p.indexTargetTile = 1
	return p
}

func (p *Policeman) Init(s *gscene.Scene[*LevelController]) {
	//todo вынести в event
	s.AddGraphics(p)
}

func (p *Policeman) Draw(screen *ebiten.Image) {
	tileOption := &ebiten.DrawImageOptions{}
	tileOption.GeoM.Translate(-float64(p.Object.Size.X)/2, -float64(p.Object.Size.Y)/2)
	// // tileOption.GeoM.Scale(1.8, 1.8)
	tileOption.GeoM.Rotate(p.angle)
	tileOption.GeoM.Translate(float64(p.Object.Position.X + 1), float64(p.Object.Position.Y - 2))
	screen.DrawImage(p.Sprite, tileOption)
	
}

func (p *Policeman) Update(delta float64) {
	p.Object.Update()
	p.targetTile = p.path[p.indexTargetTile]
	if !p.IsDeleted {
			if int(p.targetTile.GetPosX()) > int(p.Object.Position.X) {
				p.Object.Position.X += 0.1
			} else {
				p.Object.Position.X -= 0.1
			}
			if int(p.targetTile.GetPosY()) > int(p.Object.Position.Y) {
				p.Object.Position.Y += 0.1
			} else {
				p.Object.Position.Y -= 0.1
			}
		if int(p.Object.Position.X+0.001) == int(p.targetTile.GetPosX()) && int(p.Object.Position.Y+0.001) == int(p.targetTile.GetPosY()) {
			if p.indexTargetTile < len(p.path) - 1 {
				p.indexTargetTile++
				p.targetTile = p.path[p.indexTargetTile]
			}
			if p.indexTargetTile < len(p.path) - 1 {
				p.indexTargetTile++
				p.targetTile = p.path[p.indexTargetTile]
			}
		}
	}
}

func (p *Policeman) IsDisposed() bool {
	return p.IsDeleted
}