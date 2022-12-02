package server

import (
	"strings"

	"github.com/hashicorp/yamux"
)

type GTSHAgent struct {
	Id       int
	Os       string
	Ip       string
	Hostname string
	Account  string
	Session  *yamux.Session
}

var ipCounter int

func NewAgent(session *yamux.Session) GTSHAgent {
	hello := StreamSend(session, "hello")
	sessionInformations := strings.Split(hello, "--")
	ipCounter++
	return GTSHAgent{
		Id:       ipCounter,
		Os:       sessionInformations[0],
		Ip:       session.RemoteAddr().String(),
		Hostname: sessionInformations[1],
		Account:  sessionInformations[2],
		Session:  session,
	}
}
