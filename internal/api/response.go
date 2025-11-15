package api

import (
	"encoding/json"
	"net/http"
)

type successResponse struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data"`
	RequestID string      `json:"request_id"`
}

type errorResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	RequestID string `json:"request_id"`
}

func writeJSON(w http.ResponseWriter, status int, data interface{}, reqID string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(successResponse{
		Success:   true,
		Data:      data,
		RequestID: reqID,
	})
}

func writeError(w http.ResponseWriter, status int, message string, reqID string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorResponse{
		Success:   false,
		Error:     message,
		RequestID: reqID,
	})
}
