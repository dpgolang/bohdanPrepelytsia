package main

import (
	"testing"
	"errors"
)

func TestGetAnswer(t *testing.T) {
	type testStruct struct {
		square int64
		expectedNumbers string
	}

  answers := []testStruct{
    {
      square: -5,
      expectedNumbers: "",
    },
    {
      square: 0,
      expectedNumbers: "",
    },
    {
      square: 1,
      expectedNumbers: "",
    },
    {
      square: 2,
      expectedNumbers: "1",
    },
    {
      square: 9,
      expectedNumbers: "1,2",
    },
    {
      square: 24,
      expectedNumbers: "1,2,3,4",
    },
  }

  for _, answer := range answers {
    gotNumbers := GetAnswer(answer.square)
    if gotNumbers != answer.expectedNumbers {
      t.Errorf("For %s, expected: %s, got %s", answer.testName, answer.expectedNumbers, gotNumbers)
    }
  }
}
func TestArgsCheck(t *testing.T) {
  tests := []struct{
    args []string
    n    uint64
    err  error
  }{
    {
      args: []string{},
      n: 0,
      err: errors.New("HELP: This program helps to show natural numbers that have square less than inputed n. If you want to start write integer n!"),
  	},
  	{
      args: []string{"512","123"},
      n: 0,
      err: errors.New("there should be only one param n!"),
  	},
  	{
      args: []string{"1"},
      n: 1,
      err: nil,
  	},
	{
      args: []string{"a"},
      n: 0,
      err: errors.New("something wrong with your n: strconv.ParseUint: parsing \"a\": invalid syntax"),
  	},
  	{
      args: []string{"-1"},
      n: 0,
      err: errors.New("something wrong with your n: strconv.ParseUint: parsing \"-1\": invalid syntax"),
  	},
   }

  for _, test := range tests {
    v, err := ArgsCheck(test.args)
    if err != nil {
      if test.err == nil {
        t.Errorf("Expected no err, got %s err", err.Error())
      } else if test.err.Error() != err.Error() {
        t.Errorf("Expected err: %s, got %s err", test.err.Error(), err.Error())
      }
    } else {
      if v != test.n {
        t.Errorf("Expected %d, got %d", test.n, v)
      }
    }
  }
}
