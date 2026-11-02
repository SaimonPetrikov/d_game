package main

import (
	"d_game/config"
	"d_game/controls"
	"d_game/core/game"
	"d_game/core/input"
	"d_game/core/resolve_collision"
	mainmenuscene "d_game/scenes/main_menu_scene"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type state struct {
	screenId int
}

type Game struct {
	inputSystem input.System
	ctx         *game.Context
}

func main() {
	ebiten.SetWindowSize(900, 800)
	ebiten.SetWindowTitle("Ebiten UI Demo")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	// ebiten.SetScreenClearedEveryFrame(false)
	// ebiten.SetVsyncEnabled(false)

	context := game.NewContext()

	g := &Game{
		ctx: context,
	}

	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	context.Input = g.inputSystem.NewHandler(0, controls.DefaultKeymap)
	context.Space = resolve_collision.NewSpace(config.SpaceWidth, config.SpaceHigh, config.CellWidth, config.CellHigh)

	game.ChangeScene(context, mainmenuscene.NewMainMenuController(context))

	err := ebiten.RunGame(g)
	if err != nil {
		log.Print(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 340, 240
}

func (g *Game) Update() error {
	g.inputSystem.Update()
	g.ctx.CurrentScene().UpdateWithDelta(config.UpdateWithDelta)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprint(ebiten.ActualTPS()), 600, 750)
	g.ctx.CurrentScene().Draw(screen)
}
