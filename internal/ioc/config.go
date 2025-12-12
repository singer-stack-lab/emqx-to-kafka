package ioc

import (
	"github.com/singer-stack-lab/emqx-to-kafka/config"

	"github.com/spf13/viper"
)

func LoadConfig() *config.Config {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	var cfg config.Config
	if err := v.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
