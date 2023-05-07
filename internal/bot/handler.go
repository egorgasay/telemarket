package bot

import (
	"bot/internal/entity"
	"bytes"
	"fmt"
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

// handleMessage handles commands.
func (b *Bot) handleCommand(msg *tgapi.Message) {
	switch msg.Command() {
	case "start":
		b.handleStart(msg)
	}
}

// handleMessage handles messages.
func (b *Bot) handleMessage() {

}

// handleStart handles start command.
func (b *Bot) handleStart(msg *tgapi.Message) {
	msgConfig := tgapi.NewMessage(msg.Chat.ID, startMessage)

	msgConfig.ReplyMarkup = startKeyboard
	_, err := b.Send(msgConfig)
	if err != nil {
		log.Println("send error: ", err)
	}
}

// handleMessage handle callbacks from user.
func (b *Bot) handleCallbackQuery(query *tgapi.CallbackQuery) {
	markup := tgapi.NewInlineKeyboardMarkup()
	split := strings.Split(query.Data, "::")
	if len(split) == 0 {
		return
	}

	defer b.logger.Sync()

	text := split[0]

	switch text {
	case items:
		text = itemsMessage
		markup = itemsKeyboard
	case height:
		text = heightFeedbackMessage
		markup = heightKeyboard
	case address:
		markup = addressKeyboard
		text = addressFeedbackMessage
	case start:
		markup = startKeyboard
		text = startMessage
	case rate:
		if len(split) != 2 {
			return
		}

		num, err := strconv.Atoi(split[1])
		if err != nil {
			log.Println(err)
			return
		}

		b.logic.AddRate(num)

		if num > 3 {
			markup = thxFeedbackKeyboard
			text = thxFeedbackMessage
		} else {
			markup = sorryFeedbackKeyboard
			text = sorryFeedbackMessage
		}
	case feedBack:
		markup = feedBackKeyboard
		text = feedbackMessage
	case sorryHeight:
		markup = heightKeyboard
		text = sorryHeightMessage
	case size:
		if len(split) != 2 {
			markup = heightKeyboard
			text = sorryHeightMessage
			return
		}
		markup = thxFeedbackKeyboard
		text = "Ваш размер: " + split[1]
	case info:
		markup = infoKeyboard
		var bytesArray []byte

		information := entity.Information{Avg: b.logic.GetRate()}

		buf := bytes.NewBuffer(bytesArray)
		err := infoTemplate.Execute(buf, information)
		if err != nil {
			b.logger.Warn(err.Error())
			return
		}

		text = buf.String()
	case item:
		if len(split) != 2 {
			return
		}

		item, err := b.logic.GetItemByName(split[1])
		if err != nil {
			b.logger.Warn(err.Error())
			return
		}

		if item.Quantity == 0 {
			markup = soldKeyboard
		} else {
			markup = buyKeyboard
		}

		var bytesArray []byte
		buf := bytes.NewBuffer(bytesArray)
		err = itemTemplate.Execute(buf, item)
		if err != nil {
			b.logger.Warn(err.Error())
			return
		}

		text = buf.String()
	}

	msg := tgapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, text, markup)
	if _, err := b.Send(msg); err != nil {
		b.logger.Warn(fmt.Sprintf("send error: %v", err.Error()))
	}
}
