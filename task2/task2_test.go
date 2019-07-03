package main

import (
	"bytes"
	"errors"
	"testing"
)

func TestInPut(t *testing.T) {

	tests := []struct {
		input string
		err   error
	}{
		{
			"50 25\n", nil,
		},
		{
			"1 -1\n", errors.New("both of sides must be > 0"),
		},
		{
			"1 a\n", errors.New("cannot scan side b: strconv.ParseFloat: parsing \"\": invalid syntax"),
		},
	}
	for _, test := range tests {
		stdin := new(bytes.Buffer)
		stdin.Write([]byte(test.input))
		var r1 rectangle
		err := r1.InPut(stdin)
		if err != nil {
			if test.err == nil {
				t.Errorf("Expected no err, got %s err", err.Error())
			} else if test.err.Error() != err.Error() {
				t.Errorf("Expected err: %s, got %s err", test.err.Error(), err.Error())

			}
		}
	}
}

func TestValidateSides(t *testing.T) {
	tests := []struct {
		input rectangle
		err   error
	}{
		{
			rectangle{float64(3), float64(5)}, nil,
		},
		{
			rectangle{float64(0), float64(5)}, errors.New("both of sides must be > 0"),
		},
	}

	for _, test := range tests {

		err := test.input.ValidateSides()
		if err != nil {
			if test.err == nil {
				t.Errorf("Expected no err, got %s err", err.Error())
			} else if test.err.Error() != err.Error() {
				t.Errorf("Expected err: %s, got %s err", test.err.Error(), err.Error())

			}
		}
	}
}
func TestGetMax(t *testing.T) {
	tests := []struct {
		a, b     float64
		max, min float64
	}{
		{float64(3), float64(2), float64(3), float64(2)},
		{float64(2), float64(3), float64(3), float64(2)},
		{float64(0), float64(0), float64(0), float64(0)},
	}
	for _, test := range tests {
		max, min := GetMax(test.a, test.b)
		if max != test.max || min != test.min {
			t.Errorf("For %.2f, %.2f expected %.2f, %.2f,\n got %.2f, %.2f ", test.a, test.b, test.max, test.min, max, min)
		}
	}
}
func TestPutIn(t *testing.T) {
	tests := []struct {
		r1, r2 rectangle
		output bool
	}{
		{rectangle{float64(1), float64(1)}, rectangle{float64(1), float64(1)}, true},
		{rectangle{float64(3), float64(5)}, rectangle{float64(2), float64(2)}, true},
		{rectangle{float64(3), float64(1)}, rectangle{float64(2), float64(2)}, false},
	}
	for _, test := range tests {

		answ := test.r2.PutIn(test.r1)
		if answ != test.output {
			t.Errorf("For %v, %v expected %t,\n got %t ", test.r1, test.r2, test.output, answ)
		}
	}
}
