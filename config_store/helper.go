package config_store

import (
	"os/exec"
	"strings"
)

func createCommand(commandValues *string) *exec.Cmd {
	commandAndArgs := strings.Split(*commandValues, " ")
	if len(commandAndArgs) == 1 {
		return exec.Command(commandAndArgs[0])
	} else {
		return exec.Command(commandAndArgs[0], commandAndArgs[1:]...)
	}
}