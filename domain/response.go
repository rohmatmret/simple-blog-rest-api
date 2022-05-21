package domain

import (
	"encoding/json"
	"net/http"
)

type PostErrResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

func SetErrResponse(w http.ResponseWriter, data PostErrResponse, statusCode int) []byte {
	jb, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(jb))
	return jb
}

func SetResponse(w http.ResponseWriter, data []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(data))
	return
}
