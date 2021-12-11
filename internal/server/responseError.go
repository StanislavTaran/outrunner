package server

import (
	"encoding/json"
	"net/http"
)

type responseError struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

// NewResponseError creates new responseError struct and returns pointer
func (s *Server) NewResponseError(w http.ResponseWriter, message, reason string, code int) {
	err := func(w http.ResponseWriter, message, reason string, code int) error {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(code)
		body := responseError{
			Message: message,
			Reason:  reason,
		}
		res, err := json.Marshal(body)
		if err != nil {
			return err
		}

		_, err = w.Write(res)
		if err != nil {
			return err
		}

		return nil
	}(w, message, reason, code)
	if err != nil {
		s.logger.Error(err)
		http.Error(w, "internal server error.", http.StatusInternalServerError)
	}
}
