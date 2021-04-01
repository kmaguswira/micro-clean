package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/kmaguswira/micro-clean/service/file/repositories/seeder/config"
)

func main() {
	env := flag.String("e", "local", "")

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

	seeding()
}

func seeding() {
	fmt.Println(config.GetConfig())
	// readWriteRepository := repositories.NewReadWriteRepository(nil)

}
