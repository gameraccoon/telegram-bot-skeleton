package chat

import (
	"github.com/gameraccoon/telegram-bot-skeleton/dialog"
)

type Chat interface {
	SendMessage(chatId int64, message string, messageToReplace int64, preventPreview bool) int64
	SendDialog(chatId int64, dialog *dialog.Dialog, messageToReplace int64) int64
	RemoveMessage(chatId int64, messageId int64)
}
