package controls

import "d_game/core/input"

var DefaultKeymap = input.Keymap{
	ActionMoveRight: {
		input.KeyRight,        // Кнопка [>] на клавиатуре
		input.KeyD,            // Кнопка [D] на клавиатуре
		input.KeyGamepadRight, // Кнопка [>] на крестовине контроллера
	},
	ActionMoveDown: {
		input.KeyDown,
		input.KeyS,
		input.KeyGamepadDown,
	},
	ActionMoveLeft: {
		input.KeyLeft,
		input.KeyA,
		input.KeyGamepadLeft,
	},
	ActionMoveUp: {
		input.KeyUp,
		input.KeyW,
		input.KeyGamepadUp,
	},
	ActionShoot: {
		input.KeyMouseLeft,
		input.KeyGamepadX,
	},

	ActionConfirm: {
		input.KeyEnter,
		input.KeyGamepadStart,
	},
	ActionRestart: {
		input.KeyWithModifier(input.KeyR, input.ModControl),
		input.KeyGamepadBack,
	},
}