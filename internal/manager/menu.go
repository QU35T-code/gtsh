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
		Help:  "Show all sessions",
		Usage: "sessions",
		Run: func(c *grumble.Context) error {
			sessions()
			return nil
		},
	})

	grumble.Main(App)
}
