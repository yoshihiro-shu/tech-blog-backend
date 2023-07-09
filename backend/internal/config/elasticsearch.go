package config

import "fmt"

type Elasticsearch struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (e Elasticsearch) Address() string {
	return fmt.Sprintf("%s:%s", e.Host, e.Port)
}
