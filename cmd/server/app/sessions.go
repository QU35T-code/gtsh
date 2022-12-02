package app

import (
	"github.com/desertbit/grumble"
	"github.com/jedib0t/go-pretty/table"
)

func sessions(c *grumble.Context) {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.SetTitle("Sessions")
	t.AppendHeader(table.Row{"#", "OS", "IP Address", "Hostname", "Account"})
	for _, session := range AgentList {
		t.AppendRow(table.Row{session.Id, session.Os, session.Ip, session.Hostname, session.Account})
	}
	c.App.Println(t.Render())
}
