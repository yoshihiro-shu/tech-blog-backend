package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	StatusUnpublished = iota + 1
	StatusPublished
	StatusClosed
)

type Configs struct {
	User         User         `yaml:"user"`
	RelationalDB RelationalDB `yaml:"relationalDB"`
	CacheRedis   RedisCache   `yaml:"cacheRedis"`
	Twitter      Twitter      `yaml:"twitter"`
	AccessToken  AuthToken    `yaml:"access_token"`
	RefreshToken AuthToken    `yaml:"refresh_token"`
}

type User struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

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

type RedisCache struct {
	Host     string        `yaml:"host"`
	Port     string        `yaml:"port"`
	Password string        `yaml:"password"`
	DbNumber int           `yaml:"dbNumber"`
	Expires  time.Duration `yaml:"expires"`
}

type Twitter struct {
	Apikey       string `yaml:"api_key"`
	ApiKeySecret string `yaml:"api_key_secret"`
	BearerToken  string `yaml:"bearer_token"`
	UserId       string `yaml:"user_id"`
}

type AuthToken struct {
	SecretKey string        `yaml:"secret_key"`
	Expires   time.Duration `yaml:"expires"`
}

func (c Configs) MasterDB() DB {
	return c.RelationalDB.Master
}

func (c Configs) RepricaDB() []DB {
	return c.RelationalDB.Repricas
}

func (c Configs) GetCacheRedis() RedisCache {
	return c.CacheRedis
}

func (c Configs) GetUserAddr() string {
	return fmt.Sprintf("%s:%s", c.User.Host, c.User.Port)
}

func (c RedisCache) GetRedisDNS() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func New() Configs {
	var conf Configs

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
