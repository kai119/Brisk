package base

import (
	"fmt"
)

// Command is an implementation of a BRISK command
type Command struct {
	// Run runs the command.
	Run func()

	// Name is the name of the command
	Name string

	// Usage is the one-line usage message.
	Usage string

	// Short is the short description shown in the 'brisk help' output.
	Short string

	// Long is the long message shown in the 'brisk help <this-command>' output.
	Long string

	//Commands is a list of commands that are displayed in the 'brisk help' output
	Commands []*Command

	// Args is a list of arguments that are parsed to the command when it is run
	Args map[string][]string

	// ArgsList is a list of command line flags that the command accepts
	ArgsList map[string]bool

	// CmdArgs is a list of arguments that were parsed into the command
	CmdArgs []string
}

// Brisk is the default command that is used to run all other commands
var Brisk = &Command{
	Usage: "brisk <command>",
	Name:  "brisk",
	Long:  `BRISK is a tool for managing BRISK source code.`,
}

// PrintHelp shows information on the command specified along with how to use the command
func (c *Command) PrintHelp() {
	if c.Name == "brisk" {
		fmt.Printf("%s\n", c.Long)
		fmt.Printf("Usage:\n\n %s\n\n", c.Usage)
		fmt.Printf("The commands are:\n\n")
		for _, command := range c.Commands {
			fmt.Printf("\t%s\t%s\n", command.Name, command.Short)
		}
		fmt.Println("\nUse \"brisk help <command>\" for more information about that command.")
	} else {
		fmt.Printf("Usage:\t%s\n", c.Usage)
		fmt.Printf("%s\n", c.Long)
	}
}

// ProcessArgs is a function that will find the list of arguments for each command
// line flag and stores them in a map
func (c *Command) ProcessArgs(args []string) error {
	c.Args = make(map[string][]string)
	var found []string
	var err error
	isFound := false
	if len(args) == 0 {
		return nil
	}
	for {
		cmdArg := args[0]
		if len(c.ArgsList) == 0 {
			return nil
		}
		for arg, hasText := range c.ArgsList {
			if cmdArg == arg {
				isFound = true
				if hasText {
					if args[1][0] == '-' {
						return fmt.Errorf("can't use flag as an argument to a flag")
					}
					found, args, err = findArgs(args[1:], c.ArgsList)
					if c.Args[arg] != nil {
						c.Args = nil
						return fmt.Errorf("can't use command flag %s more than once", arg)
					}
					c.Args[arg] = found
				} else {
					c.Args[arg] = []string{"true"}
					args = args[1:]
					if len(args) == 0 {
						args = nil
					}
				}
			}
		}
		if err != nil {
			c.Args = nil
			return fmt.Errorf("error reading arguments: %s", err)
		}
		if !isFound {
			c.Args = nil
			return fmt.Errorf("command argument %s does not exist", args[0])
		}
		if len(args) == 0 {
			return nil
		}
	}
}

func findArgs(args []string, argsList map[string]bool) ([]string, []string, error) {
	ptr := 0
	for _, cmdArg := range args {
		for arg := range argsList {
			if cmdArg == arg {
				return args[:ptr], args[ptr:], nil
			}
		}
		if cmdArg[0] == '-' {
			return nil, nil, fmt.Errorf("command argument %s does not exist", cmdArg)
		}
		ptr++
	}
	return args, nil, nil
}
