package manager

import (
	"github.com/desertbit/grumble"
)

func Menu() {
	var App = grumble.New(&grumble.Config{
		Name:                  "gtsh",
		Description:           "Give That Shell !",
		HelpHeadlineUnderline: true,
		HelpSubCommands:       true,
	})

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
			interract(c.Args.Int("id"))
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

	grumble.Main(App)
}
