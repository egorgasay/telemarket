package bot

import (
	"bot/internal/usecase"
	"bot/pkg/client"
	"fmt"
	"go.uber.org/zap"
)

// Bot represents a bot.
type Bot struct {
	token  string
	logic  *usecase.UseCase
	logger *zap.Logger
	*client.Client
}

// New creates a new bot.
func New(token string, logic *usecase.UseCase, logger *zap.Logger) (*Bot, error) {
	allItems, err := logic.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error getting items: %w", err)
	}

	for _, item := range allItems {
		itemButtons = append(itemButtons, client.NewKeyboardRow(
			client.NewKeyboardButtonWithData(item, "item::"+item),
		))
	}

	itemButtons = append(itemButtons,
		client.NewKeyboardRow(
			client.NewKeyboardButtonWithData("Назад", start),
		),
	)

	itemsKeyboard = client.NewKeyboardWithMarkup(itemButtons...)

	return &Bot{
		token:  token,
		logic:  logic,
		Client: client.New(token),
		logger: logger,
	}, nil
}

// Start starts the bot.
func (b *Bot) Start() error {
	u := client.NewUpdate(0)
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
