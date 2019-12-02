package aoc

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	"github.com/grellyd/adventofcode/cli"
	"github.com/grellyd/adventofcode/nineteen"
)

var debug bool
var interactive bool

// AdventOfCode main
func AdventOfCode(args []string) error {
	interactive, _ = strconv.ParseBool(args[0])
	debug, _ = strconv.ParseBool(args[1])
	puzzles := []func() error{
		nineteen.One,
	}
	if interactive {
		return interaction(puzzles)
	}
	return nil
}

func interaction(puzzles []func() error) error {

	fmt.Println("2019 only so far. Which puzzle?")

	if len(puzzles) == 0 {
		fmt.Println("No puzzles solved yet.")
		fmt.Println("Goodbye")
		return nil
	}
	fmt.Printf("Select from: 1-%d or 'all' or 'exit'\n", len(puzzles))

	for {
		fmt.Printf("[AOC]:")
		inputString := cli.ReadFromStdin()

		switch inputString {
		case "all":
			for _, puzzle := range puzzles {
				puzzle()
			}
		case "exit":
			fmt.Println("Goodbye")
			return nil
		default:
			i, err := strconv.ParseInt(inputString, 10, 0)
			if err != nil {
				if err == strconv.ErrSyntax {
					fmt.Println("Unkown command. Try again")
				} else {
					return errors.Wrap(err, "unable to process puzzle selection")
				}
			}
			return puzzles[i-1]()
		}
	}
}
