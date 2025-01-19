package level_scene

import (
	"d_game/core/gobjects"
	"d_game/core/gscene"
	"d_game/core/resolve_collision"
	utils "d_game/core/utils"
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	gameObject
	angle float64
}

// реализация разброса пуль
// func randRange(min, max int) int {
//     return rand.Intn(max-min) + min
// }

func newBullet(x float64, y float64, mouseX int, mouseY int) *Bullet {
	angle := math.Atan2(float64(mouseY)-y, float64(mouseX)-x)
	object := resolve_collision.NewObject(x-5/2+math.Cos(angle)*10*2, y-5/2+math.Sin(angle)*10*2, 5, 5, "bullet")
	object.SetShape(resolve_collision.NewRectangle(0, 0, 5, 5))

	return &Bullet{
		gameObject: gobjects.Object{
			Type:      gobjects.Bullet,
			IsDeleted: false,
			Sprite:    utils.LoadImage("assets/items/bulletTile.png"),
			Object:    object,
		},
		angle: math.Atan2(float64(mouseY)-y, float64(mouseX)-x),
	}
}

func (b *Bullet) Init(s *gscene.Scene[*LevelController]) {
	s.AddGraphics(b)
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	tileOption := &ebiten.DrawImageOptions{}
	tileOption.GeoM.Translate(-float64(b.Size.X)/2, -float64(b.Size.Y)/2)
	tileOption.GeoM.Rotate(b.angle)
	tileOption.GeoM.Translate(float64(b.Position.X+5/2), float64(b.Position.Y+5/2))
	screen.DrawImage(b.Sprite, tileOption)
}

func (b *Bullet) Update(delta float64) {
	b.Object.Update()
	if b.IsDeleted {
		return
	}
	dx := math.Cos(b.angle) * 4
	dy := math.Sin(b.angle) * 4

	// var collission *resolve_collision.Collision

	if check := b.Check(math.Cos(b.angle)*4, 0, "solid"); check != nil {
		b.IsDeleted = true
		fmt.Println("solid")
		dx = check.ContactWithObject(check.Objects[0]).X
		dy = check.ContactWithObject(check.Objects[0]).Y
	}

	if check := b.Check(0, math.Sin(b.angle)*4, "solid"); check != nil {
		b.IsDeleted = true
		fmt.Println("solid")
		dx = check.ContactWithObject(check.Objects[0]).X
		dy = check.ContactWithObject(check.Objects[0]).Y
	}

	if intersection := b.Shape.Intersection(math.Cos(b.angle)*4, math.Sin(b.angle)*4, sceneState.Policeman.Shape); intersection != nil {

		// We are colliding with the stairs shape, so we can move according
		// to the delta (MTV) to get out of it.
		dx = math.Cos(b.angle) * 4
		dy = math.Sin(b.angle) * 4
		sceneState.Policeman.Sprite = utils.LoadImage("assets/policeman/policeman_dead.png")
		sceneState.Policeman.angle = b.angle + 1.5
		b.IsDeleted = true
		sceneState.Policeman.IsDeleted = true

		// You might want to move a bit less (say, 0.1) than the delta to
		// avoid "bouncing", depending on your application.

	}

	b.Position.X += dx
	b.Position.Y += dy

	ebiten.CursorPosition()
}

func (b *Bullet) IsDisposed() bool {
	return b.IsDeleted
}
