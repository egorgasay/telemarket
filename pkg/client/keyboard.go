package client

// InlineKeyboardMarkup represents an inline keyboard.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]KeyboardButton `json:"inline_keyboard"`
}

// NewKeyboardWithMarkup creates a new keyboard with the given rows.
func NewKeyboardWithMarkup(rows ...[]KeyboardButton) InlineKeyboardMarkup {
	var keyboard = make([][]KeyboardButton, 0)
	keyboard = append(keyboard, rows...)

	return InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}
}

// NewKeyboardRow creates a new row of buttons.
func NewKeyboardRow(buttons ...KeyboardButton) []KeyboardButton {
	var row = make([]KeyboardButton, 0)
	row = append(row, buttons...)

	return row
}
