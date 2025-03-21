package telegramChat

import (
	"fmt"
	"github.com/gameraccoon/telegram-bot-skeleton/dialog"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"sync"
)

type TelegramChat struct {
	bot   *tgbotapi.BotAPI
	mutex sync.Mutex
}

func MakeTelegramChat(apiToken string) (bot *TelegramChat, outErr error) {
	newBot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		outErr = err
		return
	}

	bot = &TelegramChat{
		bot: newBot,
	}

	return
}

func (telegramChat *TelegramChat) GetBot() *tgbotapi.BotAPI {
	return telegramChat.bot
}

func (telegramChat *TelegramChat) SetDebugModeEnabled(isEnabled bool) {
	telegramChat.bot.Debug = isEnabled
}

func (telegramChat *TelegramChat) GetBotUsername() string {
	return telegramChat.bot.Self.UserName
}

func makeMessage(chatId int64, message string, messageToReplace int64, preventPreview bool, markup *tgbotapi.InlineKeyboardMarkup) tgbotapi.Chattable {
	if messageToReplace == 0 {
		msg := tgbotapi.NewMessage(chatId, message)
		msg.ParseMode = "HTML"
		msg.DisableWebPagePreview = preventPreview
		if markup != nil {
			msg.ReplyMarkup = markup
		}
		return msg
	} else {
		msg := tgbotapi.NewEditMessageText(chatId, int(messageToReplace), message)
		msg.ParseMode = "HTML"
		msg.DisableWebPagePreview = preventPreview
		if markup != nil {
			msg.ReplyMarkup = markup
		}
		return msg
	}
}

func (telegramChat *TelegramChat) SendMessage(chatId int64, message string, messageToReplace int64, preventPreview bool) (messageId int64) {

	packedMessage := makeMessage(chatId, message, messageToReplace, preventPreview, nil)

	telegramChat.mutex.Lock()
	sentMessage, err := telegramChat.bot.Send(packedMessage)
	telegramChat.mutex.Unlock()

	if err == nil {
		messageId = int64(sentMessage.MessageID)
	}

	return
}

func getCommand(dialogId string, variantId string, additionalId string) string {
	if additionalId == "" {
		return fmt.Sprintf("/%s_%s", dialogId, variantId)
	} else {
		return fmt.Sprintf("/%s_%s_%s", dialogId, variantId, additionalId)
	}
}

func (telegramChat *TelegramChat) SendDialog(chatId int64, dialog *dialog.Dialog, messageToReplace int64) (messageId int64) {
	if dialog == nil {
		return
	}

	markup := tgbotapi.NewInlineKeyboardMarkup()

	currentRow := []tgbotapi.InlineKeyboardButton{}
	currentRowId := 0
	for _, variant := range dialog.Variants {
		if currentRowId != variant.RowId {
			if len(currentRow) > 0 {
				markup.InlineKeyboard = append(markup.InlineKeyboard, currentRow)
			}
			currentRow = []tgbotapi.InlineKeyboardButton{}
			currentRowId = variant.RowId
		}

		if len(variant.Url) == 0 {
			currentRow = append(currentRow, tgbotapi.NewInlineKeyboardButtonData(
				variant.Text,
				getCommand(dialog.Id, variant.Id, variant.AdditionalId),
			))
		} else {
			currentRow = append(currentRow, tgbotapi.NewInlineKeyboardButtonURL(
				variant.Text,
				variant.Url,
			))
		}
	}
	markup.InlineKeyboard = append(markup.InlineKeyboard, currentRow)

	packedMessage := makeMessage(chatId, dialog.Text, messageToReplace, true, &markup)

	telegramChat.mutex.Lock()
	sentMessage, err := telegramChat.bot.Send(packedMessage)
	telegramChat.mutex.Unlock()

	if err == nil {
		messageId = int64(sentMessage.MessageID)
	}

	return
}

func (telegramChat *TelegramChat) RemoveMessage(chatId int64, messageId int64) {
	if messageId == 0 {
		return
	}

	deleteConfig := tgbotapi.DeleteMessageConfig{
		ChatID:    chatId,
		MessageID: int(messageId),
	}

	telegramChat.mutex.Lock()
	_, err := telegramChat.bot.DeleteMessage(deleteConfig)
	telegramChat.mutex.Unlock()

	if err != nil {

	}
}

func (telegramChat *TelegramChat) LockMutex() {
	telegramChat.mutex.Lock()
}

func (telegramChat *TelegramChat) UnlockMutex() {
	telegramChat.mutex.Unlock()
}
