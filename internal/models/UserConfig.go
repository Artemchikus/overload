package models

import "time"

// TestingConfig структура конфига, который должен предоставить пользователь для стресс тестирования
type TestingConfig struct {
	ID              string         `json:"-"`
	ReqURL          string         `json:"url"`
	ReqMethod       string         `json:"url_method"`
	ReqBody         string         `json:"body"`
	RPS             int32          `json:"rps"`
	TestingDuration *time.Duration `json:"testing_duration"`
}
