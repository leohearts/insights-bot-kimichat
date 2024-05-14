package middlewares

import (
	"github.com/leohearts/insights-bot-kimichat/internal/models/chathistories"
	"github.com/leohearts/insights-bot-kimichat/pkg/bots/tgbot"
)

func SyncWithEditedMessage(chatHistories *chathistories.Model) func(c *tgbot.Context, next func()) {
	return func(c *tgbot.Context, next func()) {
		if c.Update.EditedMessage == nil {
			return
		}

		err := chatHistories.UpdateOneTelegramChatHistory(c.Update.EditedMessage)
		if err != nil {
			c.Logger.Error(err.Error())
		}

		next()
	}
}
