package app

import (
	"github.com/desertbit/grumble"
)

func interract(args grumble.ArgMap) {
	agent := AgentList[args.Int("id")]
	_ = agent

	// stream, err := agent.Session.Open()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// TODO
	// If only 1 session exists, interact without argument
	// if _, ok := ConnectionList[sessionID]; !ok {
	// 	fmt.Println("Session ID don't exists !")
	// 	return
	// }

	// sh := ishell.NewWithConfig(&readline.Config{
	// 	Prompt:      "[ # ]> ",
	// 	Stdin:       stream,
	// 	StdinWriter: stream,
	// 	Stdout:      stream,
	// 	Stderr:      stream,
	// 	VimMode:     true,
	// })

	// sh.Run()
	// agent.Session.Close()
	// os.Exit(0)
}
