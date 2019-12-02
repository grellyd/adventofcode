package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader

// Run the cli
func Run(f func([]string) error) (code int) {
	commands := ValidCommands()
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		for _, cmd := range commands {
			arg, value := getFullArg(args, i, cmd.length)
			if cmd.regex.MatchString(arg) {
				if cmd.short == "-h" {
					fmt.Println(usage(commands))
					return code
				}
				cmd.value = value
			}
		}
	}
	runArgs := convertCommands(commands)
	// Run the core program
	err := f(runArgs)
	if err != nil {
		code = 1
	}
	return code
}

// ReadFromStdin for cli interactive mode
func ReadFromStdin() string {
	reader := getReader()
	in, _ := reader.ReadBytes('\n')
	in = in[:len(in)-1]
	return string(in)
}

func getReader() *bufio.Reader {
	if reader == nil {
		reader = bufio.NewReader(os.Stdin)
	}
	return reader
}

func getFullArg(args []string, start int, length int) (string, string) {
	if length == 1 || len(args) < (start+length) {
		return args[start], "true"
	}
	return strings.Join(args[start:start+length], " "), strings.Join(args[start+1:start+length], " ")
}

func usage(commands []*Command) (usageString string) {
	var builder strings.Builder
	builder.WriteString(breakLine)
	builder.WriteString("\n")
	builder.WriteString(programName)
	builder.WriteString("\n")
	builder.WriteString(breakLine)
	builder.WriteString("\n")
	builder.WriteString(programDescription)
	builder.WriteString("\n\n")
	fmt.Fprintf(&builder, "Usage: %s\n\n", usageCommand)
	fmt.Fprint(&builder, "Valid options:\n\n")
	for _, cmd := range commands {
		fmt.Fprintf(&builder, "[ %s | %s ]: %s\n", cmd.short, cmd.long, cmd.description)
	}
	return builder.String()
}

func convertCommands(commands []*Command) (args []string) {
	args = make([]string, len(commands))
	for _, cmd := range commands {
		if cmd.Present() {
			args[cmd.order] = cmd.value
		}
	}
	return args
}
