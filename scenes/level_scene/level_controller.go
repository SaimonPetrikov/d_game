package level_scene

import (
	"d_game/core/game"
	"d_game/core/gobjects"
	"d_game/core/gscene"
)

type scene = gscene.Scene[*LevelController]
type gameObject = gobjects.Object
type ObjectType = gobjects.ObjectType

type LevelController struct {
	ctx *game.Context
}

var sceneState = SceneState{}

func NewLevelController(ctx *game.Context) *LevelController {
	return &LevelController{ctx: ctx}
}

func (c *LevelController) Init(s *gscene.RootScene[*LevelController]) {
	levelMap := newMap()
	s.AddObject(levelMap)

	player := newPlayer()
	c.ctx.Space.Add(player.Object)
	s.AddObject(player)

	
	sceneState.Policeman = newPoliceman()
	c.ctx.Space.Add(sceneState.Policeman.Object)
	s.AddObject(sceneState.Policeman)

}

func (c *LevelController) Update(delta float64) {

}