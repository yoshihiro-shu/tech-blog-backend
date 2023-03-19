package config

import "errors"

var (
	ErrReadConfigFile      = errors.New("failed read configs.yaml")
	ErrUnmarshalConfigFile = errors.New("failed Unmarshal configs.yaml")
)
