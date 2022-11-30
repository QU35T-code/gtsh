package manager

import (
	"github.com/desertbit/grumble"
	"github.com/jedib0t/go-pretty/v6/table"
)

type session struct {
	id       int
	os       string
	ip       string
	hostname string
	account  string
}

func sessions(c *grumble.Context) {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.SetTitle("Sessions")
	t.AppendHeader(table.Row{"#", "OS", "IP Address", "Hostname", "Account"})

	// TODO
	// Remove temporary table
	c1 := session{1, "Linux", "127.0.0.1", "heist", "root"}
	c2 := session{2, "Linux", "10.10.14.36", "response", "yakei"}
	c3 := session{3, "Windows", "192.168.42.132", "flight", "qu35t"}

	SessionsList := []session{c1, c2, c3}
	for _, session := range SessionsList {
		t.AppendRow(table.Row{session.id, session.os, session.ip, session.hostname, session.account})
	}
	c.App.Println(t.Render())
}
