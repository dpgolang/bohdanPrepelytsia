package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func outPut(n int) {
	if n == 1 {
		fmt.Println("")
		return
	}
	fmt.Printf("1")
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		fmt.Printf(", %d", i)
	}
	fmt.Printf("\n")
}
func argsCheck(args []string) (int, error) {
	if len(args) == 0 {
		return 0, errors.New("HELP: This program helps to show natural numbers that have square less than inputed n. If you want to start write integer n!")
	}
	if len(args) > 1 {
		return 0, errors.New("there should be only one param n!")
	}
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, errors.New("something wrong with your n! It must be digit!")
	}
	return n, nil
}
func main() {
	args := os.Args[1:]
	n, err := argsCheck(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	outPut(n)
}
