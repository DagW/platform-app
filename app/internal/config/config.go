package config

import "github.com/JeremyLoy/config"

type Config struct {
	Address string `config:"ADDR"`
	Port    string `config:"PORT"`
}

func Parse() (Config, error) {
	var c Config
	err := config.FromEnv().To(&c)
	return c, err
}
