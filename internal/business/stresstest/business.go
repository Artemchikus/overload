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

func (b *Business) Test(config *models.UserConfig) (*models.Metric, error) {
	log := b.logger.With("request_id", config.ID, "func", "test")
	log.Info("Логирую(удали меня)")

	return nil, nil
}

func (b *Business) Validate(config *models.UserConfig) error {
	log := b.logger.With("request_id", config.ID, "func", "validate")
	log.Info("Логирую(удали меня)")

	return nil
}