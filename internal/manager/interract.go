package manager

import (
	"fmt"
	"net"
	"os"

	"github.com/abiosoft/ishell"
	"github.com/abiosoft/readline"
)

var ConnectionList map[int]net.Conn

func init() {
	ConnectionList = make(map[int]net.Conn)
}

func AddConnection(conn net.Conn) {
	ConnectionList[ipCounter] = conn
}

func interract(sessionID int) {

	// TODO
	// If only 1 session exists, interact without argument
	if _, ok := ConnectionList[sessionID]; !ok {
		fmt.Println("Session ID don't exists !")
		return
	}

	// sh := ishell.NewWithConfig(&readline.Config{
	// 	Prompt:  "[ # ]> ",
	// 	VimMode: true,
	// })

	sh := ishell.NewWithConfig(&readline.Config{
		Prompt:      "[ # ]> ",
		Stdin:       ConnectionList[0],
		StdinWriter: ConnectionList[0],
		Stdout:      ConnectionList[0],
		Stderr:      ConnectionList[0],
		VimMode:     true,
	})

	sh.Run()
	os.Exit(0)
}
