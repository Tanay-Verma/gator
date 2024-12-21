package command

import (
	"errors"
	"fmt"
)

type Command struct {
	name      string
	arguments []string
}

type Commands struct {
	cmds map[string]commandHandler
}

func NewCommand(name string, args []string) Command {
	return Command{
		name:      name,
		arguments: args,
	}
}

func (c *Commands) Register(name string, f commandHandler) {
	c.cmds[name] = f
}

func NewCommands() Commands {
	return Commands{
		cmds: make(map[string]commandHandler),
	}
}

func (c *Commands) Run(s *State, cmd Command) error {
	h, ok := c.cmds[cmd.name]
	if !ok {
		errMsg := fmt.Sprintf("Command %q not found", cmd.name)
		return errors.New(errMsg)
	}

	return h(s, cmd)
}
