package manager

import (
	"fmt"

	"github.com/desertbit/grumble"
	"github.com/jedib0t/go-pretty/v6/table"
)

var AgentList map[int]sessionstruct

type sessionstruct struct {
	id       int
	os       string
	ip       string
	hostname string
	account  string
}

func init() {
	AgentList = make(map[int]sessionstruct)
}

var ipCounter int

func NewSession(informations []string) {
	fmt.Println("Adding agent to the session table !")
	fmt.Println(informations)
	agent := sessionstruct{
		id:       ipCounter,
		os:       informations[0],
		ip:       "TODO",
		hostname: informations[1],
		account:  informations[2],
	}
	AgentList[agent.id] = agent
	ipCounter++
}

func sessions(c *grumble.Context) {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.SetTitle("Sessions")
	t.AppendHeader(table.Row{"#", "OS", "IP Address", "Hostname", "Account"})
	for _, session := range AgentList {
		t.AppendRow(table.Row{session.id, session.os, session.ip, session.hostname, session.account})
	}
	c.App.Println(t.Render())
}
