package state

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
)

const (
	LotionDir = ".lotion"
)

var lotionPath string

func init() {
	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	lotionPath = filepath.Join(user.HomeDir, LotionDir)

	if !isDirExists(lotionPath) {
		err = os.Mkdir(lotionPath, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func GetNotePath(notebook, noteName string) string {
	return path.Join(lotionPath, notebook, noteName+".md")
}

func IsNotebookExist(notebook string) bool {
	notebookPath := path.Join(lotionPath, notebook)
	fmt.Println("is notebook exist: ", notebook, notebookPath, isDirExists(notebookPath))
	return isDirExists(notebookPath)
}

func CreateNotebook(notebook string) error {
	notebookPath := path.Join(lotionPath, notebook)
	return os.Mkdir(notebookPath, os.ModePerm)
}

func IsNoteExist(notebook, noteName string) bool {
	return isFileExists(GetNotePath(notebook, noteName))
}

func CreateNote(notebook, noteName string) error {
	_, err := os.Create(GetNotePath(notebook, noteName))
	return err
}

func GetLotionPath() string {
	return lotionPath
}

func GetAllNotes() (map[string][]string, error) {
	notebooks, err := os.ReadDir(lotionPath)
	notebooksToNotesMap := make(map[string][]string)
	if err != nil {
		return nil, err
	}

	for _, dirEntry := range notebooks {
		if !dirEntry.IsDir() {
			continue
		}

		path := path.Join(lotionPath, dirEntry.Name())

		notes, err := getNoteNames(path)
		if err != nil {
			continue
		}

		notebooksToNotesMap[dirEntry.Name()] = notes
	}

	return notebooksToNotesMap, nil
}

func getNoteNames(notebookPath string) ([]string, error) {
	var noteNames []string

	entries, err := os.ReadDir(notebookPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			return nil, fmt.Errorf("Error: %s can not be a directory", entry.Name())
		}

		noteNames = append(noteNames, removeNotePrefix(entry.Name()))
	}

	return noteNames, nil
}
