package services

import (
	"context"

	"github.com/leohearts/insights-bot-kimichat/internal/configs"
	"github.com/leohearts/insights-bot-kimichat/internal/datastore"
	"github.com/leohearts/insights-bot-kimichat/pkg/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/leohearts/insights-bot-kimichat/ent/slackoauthcredentials"
)

type NewServicesParam struct {
	fx.In

	Logger *logger.Logger
	Ent    *datastore.Ent
	Config *configs.Config
}

type Services struct {
	logger *logger.Logger
	Ent    *datastore.Ent
	Config *configs.Config
}

func NewServices() func(param NewServicesParam) *Services {
	return func(param NewServicesParam) *Services {
		return &Services{
			logger: param.Logger,
			Ent:    param.Ent,
			Config: param.Config,
		}
	}
}

func (b *Services) NewStoreFuncForRefresh(teamID string) func(accessToken, refreshToken string) error {
	return func(accessToken, refreshToken string) error {
		return b.CreateOrUpdateSlackCredential(teamID, accessToken, refreshToken)
	}
}

func (b *Services) CreateOrUpdateSlackCredential(teamID, accessToken, refreshToken string) error {
	affectRows, err := b.Ent.SlackOAuthCredentials.Update().
		Where(slackoauthcredentials.TeamID(teamID)).
		SetAccessToken(accessToken).
		SetRefreshToken(refreshToken).
		Save(context.Background())
	if err != nil {
		b.logger.Warn("slack: failed to update access token", zap.Error(err))
		return err
	}

	if affectRows == 0 {
		// create
		err = b.Ent.SlackOAuthCredentials.Create().
			SetTeamID(teamID).
			SetAccessToken(accessToken).
			SetRefreshToken(refreshToken).
			Exec(context.Background())
		if err != nil {
			b.logger.Warn("slack: failed to save access token", zap.Error(err))
			return err
		}
	}

	return nil
}
