package app

import (
	"fmt"

	"github.com/QU35T-code/gtsh/pkg/server"
	"github.com/desertbit/grumble"
)

func command(args grumble.ArgMap) {
	agent := AgentList[args.Int("id")]
	command := args.String("command")
	response := server.StreamSend(agent.Session, command)
	fmt.Println(response)
	// server.StreamSend()
}
