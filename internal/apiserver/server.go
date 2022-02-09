package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type server struct {
	router *mux.Router
	logger *zap.Logger
}

func newServer() *server {
	log, _ := zap.NewProduction()

	s := &server{
		router: mux.NewRouter(),
		logger: log,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/", s.handleDDOS()).Methods("POST")

}

func (s *server) handleDDOS() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}
