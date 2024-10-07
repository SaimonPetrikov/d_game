package main

import (
	"d_game/controls"
	"d_game/core/game"
	"d_game/core/input"
	"d_game/core/resolve_collision"
	mainmenuscene "d_game/scenes/main_menu_scene"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type state struct {
	screenId int
}

type Game struct {
	inputSystem input.System
	ctx *game.Context
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
	context.Space =  resolve_collision.NewSpace(int(480), int(320), 16, 16)

	game.ChangeScene(context, mainmenuscene.NewMainMenuController(context))

	err := ebiten.RunGame(g)
	if err != nil {
		log.Print(err)
	}
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	g.inputSystem.Update()
	g.ctx.CurrentScene().UpdateWithDelta(1.0 / 60.0)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprint(ebiten.ActualTPS()), 820, 750)
	g.ctx.CurrentScene().Draw(screen)

	for _, o := range g.ctx.Space.Objects() {
		if (o.HasTags("enemy")) {
			drawColor := color.RGBA{180, 100, 0, 255}
			vector.DrawFilledRect(screen, float32(o.Position.X), float32(o.Position.Y), float32(o.Size.X), float32(o.Size.Y), drawColor, false)
		}
		if (o.HasTags("bullet")) {
			drawColor := color.RGBA{180, 100, 0, 255}
			vector.DrawFilledRect(screen, float32(o.Position.X), float32(o.Position.Y), float32(o.Size.X), float32(o.Size.Y), drawColor, false)
		}
	}
}
