package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const maxCountSymbols = 9
//const billion = uint64(100000000)
const (
	thousands = iota
	millions
)
//validate os.Args
func validateArgs(args []string) error {
	if len(args) == 0 {
		return errors.New("HELP: input number that is less than one billion")
	}
	if len(args) == 1 {
		return validateInput(args[0])
	}
	return errors.New("there must be only one argument!")
}
func deleteZerosOnStart(s string) (string,error){
	n, err := strconv.ParseUint(s, 10, 32)
    if err != nil {
        return "",errors.New("it must be a number!")
    }
    return strconv.FormatUint(n,10),nil
}
//validate string for beiing a number and how correct it was written
func validateInput(s string) error {
	if len(s) > 9 {
		return errors.New("there must be a number less than one billion!")
	}
	if _, err := strconv.ParseUint(s, 10, 32); err != nil {
		return errors.New("there must be a  umber less than one billion!")
	}
	return nil
}
//fill special slice for number
//slice for 123 gonna be [       1 2 3]
func fillNumber(s string) ([]string, error) {
	number := make([]string, maxCountSymbols, maxCountSymbols)
	n := maxCountSymbols
	diff := len(number) - len(s)
	for i := n - 1; i >= diff; i-- {
		number[i] = string(s[i-diff])
	}
	return number, nil
}
func addMillon(number string) string {
	switch number {
	case "1":
		return "миллион "
	case "2", "3", "4":
		return "миллиона "
	//case "0": return ""
	default:
		return "миллионов "
	}
}
func addThousands(number string) string {
	switch number {
	case "1":
		return "тысяча "
	case "2", "3", "4":
		return "тысячи "
	//case "0": return ""
	default:
		return "тысяч "
	}
}
func addUnits(number string, typeOfNumber int) string {
	switch number {
	case "3":
		return "три "
	case "4":
		return "четыре "
	case "5":
		return "пять "
	case "6":
		return "шесть "
	case "7":
		return "семь "
	case "8":
		return "восемь "
	case "9":
		return "девять "
	default:
		if typeOfNumber == thousands {
			switch number {
			case "1":
				return "одна "
			case "2":
				return "две "
			}
		}
		if typeOfNumber == millions {
			switch number {
			case "1":
				return "один "
			case "2":
				return "два "
			}
		}

	}
	return ""
}
func addHundreds(number string) string {
	switch number {
	case "1":
		return "сто "
	case "2":
		return "двести "
	case "3":
		return "триста "
	case "4":
		return "четыреста "
	case "5":
		return "пятьсот "
	case "6":
		return "шестьсот "
	case "7":
		return "семьсот "
	case "8":
		return "восемьсот "
	case "9":
		return "девятьсот "
	default:
		return ""
	}
}
func addTensUpTo20(number string) string {
	switch number {
	case "1":
		return "одиннадцать "
	case "2":
		return "двенадцать "
	case "3":
		return "тринадцать "
	case "4":
		return "четырнадцать "
	case "5":
		return "пятнадцать "
	case "6":
		return "шестнадцать "
	case "7":
		return "семнадцать "
	case "8":
		return "восемнадцать "
	case "9":
		return "девятнадцать "
	case "0":
		return "десять "
	default:
		return ""
	}
}
func addTens(number string) string {
	switch number {
	case "2":
		return "двадцать "
	case "3":
		return "тридцать "
	case "4":
		return "сорок "
	case "5":
		return "пятьдесят "
	case "6":
		return "шестьдесят "
	case "7":
		return "семьдесят "
	case "8":
		return "восемьдесят "
	case "9":
		return "девяносто "
	case "0": return ""
	default:
		return ""
	}
}
//connects all parts of answer
//we are going to through the numbers []string
//and check number by number 
func getResult(numbers []string) string {
	answer := ""
	flag := false
	for i := 0; i < maxCountSymbols-1; i++ {
		if numbers[i] != "" {
			flag = true
		}
	}
	if numbers[8] == "0" && !flag {
		return "ноль"
	}

	if (numbers[0] != "" || numbers[1] != "" || numbers[2] != "") {//&& //check if there are empty part in number like xxx123123
																	//to skil word "миллион" 
		answer += addHundreds(numbers[0])
		if numbers[1] == "1" {
			answer += addTensUpTo20(numbers[2])
			answer += addMillon("")
		} else {
			answer += addTens(numbers[1])
			answer += addUnits(numbers[2], millions)
			answer += addMillon(numbers[2])
		}
	}

	if (numbers[3] != "" || numbers[4] != "" || numbers[5] != "") && 	//check if there are empty part in number like xxx123
		(numbers[3] != "0" || numbers[4] != "0" || numbers[5] != "0") {//check if there are 0's in number part like xxx123 
		answer += addHundreds(numbers[3])								//to skip word "тысяча" 
		if numbers[4] == "1" {
			answer += addTensUpTo20(numbers[5])
			answer += addThousands("")
		} else {
			answer += addTens(numbers[4])
			answer += addUnits(numbers[5], thousands)
			answer += addThousands(numbers[5])
		}
	}
	answer += addHundreds(numbers[6])
	if numbers[7] == "1" {
		answer += addTensUpTo20(numbers[8])

	} else {
		answer += addTens(numbers[7])
		answer += addUnits(numbers[8], millions)
	}
	return answer
}

func main() {
	args := os.Args[1:]
	var answer string
	var number []string
	var err error
	if string(args[0][0]) == "-" {
		answer += "минус "
		args[0] = args[0][1:]
	}
	if err = validateArgs(args); err != nil {
		fmt.Println(err)
		return
	}
	if args[0],err = deleteZerosOnStart(args[0]); err!=nil{
		fmt.Println(err)
	}
	if number, err = fillNumber(args[0]); err != nil {
		fmt.Println(err)
		return
	}
	answer += getResult(number)
	fmt.Println(answer)
}
