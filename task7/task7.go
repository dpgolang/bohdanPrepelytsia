package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func MakeSliceResult(n uint64)  []string{
	var answer []string
	if n == 1 {
		answer=make([]string,1)
		answer[0] = ""
		return answer
	}
	countOfElements:= int(math.Ceil(math.Sqrt(float64(n)))-1) //get count of elements (round to the least integer value greater than or equal sqrt(n) and -1)
	fmt.Println(countOfElements)
	answer=make([]string,countOfElements)
	for i := 1; i <= countOfElements; i++ {
		answer[i-1] = strconv.FormatInt(int64(i), 10)
	}
	return answer
}
func ArgsCheck(args []string) (uint64, error) {
	if len(args) == 0 {
		return 0, errors.New("HELP: This program helps to show natural numbers that have square less than inputed n. If you want to start write integer n!")
	}
	if len(args) > 1 {
		return 0, errors.New("there should be only one param n!")
	}
	n, err := strconv.ParseUint(args[0], 10, 64) 
	if err != nil {
		return 0, fmt.Errorf("something wrong with your n: %s", err)
	}
	return n, nil
}
func main() {
	args := os.Args[1:]
	n, err := ArgsCheck(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	answ:= MakeSliceResult(n)
	outPut:=strings.Join(answ,",")
	fmt.Println(outPut)
}
