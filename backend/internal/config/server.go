package config

import "fmt"

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (s Server) GetAddr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}
