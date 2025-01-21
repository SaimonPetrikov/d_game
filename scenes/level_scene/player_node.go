package level_scene

import (
	"d_game/controls"
	"d_game/core/gobjects"
	"d_game/core/resolve_collision"
	utils "d_game/core/utils"
	"fmt"
	"math"

	"d_game/core/gscene"
	"d_game/core/input"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	health int
	level  int
	s      *scene
	*input.Handler
	gameObject
}

func newPlayer() *Player {
	p := &Player{
		health: 10,
		level:  1,
		gameObject: gobjects.Object{
			Sprite: utils.LoadImage("assets/player/Handgun1.png"),
			Object: resolve_collision.NewObject(20, 20, 16, 16, "player"),
		},
	}

	return p
}

func (p *Player) Init(s *gscene.Scene[*LevelController]) {
	//todo вынести в event
	p.s = s
	p.Handler = s.Controller().ctx.Input
	s.AddGraphics(p)
}

func (p *Player) Draw(screen *ebiten.Image) {
	tileOption := &ebiten.DrawImageOptions{}

	mouseX, mouseY := ebiten.CursorPosition()

	tileOption.GeoM.Translate(-float64(16)/2, -float64(16)/2)
	// Вычисление угла между персонажем и курсором
	angle := math.Atan2(float64(mouseY)-(p.Position.Y+8), float64(mouseX)-(p.Position.X+8))

	// Отображение персонажа на экране

	// op.GeoM.Translate(-float64(characterImage.Bounds().Dx())/2, -float64(characterImage.Bounds().Dy())/2)
	tileOption.GeoM.Rotate(angle)
	tileOption.GeoM.Translate(float64(p.Position.X+8), float64(p.Position.Y+8))
	screen.DrawImage(p.Sprite, tileOption)
}

func (p *Player) Update(delta float64) {
	speed := 50 * delta

	if p.ActionIsPressed(controls.ActionMoveRight) {
		p.Speed.X = speed
		dx := p.Speed.X
		if check := p.Object.Check(p.Speed.X, 0, "solid"); check != nil {
			fmt.Println("right")
			dx = 0
			p.Speed.X = 0
		}
		p.Object.Position.X += dx
	}
	if p.ActionIsPressed(controls.ActionMoveDown) {
		p.Speed.Y = speed
		dy := p.Speed.Y
		if check := p.Object.Check(0, p.Speed.Y, "solid"); check != nil {
			fmt.Println("down")
			dy = 0
			p.Speed.Y = 0
		}
		p.Object.Position.Y += dy
	}
	if p.ActionIsPressed(controls.ActionMoveLeft) {
		p.Speed.X = speed
		dx := p.Speed.X
		if check := p.Object.Check(-p.Speed.X, 0, "solid"); check != nil {
			fmt.Println("left")
			dx = 0
			p.Speed.X = 0
		}
		p.Object.Position.X -= dx
	}
	if p.ActionIsPressed(controls.ActionMoveUp) {
		p.Speed.Y = speed
		dy := p.Speed.Y
		if check := p.Object.Check(0, -p.Speed.Y, "solid"); check != nil {
			fmt.Println("up")
			dy = 0
			p.Speed.Y = 0
		}
		p.Object.Position.Y -= dy
	}

	if p.ActionIsPressed(controls.ActionShoot) {
		//todo вынести в event
		mouseX, mouseY := ebiten.CursorPosition()
		bullet := newBullet(p.Position.X+8, p.Position.Y+8, mouseX, mouseY)

		p.s.Controller().ctx.Space.Add(bullet.Object)

		// for _, v := range p.s.Controller().ctx.Space.Objects() {
		// 	if v.HasTags("enemy") {
		// 		fmt.Println(v.BoundsToSpace(0, 0))
		// 	}
		// }
		p.s.AddObject(bullet)

		p.Sprite = utils.LoadImage("assets/player/Handgun2.png")

	} else {
		p.Sprite = utils.LoadImage("assets/player/Handgun1.png")
	}

	ebiten.CursorPosition()
}

func (p *Player) IsDisposed() bool {
	return false
}
