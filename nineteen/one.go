package nineteen

import (
	"fmt"
	"math"

	"github.com/pkg/errors"

	"github.com/grellyd/adventofcode/util"
)

// One puzzle solution
func One() error {
	fmt.Println("-----------")
	fmt.Println("Puzzle one")
	fmt.Println("-----------")

	data, err := util.ImportInts("nineteen/one_data.txt", "\n")
	if err != nil {
		return errors.Wrap(err, "unable to import data")
	}

	var sum int

	for _, module := range data {
		fuel, err := CalculateFuel(module)
		if err != nil {
			return errors.Wrap(err, "unable to calculate fuel")
		}
		sum += fuel
	}
	fmt.Printf("sum: %d\n", sum)

	sum = 0

	for _, module := range data {
		moduleSum, err := CalculateCombinedFuel(module)
		if err != nil {
			return errors.Wrap(err, "unable to calculate combined fuel")
		}
		sum += moduleSum
	}
	fmt.Printf("combined sum: %d\n", sum)

	fmt.Println("-----------")
	return nil
}

// CalculateFuel for a mass
func CalculateFuel(mass int) (fuel int, err error) {
	third := float64(mass) / 3
	rounded := math.Trunc(third)
	fuel = int(rounded) - 2
	if fuel < 0 {
		fuel = 0
	}
	return fuel, nil
}

// CalculateCombinedFuel including fuel mass
func CalculateCombinedFuel(mass int) (int, error) {
	sum := 0
	for {
		fuel, err := CalculateFuel(mass)
		if err != nil {
			return -1, errors.Wrap(err, "unable to calculate fuel")
		}
		if fuel == 0 {
			break
		}
		sum += fuel
		mass = fuel
	}
	return sum, nil
}
