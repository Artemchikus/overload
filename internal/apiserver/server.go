package apiserver

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type ctxKey int8

const (
	ctxKeyRequestID ctxKey = iota
)

type server struct {
	router *mux.Router
	logger *zap.SugaredLogger
}

func newServer() *server {
	log, _ := zap.NewProduction()
	defer log.Sync()
	sugar := log.Sugar()

	s := &server{
		router: mux.NewRouter(),
		logger: sugar,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)

	s.router.HandleFunc("/setConfig", s.handleDDOS()).Methods("POST")
	s.router.HandleFunc("/getInfo", s.handleInfo()).Methods("POST")
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sugar := s.logger.With(
			"remoute_addr", r.RemoteAddr,
			"request_id", r.Context().Value(ctxKeyRequestID),
		)
		sugar.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		sugar.Infof("complited with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Since(start))
	})
}

func (s *server) handleDDOS() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) handleInfo() http.HandlerFunc {
	type request struct {
		UserID string `json:"user_id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
