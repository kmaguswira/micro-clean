package config

import "github.com/kmaguswira/micro-clean/utils"

type Config struct {
	Name         string       `json:"name"`
	Server       server       `json:"server"`
	Repositories repositories `json:"repositories"`
}

type server struct {
	Name   string `json:"name"`
	Port   string `json:"port"`
	Secret string `json:"secret"`
}

type repositories struct {
	Read      DB `json:"read"`
	ReadWrite DB `json:"readWrite"`
}

type DB struct {
	Driver   string `json:"driver"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Drop     bool   `json:"drop"`
}

var config Config

func Init(env string) {
	conf := utils.Config{}

	conf.Parse(env, &config)
}

func GetConfig() Config {
	return config
}
