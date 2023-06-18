package config

import "fmt"

type RelationalDB struct {
	Master   DB   `yaml:"master"`
	Repricas []DB `yaml:"repricas"`
}

type DB struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	Sslmode  string `yaml:"sslMode"`
}

func (db DB) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		db.Host,
		db.User,
		db.Password,
		db.Name,
		db.Port,
		db.Sslmode,
	)
}
