package config

type Twitter struct {
	Apikey       string `yaml:"api_key"`
	ApiKeySecret string `yaml:"api_key_secret"`
	BearerToken  string `yaml:"bearer_token"`
	UserId       string `yaml:"user_id"`
}
