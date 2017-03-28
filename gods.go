package main

import (
	"os/exec"
	"strings"
	"time"
)

func main() {
	for {
		var status = []string{
			time.Now().Local().Format("Mon Jan 01 | 15:04:05"),
		}
		exec.Command("xsetroot", "-name", strings.Join(status, " | ")).Run()

		var now = time.Now()
		time.Sleep(now.Truncate(time.Second).Add(time.Second).Sub(now))
	}
}
