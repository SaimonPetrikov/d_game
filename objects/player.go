package objects

import (
	"d_game/config"
	"d_game/controls"
	"d_game/core/gmap"
	"d_game/core/gobjects"
	"d_game/core/gscene"
	"d_game/core/input"
	"d_game/core/resolve_collision"
	utils "d_game/core/utils"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

//todo сделать круг

type squareCamera struct {
	minX, minY, maxX, maxY float64
}

type Player[ControllerAccessor any] struct {
	health int
	level  int
	TileId int
	scene  *gscene.Scene[*ControllerAccessor]
	state  *State[ControllerAccessor]
	*input.Handler
	gameObject
	angle float64
	squareCamera squareCamera
}

func newPlayer[ControllerAccessor any](handler *input.Handler, state *State[ControllerAccessor]) *Player[ControllerAccessor] {
	p := &Player[ControllerAccessor]{
		health: 10,
		level:  1,
		gameObject: gobjects.Object{
			Sprite: utils.LoadImage("assets/player/Handgun1.png"),
			Object: resolve_collision.NewObject(20, 20, 16, 16, "player"),
		},
		Handler: handler,
		state: state,
	}

	p.squareCamera = squareCamera{
		minX: p.Position.X - 50,
		minY: p.Position.Y - 50,
		maxX: p.Position.X + 50,
		maxY: p.Position.Y + 50,
	}

	// sceneState.Player = p

	return p
}

func (p *Player[ControllerAccessor]) Init(s *gscene.Scene[*ControllerAccessor]) {
	//todo вынести в event
	p.scene = s
	s.AddGraphics(p)
}

func (p *Player[ControllerAccessor]) Draw(screen *ebiten.Image) {
	tileOption := &ebiten.DrawImageOptions{}

	tileOption.GeoM.Translate(-p.Size.X/2, -p.Size.Y/2)

	mouseVector := utils.CursorPosition()
	vectorPlayer := utils.Vector{
		X: p.Position.X + p.Size.X/2 + p.state.Cam.X,
		Y: p.Position.Y + p.Size.Y/2 + p.state.Cam.Y,
	}
	distance := utils.GetDistance(mouseVector, vectorPlayer)
	p.angle = utils.Atan2(distance)
	tileOption.GeoM.Rotate(p.angle)
	
	tileOption.GeoM.Translate(p.Position.X+p.Size.X/2, p.Position.Y+p.Size.Y/2)
	tileOption.GeoM.Translate(p.state.Cam.X, p.state.Cam.Y)
	screen.DrawImage(p.Sprite, tileOption)
}

func (p *Player[ControllerAccessor]) Update(delta float64) {
	speed := 50 * delta

	if p.ActionIsPressed(controls.ActionMoveRight) {
		p.Speed.X = speed
		dx := p.Speed.X
		if p.Object.Check(p.Speed.X, 0, config.Solid) != nil {
			dx = 0
			p.Speed.X = 0
		}
		p.Object.Position.X += dx
	}
	if p.ActionIsPressed(controls.ActionMoveDown) {
		p.Speed.Y = speed
		dy := p.Speed.Y
		if p.Object.Check(0, p.Speed.Y, config.Solid) != nil {
			dy = 0
			p.Speed.Y = 0
		}
		p.Object.Position.Y += dy
	}
	if p.ActionIsPressed(controls.ActionMoveLeft) {
		p.Speed.X = speed
		dx := p.Speed.X
		if p.Object.Check(-p.Speed.X, 0, config.Solid) != nil {
			dx = 0
			p.Speed.X = 0
		}
		p.Object.Position.X -= dx
	}
	if p.ActionIsPressed(controls.ActionMoveUp) {
		p.Speed.Y = speed
		dy := p.Speed.Y
		if p.Object.Check(0, -p.Speed.Y, config.Solid) != nil {
			dy = 0
			p.Speed.Y = 0
		}
		p.Object.Position.Y -= dy
	}

	if p.ActionIsJustPressed(controls.ActionShoot) {
		//todo вынести в event
		bullet := newBullet(p.Position.X + 8, p.Position.Y + 8, p.angle, p.state)

		p.state.AddToSpace(bullet.Object)

		// for _, v := range p.s.Controller().ctx.Space.Objects() {
		// 	if v.HasTags("enemy") {
		// 		fmt.Println(v.BoundsToSpace(0, 0))
		// 	}
		// }
		p.scene.AddObject(bullet)

		p.Sprite = utils.LoadImage("assets/player/Handgun2.png")

	} else {
		p.Sprite = utils.LoadImage("assets/player/Handgun1.png")
	}

	if p.Position.X < p.squareCamera.minX || p.Position.Y < p.squareCamera.minY {
		p.state.Cam.FollowTarget(p.Position.X+p.Size.X/2,p.Position.Y+p.Size.Y/2, 320, 240)
		p.squareCamera = squareCamera{
			minX: p.Position.X - 50,
			minY: p.Position.Y - 50,
			maxX: p.Position.X + 50,
			maxY: p.Position.Y + 50,
		}
	}

	if p.Position.X > p.squareCamera.maxX || p.Position.Y > p.squareCamera.maxY {
		p.state.Cam.FollowTarget(p.Position.X+p.Size.X/2,p.Position.Y+p.Size.Y/2, 320, 240)
		p.squareCamera = squareCamera{
			minX: p.Position.X - 50,
			minY: p.Position.Y - 50,
			maxX: p.Position.X + 50,
			maxY: p.Position.Y + 50,
		}
	}

	p.state.Cam.Update()

	if check := p.Check(p.Size.X/2, p.Size.Y/2, "floor"); check != nil {
		p.TileId = gmap.GetTileInSlice(check.Objects[0].Position.X, check.Objects[0].Position.Y)
		fmt.Println(p.TileId)
	}
}


// func(p *Player) Move(speed float64, direct string) {
// 	p.Speed.(direct) = speed
// 	dx := p.Speed.X
// 	if check := p.Object.Check(p.Speed.X, p.Speed.Y, "solid"); check != nil {
// 		fmt.Println("right")
// 		dx = 0
// 		p.Speed.X = 0
// 	}
// 	p.Object.Position.X += dx
// }

func (p *Player[ControllerAccessor]) IsDisposed() bool {
	return false
}
