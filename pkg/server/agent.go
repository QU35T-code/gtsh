package server

import "github.com/hashicorp/yamux"

type GTSHAgent struct {
	Id       int
	Os       string
	Ip       string
	Hostname string
	Account  string
	Session  *yamux.Session
}

func NewAgent(session *yamux.Session) GTSHAgent {
	return GTSHAgent{
		Id:       1,
		Os:       "test",
		Ip:       "10.10.10.10",
		Hostname: "test-hostname",
		Account:  "root",
		Session:  session,
	}
	// TODO
	// Change this by multiple commands ? Better ?
	// hello := streamSend(session, "hello")
	// info := strings.Split(hello, "--")

	// Add conn to the list
	// ConnectionList = make(map[int]net.Conn)
	// manager.AddConnection(conn)
	// go manager.NewSession(info, conn.RemoteAddr().String())
}
