package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (c Context) MustBind(r *http.Request, i interface{}) error {
	err := c.bind(r, i)
	if err != nil {
		return err
	}

	if err = c.validateStruct(i); err != nil {
		return err
	}

	return nil
}

func (c Context) bind(r *http.Request, i interface{}) error {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, i)
	if err != nil {
		return err
	}

	return nil
}
