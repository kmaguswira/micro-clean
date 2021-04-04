package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"text/tabwriter"

	"github.com/kmaguswira/micro-clean/service/file/config"
	"github.com/kmaguswira/micro-clean/service/file/repositories"
	"github.com/kmaguswira/micro-clean/service/file/repositories/entity"
)

func main() {
	env := flag.String("e", "local", "")
	isDrop := flag.Bool("d", false, "")

	flag.Usage = func() {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
		fmt.Fprintln(w, "FLAG \t DESCRIPTION \t DEFAULT \t OPTIONS")
		fmt.Fprintln(w, "-d \t Drop Tables \t false \t false/true")
		fmt.Fprintln(w, "-e \t Environments \t local \t local/production")

		w.Flush()
		os.Exit(1)
	}

	flag.Parse()
	config.Init(*env)
	cfg := config.GetConfig()

	tables := []interface{}{
		entity.Document{},
		entity.Image{},
	}

	repo := repositories.NewDB(cfg.Repositories.ReadWrite)

	for _, model := range tables {

		if *isDrop {
			repo.Migrator().DropTable(model)
		}

		if !repo.Migrator().HasTable(model) {
			repo.Migrator().CreateTable(model)
		} else {
			log.Println("Table", reflect.TypeOf(model).Name(), "already exists")
		}

		repo.AutoMigrate(model)
	}
}
