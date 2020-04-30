package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kmaguswira/micro-clean/service/file/config"
	"github.com/kmaguswira/micro-clean/service/file/handler"

	// "github.com/kmaguswira/micro-clean/service/file/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"

	file "github.com/kmaguswira/micro-clean/service/file/proto/file"
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
		micro.Name("kmaguswira.srv.file"),
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
	fileHandler := handler.NewFile()
	file.RegisterFileHandler(service.Server(), fileHandler)

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("kmaguswira.srv.file", service.Server(), new(subscriber.file))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("kmaguswira.srv.file", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
