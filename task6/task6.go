package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const ticketSize = 6
const (
	ticketTypeMoskow = iota + 1
	ticketTypePiter
)

func readLines(path string, ticketType int) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return -1, errors.New("something wrong with the file (maybe it is not exist)")
	}
	defer file.Close()
	ticketCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) != ticketSize {
			continue
		}
		if ticketType == ticketTypeMoskow && сheckMoskow(scanner.Text()) {
			ticketCount++
		} else if ticketType == ticketTypePiter && сheckPiter(scanner.Text()) {
			ticketCount++
		}
	}
	if err := scanner.Err(); err != nil {
		return -1, errors.New("one of the lines was too long!")
	}

	return ticketCount, scanner.Err()
}

func сheckPiter(s string) bool {
	countPart1 := 0
	countPart2 := 0
	var tmp int
	var err error

	for i, _ := range s {
		if tmp, err = strconv.Atoi(string(s[i])); err != nil {
			return false
		}
		if i%2 == 0 {
			countPart1 += tmp
		} else {
			countPart2 += tmp
		}
	}
	return countPart1 == countPart2
}
func сheckMoskow(s string) bool {
	countPart1 := 0
	countPart2 := 0
	var tmp int
	var err error
	for i, _ := range s {
		if tmp, err = strconv.Atoi(string(s[i])); err != nil {
			return false
		}
		if i < len(s)/2 {
			countPart1 += tmp
		} else {
			countPart2 += tmp
		}
	}
	return countPart1 == countPart2
}
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("HELP: This program count lucky tickets in the file tickets.txt. Print 1 to use 'Moskow' rule or 2 to use 'Piter' rule")
		return
	}
	n, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Argument must be digit!")
		return
	}
	answer, err := readLines("tickets.txt", n)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("There are %d lucky tickets \n", answer)
}
