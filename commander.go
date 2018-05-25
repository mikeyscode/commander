package commander

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// The Runnable interface defines the necessary requirements to run shell commands
type Runnable interface {
	Run([]string) ([]byte, error)
}

// Console acts as a storage unit for all actions to be performed on a single shell
type Console struct {
	Commands []string
	os       Runnable
}

// Add will append a command onto the command list
func (c *Console) Add(cmd string) {
	c.Commands = append(c.Commands, cmd)
}

// Run will execute the commands against the internal shell
func (c *Console) Run() ([]byte, error) {
	return c.os.Run(c.Commands)
}

var consoles = map[string]Runnable{
	"linux":   LinuxConsole{},
	"windows": WindowsConsole{},
}

// NewConsole creates a console instance dependent on OS at runtime
func NewConsole() (*Console, error) {
	if t, ok := consoles[runtime.GOOS]; ok {
		return &Console{os: t}, nil
	}

	return nil, fmt.Errorf("unknown runetime, could not create console instance")
}

// LinuxConsole acts as a shell instance for the linux terminal
type LinuxConsole struct{}

// Run will execute all current commands within the console on the same shell instance
func (l LinuxConsole) Run(commands []string) ([]byte, error) {
	commandSequence := strings.Join(commands, ";")
	command := exec.Command("/bin/sh", "-c", commandSequence)

	result, err := command.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// WindowsConsole acts as a shell instance for Windows cmd
type WindowsConsole struct{}

// Run will execute all current commands within the console on the same shell instance
func (w WindowsConsole) Run(commands []string) ([]byte, error) {
	commandSequence := strings.Join(commands, " && ")
	command := exec.Command("cmd", "/c", commandSequence)

	result, err := command.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return result, nil
}
