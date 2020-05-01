package dialogFactory

import (
	"github.com/gameraccoon/telegram-bot-skeleton/dialog"
	"github.com/gameraccoon/telegram-bot-skeleton/processing"
	"github.com/nicksnyder/go-i18n/i18n"
)

type DialogFactory interface {
	MakeDialog(id int64, trans i18n.TranslateFunc, staticData *processing.StaticProccessStructs, customData interface{}) *dialog.Dialog
	ProcessVariant(variantId string, additionalId string, data *processing.ProcessData) bool
}
