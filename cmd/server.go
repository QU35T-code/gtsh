package cmd

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
	Keys   string `short:"k" long:"keys" description:"Path to folder with server.{pem,key}" default:"./certs" required:"true"`
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

func Server() {
	var listener net.Listener
	var err error

	if opts.Socket == "" {
		listener, err = newTLSListener()
	} else {
		listener, err = net.Listen("unix", opts.Socket)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		defer conn.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
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

func handleConnection(conn net.Conn) {
	reader, writer := bufio.NewReader(conn), bufio.NewWriter(conn)

	// A.B.C - Always Be Checking if there's new data to pull down on the wire
	go func() {
		for {
			out, err := reader.ReadByte()
			if err != nil {
				log.Fatal(err)
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
