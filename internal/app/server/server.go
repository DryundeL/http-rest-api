package server

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	if err := s.ConfigureLogger(); err != nil {
		return err
	}

	s.ConfigureRouter()

	s.logger.Info("Starting API Server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) ConfigureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *Server) ConfigureRouter() {
	s.router.HandleFunc("/hello", s.HandleHello())
}

func (s *Server) HandleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}