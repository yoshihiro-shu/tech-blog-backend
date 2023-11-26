package config

import "time"

type AuthToken struct {
	SecretKey string        `yaml:"secret_key"`
	Expires   time.Duration `yaml:"expires"`
}
