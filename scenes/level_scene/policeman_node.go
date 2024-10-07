package level_scene

import (
	"d_game/core/resolve_collision"
	utils "d_game/core/utils"

	"d_game/core/gscene"

	"github.com/hajimehoshi/ebiten/v2"
)

type Policeman struct {
	health int
	angle float64
	sprite *ebiten.Image
	IsDeleted bool
	Object *resolve_collision.Object
}

func newPoliceman() *Policeman {
	object := resolve_collision.NewObject(50-2, 50-4, 5, 5, "enemy")
	object.SetShape(resolve_collision.NewRectangle(0, 0, 5, 5))
	return &Policeman{
		health: 1,
		sprite: utils.LoadImage("assets/policeman/policeman.png"),
		Object: object,
	}
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
	screen.DrawImage(p.sprite, tileOption)
	
}

func (p *Policeman) Update(delta float64) {
	p.Object.Update()
	if !p.IsDeleted {
		p.Object.Position.X += 0.1
		p.Object.Position.Y += 0.1
	}

	// if check := p.Object.Check(0, 0, "bullet"); check != nil {
	// 	sceneState.Policeman.sprite = utils.LoadImage("assets/policeman/policeman_dead.png")
	// 	// sceneState.Policeman.angle = b.angle + 1.5
	// 	// b.IsDeleted = true
	// }

}

func (p *Policeman) IsDisposed() bool {
	return p.IsDeleted
}