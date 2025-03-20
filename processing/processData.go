package processing

import (
	"github.com/gameraccoon/telegram-bot-skeleton/dialog"
	"github.com/nicksnyder/go-i18n/i18n"
)

type ProcessData struct {
	Static            *StaticProccessStructs
	Command           string // first part of command without slash(/)
	Message           string // parameters of command or plain message
	ChatId            int64
	UserId            int64
	Trans             i18n.TranslateFunc
	AnsweredMessageId int64
	UserSystemLang    string
	UserSystemName    string
}

func (data *ProcessData) SubstituteMessage(message string) int64 {
	return data.Static.Chat.SendMessage(data.ChatId, message, data.AnsweredMessageId, false)
}

func (data *ProcessData) SubstituteDialog(dialog *dialog.Dialog) int64 {
	return data.Static.Chat.SendDialog(data.ChatId, dialog, data.AnsweredMessageId)
}

func (data *ProcessData) SendMessage(message string, preventPreview bool) int64 {
	return data.Static.Chat.SendMessage(data.ChatId, message, 0, preventPreview)
}

func (data *ProcessData) SendDialog(dialog *dialog.Dialog) int64 {
	return data.Static.Chat.SendDialog(data.ChatId, dialog, 0)
}
