// main.go
package main

import (
	"os"
	"os/exec"
	"strings"
)

func main() {
	ipAddr := "*"
	if len(os.Args) >= 2 {
		ipAddr = os.Args[1]
	}

	out, _ := exec.Command("ping", ipAddr).Output()
	if strings.Contains(string(out), "Request timed out") {
		start("IT'S DOWN")
	} else {
		start("IT'S ALIVEEE")
	}
	//fmt.Println(string(out))
	//start(ipAddr)
}
