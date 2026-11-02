package level_scene

import (
	"d_game/camera"
	"d_game/core/astar"
	"d_game/core/game"
	"d_game/core/gobjects"
	"d_game/core/gscene"
	"d_game/objects"
)

type scene = gscene.Scene[*LevelController]
type gameObject = gobjects.Object
type ObjectType = gobjects.ObjectType

type LevelController struct {
	ctx *game.Context
}

var sceneState = &objects.State[LevelController]{}

func NewLevelController(ctx *game.Context) *LevelController {
	return &LevelController{ctx: ctx}
}

func (c *LevelController) Init(s *gscene.RootScene[*LevelController]) {
	sceneState.PathMap = make(map[string][]astar.Pather)
	sceneState.Cam = camera.NewCamera(20, 20)

	levelMap := newMap()
	sceneState.Tiles = levelMap.levelMap.TileSet
	s.AddObject(levelMap)


	factory := objects.NewFactory(s)
	factory.State = sceneState
	factory.State.Space = c.ctx.Space

	player := factory.CreatePlayer(c.ctx.Input)
	factory.State.AddToSpace(player.Object)
	s.AddObject(player)

	factory.State.PlayerPosition = &player.Position
	factory.State.PlayerTile = &player.TileId

	npc1 := factory.CreatePoliceman(41, 31)
	factory.State.NpcLst = append(factory.State.NpcLst, npc1)
	factory.State.AddToSpace(npc1.Object)
	s.AddObject(npc1)

	npc2 := factory.CreatePoliceman(41, 31)
	factory.State.NpcLst = append(factory.State.NpcLst, npc2)
	factory.State.AddToSpace(npc2.Object)
	s.AddObject(npc2)
}

func (c *LevelController) Update(delta float64) {

}