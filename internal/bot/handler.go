package bot

import (
	"bot/internal/entity"
	"bytes"
	"fmt"
	api "gopkg.in/telegram-bot-api.v4"
	"log"
	"strconv"
	"strings"
)

// handleMessage handles commands.
func (b *Bot) handleCommand(msg *api.Message) {
	switch msg.Text {
	case "/start":
		b.handleStart(msg)
	}
}

// handleMessage handles messages.
func (b *Bot) handleMessage() {}

// handleStart handles start command.
func (b *Bot) handleStart(msg *api.Message) {
	msgConfig := api.NewPhotoUpload(msg.Chat.ID, startImage)

	msgConfig.ReplyMarkup = startKeyboard
	msgConfig.Caption = startMessage

	_, err := b.Send(msgConfig)
	if err != nil {
		log.Println("send error: ", err)
	}
}

// handleMessage handle callbacks from user.
func (b *Bot) handleCallbackQuery(query *api.CallbackQuery) {
	if b.logic.IsPaused() {
		return
	}

	markup := api.NewInlineKeyboardMarkup()
	split := strings.Split(query.Data, "::")
	if len(split) == 0 {
		return
	}

	defer b.logger.Sync()

	text := split[0]
	pathToFile := ""

	switch text {
	case items:
		err := b.formItems()
		if err != nil {
			b.logger.Warn(err.Error())
		}
		pathToFile = allItemsImage
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
		pathToFile = startImage
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
			break
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

		if item.GetQuantity() == 0 {
			markup = soldKeyboard
		} else {
			markup = buyKeyboard
		}

		pathToFile = item.GetImage()

		var bytesArray []byte
		buf := bytes.NewBuffer(bytesArray)
		err = itemTemplate.Execute(buf, item)
		if err != nil {
			b.logger.Warn(err.Error())
			return
		}

		text = buf.String()
	}

	var msg api.Chattable

	if pathToFile != "" {
		msgWithPhoto := api.NewPhotoUpload(query.Message.Chat.ID, pathToFile)
		msgWithPhoto.Caption = text
		msgWithPhoto.ReplyMarkup = markup
		msg = msgWithPhoto
	} else {
		msgText := api.NewMessage(query.Message.Chat.ID, text)
		msgText.ReplyMarkup = markup
		msg = msgText
	}

	if _, err := b.Send(msg); err != nil {
		b.logger.Warn(fmt.Sprintf("send error: %v", err.Error()))
	}
}
