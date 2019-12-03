package nineteen_test

import (
	"testing"

	"github.com/grellyd/adventofcode/nineteen"
)

func TestOnePartOne(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, test := range tests {
		result, err := nineteen.CalculateFuel(test.input)
		if err != nil {
			t.Errorf("failed to run test: %s", err.Error())
		}
		if result != test.want {
			t.Errorf("output does not match: '%d' != '%d'", result, test.want)
		}
	}
}

func TestOnePartTwo(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, test := range tests {
		result, err := nineteen.CalculateCombinedFuel(test.input)
		if err != nil {
			t.Errorf("failed to run test: %s", err.Error())
		}
		if result != test.want {
			t.Errorf("output does not match: '%d' != '%d'", result, test.want)
		}
	}
}
