package gobjects

import (
	"d_game/core/resolve_collision"
	"github.com/hajimehoshi/ebiten/v2"
)

type CollisionObject *resolve_collision.Object

//todo solve 
type Triger interface {
	Update()
}

type Object struct {
	Type      ObjectType
	IsDeleted bool
	Sprite    *ebiten.Image
	Speed     resolve_collision.Vector
	*resolve_collision.Object
}

func (o Object) Update() {

}

