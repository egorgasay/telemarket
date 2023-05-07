package client

// KeyboardButton represents a button of the keyboard.
type KeyboardButton struct {
	Text                         string  `json:"text"`
	URL                          *string `json:"url,omitempty"`
	CallbackData                 *string `json:"callback_data,omitempty"`
	SwitchInlineQuery            *string `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat *string `json:"switch_inline_query_current_chat,omitempty"`
}

// NewKeyboardButtonWithData creates a new KeyboardButton.
func NewKeyboardButtonWithData(text, data string) KeyboardButton {
	return KeyboardButton{
		Text:         text,
		CallbackData: &data,
	}
}

// NewKeyboardButtonURL creates a new KeyboardButton with URL.
func NewKeyboardButtonURL(text, url string) KeyboardButton {
	return KeyboardButton{
		Text: text,
		URL:  &url,
	}
}
