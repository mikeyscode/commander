package commander

import (
	"fmt"
	"os/exec"
	"strings"
)

// Console acts as a storage unit for all actions to be performed on a single shell
type Console struct {
	Commands []string
}

// os console standard properties
type osConsole struct {
	cmd,
	arg,
	sep string
}

// windows acts as a shell instance for Windows Cmd
type windows struct {
	osConsole
}

// linux acts as a shell instance for Linux Cmd
type linux struct {
	osConsole
}

func (o osConsole) GetCommand() string {
	return o.cmd
}

func (o osConsole) GetArg() string {
	return o.arg
}

func (o osConsole) GetSeperator() string {
	return o.arg
}

type runnable interface {
	GetCommand() string
	GetArg() string
	GetSeperator() string
}

var consoles = map[string]runnable{
	"linux": linux{
		osConsole: osConsole{
			cmd: "/bin/sh",
			arg: "-c",
			sep: ";",
		},
	},
	"windows": windows{
		osConsole: osConsole{
			cmd: "Cmd",
			arg: "/c",
			sep: "&&",
		},
	},
}

// Add will append a command onto the command list
func (c *Console) Add(cmd string) {
	c.Commands = append(c.Commands, cmd)
}

// Run will execute the commands against the internal shell
func (c *Console) Run(goos string) ([]byte, error) {
	if _, ok := consoles[goos]; !ok {
		return nil, fmt.Errorf("unknown runtime, could not create console instance")
	}

	return c.run(consoles[goos], c.Commands)
}

// Run will execute all current commands within the console on the same shell instance
func (c *Console) run(r runnable, commands []string) ([]byte, error) {
	commandSequence := strings.Join(commands, r.GetSeperator())

	cmd := r.GetCommand()
	arg := r.GetArg()

	return exec.Command(cmd, arg, commandSequence).CombinedOutput()
}
