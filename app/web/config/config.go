package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Name           string         `json:"name"`
	Server         server         `json:"server"`
	Infrastructure infrastructure `json:"infrastructure"`
	Security       security       `json:"security"`
}

type server struct {
	Port   string `json:"port"`
	Secret string `json:"secret"`
}

type infrastructure struct {
	Persistance persistance `json:"persistance"`
}

type persistance struct {
	ReadWrite db `json:"readWrite"`
}

type db struct {
	Driver   string `json:"driver"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type security struct {
	Cors cors `json:"cors"`
}

type cors struct {
	AllowOrigin      []string `json:"allowOrigin"`
	MaxAge           int      `json:"maxAge"`
	AllowMethod      []string `json:"allowMethod"`
	AllowHeaders     []string `json:"allowHeaders"`
	ExposeHeaders    []string `json:"exposeHeaders"`
	AllowCredentials bool     `json:"allowCredentials"`
}

var config *Config

func Init(env string) {
	jsonFile, err := os.Open("config/" + env + ".json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &config)
}

func GetConfig() *Config {
	return config
}
