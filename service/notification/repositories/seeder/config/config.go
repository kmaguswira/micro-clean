package config

import (
	"github.com/asim/go-micro/v3/config"
)

type Config struct {
	Name         string       `json:"name"`
	Repositories repositories `json:"repositories"`
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

func Init(env string) {
	config.LoadFile("config/" + env + ".json")
}

func GetConfig() *Config {
	var cfg Config
	config.Scan(&cfg)
	return &cfg
}
