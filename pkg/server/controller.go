package server

import (
	"io"
	"log"
	"net"

	"github.com/hashicorp/yamux"
)

func ListenAndServer(listenInterface string) {
	l, err := net.Listen("tcp4", listenInterface)
	if err != nil {
		log.Fatalf("TCP server: %s", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("TCP accept: %s", err)
		}
		Connection <- conn
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

func StreamSend(session *yamux.Session, data string) string {
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

var ConnectionList map[int]net.Conn
var Connection chan net.Conn

func New() {
	Connection = make(chan net.Conn, 1024)
	ConnectionList = make(map[int]net.Conn)
}
