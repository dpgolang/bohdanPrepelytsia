package main

import (
	//"flag"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func validateArgs(args []string) (uint64, uint64, error) {
	var w, h uint64
	var err error

	if len(args) == 0 {
		return uint64(0), uint64(0), errors.New("HELP: you should set 2 params: width and height, it must be 0 < dighit <= 100")
	}
	if len(args) != 2 {
		return uint64(0), uint64(0), errors.New("two params should be set")
	}

	if w, err = strconv.ParseUint(args[0], 10, 32); err != nil {
		return uint64(0), uint64(0), errors.New("something wrong with width!")
	}
	if h, err = strconv.ParseUint(args[1], 10, 32); err != nil {
		return uint64(0), uint64(0), errors.New("something wrong with height!")
	}
	if validateSize(w, h) {
		return w, h, err
	}
	return uint64(0), uint64(0), errors.New("params must be 0 < dighit <= 100")
}
func validateSize(w, h uint64) bool {
	return w <= 100 && h <= 100
}

func main() {
	args := os.Args[1:]
	var w,h uint64
	var err error
	if w, h, err = validateArgs(args); err != nil {
		fmt.Println(err)
		return
	}
	outPut(w, h)
}

func outPut(width, height uint64) {
	for i := uint64(0); i < height; i++ {
		if i%2 != 0 {
			fmt.Printf(" ")
		}
		for j := uint64(0); j < width; j++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}
