package app

import (
	"fmt"

	"github.com/QU35T-code/gtsh/pkg/server"
	"github.com/desertbit/grumble"
)

var App = grumble.New(&grumble.Config{
	Name:                  "gtsh",
	Description:           "Give That Shell !",
	HelpHeadlineUnderline: true,
	HelpSubCommands:       true,
})

var AgentList map[int]server.GTSHAgent

func RegisterAgent(agent server.GTSHAgent) {
	fmt.Printf("INFO[0001] Agent joined from %s\n", agent.Session.RemoteAddr().String())
	AgentList[agent.Id] = agent
}

func Run() {
	AgentList = make(map[int]server.GTSHAgent)

	App.AddCommand(&grumble.Command{
		Name:  "sessions",
		Help:  "Display all sessions",
		Usage: "sessions",
		Run: func(c *grumble.Context) error {
			sessions(c)
			return nil
		},
	})

	App.AddCommand(&grumble.Command{
		Name:  "interract",
		Help:  "Interract with a session",
		Usage: "interract",
		Args: func(a *grumble.Args) {
			a.Int("id", "session id")
		},
		Run: func(c *grumble.Context) error {
			interract(c.Args)
			return nil
		},
	})

	App.AddCommand(&grumble.Command{
		Name:  "command",
		Help:  "Execute a shell command",
		Usage: "command",
		Args: func(a *grumble.Args) {
			a.Int("id", "session id")
			a.String("command", "command")
		},
		Run: func(c *grumble.Context) error {
			command(c.Args)
			return nil
		},
	})

	App.AddCommand(&grumble.Command{
		Name:  "kill",
		Help:  "Kill a session",
		Usage: "kill",
		Args: func(a *grumble.Args) {
			a.Int("id", "session id")
		},
		Run: func(c *grumble.Context) error {
			kill(c.Args)
			return nil
		},
	})
}
