package main

import (
	// "flag"
	"fmt"
	"lotion/internal/commands"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: lotion [new|open] --notebook=[notebook name] [note name]")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmd, err := commands.CreateCommand(os.Args[2:], commands.CommandName(cmdName))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cmd.Debug()
	err = cmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
