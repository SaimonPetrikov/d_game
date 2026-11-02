package objects

import (
	"d_game/core/gobjects"
	"d_game/core/gscene"
	"d_game/core/resolve_collision"
	utils "d_game/core/utils"
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet[ControllerAccessor any] struct {
	gameObject
	angle float64
	posX float64
	posY float64
	angleImage float64
	scene *gscene.Scene[*ControllerAccessor]
	state *State[ControllerAccessor]
}

// реализация разброса пуль
// func randRange(min, max int) int {
//     return rand.Intn(max-min) + min
// }

func newBullet[ControllerAccessor any](x float64, y float64, angle float64, state *State[ControllerAccessor]) *Bullet[ControllerAccessor] {	
	angleImage := angle
	object := resolve_collision.NewObject(x-5/2+math.Cos(angleImage)*10*2, y-5/2+math.Sin(angleImage)*10*2, 5, 5, "bullet")
	object.SetShape(resolve_collision.NewRectangle(0, 0, 5, 5))

	return &Bullet[ControllerAccessor]{
		gameObject: gobjects.Object{
			Type:      gobjects.Bullet,
			IsDeleted: false,
			Sprite:    utils.LoadImage("assets/items/bulletTile.png"),
			Object:    object,
		},
		angle: angleImage,
		posX: x - 5/2+math.Cos(angleImage)*10*2,
		posY: y - 5/2+math.Sin(angleImage)*10*2,
		angleImage: angleImage,
		state: state,
	}
}

func (b *Bullet[ControllerAccessor]) Init(s *gscene.Scene[*ControllerAccessor]) {
	b.scene = s
	s.AddGraphics(b)
}

func (b *Bullet[ControllerAccessor]) Draw(screen *ebiten.Image) {
	tileOption := &ebiten.DrawImageOptions{}
	tileOption.GeoM.Translate(-float64(b.Size.X)/2, -float64(b.Size.Y)/2)
	
	tileOption.GeoM.Rotate(b.angleImage)
	
	tileOption.GeoM.Translate(float64(b.posX+5/2), float64(b.posY+5/2))
	tileOption.GeoM.Translate(b.state.Cam.X, b.state.Cam.Y)
	screen.DrawImage(b.Sprite, tileOption)
}

func (b *Bullet[ControllerAccessor]) Update(delta float64) {
	b.Object.Update()
	if b.IsDeleted {
		return
	}
	dx := math.Cos(b.angle) * 4
	dy := math.Sin(b.angle) * 4

	if check := b.Check(math.Cos(b.angle)*4, 0, "solid"); check != nil {
		b.IsDeleted = true
		fmt.Println(1, math.Cos(b.angle)*4)
		if check.Objects[0].Sprite != nil {
			check.Objects[0].TriggerChangeSprite(0.1)
		}

		// angle := 3.0

		// if math.Cos(b.angle) < 0 {
		// 	angle = 0.0
		// }

		// fragments := NewFragments(
		// 	utils.LoadImage("assets/BulletProjectile.png"), 10, b.Position.X + dx, b.Position.Y + dy, angle)
		// b.scene.AddObject(fragments)
	}

	if check := b.Check(0, math.Sin(b.angle)*4, "solid"); check != nil {
		if check.Objects[0].Sprite != nil {
			check.Objects[0].TriggerChangeSprite(0.1)
		}
		b.IsDeleted = true

		// angle := 4.5

		// if math.Sin(b.angle) < 0 {
		// 	angle = 1.5
		// }

		// fragments := NewFragments(
		// 	utils.LoadImage("assets/BulletProjectile.png"), 10, b.Position.X + dx, b.Position.Y + dy, angle)
		// b.scene.AddObject(fragments)
	}

	for _, npc := range b.state.NpcLst {
		if intersection := b.Shape.Intersection(math.Cos(b.angle)*4, math.Sin(b.angle)*4, npc.Shape); intersection != nil {
			fmt.Println("got it")
			npc.Sprite = utils.LoadImage("assets/policeman/policeman_dead.png")
			// npc.angle = b.angle + 1.5
			b.IsDeleted = true
			npc.IsDeleted = true
		}
	}

	b.Position.X += dx
	b.Position.Y += dy
	b.posX += math.Cos(b.angleImage) * 4
	b.posY += math.Sin(b.angleImage) * 4
}

func (b *Bullet[ControllerAccessor]) IsDisposed() bool {
	return b.IsDeleted
}
