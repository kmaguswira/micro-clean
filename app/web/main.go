package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kmaguswira/micro-clean/app/web/config"
	"github.com/kmaguswira/micro-clean/app/web/server"
)

func main() {
	env := flag.String("e", "local", "")

	flag.Usage = func() {
		fmt.Println("FLAG	DESCRIPTION		DEFAULT			OPTIONS")
		fmt.Println("-e		Environment		local			development/local/local_test/staging/production/test/migrate")
		os.Exit(1)
	}
	flag.Parse()

	config.Init(*env)
	bootstrap()
}

func bootstrap() {
	server.Init()
}
