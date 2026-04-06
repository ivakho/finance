package category

import (
	"strconv"

	"finance/internal/service/tg_bot/api/category"
	"finance/internal/service/tg_bot/keyboard"
	"finance/internal/service/tg_bot/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleDeleteCategory(service *category.Service, userState *state.UserState, chatID int64, text string) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig

	if userState.Step != "delete_category_confirm" {
		return msg
	}

	id, _ := strconv.Atoi(userState.TempData["category_id"])

	switch text {
	case "Yes":
		_ = service.DeleteCategory(id)
		msg = tgbotapi.NewMessage(chatID, "Category deleted")

	case "No":
		msg = tgbotapi.NewMessage(chatID, "Cancelled")

	default:
		msg = tgbotapi.NewMessage(chatID, "Choose Yes or No")
		msg.ReplyMarkup = keyboard.Confirm()
		return msg
	}

	msg.ReplyMarkup = keyboard.CategoryMenu()
	userState.Step = "category_menu"
	userState.TempData = map[string]string{}

	return msg
}