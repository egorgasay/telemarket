package bot

import (
	api "gopkg.in/telegram-bot-api.v4"
	"text/template"
)

// Group of constants for bot messages
const (
	startMessage           = "👋 Привет, меня зовут Космос! \n Я бот, который поможет тебе купить футболку:)"
	feedbackMessage        = "Оцени магазин и качество вещей по пятибалльной шкале:"
	thxFeedbackMessage     = "Спасибо! <3"
	sorryFeedbackMessage   = "Нам очень жаль, что вам не понравилось, мы постараемся стать лучше!"
	heightFeedbackMessage  = "Чтобы \"Предмет\" смотрелся как задумано, выбери свой дипазон роста:"
	addressFeedbackMessage = "🇷🇺 Россия г. Санкт-Петербург"
	sorryHeightMessage     = "У нас пока что нет таких размеров, но мы уже стараемся исправить эту проблему!"
	itemsMessage           = "Выберите товар:"
	infoMessage            = "Средняя оценка: {{ .Avg }} ⭐️\n"
	itemMessage            = "{{ .Name }} \n{{ .Price }}р.\n{{ .Description }}"
)

// Group of constants for handling messages from user.
const (
	height        = "Рост"
	start         = "start"
	address       = "Адрес"
	feedBack      = "Оставить отзыв"
	thxFeedback   = "Спасибо!"
	sorryFeedback = "Мы стараемся!"
	sorryHeight   = "Неверный размер"
	size          = "размер"
	items         = "Предметы"
	item          = "item"
	info          = "info"
	rate          = "rate"
)

// itemButtons array of items. Automatically fulfilled from storage when bot starts.
var itemButtons = make([][]api.InlineKeyboardButton, 0)

var (
	allItemsImage = "src/images/allItems.png"
	startImage    = "src/images/ledda.png"
)

// Group of variables that are keyboard buttons.
var (
	startKeyboard = api.NewInlineKeyboardMarkup(
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("Купить 🛒", items),
			api.NewInlineKeyboardButtonData("Адрес 📍", address),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("Отзыв ⭐️", feedBack),
			api.NewInlineKeyboardButtonURL("VK 💙", "https://vk.com/ledda.store"),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("Узнать размер ❔", height),
			api.NewInlineKeyboardButtonData("О магазине ℹ️", info),
		),
	)

	addressKeyboard = api.NewInlineKeyboardMarkup(
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("Назад", start),
		),
	)

	itemsKeyboard = api.NewInlineKeyboardMarkup()

	feedBackKeyboard = api.NewInlineKeyboardMarkup(
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("1", "rate::1"),
			api.NewInlineKeyboardButtonData("2", "rate::2"),
			api.NewInlineKeyboardButtonData("3", "rate::3"),
			api.NewInlineKeyboardButtonData("4", "rate::4"),
			api.NewInlineKeyboardButtonData("5", "rate::5"),
		),

		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("Назад", start),
		),
	)

	thxFeedbackKeyboard = api.NewInlineKeyboardMarkup(
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("На главную", start),
		),
	)

	sorryFeedbackKeyboard = api.NewInlineKeyboardMarkup(
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("Изменить отзыв", feedBack),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("На главную", start),
		),
	)

	heightKeyboard = api.NewInlineKeyboardMarkup(
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(" - 158", sorryHeight),
			api.NewInlineKeyboardButtonData("159 - 170", size+"::S"),
			api.NewInlineKeyboardButtonData("171 - 180", size+"::M"),
			api.NewInlineKeyboardButtonData("181 - 188", size+"::L"),
			api.NewInlineKeyboardButtonData("189 - ", sorryHeight),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(" - 158", sorryHeight),
		), api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("159 - 170", size+"::S"),
		), api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("171 - 180", size+"::M"),
		), api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("189 - ", sorryHeight),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("Назад", start),
		),
	)

	soldKeyboard = api.NewInlineKeyboardMarkup(
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("Нет в наличии 💔", items),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("Назад", items),
		),
	)

	buyKeyboard = api.NewInlineKeyboardMarkup(
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("Купить 🛒", items),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("Назад", start),
		),
	)

	infoKeyboard = api.NewInlineKeyboardMarkup(
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("Оставить отзыв 💫", feedBack),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData("Назад", start),
		),
	)
)

// Group of templates for messages.
var (
	itemTemplate = template.Must(template.New("items").Parse(itemMessage))
	infoTemplate = template.Must(template.New("info").Parse(infoMessage))
)
