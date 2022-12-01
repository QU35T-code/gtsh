package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/QU35T-code/gtsh/internal/manager"
	"github.com/hashicorp/yamux"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	Iface string `short:"i" long:"host" description:"Interface address on which to bind" default:"127.0.0.1" required:"true"`
	Port  string `short:"p" long:"port" description:"Port on which to bind" default:"9000" required:"true"`
}

func init() {
	_, err := flags.Parse(&opts)
	if err != nil {
		if err.(*flags.Error).Type == flags.ErrHelp {
			os.Exit(0)
		}
		log.Fatal(err)
	}
}

func streamReceive(conn net.Conn) string {
	buff := make([]byte, 0xff)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Stream read error: %s", err)
			break
		}
		data := string(buff[:n])
		return data
	}
	return ""
}

func streamSend(session *yamux.Session, data string) string {
	stream, err := session.Open()
	if err != nil {
		log.Fatal(err)
	}

	n, err := stream.Write([]byte(data))
	_ = n
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := session.Accept()
		if err != nil {
			if session.IsClosed() {
				log.Printf("TCP closed")
				break
			}
			log.Printf("Yamux accept: %s", err)
			continue
		}
		data := streamReceive(conn)
		return data
	}
	return ""
}

// var ConnectionList map[int]net.Conn

// var ConnectionList []net.Conn

func NewAgent(session *yamux.Session, conn net.Conn) {
	// TODO
	// Change this by multiple commands ? Better ?
	hello := streamSend(session, "hello")
	info := strings.Split(hello, "--")

	// Add conn to the list
	// ConnectionList = make(map[int]net.Conn)
	manager.AddConnection(conn)
	go manager.NewSession(info, conn.RemoteAddr().String())
}

func handle(conn net.Conn) {
	session, err := yamux.Server(conn, nil)
	if err != nil {
		log.Fatalf("Yamux server: %s", err)
	}
	NewAgent(session, conn)
}

func Server() {
	l, err := net.Listen("tcp4", fmt.Sprintf("%s:%s", opts.Iface, opts.Port))
	if err != nil {
		log.Fatalf("TCP server: %s", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("TCP accept: %s", err)
		}
		go handle(conn)
	}
}
