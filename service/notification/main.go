package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/kmaguswira/micro-clean/service/notification/config"
	"github.com/kmaguswira/micro-clean/service/notification/handler"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/server"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/kmaguswira/micro-clean/service/notification/subscriber"

	notification "github.com/kmaguswira/micro-clean/service/notification/proto/notification"
)

func main() {
	env := flag.String("e", "local", "")

	flag.Usage = func() {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
		fmt.Fprintln(w, "FLAG \t DESCRIPTION \t DEFAULT \t OPTIONS")
		fmt.Fprintln(w, "-e \t Environments \t local \t local/production")

		w.Flush()
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
		micro.Name("notification"),
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
	notificationSubscriber := subscriber.NewNotification()
	micro.RegisterSubscriber("kmaguswira.srv.notification.send-email", service.Server(), notificationSubscriber)

	// // Register Function as Subscriber
	// micro.RegisterSubscriber("kmaguswira.srv.notification.send-email", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
