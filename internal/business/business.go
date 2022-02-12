package business

import "overload/internal/models"

type Tester interface {
	Test(config *models.TestingConfig) (*models.Metric, error)
	Validate(config *models.TestingConfig) error
}
