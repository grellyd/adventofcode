package cli

import "regexp"

const (
	breakLine          = "===================="
	programName        = "Advent Of Code"
	usageCommand       = "go run main.go"
	programDescription = "Go implementation of the Advent of Code"
)

// A Command is an application instruction.
type Command struct {
	order       int
	long        string
	short       string
	description string
	regex       regexp.Regexp
	length      int
	value       string
}

// Present checks if a Command is present
func (c *Command) Present() bool {
	return c.value != ""
}

// ValidCommands to this program
func ValidCommands() (commands []*Command) {
	return []*Command{debug(), interactive(), help(), server()}
}

func server() *Command {
	return &Command{
		order:       2,
		long:        "--server",
		short:       "-s",
		description: "Server IP and port",
		regex:       *regexp.MustCompile("(-s|--server) ([0-9]{1,3}.?){4}:[0-9]{1,5}"),
		length:      2,
	}
}

func interactive() *Command {
	return &Command{
		order:       0,
		long:        "--interactive",
		short:       "-i",
		description: "Run the program in interactive mode",
		regex:       *regexp.MustCompile("-i|--interactive"),
		length:      1,
	}
}

func debug() *Command {
	return &Command{
		order:       1,
		long:        "--debug",
		short:       "-d",
		description: "Run the program in debug mode, with extra logging",
		regex:       *regexp.MustCompile("-d|--debug"),
		length:      1,
	}
}

func help() *Command {
	return &Command{
		// Help is handled by the cli wrapper
		order:       -1,
		long:        "--help",
		short:       "-h",
		description: "Display this help dialogue and exit",
		regex:       *regexp.MustCompile("-h|--help"),
		length:      1,
	}
}
