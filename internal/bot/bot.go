package bot

import (
	"bot/internal/usecase"
	"fmt"
	"github.com/egorgasay/gotils"
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	token string
	logic *usecase.UseCase
	*tgapi.BotAPI
}

func New(token string, logic *usecase.UseCase) (*Bot, error) {
	bot, err := tgapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("error creating bot: %w", err)
	}

	bot.Debug = true // TODO: remove

	allItems, err := logic.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error getting items: %w", err)
	}

	for _, item := range allItems {
		itemButtons = append(itemButtons, tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData(item, "item::"+item),
		))
	}
	gotils.Reverse(itemButtons)
	return &Bot{
		token:  token,
		logic:  logic,
		BotAPI: bot,
		// TODO: add logger
	}, nil
}

func (b *Bot) Start() error {
	u := tgapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.GetUpdatesChan(u)
	for update := range updates {
		if update.CallbackQuery != nil {
			b.handleCallbackQuery(update.CallbackQuery)
			continue
		}

		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}

		b.handleMessage()
	}

	return nil
}

func (b *Bot) Stop() {
	b.StopReceivingUpdates()
}
