package fcmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Run executes command with arguments
func Run(s string) {
	parts := strings.Split(s, " ")
	command := parts[0]
	args := parts[1:]

	out, err := exec.Command(command, args...).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error executing command", err)
		os.Exit(1)
	}
	fmt.Print(string(out))
}
