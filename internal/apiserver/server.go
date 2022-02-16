package apiserver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"overload/internal/business"
	"overload/internal/business/stresstest"
	"overload/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type ctxKey int8 // Новый тип данных необходимый для работы с контекстом

const (
	ctxKeyRequestID ctxKey = iota // Контекстная ферменная для доступа к id запроса
)

// server struct
type server struct {
	router *mux.Router
	logger *zap.SugaredLogger
	tester business.Tester
}

// Функция возвращения сконфигурированного сервера
func newServer() *server {
	// инициализация "сахарного" логгера
	log, _ := zap.NewProduction()
	defer log.Sync()
	sugar := log.Sugar()

	test := stresstest.New() // создание экземпляра бизнес логики по стресс тестам

	// конфигурация сервера стресс тестов
	s := &server{
		router: mux.NewRouter(),
		logger: sugar,
		tester: test,
	}

	s.configureRouter() // конфигурация всех URL оработчиков запросов

	return s
}

// переопределения метода serveHTTP чтобы сервер соответсвтовал интерфейсу http.Handler
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// удобная фунция для ответов с информацией об ошибке
func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

// удобная функция для штатных ответов
func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// самая главная функция, где соединяются URL-ы и функции, которые обрабатывают запросы на них приходящие
func (s *server) configureRouter() {
	s.router.Use(s.setRequestID) // прописывание уникального Id каждому запросу
	s.router.Use(s.logRequest)   // логирование каждого запроса

	// обработка URL-ов
	s.router.HandleFunc("/setConfig", s.handleDDOS()).Methods("POST")
	s.router.HandleFunc("/getInfo", s.handleInfo()).Methods("POST")
}

// функция прописывания уникального id запроса в заголовке ответа
func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

// функция логирования запросов
func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sugar := s.logger.With(
			"remoute_addr", r.RemoteAddr,
			"request_id", r.Context().Value(ctxKeyRequestID),
		) // добавление полей в логгер
		sugar.Infof("started %s %s", r.Method, r.RequestURI) // лог прихода запроса

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		sugar.Infof("complited with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Since(start)) // лог отпрваки ответа
	})
}

// функция обработки запросов на стресс тестирование
func (s *server) handleDDOS() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.TestingConfig{
			ID: fmt.Sprint(ctxKeyRequestID),
		}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		} // декодирование тела json запроса в структуру

		req.ID = fmt.Sprintf("%s", r.Context().Value(ctxKeyRequestID))

		if err := s.tester.Validate(req); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		} // валидация данных

		metric, err := s.tester.
			Test(req) // получение метрик
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		//Returning Metrics if there is no db...

		s.respond(w, r, http.StatusOK, metric)
	}
}

// функция для получения инфы о результатах тестирования, если добавим бд
func (s *server) handleInfo() http.HandlerFunc {

	// получение результатов будет по уникальному id (ожидаемый формат json)
	type request struct {
		UserID string `json:"user_id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		} // декодирование тела json запроса в структкру

		s.respond(w, r, http.StatusOK, nil)
	}
}
