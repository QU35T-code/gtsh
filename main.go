package main

import (
	"fmt"
	"log"
	"os"

	"github.com/QU35T-code/gtsh/internal/manager"
	"github.com/briandowns/spinner"
	"github.com/jessevdk/go-flags"
)

var spin *spinner.Spinner

var opts struct {
	Iface  string `short:"i" long:"host" description:"Interface address on which to bind" default:"127.0.0.1" required:"true"`
	Port   string `short:"p" long:"port" description:"Port on which to bind" default:"9000" required:"true"`
	Socket string `short:"s" long:"socket" description:"Domain socket from which the program reads"`
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
	fmt.Println(fmt.Sprintf("Listening on : %s:%s", opts.Iface, opts.Port))
	// go cmd.Server()
	manager.Menu()
}
