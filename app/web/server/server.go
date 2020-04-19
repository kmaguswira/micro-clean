package server

import (
	"github.com/kmaguswira/micro-clean/app/web/config"
)

func Init() {
	config := config.GetConfig()

	r := SetupRouter(config.Name)

	r.Run(config.Server.Port)
}
