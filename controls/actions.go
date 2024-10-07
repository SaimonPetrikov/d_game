package controls

import "d_game/core/input"

const (
	ActionNone input.Action = iota

	ActionMoveRight
	ActionMoveDown
	ActionMoveLeft
	ActionMoveUp
	ActionShoot

	// Эти действия понадобятся позднее.
	ActionConfirm
	ActionRestart
)