package commands

import (
	"flag"
	"fmt"
	"io"
	"lotion/internal/state"
	"os"
	"os/exec"
	"path"
	"strings"
)

type OpenCommand struct {
	query string
}

func (cmd OpenCommand) Execute() error {
	return cmd.execute()
}

func (cmd OpenCommand) Debug() {
	fmt.Printf("Query: %s\n", cmd.query)
}

func CreateOpenCommand(args []string) (Command, error) {
	var cmd OpenCommand

	flag := flag.NewFlagSet(string(OpenNote), flag.ContinueOnError)

	flag.StringVar(&cmd.query, "query", "", "query to search")

	err := flag.Parse(args)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd OpenCommand) execute() error {
	_, err := exec.LookPath("fzf")
	if err != nil {
		return err
	}

	r, w, err := os.Pipe()
	if err != nil {
		return err
	}

	notebooksMap, err := state.GetAllNotes()
	if err != nil {
		return err
	}

	fzfInput := prepareFzfNotesListInput(notebooksMap)
	fmt.Println(fzfInput)

	fzf := exec.Command("fzf")
	fzf.Stdin = strings.NewReader(fzfInput)
	fzf.Stdout = w

	err = fzf.Run()
	if err != nil {
		return err
	}

	w.Close()

	bytes, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	splitted := strings.Split(string(bytes), "/")
	notebook := splitted[0]
	noteName := splitted[1]

	notePath := state.GetNotePath(notebook, noteName)
	err = runCmdFromTerminal(getEditorCmd(notePath))
	if err != nil {
		return err
	}

	return nil
}

func prepareFzfNotesListInput(notebooksMap map[string][]string) string {
	var input string
	for key, notes := range notebooksMap {
		for _, note := range notes {
			input += path.Join(key, strings.TrimSuffix(note, "\n")) + "\n"
		}
	}

	return input
}
