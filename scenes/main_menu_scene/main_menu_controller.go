package main_menu_scene

import (
	"d_game/core/game"
	"d_game/core/gscene"
	"d_game/scenes/level_scene"
)

type scene = gscene.Scene[*MainMenuController]

type MainMenuController struct {
	ctx *game.Context
}

func NewMainMenuController(ctx *game.Context) *MainMenuController {
	return &MainMenuController{ctx: ctx}
}

func (c *MainMenuController) Init(s *gscene.RootScene[*MainMenuController]) {
	mainMenu := newMainMenu("play", c.ctx.UI)
	s.AddObject(mainMenu)
}

func (c *MainMenuController) Update(delta float64) {

}

func (c *MainMenuController) ChangeScene() {
	game.ChangeScene(c.ctx, level_scene.NewLevelController(c.ctx))
}
