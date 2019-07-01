package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func argsCheck(args []string) (int, int, error) {
	if len(args) == 0 {
		return 0, 0, errors.New("HELP: This program helps to show fibonacci series from n to m (n>=0, m>=n).If you want to start write integers n and m!")
	}
	if len(args) > 2 {
		return 0, 0, errors.New("there should be only one param n!")
	}
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, 0, errors.New("something wrong with your n! It must be digit!")
	}
	m, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, 0, errors.New("something wrong with your m! It must be digit!")
	}
	if n < 0 {
		return 0, 0, errors.New("numbers must be>=0!")
	}
	if m < n {
		return 0, 0, errors.New("m cannot be < n!")
	}

	return n, m, nil
}

func bineFunc(n int) int {
	//var i float64
	sqrt5 := math.Sqrt(5)
	phi := (sqrt5 + 1) / 2
	return int((math.Pow(phi, float64(n)))/sqrt5 + 0.5)
}
func fib(n int) int {
	x := 1
	y := 0
	for i := 0; i < n; i++ {
		x += y
		y = x - y
	}
	return y
}

func fib_rec(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib_rec(n-1) + fib_rec(n-2)
}
//output of result
//we need flag there to check if fib(i) 
//is the first element of our set
//to write commas correct
func outPut(n, m int) { 
	flag := true
	fmt.Printf("fibbo[%d,%d]:", n, m)
	for i := 0; fib(i) <= m; i++ {
		if fib(i) >= n && flag {
			flag = false
			fmt.Printf("%d", fib(i))
		} else if fib(i) > n {
			fmt.Printf(", %d", fib(i))
		}

	}
	fmt.Printf("\n")
}

func main() {
	args := os.Args[1:]
	n, m, err := argsCheck(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	outPut(n, m)
}
