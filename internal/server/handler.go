package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type queryInfo struct {
	Table string `json:"table"`
	Query map[string]interface{}
}

// mySQLGetRecords returns list of records by passed query in request body
func (s *Server) mySQLGetRecords() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		vars := mux.Vars(request)
		if _, ok := s.config.MySQL[vars["dbName"]]; !ok {
			http.Error(writer, "MySql db with such name not configured", http.StatusBadRequest)
			return
		}

		data, err := (*s.mySQL)[vars["dbName"]].GetRecords()
		if err != nil {
			s.logger.Error(err)
			http.Error(writer, "Something went wrong 1", http.StatusInternalServerError)
		} else {
			j, err := json.Marshal(data)
			if err != nil {
				http.Error(writer, "Something went wrong 2", http.StatusInternalServerError)
			}

			writer.Header().Set("Content-Type", "application/json")
			_, err = writer.Write(j)
			if err != nil {
				http.Error(writer, "Something went wrong 3", http.StatusInternalServerError)
			}
		}
	}
}
