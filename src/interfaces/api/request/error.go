package request

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type errorResponse struct {
	Status int    `json:"status"`
	Err    string `json:"error"`
}

func (c Context) Error(w http.ResponseWriter, status int, err error) error {
	res := errorResponse{
		Status: status,
		Err:    err.Error(),
	}

	b, err := json.Marshal(res)
	if err != nil {
		c.Logger.Error("failed at convert responset to json.", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, err = w.Write(b)
	if err != nil {
		c.Logger.Error("failed at write response.", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return nil
}

func Error(w http.ResponseWriter, status int, err error) error {
	res := errorResponse{
		Status: status,
		Err:    err.Error(),
	}

	b, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, err = w.Write(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return nil
}
