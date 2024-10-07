package game

import (
	"d_game/core/gscene"
	"d_game/core/input"
	"d_game/core/resolve_collision"

	"github.com/ebitenui/ebitenui"
	"github.com/quasilyte/gmath"
)

type Context struct {
	Rand gmath.Rand
	UI *ebitenui.UI
	Input  *input.Handler
	Space *resolve_collision.Space

	WindowWidth  int
	WindowHeight int

	scene gscene.GameRunner
}

func ChangeScene[ControllerAccessor any](ctx *Context, c gscene.Controller[ControllerAccessor]) {
	s := gscene.NewRootScene[ControllerAccessor](c)
	ctx.scene = s
}

func NewContext() *Context {
	return &Context{}
}

func (ctx *Context) CurrentScene() gscene.GameRunner {
	return ctx.scene
}
