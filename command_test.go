package commander_test

import (
	"testing"

	"github.com/mikeyscode/commander"
	"github.com/stretchr/testify/assert"
)

func TestRunConsoleWithMultipleCommands(t *testing.T) {
	c := commander.Console{}
	c.Add("echo 'Hello'")
	c.Add("echo 'World'")

	res, err := c.Run("linux")
	if err != nil {
		t.Errorf("run failed %s", err)
	}

	assert.Contains(t, string(res), "Hello")
	assert.Contains(t, string(res), "World")
}

func TestRunConsoleWithInvalidOperatingSystem(t *testing.T) {
	c := commander.Console{}
	c.Add("echo 'Hello World'")

	_, err := c.Run("foo")

	assert.NotEmpty(t, err)
}
