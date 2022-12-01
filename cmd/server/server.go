package server

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/mattn/go-tty"
)

var opts struct {
	Iface  string `short:"i" long:"host" description:"Interface address on which to bind" default:"127.0.0.1" required:"true"`
	Port   string `short:"p" long:"port" description:"Port on which to bind" default:"9000" required:"true"`
	Socket string `short:"s" long:"socket" description:"Domain socket from which the program reads"`
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

// var messages chan string

func Server() {
	var listener net.Listener
	var err error

	// messages = make(chan string)

	// go func() {
	// 	for {
	// 		message := <-messages
	// 		fmt.Println(message)
	// 	}
	// }()

	if opts.Socket == "" {
		listener, err = newTLSListener()
	} else {
		listener, err = net.Listen("unix", opts.Socket)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	waitingForConnections(listener)
}

func waitingForConnections(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
		}
		// defer conn.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}
		handleConnection(conn, listener)
	}
}

func newTLSListener() (net.Listener, error) {
	pem := "./certs/server.pem"
	key := "./certs/server.key"
	cer, err := tls.LoadX509KeyPair(pem, key)
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	connStr := fmt.Sprintf("%s:%s", opts.Iface, opts.Port)
	return tls.Listen("tcp", connStr, config)
}

func registerAgent(conn net.Conn) {
	// ip := conn.RemoteAddr().String()
	// fmt.Println(ip)
}

func handleConnection(conn net.Conn, listener net.Listener) {
	fmt.Printf("INFO[0001] Agent joined from %s\n", conn.RemoteAddr().String())
	registerAgent(conn)
	// messages <- "INFO[0001] Agent joined."
	reader, writer := bufio.NewReader(conn), bufio.NewWriter(conn)

	go func() {
		for {
			out, err := reader.ReadByte()
			if err != nil {
				fmt.Println("\nINFO[0002] Lost connection with an agent.")
				// messages <- "INFO[0002] Lost connection with an agent."
				waitingForConnections(listener)
			}
			fmt.Printf(string(out))
		}
	}()

	teaTeeWhy, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer teaTeeWhy.Close()

	go func() {
		for ws := range teaTeeWhy.SIGWINCH() {
			fmt.Println("Resized", ws.W, ws.H)
		}
	}()

	for {
		key, err := teaTeeWhy.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		writer.WriteRune(key)
		writer.Flush()
	}
}
