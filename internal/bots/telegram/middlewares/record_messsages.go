package middlewares

import (
	"github.com/leohearts/insights-bot-kimichat/internal/models/chathistories"
	"github.com/leohearts/insights-bot-kimichat/internal/models/tgchats"
	"github.com/leohearts/insights-bot-kimichat/pkg/bots/tgbot"
	"github.com/leohearts/insights-bot-kimichat/pkg/types/telegram"
	"github.com/samber/lo"
)

func RecordMessage(chatHistories *chathistories.Model, tgchats *tgchats.Model) func(c *tgbot.Context, next func()) {
	return func(c *tgbot.Context, next func()) {
		if c.Update.Message == nil {
			return
		}

		chatType := telegram.ChatType(c.Update.Message.Chat.Type)
		if !lo.Contains([]telegram.ChatType{telegram.ChatTypeGroup, telegram.ChatTypeSuperGroup, telegram.ChatTypePrivate}, chatType) {
			return
		}
		if lo.Contains([]telegram.ChatType{telegram.ChatTypeGroup, telegram.ChatTypeSuperGroup}, chatType) {
			enabled, err := tgchats.HasChatHistoriesRecapEnabledForGroups(c.Update.Message.Chat.ID, c.Update.Message.Chat.Title)
			if err != nil {
				c.Logger.Error(err.Error())
				return
			}
			if !enabled {
				return
			}

			err = chatHistories.SaveOneTelegramChatHistory(c.Update.Message)
			if err != nil {
				c.Logger.Error(err.Error())
				return
			}
		}
		if lo.Contains([]telegram.ChatType{telegram.ChatTypePrivate}, chatType) {
			err := chatHistories.SaveOneTelegramPrivateForwardedReplayChatHistory(c.Update.Message)
			if err != nil {
				c.Logger.Error(err.Error())
				return
			}
		}

		next()
	}
}
