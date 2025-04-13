package commands

import (
	"flag"
	"fmt"
	"io"
	"lotion/internal/state"
	"lotion/internal/utils"
	"os"
	"os/exec"
	"path"
	"strings"
)

type SearchCommand struct {
	query string
}

func (cmd SearchCommand) Execute() error {
	return cmd.execute()
}

func (cmd SearchCommand) Debug() {
	fmt.Printf("Query: %s\n", cmd.query)
}

func CreateSearchCommand(args []string) (Command, error) {
	var cmd SearchCommand

	flag := flag.NewFlagSet(string(SearchNote), flag.ContinueOnError)

	flag.StringVar(&cmd.query, "query", "", "query to search")

	err := flag.Parse(args)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd SearchCommand) execute() error {
	notebooksMap, err := state.GetAllNotes()
	if err != nil {
		return err
	}

	output, err := runExternalTool("fzf",
		fmt.Sprintf("-q %s", cmd.query),
		prepareFzfNotesListInput(notebooksMap))
	if err != nil {
		return err
	}

	noteName, notebook := extractNoteAndNotebookNames(output)
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

func extractNoteAndNotebookNames(note string) (string, string) {
	splitted := strings.Split(note, "/")
	notebook := splitted[0]
	noteName := splitted[1]

	return noteName, notebook
}

func runExternalTool(tool string, extraFlags string, input string) (string, error) {
	if !utils.IsToolExists(tool) {
		return "", fmt.Errorf("%s does not exist!", tool)
	}

	r, w, err := os.Pipe()
	if err != nil {
		return "", err
	}

	toolCmd := exec.Command(tool, extraFlags)
	toolCmd.Stdin = strings.NewReader(input)
	toolCmd.Stdout = w

	err = toolCmd.Run()
	if err != nil {
		return "", err
	}
	w.Close()

	fmt.Println("what tthe fuck")
	bytes, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	r.Close()

	return string(bytes), nil
}
