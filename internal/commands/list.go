package commands

import (
	"fmt"
	"lotion/internal/state"
	"path"
)

type ListCommand struct{}

func (cmd ListCommand) Execute() error {
	return cmd.execute()
}

func (cmd ListCommand) Debug() {}

func CreateListCommand(args []string) (Command, error) {
	return ListCommand{}, nil
}

func (cmd ListCommand) execute() error {
	lotionPath := state.GetLotionPath()

	notebooksMap, err := state.GetAllNotes()
	if err != nil {
		return err
	}

	var out string
	for notebookName, notes := range notebooksMap {
		for _, noteName := range notes {
			noteFullPath := path.Join(lotionPath, notebookName, noteName)
			out += noteFullPath + "\n"
		}
	}

	fmt.Print(out)
	return nil
}
