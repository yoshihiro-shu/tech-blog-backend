package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Configs struct {
	User       User       `yaml:"user"`
	Db         DB         `yaml:"db"`
	CacheRedis RedisCache `yaml:"cacheRedis"`
	Twitter    Twitter    `yaml:"twitter"`
}

type User struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
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

type RedisCache struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DbNumber int    `yaml:"dbNumber"`
}

type Twitter struct {
	Apikey       string `yaml:"api_key"`
	ApiKeySecret string `yaml:"api_key_secret"`
	BearerToken  string `yaml:"bearer_token"`
	UserId       string `yaml:"user_id"`
}

const (
	StatusUnpublished = iota + 1
	StatusPublished
	StatusClosed
)

func New() Configs {
	conf := Configs{}

	b, err := os.ReadFile("./configs.yaml")
	if err != nil {
		log.Fatalf("failed read configs.yaml. err :%s", err.Error())
	}
	err = yaml.Unmarshal(b, &conf)
	if err != nil {
		log.Fatalf("failed Unmarshal configs.yaml. err :%s", err.Error())
	}

	return conf
}

func (c Configs) GetDb() DB {
	return c.Db
}

func (c Configs) GetCacheRedis() RedisCache {
	return c.CacheRedis
}

func (c Configs) GetRedisDNS() string {
	return fmt.Sprintf("%s:%s", c.CacheRedis.Host, c.CacheRedis.Port)
}

func (c Configs) GetUserAddr() string {
	return fmt.Sprintf("%s:%s", c.User.Host, c.User.Port)
}
