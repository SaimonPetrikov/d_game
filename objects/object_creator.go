package objects

import (
	"d_game/core/gscene"
	"d_game/core/input"
)

type FactoryObject[ControllerAccessor any] struct{
	State *State[ControllerAccessor]
	scene *gscene.RootScene[*ControllerAccessor]
}

func NewFactory[ControllerAccessor any](scene *gscene.RootScene[*ControllerAccessor]) FactoryObject[ControllerAccessor] {
	return FactoryObject[ControllerAccessor]{
		State: &State[ControllerAccessor]{},
		scene: scene,
	}
}

func (f FactoryObject[ControllerAccessor]) CreatePoliceman(idStart, idFinish int) *Npc[ControllerAccessor] {
	return newNpc(idStart, idFinish, f.State)
}

func (f FactoryObject[ControllerAccessor]) CreateBullet(x float64, y float64, angle float64) *Bullet[ControllerAccessor] {
	return newBullet(x, y, angle, f.State)
}

func (f FactoryObject[ControllerAccessor]) CreatePlayer(handler *input.Handler) *Player[ControllerAccessor] {
	return newPlayer(handler, f.State)
}

