package objects

import (
	"d_game/core/astar"
	"d_game/core/gmap"
	"d_game/core/gobjects"
	"d_game/core/gscene"
	"d_game/core/resolve_collision"
	utils "d_game/core/utils"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

// type Npc interface {
// 	pathFind()
// 	setTarget()
// }

// type PatrolledArea struct {
// 	tiles     []maps.TileMapI
// 	startArea maps.TileMapI
// 	nextArea  maps.TileMapI
// }

type Npc[ControllerAccessor any] struct {
	gameObject
	health          int
	path            []astar.Pather
	targetTile      astar.Pather
	indexTargetTile int
	speed           float64
	needFindTarget  bool
	state 			*State[ControllerAccessor]
}

func newNpc[ControllerAccessor any](idStart, idFinish int, state *State[ControllerAccessor]) *Npc[ControllerAccessor] {
	//finish
	from := state.Tiles[idFinish]
	//start
	to := state.Tiles[idStart]
	
	path, _, _ := astar.Path(from.GetTile(), to.GetTile())

	object := resolve_collision.NewObject(path[0].GetPosX(), path[0].GetPosY(), 8, 8, "enemy")
	object.SetShape(resolve_collision.NewRectangle(path[0].GetPosX(), path[0].GetPosY(), 8, 8))
	p := &Npc[ControllerAccessor]{
		health: 1,
		path:   path,
		gameObject: gobjects.Object{
			Sprite: utils.LoadImage("assets/policeman/policeman.png"),
			Object: object,
		},
		speed: 4,
		needFindTarget: true,
		state: state,
	}
	p.targetTile = p.path[0]
	p.indexTargetTile = 1
	return p
}

func (p *Npc[ControllerAccessor]) Init(s *gscene.Scene[*ControllerAccessor]) {
	//todo вынести в event
	s.AddGraphics(p)
}

func (p *Npc[ControllerAccessor]) Draw(screen *ebiten.Image) {
	tileOption := &ebiten.DrawImageOptions{}
	tileOption.GeoM.Translate(-float64(16)/2, -float64(16)/2)

	angle := math.Atan2(p.targetTile.GetPosY()-p.Object.Position.Y, p.targetTile.GetPosX()-p.Object.Position.X)
	tileOption.GeoM.Rotate(angle)

	tileOption.GeoM.Translate(float64(p.Object.Position.X+4), float64(p.Object.Position.Y+4))
	tileOption.GeoM.Translate(p.state.Cam.X, p.state.Cam.Y)

	screen.DrawImage(p.Sprite, tileOption)
}

func (p *Npc[ControllerAccessor]) Update(delta float64) {
	p.Object.Update()
	line := resolve_collision.NewLineFromVector(
		p.Position,
		*p.state.PlayerPosition,
	)
	notSeeYou := false;
	// TODO запускать проверку один раз на тайл в маршруте 
	// возможно ограничивать линию максимальной дистанцией 
	for _, t := range p.state.Tiles {
		if t.IsResolveCollision() {
			if poly, isPoly := t.GetObject().Shape.(*resolve_collision.ConvexPolygon); isPoly {
				for _, otherLine := range poly.Lines() {
					_, ok := line.IntersectionPointsLine(otherLine);
					if !notSeeYou && ok {
						notSeeYou = true
					}
				}
			}
		}
	}

	if !notSeeYou  {
		p.needFindTarget = false
		p.updateTargetTile()
	}

	if notSeeYou { 
		p.needFindTarget = true
	}
	p.targetTile = p.path[p.indexTargetTile]
	if !p.IsDeleted {
		xDiff := math.Abs(p.Object.Position.X - p.targetTile.GetPosX())
		yDiff := math.Abs(p.Object.Position.Y - p.targetTile.GetPosY())
		if xDiff < 3 && yDiff < 3 {
			if p.indexTargetTile < len(p.path)-1 {
				p.indexTargetTile++
				p.targetTile = p.path[p.indexTargetTile]
			}
		}

		if int(p.targetTile.GetPosX()) > int(p.Object.Position.X) {
			p.Object.Position.X += 0.1*p.speed
		} else {
			p.Object.Position.X -= 0.1*p.speed
		}
		if int(p.targetTile.GetPosY()) > int(p.Object.Position.Y) {
			p.Object.Position.Y += 0.1*p.speed
		} else {
			p.Object.Position.Y -= 0.1*p.speed
		}
	}
}



func (p *Npc[ControllerAccessor]) updateTargetTile() {
	npcTileId := gmap.GetTileInSlice(p.Position.X, p.Position.Y)
	playerTileId := *p.state.PlayerTile
	index := strconv.Itoa(npcTileId) + strconv.Itoa(playerTileId)

	path, ok := p.state.PathMap[index]
	if ok {
		p.path = path
		p.targetTile = path[0]
		p.indexTargetTile = 1
		return
	} else {
		from := p.state.Tiles[playerTileId]
		to := p.state.Tiles[npcTileId]
		path, _, _ = astar.Path(from.GetTile(), to.GetTile())
	
		p.path = path
		p.targetTile = path[0]
		p.indexTargetTile = 1
		p.state.PathMap[index] = path
	}
}

func (p *Npc[ControllerAccessor]) IsDisposed() bool {
	return p.IsDeleted
}