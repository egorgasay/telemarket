package bot

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (b *Bot) handleCommand(msg *tgapi.Message) {
	switch msg.Command() {
	case "start":
		b.handleStart(msg)
	}
}

func (b *Bot) handleMessage() {

}

func (b *Bot) handleStart(msg *tgapi.Message) {
	msgConfig := tgapi.NewMessage(msg.Chat.ID, startMessage)

	msgConfig.ReplyMarkup = startKeyboard
	_, err := b.Send(msgConfig)
	if err != nil {
		log.Println("send error: ", err)
	}
}

func (b *Bot) handleCallbackQuery(query *tgapi.CallbackQuery) {
	//msg.ParseMode = tgapi.ModeHTML

	markup := tgapi.NewInlineKeyboardMarkup()
	text := query.Data

	switch text {
	case items:
		text = itemsMessage
		markup = tgapi.NewInlineKeyboardMarkup(itemButtons...)
	case height:
		text = heightFeedbackMessage
		markup = heightKeyboard
	case address:
		markup = addressKeyboard
		text = addressFeedbackMessage
	case start:
		markup = startKeyboard
		text = startMessage
	case feedBack:
		markup = feedBackKeyboard
		text = feedbackMessage
	case thxFeedback:
		markup = thxFeedbackKeyboard
		text = thxFeedbackMessage
	case sorryFeedback:
		markup = sorryFeedbackKeyboard
		text = sorryFeedbackMessage
	case sorryHeight:
		markup = heightKeyboard
		text = sorryHeightMessage
	case sSize, mSize, lSize:
		markup = thxFeedbackKeyboard
		text = "Ваш " + text
	}

	msg := tgapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, text, markup)
	if _, err := b.Send(msg); err != nil {
		log.Println("send error: ", err)
	}
}
