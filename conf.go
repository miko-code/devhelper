package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	RedirectURL  string   `mapstructure:REDIRECT_URL`
	ClientID     string   `mapstructure:CLIENT_ID`
	ClientSecret string   `mapstructure:CLIENT_SECRET`
	Scopes       []string `mapstructure:SCOPES`
}

func LoadConf() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var c Config
	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil

}
