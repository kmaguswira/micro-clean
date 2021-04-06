package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/kmaguswira/micro-clean/service/file/config"
	"github.com/kmaguswira/micro-clean/service/file/handler"

	// "github.com/kmaguswira/micro-clean/service/file/subscriber"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/server"
	"github.com/asim/go-micro/v3/util/log"

	file "github.com/kmaguswira/micro-clean/service/file/proto/file"
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
		micro.Name("file"),
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
	// micro.RegisterSubscriber("file", service.Server(), new(subscriber.file))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("file", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
