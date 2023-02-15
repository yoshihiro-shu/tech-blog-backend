package request

import (
	"encoding/json"
	"log"
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

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, err = w.Write(b)
	if err != nil {
		log.Fatalf("failed at write response. err is %s\n", err.Error())
	}

	return nil
}
