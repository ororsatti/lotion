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

	fmt.Println(lotionPath)
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
