package commander_test

import (
	"testing"

	"github.com/mikeyscode/commander"
	"github.com/stretchr/testify/assert"
)

func TestAddToConsole(t *testing.T) {
	c, _ := commander.NewConsole()
	c.Add("foo bar")

	assert.NotEmpty(t, c.Commands)
	assert.Contains(t, c.Commands, "foo bar")
}

func TestRunConsoleWithSingleCommand(t *testing.T) {
	c, _ := commander.NewConsole()
	c.Add("echo 'Hello'")

	res, err := c.Run()
	if err != nil {
		t.Errorf("run failed %s", err)
	}

	assert.Contains(t, string(res), "Hello")
}

func TestRunConsoleWithMultipleCommands(t *testing.T) {
	c, _ := commander.NewConsole()
	c.Add("echo 'Hello'")
	c.Add("echo 'World'")

	res, err := c.Run()
	if err != nil {
		t.Errorf("run failed %s", err)
	}

	assert.Contains(t, string(res), "Hello")
	assert.Contains(t, string(res), "World")
}

func TestRunConsoleWithInvalidCommand(t *testing.T) {
	c, _ := commander.NewConsole()
	c.Add("foo")

	_, err := c.Run()

	assert.NotEmpty(t, err)
}
