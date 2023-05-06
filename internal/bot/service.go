package bot

import tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	startMessage           = "👋 Привет, меня зовут Космос! \n Я бот, который поможет тебе купить футболку:)"
	feedbackMessage        = "Оцени магазин и качество вещей по пятибалльной шкале:"
	thxFeedbackMessage     = "Спасибо! <3"
	sorryFeedbackMessage   = "Нам очень жаль, что вам не понравилось, мы постараемся стать лучше!"
	heightFeedbackMessage  = "Чтобы \"Предмет\" смотрелся как задумано, выбери свой дипазон роста:"
	addressFeedbackMessage = "Россия г. Санкт-Петербург"
	sorryHeightMessage     = "У нас пока что нет таких размеров, но мы уже стараемся исправить эту проблему!"
	itemsMessage           = "Выберите товар:"
)

type keyboardAndMessage struct {
	keyboard tgapi.InlineKeyboardMarkup
	message  string
}

const (
	height        = "Рост"
	start         = "start"
	address       = "Адрес"
	feedBack      = "Оставить отзыв"
	thxFeedback   = "Спасибо!"
	sorryFeedback = "Мы стараемся!"
	sorryHeight   = "Неверный размер"
	sSize         = "размер S"
	mSize         = "размер M"
	lSize         = "размер L"
	items         = "Предметы"
)

//var keyboardsAndMessages = map[string]keyboardAndMessage{
//	start: {keyboard: startKeyboard, message: startMessage},
//	start: {keyboard: feedBackKeyboard, message: startMessage},
//}

//
//var previous = map[string]string{
//	"aaa": start,
//}

var itemButtons = [][]tgapi.InlineKeyboardButton{
	tgapi.NewInlineKeyboardRow(
		tgapi.NewInlineKeyboardButtonData("Назад", start),
	),
}

var (
	startKeyboard = tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("Купить 🛒", items),
			tgapi.NewInlineKeyboardButtonData("Адрес 📍", address),
			tgapi.NewInlineKeyboardButtonData("Отзыв ⭐️", feedBack),
			tgapi.NewInlineKeyboardButtonURL("VK 💙", "https://vk.com/ledda.store"),
		),
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("Узнать размер ❔", height),
		),
	)

	addressKeyboard = tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("Назад", start),
		),
	)

	itemsKeyboard = tgapi.NewInlineKeyboardMarkup(itemButtons...)

	feedBackKeyboard = tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("1", sorryFeedback),
			tgapi.NewInlineKeyboardButtonData("2", sorryFeedback),
			tgapi.NewInlineKeyboardButtonData("3", sorryFeedback),
			tgapi.NewInlineKeyboardButtonData("4", thxFeedback),
			tgapi.NewInlineKeyboardButtonData("5", thxFeedback),
		),

		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("Назад", start),
		),
	)

	thxFeedbackKeyboard = tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("На главную", start),
		),
	)

	sorryFeedbackKeyboard = tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("Изменить отзыв", feedBack),
		),
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("На главную", start),
		),
	)

	heightKeyboard = tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData(" - 158", sorryHeight),
			tgapi.NewInlineKeyboardButtonData("159 - 170", sSize),
			tgapi.NewInlineKeyboardButtonData("171 - 180", mSize),
			tgapi.NewInlineKeyboardButtonData("181 - 188", lSize),
			tgapi.NewInlineKeyboardButtonData("189 - ", sorryHeight),
		),

		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("Назад", start),
		),
	)
)
