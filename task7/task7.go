package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func GetAnswer(n int64) string {
	if n < 1 {
		return ""
	}
	maxRoot := int(math.Ceil(math.Sqrt(float64(n))) - 1)
	rootNambers := []string{}
	for i := 0; i < maxRoot; i++ {
		rootNambers = append(rootNambers, strconv.FormatInt(int64(i+1), 10))
	}
	return strings.Join(rootNambers, ",")
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
	fmt.Println(GetAnswer(int64(n)))
}
