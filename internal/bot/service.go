package bot

import (
	"bot/pkg/client"
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
var itemButtons = make([][]client.KeyboardButton, 0)

// Group of variables that are keyboard buttons.
var (
	startKeyboard = client.NewKeyboardWithMarkup(
		client.NewKeyboardRow(
			client.NewKeyboardButtonWithData("Купить 🛒", items),
			client.NewKeyboardButtonWithData("Адрес 📍", address),
			client.NewKeyboardButtonWithData("Отзыв ⭐️", feedBack),
			client.NewKeyboardButtonURL("VK 💙", "https://vk.com/ledda.store"),
		),
		client.NewKeyboardRow(
			client.NewKeyboardButtonWithData("Узнать размер ❔", height),
			client.NewKeyboardButtonWithData("О магазине ℹ️", info),
		),
	)

	addressKeyboard = client.NewKeyboardWithMarkup(
		client.NewKeyboardRow(
			client.NewKeyboardButtonWithData("Назад", start),
		),
	)

	itemsKeyboard = client.NewKeyboardWithMarkup()

	feedBackKeyboard = client.NewKeyboardWithMarkup(
		client.NewKeyboardRow(
			client.NewKeyboardButtonWithData("1", "rate::1"),
			client.NewKeyboardButtonWithData("2", "rate::2"),
			client.NewKeyboardButtonWithData("3", "rate::3"),
			client.NewKeyboardButtonWithData("4", "rate::4"),
			client.NewKeyboardButtonWithData("5", "rate::5"),
		),

		client.NewKeyboardRow(
			client.NewKeyboardButtonWithData("Назад", start),
		),
	)

	thxFeedbackKeyboard = client.NewKeyboardWithMarkup(
		client.NewKeyboardRow(
			client.NewKeyboardButtonWithData("На главную", start),
		),
	)

	sorryFeedbackKeyboard = client.NewKeyboardWithMarkup(
		client.NewKeyboardRow(
			client.NewKeyboardButtonWithData("Изменить отзыв", feedBack),
		),
		client.NewKeyboardRow(
			client.NewKeyboardButtonWithData("На главную", start),
		),
	)

	heightKeyboard = client.NewKeyboardWithMarkup(
		client.NewKeyboardRow(
			client.NewKeyboardButtonWithData(" - 158", sorryHeight),
			client.NewKeyboardButtonWithData("159 - 170", size+"::S"),
			client.NewKeyboardButtonWithData("171 - 180", size+"::M"),
			client.NewKeyboardButtonWithData("181 - 188", size+"::L"),
			client.NewKeyboardButtonWithData("189 - ", sorryHeight),
		),

		client.NewKeyboardRow(
			client.NewKeyboardButtonWithData("Назад", start),
		),
	)

	soldKeyboard = client.NewKeyboardWithMarkup(
		client.NewKeyboardRow(
			client.NewKeyboardButtonWithData("Нет в наличии 💔", items),
		),
		client.NewKeyboardRow(
			client.NewKeyboardButtonWithData("Назад", items),
		),
	)

	buyKeyboard = client.NewKeyboardWithMarkup(
		client.NewKeyboardRow(
			client.NewKeyboardButtonWithData("Купить 🛒", items),
		),
		client.NewKeyboardRow(
			client.NewKeyboardButtonWithData("Назад", start),
		),
	)

	infoKeyboard = client.NewKeyboardWithMarkup(
		client.NewKeyboardRow(
			client.NewKeyboardButtonWithData("Оставить отзыв 💫", feedBack),
		),
		client.NewKeyboardRow(
			client.NewKeyboardButtonWithData("Назад", start),
		),
	)
)

// Group of templates for messages.
var (
	itemTemplate = template.Must(template.New("items").Parse(itemMessage))
	infoTemplate = template.Must(template.New("info").Parse(infoMessage))
)
