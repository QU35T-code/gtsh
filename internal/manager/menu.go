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
		Name:  "status",
		Help:  "Shows the status of modules",
		Usage: "status",
		Run: func(c *grumble.Context) error {
			status()
			return nil
		},
	})

	grumble.Main(App)
}
