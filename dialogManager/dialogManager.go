package dialogManager

import (
	"github.com/gameraccoon/telegram-bot-skeleton/dialog"
	"github.com/gameraccoon/telegram-bot-skeleton/processing"
	"github.com/gameraccoon/telegram-bot-skeleton/dialogFactory"
	"github.com/nicksnyder/go-i18n/i18n"
)

type DialogManager struct {
	dialogs map[string]dialogFactory.DialogFactory
	textProcessors TextInputProcessorManager
}

func (dialogManager *DialogManager) RegisterDialogFactory(id string, factory dialogFactory.DialogFactory) {
	if dialogManager.dialogs == nil {
		dialogManager.dialogs = make(map[string]dialogFactory.DialogFactory)
	}

	dialogManager.dialogs[id] = factory
}

func (dialogManager *DialogManager) RegisterTextInputProcessorManager(textInputProcessorManager TextInputProcessorManager) {
	dialogManager.textProcessors = textInputProcessorManager
}

func (dialogManager *DialogManager) MakeDialog(dialogId string, id int64, trans i18n.TranslateFunc, staticData *processing.StaticProccessStructs) (dialog *dialog.Dialog) {
	factory := dialogManager.getDialogFactory(dialogId)
	if factory != nil {
		dialog = factory.MakeDialog(id, trans, staticData)
		dialog.Id = dialogId
	}
	return
}

func (dialogManager *DialogManager) ProcessVariant(dialogId string, variantId string, additionalId string, data *processing.ProcessData) (processed bool) {
	factory := dialogManager.getDialogFactory(dialogId)
	if factory != nil {
		processed = factory.ProcessVariant(variantId, additionalId, data)
	}
	return
}

func (dialogManager *DialogManager) ProcessText(data *processing.ProcessData) bool {
	return dialogManager.textProcessors.processText(data)
}

func (dialogManager *DialogManager) getDialogFactory(id string) dialogFactory.DialogFactory {
	dialogFactory, ok := dialogManager.dialogs[id]
	if ok && dialogFactory != nil {
		return dialogFactory
	} else {
		return nil
	}
}
