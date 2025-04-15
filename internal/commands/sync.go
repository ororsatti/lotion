package commands

import (
	"errors"
	"fmt"
	"lotion/internal/state"

	"github.com/go-git/go-git/v5"
)

type SyncCommand struct {
	repo *git.Repository
}

func (cmd SyncCommand) Execute() error {
	return cmd.execute()
}

func (cmd SyncCommand) Debug() {
	fmt.Printf("Running Sync command")
}

func (cmd SyncCommand) execute() error {
	w, err := cmd.repo.Worktree()
	if err != nil {
		return err
	}

	s, err := w.Status()
	if err != nil {
		return err
	}

	for filePath, status := range s {
		if !isUnstagedStatus(status.Staging) {
			continue
		}

		w.Add(filePath)
	}

	commitMsg := fmt.Sprintf("changes in %d notes", len(s))

	commitHash, err := w.Commit(commitMsg, &git.CommitOptions{All: true})
	if err != nil {
		return err
	}

	fmt.Println(commitHash.String())

	return nil
}

func CreateSyncCommand(args []string) (Command, error) {
	syncCmd, err := initSyncCommand()
	if err != nil {
		return nil, err
	}

	return *syncCmd, nil
}

func initSyncCommand() (*SyncCommand, error) {
	var syncCmd SyncCommand

	repo, err := openOrInitRepo(state.GetLotionPath())
	if err != nil {
		return nil, err
	}

	syncCmd.repo = repo

	return &syncCmd, nil
}

func openOrInitRepo(path string) (*git.Repository, error) {
	var repo *git.Repository

	repo, err := git.PlainOpen(state.GetLotionPath())
	if errors.Is(err, git.ErrRepositoryNotExists) {
		repo, err = git.PlainInit(path, false)
		if err != nil {
			return nil, err
		}

	} else if err != nil {
		return nil, err
	}

	return repo, nil
}

func isUnstagedStatus(s git.StatusCode) bool {
	return s == git.Untracked || s == git.Modified
}
