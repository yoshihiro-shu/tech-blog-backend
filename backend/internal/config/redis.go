package config

import (
	"fmt"
	"time"
)

type RedisCache struct {
	Host     string        `yaml:"host"`
	Port     string        `yaml:"port"`
	Password string        `yaml:"password"`
	DbNumber int           `yaml:"dbNumber"`
	Expires  time.Duration `yaml:"expires"`
}

func (c RedisCache) GetRedisDNS() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
