package commands

import "fmt"

type SyncCommand struct{}

func (cmd SyncCommand) Execute() error {
	return cmd.execute()
}

func (cmd SyncCommand) Debug() {
	fmt.Printf("Running Sync command")
}

func (cmd SyncCommand) execute() error {
	return nil
}

func CreateSyncCommand(args []string) (Command, error) {
	return SyncCommand{}, nil
}
