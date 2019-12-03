package nineteen_test

import (
	"testing"

	"github.com/grellyd/adventofcode/nineteen"
)

func TestTwoPartOne(t *testing.T) {
	var tests = []struct {
		input  []int
		ip     int
		output []int
	}{
		{[]int{1, 0, 0, 0, 99}, 0, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, 0, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, 0, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, 0, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, test := range tests {
		result, err := nineteen.Calculate(test.input, test.ip)
		if err != nil {
			t.Errorf("unable to calculate: %s", err.Error())
		}
		if len(result) != len(test.output) {
			t.Errorf("length of output does not match")
		} else {
			for i, v := range result {
				if v != test.output[i] {
					t.Log(test.output)
					t.Log(result)
					t.Errorf("output at %d does not match: '%d' != '%d'", i, result[i], test.output[i])
				}
			}
		}
	}
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		input  []int
		ip     int
		output []int
	}{
		{[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, 0, []int{1, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
	}

	for _, test := range tests {
		result, _ := nineteen.Add(test.input, test.ip)
		if len(result) != len(test.output) {
			t.Errorf("length of output does not match")
		} else {
			for i, v := range result {
				if v != test.output[i] {
					t.Errorf("output does not match: '%d' != '%d'", result[i], test.output[i])
				}
			}
		}
	}
}

func TestMultiply(t *testing.T) {
	var tests = []struct {
		input  []int
		ip     int
		output []int
	}{
		{[]int{1, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, 4, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
		{[]int{2, 0, 0, 0, 99}, 0, []int{4, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, 0, []int{2, 3, 0, 6, 99}},
	}

	for _, test := range tests {
		result, _ := nineteen.Multiply(test.input, test.ip)
		if len(result) != len(test.output) {
			t.Errorf("length of output does not match")
		} else {
			for i, v := range result {
				if v != test.output[i] {
					t.Errorf("output does not match: '%d' != '%d'", result[i], test.output[i])
				}
			}
		}
	}
}
