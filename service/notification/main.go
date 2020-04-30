package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kmaguswira/micro-clean/service/notification/config"
	"github.com/kmaguswira/micro-clean/service/notification/handler"

	// "github.com/kmaguswira/micro-clean/service/notification/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"

	notification "github.com/kmaguswira/micro-clean/service/notification/proto/notification"
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
		micro.Name("kmaguswira.srv.notification"),
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
	notificationHandler := handler.NewNotification()
	notification.RegisterNotificationHandler(service.Server(), notificationHandler)

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("kmaguswira.srv.notification", service.Server(), new(subscriber.Notification))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("kmaguswira.srv.notification", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
