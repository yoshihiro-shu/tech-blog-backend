package elastic

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/config"
)

func New(c config.Elasticsearch) (*elasticsearch.Client, error) {
	return elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			c.Address(),
		},
	})
}
