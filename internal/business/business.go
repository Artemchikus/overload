package business

import "overload/internal/models"

type Business interface {
	Test(*models.UserConfig) (*models.Metric, error)
	Validate(*models.UserConfig) error
}
