package models

import "time"

// стурктура конфига, который должен предоставить пользователь для стресс темтирования
type UserConfig struct {
	ID        string    `json:"-"`
	ReqURL    string    `json:"url"`
	ReqMethod string    `json:"url_method"`
	ReqBody   string    `json:"body"`
	NumOfReq  int       `json:"number_of_requests"`
	ReqPerMin int       `json:"requests_per_min"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
}
