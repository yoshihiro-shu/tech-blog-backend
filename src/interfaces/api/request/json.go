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
		c.Logger.Error("failed at convert response to json.", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	w.WriteHeader(status)
	_, err = w.Write(b)
	if err != nil {
		c.Logger.Error("failed at write response.", zap.Error(err))
	}

	return nil
}
