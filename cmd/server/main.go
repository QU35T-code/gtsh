package main

import (
	"flag"
	"fmt"

	"github.com/QU35T-code/gtsh/cmd/server/app"
	"github.com/QU35T-code/gtsh/pkg/server"
	"github.com/desertbit/grumble"
	"github.com/hashicorp/yamux"
)

func main() {
	var listenInterface = flag.String("i", "0.0.0.0:9000", "Listening address")
	flag.Parse()
	fmt.Printf("Listening on : %s\n", *listenInterface)

	app.Run()
	server.New()
	go server.ListenAndServer(*listenInterface)

	go func() {
		for {
			remoteConn := <-server.Connection
			yamuxConn, err := yamux.Server(remoteConn, nil)
			if err != nil {
				panic(err)
			}
			agent := server.NewAgent(yamuxConn)
			app.RegisterAgent(agent)
		}
	}()
	grumble.Main(app.App)
}
