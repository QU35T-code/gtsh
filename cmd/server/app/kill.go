package app

import "github.com/desertbit/grumble"

func killSession(args grumble.ArgMap) {
	agent := AgentList[args.Int("id")]
	agent.Session.Close()

	// TODO
	// Remove from table
}
