package lib

import (
	"go.uber.org/fx"

	"github.com/leohearts/insights-bot-kimichat/internal/configs"
	"github.com/leohearts/insights-bot-kimichat/pkg/i18n"
	"github.com/leohearts/insights-bot-kimichat/pkg/logger"
)

type NewI18nParams struct {
	fx.In

	Configs *configs.Config
	Logger  *logger.Logger
}

func NewI18n() func(NewI18nParams) (*i18n.I18n, error) {
	return func(params NewI18nParams) (*i18n.I18n, error) {
		return i18n.NewI18n(
			i18n.WithLocalesDir(params.Configs.LocalesDir),
			i18n.WithLogger(params.Logger),
		)
	}
}
