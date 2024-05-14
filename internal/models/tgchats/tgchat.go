package tgchats

import (
	"github.com/leohearts/insights-bot-kimichat/internal/configs"
	"github.com/leohearts/insights-bot-kimichat/internal/datastore"
	"github.com/leohearts/insights-bot-kimichat/pkg/logger"
	"go.uber.org/fx"
)

type NewModelParams struct {
	fx.In

	Config *configs.Config
	Ent    *datastore.Ent
	Digger *datastore.AutoRecapTimeCapsuleDigger
	Logger *logger.Logger
}

type Model struct {
	config *configs.Config
	ent    *datastore.Ent
	logger *logger.Logger
	digger *datastore.AutoRecapTimeCapsuleDigger
}

func NewModel() func(NewModelParams) (*Model, error) {
	return func(param NewModelParams) (*Model, error) {
		return &Model{
			config: param.Config,
			ent:    param.Ent,
			logger: param.Logger,
			digger: param.Digger,
		}, nil
	}
}
