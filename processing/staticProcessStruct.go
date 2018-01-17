package processing

import (
	"github.com/gameraccoon/telegram-bot-skeleton/chat"
	"github.com/gameraccoon/telegram-bot-skeleton/database"
	"github.com/gameraccoon/telegram-bot-skeleton/dialog"
	"github.com/nicksnyder/go-i18n/i18n"
)

type AwaitingTextProcessorData struct {
	ProcessorId string
	AdditionalId int64
}

type UserState struct {
	awaitingTextProcessor *AwaitingTextProcessorData
	currentPage int // temporary data for lists handling
	// temporary data used for consecutive menues
	customData map[string]interface{}
}

type StaticProccessStructs struct {
	Chat chat.Chat
	Db *database.Database
	Config interface{}
	Trans map[string]i18n.TranslateFunc
	MakeDialogFn func(string, int64, i18n.TranslateFunc, *StaticProccessStructs)*dialog.Dialog
	userStates map[int64]UserState
	// custom data that can be used in different places
	customData map[string]interface{}
}

func (staticData *StaticProccessStructs) Init() {
	staticData.userStates = make(map[int64]UserState)
}

func (staticData *StaticProccessStructs) SetUserStateTextProcessor(userId int64, processor *AwaitingTextProcessorData) {
	state := staticData.userStates[userId]
	state.awaitingTextProcessor = processor
	staticData.userStates[userId] = state
}

func (staticData *StaticProccessStructs) GetUserStateTextProcessor(userId int64) *AwaitingTextProcessorData {
	if state, ok := staticData.userStates[userId]; ok {
		return state.awaitingTextProcessor
	} else {
		return nil
	}
}

func (staticData *StaticProccessStructs) SetUserStateCurrentPage(userId int64, page int) {
	state := staticData.userStates[userId]
	state.currentPage = page
	staticData.userStates[userId] = state
}

func (staticData *StaticProccessStructs) GetUserStateCurrentPage(userId int64) int {
	if state, ok := staticData.userStates[userId]; ok {
		return state.currentPage
	} else {
		return 0
	}
}

func (staticData *StaticProccessStructs) SetUserStateValue(userId int64, key string, value interface{}) {
	state := staticData.userStates[userId]
	if state.customData == nil {
		state.customData = map[string]interface{}{}
	}
	state.customData[key] = value
	staticData.userStates[userId] = state
}

func (staticData *StaticProccessStructs) GetUserStateValue(userId int64, key string) interface{} {
	if state, ok := staticData.userStates[userId]; ok {
		if state.customData != nil {
			return state.customData[key]
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (staticData *StaticProccessStructs) SetCustomValue(key string, value interface{}) {
	if staticData.customData == nil {
		staticData.customData = map[string]interface{}{}
	}

	staticData.customData[key] = value
}

func (staticData *StaticProccessStructs) GetCustomValue(key string) interface{} {
	if staticData.customData == nil {
		return nil
	}

	return staticData.customData[key]
}

