package commander

import (
	"os/exec"
	"strings"
)

// Console acts as a storage unit for all actions to be performed on a single shell
type Console struct {
	Commands []string
}

// Add will append a command onto the command list
func (c *Console) Add(cmd string) {
	c.Commands = append(c.Commands, cmd)
}

// Run will execute all current commands within the console on the same shell instance
func (c *Console) Run() ([]byte, error) {
	commandSequence := strings.Join(c.Commands, ";")
	command := exec.Command("/bin/sh", "-c", commandSequence)

	result, err := command.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return result, nil
}
