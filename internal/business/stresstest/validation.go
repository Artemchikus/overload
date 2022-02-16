package stresstest

import (
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/go-ozzo/ozzo-validation/v4"
	"overload/internal/models"
)

func (b *Tester) Validate(config *models.TestingConfig) error {
	log := b.logger.With("request_id", config.ID, "func", "validate")
	err := validation.ValidateStruct(config,
		validation.Field(
			&config.DurationUNIX,
			validation.Required,
			validation.Max(30*60),
			validation.Min(1),
		),
		validation.Field(&config.ReqBody, is.JSON),
		validation.Field(&config.ReqMethod, validation.Required, validation.In("POST", "GET")),
		validation.Field(&config.ReqURL, validation.Required, is.URL),
		validation.Field(&config.RPS, validation.Required, validation.Min(2), validation.Max(10000)),
	)
	if err != nil {
		log.Errorf("ошибка при валидации конфига: %v", err)
		return err
	}
	return nil
}
