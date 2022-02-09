package models

import "time"

type UserConfig struct {
	ID        int       `json:"id"`
	ReqURL    string    `json:"url"`
	ReqMethod string    `json:"url_method"`
	ReqBody   string    `json:"body"`
	NumOfReq  int       `json:"number_of_requests"`
	ReqPerMin int       `json:"requests_per_min"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
}
