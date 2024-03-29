package cache_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/mock_test"
)

type data struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func TestRedis(t *testing.T) {
	testRedis := mock_test.MockRedis(t)
	tests := []struct {
		Key  string
		Data data
	}{
		{
			Key: "test1",
			Data: data{
				Id:        11111,
				Name:      "name1",
				CreatedAt: time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC),
			},
		}, {
			Key: "test2",
			Data: data{
				Id:        22222,
				Name:      "name2",
				CreatedAt: time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC).Add(time.Hour * 24),
			},
		}, {
			Key: "test3",
			Data: data{
				Id:        33333,
				Name:      "name3",
				CreatedAt: time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC).Add(time.Hour * 48),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Key, func(t *testing.T) {
			err := testRedis.SET(test.Key, test.Data)
			if err != nil {
				t.Fatalf("Error occured at testRedis.SET(). %s", err)
			}
			var testData data
			err = testRedis.GET(test.Key, &testData)
			if err != nil {
				t.Fatalf("Error occured at testRedis.GET(). %s", err)
			}

			if diff := cmp.Diff(test.Data, testData); diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
