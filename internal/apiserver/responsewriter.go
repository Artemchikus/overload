package apiserver

import "net/http"

// кастомный писатель ответов
type responseWriter struct {
	http.ResponseWriter
	code int
}

// функция заполнения необходимых заголовков ответов сервера
func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}