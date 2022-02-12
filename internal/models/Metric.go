package models

import "time"

// Metric структура метрик, которые мы будем возвращать пользователям
type Metric struct {
	ID string `json:"id"`
}

type ResponseInfo struct {
	Status    int
	TotalTime time.Duration
	Time      time.Time
}
