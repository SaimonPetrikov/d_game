package main_menu_scene

import (
	"d_game/ui_container"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
)

type MainMenu struct {
	textButton string
	ui  *ebitenui.UI
}

func newMainMenu(textButton string, ui *ebitenui.UI) *MainMenu {
	return &MainMenu{textButton: textButton, ui: ui}
}

func (menu *MainMenu) Init(s *scene) {
	s.AddGraphics(menu)

	ui, closeUI, err := ui_container.CreateMainMenu(s.Controller().ctx)
	if err != nil {
		log.Fatal(err)
	}

	menu.ui = ui

	defer closeUI()
}

func (menu *MainMenu) Draw(screen *ebiten.Image) {
	menu.ui.Draw(screen)
}

func (menu *MainMenu) IsDisposed() bool {
	return false
}

func (menu *MainMenu) Update(delta float64) {
	menu.ui.Update()
}

func (menu *MainMenu) ChangeScene(s *scene) {
	s.Controller().ChangeScene()
}