package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.ibm.com/Kai-Mumford-CIC-UK/brisk/src/brisk/base"
	"github.ibm.com/Kai-Mumford-CIC-UK/brisk/src/evaluator"
	"github.ibm.com/Kai-Mumford-CIC-UK/brisk/src/evaluator/object"
	"github.ibm.com/Kai-Mumford-CIC-UK/brisk/src/lexer"
	"github.ibm.com/Kai-Mumford-CIC-UK/brisk/src/lexer/token"
	"github.ibm.com/Kai-Mumford-CIC-UK/brisk/src/parser"
)

// PROMPT is the string that shall appear on each line of the REPL before user input
const (
	PROMPT = ">> "
	EXIT   = "exit"
)

// CmdRepl is the implementation of the base command struct for the REPL
var CmdRepl = &base.Command{
	Usage: "brisk repl [-c command list [-i]]",
	Name:  "repl",
	Short: "Start the BRISK REPL",
	Long: `
REPL starts the BRISK REPL
the REPL allows you to start a BRISK shell from the command
line. This will take BRISK code one line at a time, evaluate
the result of that line, and return the result to the console.
To exit the REPL, type "exit" or press Ctl+C

Command Arguments:
	-c [command list]:	A list of commands to run inside of the repl
	-i:	REPL won't exit after -c commands are processed

`,
}

func init() {
	CmdRepl.ArgsList = map[string]bool{
		"-c": true,
		"-i": false,
	}
	CmdRepl.Run = runRepl
}

func runRepl() {
	fmt.Printf("Hello! This is the BRISK programming language!\n")
	fmt.Printf("Type in any BRISK commands below\n\n")
	Start(os.Stdin, os.Stdout)
}

// Start will start the REPL with the specified reader and writer, exit the REPL
// by pressing Ctl+C
func Start(in io.Reader, out io.Writer) {
	output, isContinuing := processCommands(CmdRepl.Args)
	for _, out := range output {
		fmt.Printf("%+v\n", out)
	}
	if !isContinuing {
		os.Exit(0)
	}

	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)
		p, isContinuing := createParser(scanner)
		if !isContinuing {
			break
		}

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			_, err := io.WriteString(out, evaluated.Inspect())
			if err != nil {
				fmt.Printf("error writing string: %s\n", err)
			}
			_, err = io.WriteString(out, "\n")
			if err != nil {
				fmt.Printf("error writing string: %s\n", err)
			}
		}
	}
}

func processCommands(args map[string][]string) ([]token.Token, bool) {
	var commands []string
	var tokens []token.Token
	isContinuing := true
	if args["-c"] != nil {
		commands = args["-c"]
	} else if args["-i"] != nil {
		fmt.Printf("-i flag cannot be used without the -c flag. Starting REPL normally...\n")
	}
	if len(commands) > 0 {
		for _, arg := range commands {
			if arg == EXIT {
				return tokens, false
			}

			l := lexer.New(arg)

			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				tokens = append(tokens, tok)
			}
		}
		if args["-i"] == nil {
			isContinuing = false
		}
	}
	return tokens, isContinuing
}

func createParser(scanner *bufio.Scanner) (*parser.Parser, bool) {
	scanned := scanner.Scan()
	if !scanned {
		return nil, true
	}

	line := scanner.Text()
	if line == EXIT {
		return nil, false
	}
	if !strings.HasSuffix(line, ";") && !strings.HasSuffix(line, "{") {
		line += ";"
	}
	fullText := line

	curlyBrackets := strings.Count(line, "{") - strings.Count(line, "}")
	for curlyBrackets != 0 {
		fmt.Print(strings.Repeat("\t", curlyBrackets))
		scanned = scanner.Scan()
		if !scanned {
			return nil, true
		}
		line := scanner.Text()
		if line == EXIT {
			return nil, false
		}
		if !strings.HasSuffix(line, ";") && !strings.HasSuffix(line, "{") {
			line += ";"
		}
		fullText += "\n" + line
		curlyBrackets += strings.Count(line, "{") - strings.Count(line, "}")
	}
	l := lexer.New(fullText)
	return parser.New(l), true
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		_, err := io.WriteString(out, "\t"+msg+"\n")
		if err != nil {
			fmt.Printf("error writing string: %s\n", err)
		}
	}
}
