package commands

import (
	"errors"
	"flag"
	"fmt"
	"lotion/internal/state"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

const defaultRemoteName = "origin"

type SyncCommand struct {
	repo      *git.Repository
	remoteURL string
}

func (cmd SyncCommand) Execute() error {
	return cmd.execute()
}

func (cmd SyncCommand) Debug() {
}

func (cmd SyncCommand) execute() error {
	if cmd.remoteURL != "" {
		return cmd.setRemote(cmd.remoteURL)
	}

	return cmd.performSync()
}

func CreateSyncCommand(args []string) (Command, error) {
	syncCmd, err := initSyncCommand()

	fs := flag.NewFlagSet(string(Sync), flag.ContinueOnError)
	fs.StringVar(&syncCmd.remoteURL, "remote", "", "set remote URL")

	err = fs.Parse(args)
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

func (cmd SyncCommand) setRemote(remoteURL string) error {
	_, err := cmd.repo.CreateRemote(&config.RemoteConfig{
		Name: defaultRemoteName,
		URLs: []string{remoteURL},
	})
	if err != nil {
		return err
	}

	return nil
}

func (cmd SyncCommand) performSync() error {
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

	_, err = w.Commit(commitMsg, &git.CommitOptions{All: true})
	if err != nil {
		return err
	}

	err = cmd.repo.Push(&git.PushOptions{RemoteName: defaultRemoteName})
	if errors.Is(err, git.NoErrAlreadyUpToDate) {
		return nil
	}

	return err
}

func isUnstagedStatus(s git.StatusCode) bool {
	return s == git.Untracked || s == git.Modified
}
