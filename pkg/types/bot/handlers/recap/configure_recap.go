package recap

import (
	"github.com/leohearts/insights-bot-kimichat/ent/feedbackchathistoriesrecapsreactions"
	"github.com/leohearts/insights-bot-kimichat/ent/feedbacksummarizationsreactions"
	"github.com/leohearts/insights-bot-kimichat/pkg/types/tgchat"
)

type ConfigureRecapToggleActionData struct {
	Status bool  `json:"status"`
	ChatID int64 `json:"chatId"`
	FromID int64 `json:"fromId"`
}

type ConfigureRecapAssignModeActionData struct {
	Mode   tgchat.AutoRecapSendMode `json:"mode"`
	ChatID int64                    `json:"chatId"`
	FromID int64                    `json:"fromId"`
}

type ConfigureRecapCompleteActionData struct {
	ChatID int64 `json:"chatId"`
	FromID int64 `json:"fromId"`
}

type ConfigureAutoRecapRatesPerDayActionData struct {
	Rates  int   `json:"rates"`
	ChatID int64 `json:"chatId"`
	FromID int64 `json:"fromId"`
}

type FeedbackRecapReactionActionData struct {
	ChatID int64                                     `json:"chatId"`
	LogID  string                                    `json:"logId"`
	Type   feedbackchathistoriesrecapsreactions.Type `json:"type"`
}

type FeedbackSummarizationReactionActionData struct {
	ChatID int64                                `json:"chatId"`
	LogID  string                               `json:"logId"`
	Type   feedbacksummarizationsreactions.Type `json:"type"`
}
