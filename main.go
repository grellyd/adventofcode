package main

import (
	"os"

	"github.com/grellyd/adventofcode/aoc"
	"github.com/grellyd/adventofcode/cli"
)

func main() {
	os.Exit(cli.Run(aoc.AdventOfCode))
}
