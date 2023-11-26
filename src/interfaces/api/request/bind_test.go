package request_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/mock_test"
)

type requestBody map[string]interface{}
type Expected struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age"`
}

func TestMustBind(t *testing.T) {
	// Create a mock request with JSON payload
	tests := []struct {
		Key      string
		Data     requestBody
		IsError  bool
		Expected Expected
	}{
		{
			Key: "test1",
			Data: requestBody{
				"name": "John",
				"age":  30,
			},
			IsError: false,
			Expected: Expected{
				Name: "John",
				Age:  30,
			},
		},
		{
			Key: "test2",
			Data: requestBody{
				"name": "John",
			},
			IsError: false,
			Expected: Expected{
				Name: "John"},
		},
		{
			Key: "test3",
			Data: requestBody{
				"age": 30,
			},
			IsError: true,
			Expected: Expected{
				Age: 30,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Key, func(t *testing.T) {
			json, _ := json.Marshal(test.Data)
			req, err := http.NewRequest("POST", "/test", bytes.NewBuffer(json))
			if err != nil {
				t.Fatal(err)
			}
			// Create a test context
			ctx := mock_test.NewMinContext()

			// Call MustBind to bind the request and validate the struct
			var testObj Expected
			err = ctx.MustBind(req, &testObj)
			if !test.IsError {
				if err != nil {
					t.Logf("MustBind returned an error: %v", err)
				}
			}
			if !reflect.DeepEqual(testObj, test.Expected) {
				t.Errorf("MustBind did not bind the request correctly. Got: %+v, Expected: %+v", testObj, test.Expected)
			}
		})
	}
}
