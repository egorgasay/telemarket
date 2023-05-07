package bot

import (
	"bot/internal/usecase"
	"fmt"
	"go.uber.org/zap"

	// telegram SDK
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	token  string
	logic  *usecase.UseCase
	logger *zap.Logger

	*tgapi.BotAPI
}

// New creates a new bot.
func New(token string, logic *usecase.UseCase, logger *zap.Logger) (*Bot, error) {
	bot, err := tgapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("error creating bot: %w", err)
	}

	allItems, err := logic.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error getting items: %w", err)
	}

	for _, item := range allItems {
		itemButtons = append(itemButtons, tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData(item, "item::"+item),
		))
	}

	itemButtons = append(itemButtons,
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("Назад", start),
		),
	)

	itemsKeyboard = tgapi.NewInlineKeyboardMarkup(itemButtons...)

	return &Bot{
		token:  token,
		logic:  logic,
		BotAPI: bot,
		logger: logger,
	}, nil
}

// Start starts the bot.
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

// Stop stops the bot.
func (b *Bot) Stop() {
	b.StopReceivingUpdates()
}
