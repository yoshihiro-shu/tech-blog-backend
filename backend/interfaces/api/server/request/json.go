package request

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type JSONResponce struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func (c Context) JSON(w http.ResponseWriter, status int, data interface{}) error {
	res := JSONResponce{
		Status: status,
		Data:   data,
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
		c.Logger.Error("failed at write response.", zap.Error(err))
	}

	return nil
}
