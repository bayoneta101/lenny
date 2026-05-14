package commands

import (
	"fmt"
	"os"
	"os/exec"
)

type listCmd struct{}

// what gets matched to the user input
func (listCmd) Name() string { return "list" }

func (listCmd) Execute(args []string) error {
	cmd := exec.Command("pacman", "-Qq")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("command failed", err)
	}

	fmt.Println(cmd.Stdout)
	return nil
}

// adds command to command map
func init() {
	RegisterCommand(listCmd{})
}
