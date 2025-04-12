package commands

import (
	"fmt"
)

type CommandName string

const (
	NewNote   CommandName = "new"
	ListNotes CommandName = "ls"
	OpenNote  CommandName = "open"
)

type Command interface {
	Execute() error
	Debug()
}

var commandConstructors = map[CommandName]func(args []string) (Command, error){
	NewNote: CreateNewCommand,
}

func CreateCommand(argsWithoutSubcommand []string, name CommandName) (Command, error) {
	constructor, found := commandConstructors[name]

	if !found {
		return nil, fmt.Errorf("Unsupported command: %s", name)
	}

	return constructor(argsWithoutSubcommand)
}
