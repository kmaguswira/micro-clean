package utils

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct{}

func (c *Config) Parse(env string, config interface{}) {
	var err error
	vp := viper.New()
	vp.SetConfigType("json")
	vp.SetConfigName("config")
	vp.AddConfigPath("config/")
	vp.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	vp.AutomaticEnv()

	if env == "local" {
		vp.SetConfigName("config." + env)

	}

	err = vp.ReadInConfig()
	if err != nil {
		log.Fatal("Parsing error", err)
	}

	vp.WatchConfig()
	vp.Unmarshal(config)
}
