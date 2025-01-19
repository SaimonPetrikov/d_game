package level_scene

import (
	"d_game/core/astar"
	"d_game/core/gobjects"
	"d_game/core/gscene"
	"d_game/core/resolve_collision"
	utils "d_game/core/utils"
	"fmt"

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
	path []string
	targetTile string
	indexTargetTile int
}

func newPoliceman() *Policeman {
	path := astar.FindPath("2.3", "14.3")
	fmt.Println(path)
	indexStartTile := path[0]
	object := resolve_collision.NewObject(float64(astar.Graph[indexStartTile].Sx), float64(astar.Graph[indexStartTile].Sy), 5, 5, "enemy")
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
			if astar.Graph[p.targetTile].Sx > int(p.Object.Position.X) {
				p.Object.Position.X += 0.1
			} else {
				p.Object.Position.X -= 0.1
			}
			if astar.Graph[p.targetTile].Sy > int(p.Object.Position.Y) {
				p.Object.Position.Y += 0.1
			} else {
				p.Object.Position.Y -= 0.1
			}
		if int(p.Object.Position.X+0.001) == astar.Graph[p.targetTile].Sx && int(p.Object.Position.Y+0.001) == astar.Graph[p.targetTile].Sy {
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