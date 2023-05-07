package client

import (
	"encoding/json"
	"fmt"
)

// Chat describes the chat with user.
type Chat struct {
	ID int `json:"id"`
}

// From describes the user who sent the message.
type From struct {
	Username string `json:"username"`
}

// IncomingMessage describes the incoming message.
type IncomingMessage struct {
	Text string `json:"text"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
}

// UpdateConfig describes the update configuration.
type UpdateConfig struct {
	Offset  int
	Limit   int
	Timeout int
}

// Update describes the update event.
type Update struct {
	Message       *Message       `json:"message,omitempty"`
	EditedMessage *Message       `json:"edited_message,omitempty"`
	InlineQuery   *InlineQuery   `json:"inline_query,omitempty"`
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
	UpdateID      int            `json:"update_id"`
}

// InlineQuery describes the inline query.
type InlineQuery struct {
	ID       string `json:"id"`
	From     *From  `json:"from"`
	Query    string `json:"query"`
	Offset   string `json:"offset"`
	ChatType string `json:"chat_type"`
}

// CallbackQuery describes the callback query.
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *From    `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID string   `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data,omitempty"`
	GameShortName   string   `json:"game_short_name,omitempty"`
}

// Message describes the message.
type Message struct {
	MessageID   int                   `json:"message_id"`
	From        *From                 `json:"from,omitempty"`
	Date        int                   `json:"date"`
	Chat        *Chat                 `json:"chat"`
	Text        string                `json:"text,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// IsCommand returns true if the message is a command.
func (m Message) IsCommand() bool {
	return m.Text != "" && m.Text[0] == '/'
}

// Response describes the response from telegram API.
type Response struct {
	Ok        bool            `json:"ok"`
	Result    json.RawMessage `json:"result,omitempty"`
	ErrorCode int             `json:"error_code,omitempty"`
}

// EditMessageTextConfig describes the edit of message text configuration.
type EditMessageTextConfig struct {
	BaseEdit
	Text                  string
	ParseMode             string
	Entities              []MessageEntity
	DisableWebPagePreview bool
}

// getEndpoint returns the endpoint of the telegram API.
func (e EditMessageTextConfig) getEndpoint() string {
	return "editMessageText"
}

// getParams returns the parameters of the telegram API.
func (e EditMessageTextConfig) getParams() map[string]string {
	params := make(map[string]string)
	params["text"] = e.Text
	params["parse_mode"] = e.ParseMode
	params["disable_web_page_preview"] = fmt.Sprintf("%v", e.DisableWebPagePreview)
	params["reply_markup"] = jsonAnything(e.ReplyMarkup)
	params["entities"] = jsonAnything(e.Entities)
	params["chat_id"] = fmt.Sprintf("%d", e.ChatID)
	params["message_id"] = fmt.Sprintf("%d", e.MessageID)

	return params
}

// BaseEdit is base type of all chat edits.
type BaseEdit struct {
	ChatID          int64
	ChannelUsername string
	MessageID       int
	InlineMessageID string
	ReplyMarkup     *InlineKeyboardMarkup
}

// MessageEntity represents one special entity in a text message.
type MessageEntity struct {
	Type     string `json:"type"`
	Offset   int    `json:"offset"`
	Length   int    `json:"length"`
	URL      string `json:"url,omitempty"`
	User     *From  `json:"user,omitempty"`
	Language string `json:"language,omitempty"`
}

// MessageConfig describes the message configuration.
type MessageConfig struct {
	BaseChat
	Text                  string
	ParseMode             string
	Entities              []MessageEntity
	DisableWebPagePreview bool
}

// getEndpoint returns the endpoint of the telegram API.
func (m MessageConfig) getEndpoint() string {
	return "sendMessage"
}

// getParams returns the parameters of the telegram API.
func (m MessageConfig) getParams() map[string]string {
	params := map[string]string{
		"text":                 m.Text,
		"chat_id":              fmt.Sprintf("%d", m.ChatID),
		"reply_markup":         jsonAnything(m.BaseChat.ReplyMarkup),
		"disable_notification": fmt.Sprintf("%v", m.DisableNotification),
		"entities":             jsonAnything(m.Entities),
	}

	return params
}

// BaseChat is base type of all chat edits.
type BaseChat struct {
	ChatID                   int64
	ChannelUsername          string
	ReplyToMessageID         int
	ReplyMarkup              interface{}
	DisableNotification      bool
	AllowSendingWithoutReply bool
}
