package request

import (
	"encoding/json"
	"net/http"
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

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

	return nil
}
