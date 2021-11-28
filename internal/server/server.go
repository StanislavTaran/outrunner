package server

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Server - connector server struct
type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	mySQL  *MySQL
}

// New - initialize new connector server
func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Run connector server
func (s *Server) Run() error {
	if err := s.ConfigureLogger(); err != nil {
		return err
	}

	s.logger.Info("Starting API server on port", s.config.BindAddr)

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// ConfigureLogger - configure logger for REST API connector server
func (s *Server) ConfigureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

// InitRoutes for REST API connector server
func (s *Server) InitRoutes() {
	// MySQL routes group
}
