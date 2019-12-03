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

	data[1] = 12
	data[2] = 2

	data, err = Calculate(data, 0)
	if err != nil {
		return errors.Wrap(err, "unable to calculate")
	}
	fmt.Println(data[0])
	fmt.Println("-----------")
	return nil
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
