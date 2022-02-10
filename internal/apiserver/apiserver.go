package apiserver

import "net/http"

// функция начала работы сервера
func Start(config *Config) error {
	srv := newServer() // получение сконфигурированного сервера

	return http.ListenAndServe(config.BindAddr, srv) // запуск сервера на порту
}
