package stresstest

import (
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/go-ozzo/ozzo-validation/v4"
	"overload/internal/models"
	"time"
)

func (b *Tester) Validate(config *models.TestingConfig) error {
	log := b.logger.With("request_id", config.ID, "func", "validate")
	//log.Info("Логирую(удали меня)")
	err := validation.ValidateStruct(config,
		validation.Field(
			config.TestingDuration,
			validation.Required,
			validation.Max(time.Minute*30),
			validation.Min(time.Second),
		),
		validation.Field(config.ReqBody, is.JSON),
		validation.Field(config.ReqMethod, validation.Required, validation.In("POST", "GET", "PATCH")),
		validation.Field(config.ReqURL, validation.Required, is.URL),
		validation.Field(config.RPS, validation.Required, validation.Min(2), validation.Max(10000)),
	)
	if err != nil {
		log.Errorf("ошибка при валидации конфига: %v", err)
		return err
	}
	return nil
}
