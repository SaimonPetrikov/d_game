package enemys

import (
	"d_game/core/astar"
	"d_game/core/gobjects"
)

type Weapon struct {
	damage int
}

type Enemy struct {
    Weapon
    gobjects.Object
    health int
    angle float64
    path []astar.Pather
    targetTile astar.Pather
    indexTargetTile int
}

func NewEnemy(damage int) *Enemy {

	return &Enemy{
        Weapon: Weapon{damage: damage},
        //...
    }
}