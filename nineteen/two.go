package nineteen

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/grellyd/adventofcode/util"
)

// Two puzzle solution
func Two() error {
	fmt.Println("-----------")
	fmt.Println("Puzzle two")
	fmt.Println("-----------")

	data, err := util.ImportInts("nineteen/two_data.txt", ",")
	if err != nil {
		return errors.Wrap(err, "unable to import data")
	}

	output, err := RunCalculation(data, 12, 2)
	if err != nil {
		return errors.Wrap(err, "failed to run calculation")
	}

	fmt.Printf("part1: %d\n", output)

	data, err = util.ImportInts("nineteen/two_data.txt", ",")
	if err != nil {
		return errors.Wrap(err, "unable to import data")
	}

	noun, verb := FindInputs(data, 19690720)
	fmt.Printf("part2: (%d * %d) + %d = %d\n", 100, noun, verb, (100*noun)+verb)
	fmt.Println("-----------")
	return nil
}

// FindInputs on the calculation engine
func FindInputs(data []int, result int) (int, int) {
	dataCopy := make([]int, len(data))

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			copy(dataCopy, data)
			output, err := RunCalculation(dataCopy, i, j)
			if err != nil {
			} else {
				fmt.Printf("%d and %d produce: %d\n", i, j, output)
				if output == result {
					return i, j
				}
			}
		}
	}

	return -1, -1
}

// RunCalculation on data with a noun and verb
func RunCalculation(data []int, noun int, verb int) (int, error) {
	data[1] = noun
	data[2] = verb

	data, err := Calculate(data, 0)
	if err != nil {
		return 0, errors.Wrap(err, "unable to calculate")
	}
	return data[0], nil
}

// Calculate computer output
func Calculate(data []int, ip int) ([]int, error) {
	for {
		switch data[ip] {
		case 1:
			data, ip = Add(data, ip)
		case 2:
			data, ip = Multiply(data, ip)
		case 99:
			return data, nil
		default:
			return nil, errors.New("unknown opcode")
		}
	}
}

// Add for op 1
func Add(data []int, ip int) ([]int, int) {
	a := data[data[ip+1]]
	b := data[data[ip+2]]
	dest := data[ip+3]
	result := a + b
	data[dest] = result
	ip += 4
	return data, ip
}

// Multiply for op 2
func Multiply(data []int, ip int) ([]int, int) {
	a := data[data[ip+1]]
	b := data[data[ip+2]]
	dest := data[ip+3]
	result := a * b
	data[dest] = result
	ip += 4
	return data, ip
}
