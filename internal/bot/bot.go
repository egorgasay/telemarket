package bot

import (
	"bot/internal/usecase"
	"fmt"
	"go.uber.org/zap"
	api "gopkg.in/telegram-bot-api.v4"
)

// Bot represents a bot.
type Bot struct {
	token  string
	logic  *usecase.UseCase
	logger *zap.Logger
	*api.BotAPI
}

// New creates a new bot.
func New(token string, logic *usecase.UseCase, logger *zap.Logger) (*Bot, error) {

	b, err := api.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("error creating bot api: %w", err)
	}

	bot := &Bot{
		token:  token,
		logic:  logic,
		BotAPI: b,
		logger: logger,
	}

	err = bot.formItems()
	if err != nil {
		return nil, fmt.Errorf("error while form items: %w", err)
	}

	return bot, nil
}

func (b *Bot) formItems() error {
	allItems, err := b.logic.GetAll()
	if err != nil {
		return fmt.Errorf("error getting items: %w", err)
	}

	itemButtons = itemButtons[:0]
	for _, item := range allItems {
		itemButtons = append(itemButtons, api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(item, "item::"+item),
		))
	}

	itemButtons = append(itemButtons,
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("Назад", start),
		),
	)

	itemsKeyboard = api.NewInlineKeyboardMarkup(itemButtons...)

	return nil
}

// Start starts the bot.
func (b *Bot) Start() error {
	u := api.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.GetUpdatesChan(u)
	if err != nil {
		panic(err) // TODO: REMOVE THIS
	}

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
