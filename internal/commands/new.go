package commands

import (
	"errors"
	"flag"
	"fmt"
	"lotion/internal/state"
	"os"
	"os/exec"
)

type NewCommand struct {
	noteName string
	noteBook string
}

func (cmd NewCommand) Execute() error {
	return cmd.execute()
}

func (cmd NewCommand) Debug() {
	fmt.Printf("noteName: %s, noteBook: %s\n", cmd.noteName, cmd.noteBook)
}

func CreateNewCommand(args []string) (Command, error) {
	var cmd NewCommand

	flag := flag.NewFlagSet(string(NewNote), flag.ContinueOnError)
	flag.StringVar(&cmd.noteBook, "notebook", "", "choose the notebook to create the note in (will use the last use notebook by default).")

	err := flag.Parse(args)
	if err != nil {
		return nil, err
	}

	cmd.noteName = flag.Arg(0)
	if cmd.noteName == "" {
		return nil, errors.New("note name can not be nil")
	}

	return cmd, nil
}

func (cmd NewCommand) execute() error {
	if err := cmd.createNote(); err != nil {
		return err
	}

	if err := cmd.openNoteInEditor(cmd.getPath()); err != nil {
		return err
	}

	return nil
}

func (cmd NewCommand) createNote() error {
	if !state.IsNotebookExist(cmd.noteBook) {
		err := state.CreateNotebook(cmd.noteBook)
		if err != nil {
			return err
		}
	}

	if !state.IsNoteExist(cmd.noteBook, cmd.noteName) {
		err := state.CreateNote(cmd.noteBook, cmd.noteName)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cmd NewCommand) openNoteInEditor(notePath string) error {
	editorCmd := getEditorCmd(notePath)
	return runCmdFromTerminal(editorCmd)
}

func (cmd NewCommand) getPath() string {
	return state.GetNotePath(cmd.noteBook, cmd.noteName)
}

func getEditorCmd(notePath string) *exec.Cmd {
	return exec.Command("nvim", notePath)
}

func runCmdFromTerminal(cmd *exec.Cmd) error {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
