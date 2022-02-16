package models

import "time"

// Metric структура метрик, которые мы будем возвращать пользователям
type Metric struct {
	CodeToAmount       map[int]uint32 `json:"codes_ratio"`
	TotalRequests      uint64         `json:"total_requests"`
	MedianResponseTime uint64         `json:"median_time"`
	TestingTime        float64        `json:"testing_time"`
}

type ResponseInfo struct {
	Status    int
	TotalTime time.Duration
	Time      time.Time
}
