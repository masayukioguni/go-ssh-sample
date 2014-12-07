package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ssh", "root@xxx.xxx.xxx.xxx")
	cmd.Stdin = os.Stdin

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

}
