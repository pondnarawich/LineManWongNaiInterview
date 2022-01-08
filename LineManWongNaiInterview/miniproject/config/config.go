package config

import "github.com/caarlos0/env/v6"

type Config struct {
	PublicIP string `env:"PUBLIC_API" envDefault:"http://static.wongnai.com/devinterview/covid-cases.json"`
}

func Get() *Config {
	appConfig := &Config{}
	if err := env.Parse(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
