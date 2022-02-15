package models

// TestingConfig структура конфига, который должен предоставить пользователь для стресс тестирования
type TestingConfig struct {
	ID           string `json:"-"`
	ReqURL       string `json:"url"`
	ReqMethod    string `json:"method"`
	ReqBody      string `json:"body"`
	RPS          int32  `json:"rps"`
	DurationUNIX int64  `json:"testing_duration"`
}
