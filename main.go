package main

import (
	"fmt"
	"log"
	"os"

	"github.com/QU35T-code/gtsh/cmd/server"
	"github.com/QU35T-code/gtsh/internal/manager"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	Iface string `short:"i" long:"host" description:"Interface address on which to bind" default:"127.0.0.1" required:"true"`
	Port  string `short:"p" long:"port" description:"Port on which to bind" default:"9000" required:"true"`
}

func init() {
	_, err := flags.Parse(&opts)
	if err != nil {
		if err.(*flags.Error).Type == flags.ErrHelp {
			os.Exit(0)
		}
		log.Fatal(err)
	}
}

func main() {
	go server.Server()
	fmt.Printf("Listening on : %s:%s\n", opts.Iface, opts.Port)
	manager.Menu()
}
