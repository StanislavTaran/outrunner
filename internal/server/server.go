package server

import (
	"errors"
	"fmt"
	"github.com/StanislavTaran/outrunner/internal/mongodb"
	"github.com/StanislavTaran/outrunner/internal/mysql"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// Server - connector server struct
type Server struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	mySQL   *map[string]*mysql.MySQL
	mongodb *map[string]*mongodb.Mongodb
}

// New - initialize new connector server
func New(config *Config) *Server {
	return &Server{
		config:  config,
		logger:  logrus.New(),
		router:  mux.NewRouter(),
		mySQL:   new(map[string]*mysql.MySQL),
		mongodb: new(map[string]*mongodb.Mongodb),
	}
}

// Run connector server
func (s *Server) Run() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.configureMysqlStore(); err != nil {
		return err
	}
	if err := s.configureMongoStore(); err != nil {
		return err
	}
	s.initRoutes()

	s.logger.Info("Starting API server on port", s.config.BindAddr)

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	s.logger.SetOutput(os.Stdout)
	s.logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000",
		FullTimestamp:   true,
	})
	return nil
}

func (s *Server) initRoutes() {
	s.router.HandleFunc("/api/v1/mysql/{dbName}/get", s.mySQLGetRecords())
	s.router.HandleFunc("/api/v1/mysql/{dbName}/create", s.mySQLCreateRecord())

	s.router.HandleFunc("/api/v1/mongo/{dbName}/get", s.mongoGetRecords())
	s.router.HandleFunc("/api/v1/mongo/{dbName}/create", s.mongoCreateRecord())
}

// configureMysqlStore - setup all your MySQL connections.
// This method uses connection info passed into server config (server.json)
func (s *Server) configureMysqlStore() error {
	for k, v := range s.config.MySQL {
		if (*s.mySQL) == nil {
			*s.mySQL = map[string]*mysql.MySQL{}
		}
		(*s.mySQL)[k] = mysql.New(&v)

		if err := (*s.mySQL)[k].Open(); err != nil {
			e := fmt.Errorf("MySql : %s, \n%w", k, err)
			return errors.New(e.Error())
		}
	}

	return nil
}

// configureMongoStore - setup all your mongodb connections.
// This method uses connection info passed into server config (server.json)
func (s *Server) configureMongoStore() error {
	for k, v := range s.config.Mongodb {
		if (*s.mongodb) == nil {
			*s.mongodb = map[string]*mongodb.Mongodb{}
		}
		(*s.mongodb)[k] = mongodb.New(&v)

		if err := (*s.mongodb)[k].Open(); err != nil {
			e := fmt.Errorf("MongoDB : %s, \n%w", k, err)
			return errors.New(e.Error())
		}
	}

	return nil
}
