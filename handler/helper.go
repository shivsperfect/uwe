package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type APIError struct {
	StatusCode int    `json:"status_code"`
	Msg        string `json:"msg"`
}

func (e APIError) Error() string {
	return e.Msg
}

func ShowAPIError(status int, err error) APIError {
	return APIError{
		StatusCode: status,
		Msg:        err.Error(),
	}
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func ServeHTTP(h APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if apiErr, ok := err.(APIError); ok {
				writeJSON(w, apiErr.StatusCode, apiErr)
			} else {
				errResp := APIError{
					StatusCode: http.StatusInternalServerError,
					Msg:        "Internal Server Error",
				}
				writeJSON(w, errResp.StatusCode, errResp)
			}
			slog.Error("HTTP API error", "err", err.Error(), "path", r.URL.Path)
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
