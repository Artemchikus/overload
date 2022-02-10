package stresstest

import (
	"overload/internal/models"

	"go.uber.org/zap"
)

type Business struct {
	logger *zap.SugaredLogger
}

func New() *Business {
	log, _ := zap.NewProduction()
	defer log.Sync()
	sugar := log.Sugar()

	return &Business{
		logger: sugar,
	}
}

func (b *Business) Test(*models.UserConfig) (*models.UserConfig, error) {
	return nil, nil
}