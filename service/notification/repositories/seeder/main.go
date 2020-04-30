package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kmaguswira/micro-clean/service/notification/repositories"
	"github.com/kmaguswira/micro-clean/service/notification/repositories/seeder/config"
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

	seeding()
}

func seeding() {
	fmt.Println(config.GetConfig())
	readWriteRepository := repositories.NewReadWriteRepository(nil)

}
