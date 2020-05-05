package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kmaguswira/micro-clean/service/account/config"
	"github.com/kmaguswira/micro-clean/service/account/handler"

	// "github.com/kmaguswira/micro-clean/service/account/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"

	account "github.com/kmaguswira/micro-clean/service/account/proto/account"
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
	// New Service
	cfg := config.GetConfig()
	service := micro.NewService(
		micro.Name("kmaguswira.srv.account"),
		micro.Version("latest"),
		micro.Server(
			server.NewServer(
				server.Name(cfg.Server.Name),
				server.Address(":"+cfg.Server.Port),
			),
		),
	)

	// Initialise service
	service.Init()

	// Register Handler
	accountHandler := handler.NewAccount(service)
	account.RegisterAccountHandler(service.Server(), accountHandler)

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("kmaguswira.srv.account", service.Server(), new(subscriber.Account))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("kmaguswira.srv.account", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
