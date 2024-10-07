package ui_container

import (
	"d_game/core/game"
	"d_game/scenes/level_scene"
	"fmt"
	_ "image/png"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
)

func newRowLayoutContainerWithMinWidth(minWidth, spacing int, rowScale []bool) *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
			StretchVertical:   true,
		})),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.MinSize(minWidth, 0)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch([]bool{true}, rowScale),
			widget.GridLayoutOpts.Spacing(spacing, spacing),
		)),
	)
}

func CreateMainMenu(context *game.Context) (*ebitenui.UI, func(), error) {
	res, err := NewUIResources()
	if err != nil {
		return nil, nil, err
	}

	rootContainer := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
		})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()))

	rowContainer := newRowLayoutContainerWithMinWidth(260, 10, nil)
	rootContainer.AddChild(rowContainer)

	rowContainer.AddChild(NewButton(res, "play", func() {
		fmt.Println("play")
		game.ChangeScene(context, level_scene.NewLevelController(context))
	}))

	footerContainer := widget.NewContainer(widget.ContainerOpts.Layout(widget.NewRowLayout(
		widget.RowLayoutOpts.Padding(widget.Insets{
			Left:  25,
			Right: 25,
		}),
	)))
	rowContainer.AddChild(footerContainer)

	ui := &ebitenui.UI{
		Container: rootContainer,
	}

	return ui, func() {
		res.close()
	}, nil
}
