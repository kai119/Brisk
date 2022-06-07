package main

import (
	"fmt"
	"os"

	"github.com/kai119/Brisk/src/brisk/base"
	"github.com/kai119/Brisk/src/brisk/repl"
)

func init() {
	base.Brisk.Commands = []*base.Command{
		repl.CmdRepl,
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("please enter a command.\n\n")
		base.Brisk.PrintHelp()
		os.Exit(1)
	} else {
		if args[0] == "help" {
			if len(args) == 1 {
				base.Brisk.PrintHelp()
			} else {
				command := findCommand(args[1])
				command.PrintHelp()
			}
		} else {
			command := findCommand(args[0])
			err := command.ProcessArgs(args[1:])
			command.CmdArgs = args[1:]
			if err != nil {
				fmt.Printf("could not process command line flags: %s\n\n", err)
				command.PrintHelp()
				os.Exit(0)
			}
			command.Run()
		}
	}
}

func findCommand(com string) *base.Command {
	isFound := false
	var commandName *base.Command
	for _, command := range base.Brisk.Commands {
		if com == command.Name {
			isFound = true
			commandName = command
		}
	}
	if !isFound {
		fmt.Printf("\"%s\" is not a command.\n\n", com)
		base.Brisk.PrintHelp()
		os.Exit(1)
	}
	return commandName
}
