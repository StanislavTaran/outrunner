package server

import (
	"encoding/json"
	"github.com/StanislavTaran/outrunner/internal/mongodb"
	"github.com/StanislavTaran/outrunner/internal/mysql"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

// mySQLGetRecords returns list of records by passed query in request body
func (s *Server) mySQLGetRecords() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		if _, ok := s.config.MySQL[vars["dbName"]]; !ok {
			s.NewResponseError(
				writer,
				"MySql db with such name not configured. ",
				"Check your config.",
				http.StatusBadRequest,
			)
			return
		}

		b, err := ioutil.ReadAll(request.Body)
		if err != nil {
			s.NewResponseError(
				writer,
				"Something went wrong.",
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		q := mysql.QueryInfo{}
		if err := json.Unmarshal(b, &q); err != nil {
			s.NewResponseError(
				writer,
				"Something went wrong.",
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		data, err := (*s.mySQL)[vars["dbName"]].GetRecords(q)
		if err != nil {
			s.logger.Error(err)
			s.NewResponseError(
				writer,
				"Something went wrong.",
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		} else {
			j, err := json.Marshal(data)
			if err != nil {
				s.NewResponseError(
					writer,
					"Something went wrong.",
					err.Error(),
					http.StatusInternalServerError,
				)
				return
			}

			writer.Header().Set("Content-Type", "application/json")
			_, err = writer.Write(j)
			if err != nil {
				s.NewResponseError(
					writer,
					"Something went wrong.",
					err.Error(),
					http.StatusInternalServerError,
				)
				return
			}
		}
	}
}

// mySQLCreateRecord creates records according to passed query
func (s *Server) mySQLCreateRecord() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		if _, ok := s.config.MySQL[vars["dbName"]]; !ok {
			s.NewResponseError(
				writer,
				"MySql db with such name not configured. ",
				"Check your config.",
				http.StatusInternalServerError,
			)
			return
		}

		b, err := ioutil.ReadAll(request.Body)
		if err != nil {
			s.NewResponseError(
				writer,
				"Something went wrong.",
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		q := mysql.QueryInfo{}
		if err := json.Unmarshal(b, &q); err != nil {
			s.NewResponseError(
				writer,
				"Something went wrong.",
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		_, err = (*s.mySQL)[vars["dbName"]].CreateRecord(q)
		if err != nil {
			s.NewResponseError(
				writer,
				"Something went wrong.",
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		} else {
			j, err := json.Marshal(map[string]string{
				"result": "ok",
			})
			if err != nil {
				http.Error(writer, "Something went wrong", http.StatusInternalServerError)
				return
			}

			writer.Header().Set("Content-Type", "application/json")
			_, err = writer.Write(j)
			if err != nil {
				s.NewResponseError(
					writer,
					"Something went wrong.",
					err.Error(),
					http.StatusInternalServerError,
				)
				return
			}
		}
	}
}

// mongoGetRecords returns list of records by passed query in request body
func (s *Server) mongoGetRecords() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		if _, ok := s.config.Mongodb[vars["dbName"]]; !ok {
			s.NewResponseError(
				writer,
				"Mongo db with such name not configured. ",
				"Check your config.",
				http.StatusBadRequest,
			)
			return
		}

		b, err := ioutil.ReadAll(request.Body)
		if err != nil {
			s.NewResponseError(
				writer,
				"Something went wrong.",
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		q := mongodb.QueryGet{}
		if err := json.Unmarshal(b, &q); err != nil {
			s.NewResponseError(
				writer,
				"Something went wrong.",
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		data, err := (*s.mongodb)[vars["dbName"]].GetRecords(q)
		if err != nil {
			s.logger.Error(err)
			s.NewResponseError(
				writer,
				"Something went wrong.",
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		} else {
			j, err := json.Marshal(data)
			if err != nil {
				s.NewResponseError(
					writer,
					"Something went wrong.",
					err.Error(),
					http.StatusInternalServerError,
				)
				return
			}

			writer.Header().Set("Content-Type", "application/json")
			_, err = writer.Write(j)
			if err != nil {
				s.NewResponseError(
					writer,
					"Something went wrong.",
					err.Error(),
					http.StatusInternalServerError,
				)
				return
			}
		}
	}
}

// mongoCreateRecord creates records according to passed query
func (s *Server) mongoCreateRecord() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		if _, ok := s.config.Mongodb[vars["dbName"]]; !ok {
			s.NewResponseError(
				writer,
				"Mongo db with such name not configured. ",
				"Check your config.",
				http.StatusInternalServerError,
			)
			return
		}

		b, err := ioutil.ReadAll(request.Body)
		if err != nil {
			s.NewResponseError(
				writer,
				"Something went wrong.",
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		q := mongodb.QueryInsert{}
		if err := json.Unmarshal(b, &q); err != nil {
			s.NewResponseError(
				writer,
				"Something went wrong.",
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		_, err = (*s.mongodb)[vars["dbName"]].CreateRecords(q)
		if err != nil {
			s.NewResponseError(
				writer,
				"Something went wrong.",
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		} else {
			j, err := json.Marshal(map[string]string{
				"result": "ok",
			})
			if err != nil {
				http.Error(writer, "Something went wrong", http.StatusInternalServerError)
				return
			}

			writer.Header().Set("Content-Type", "application/json")
			_, err = writer.Write(j)
			if err != nil {
				s.NewResponseError(
					writer,
					"Something went wrong.",
					err.Error(),
					http.StatusInternalServerError,
				)
				return
			}
		}
	}
}
