package request

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
}

func (c Context) Error(w http.ResponseWriter, status int, err error) error {
	res := JSONResponce{
		Status: status,
		Data:   err.Error(),
	}

	b, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
	w.WriteHeader(status)
	return nil
}
