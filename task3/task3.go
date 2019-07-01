package main

import (
	"bufio"
	"errors"
	//"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type triangle struct {
	name   string
	Square float64
}

func main() {
	workCycle()
}
func getSquare(a, b, c float64) float64 {
	p := (a + b + c) / float64(2)
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}
func outPutTriangleList(list []triangle) {
	fmt.Println("==========Triangleslist:============")
	for i, value := range list {
		fmt.Printf("%d. [%s]: %.2f cm \n", i+1, value.name, value.Square)
	}
}
func stringValidation(str string) ([]string, error) {
	str = strings.Trim(str, "\n")
	splitedStr := strings.Split(str, ",")
	if len(splitedStr) != 4 {
		return nil, errors.New("you shoud write 4 params: name, a,b,c (a,b,c - sides of triangle)")
	}
	return splitedStr, nil
}
func triangleValidation(a, b, c float64) error {
	if a < b+c && b < a+c && c < a+b { //check if it is possible to build a triangle
		return nil
	} else {
		return errors.New("we cannot build triangle with these sides!")
	}
}
func sidesValidation(splitedStr []string) (float64, float64, float64, error) {
	var arr [3]float64
	for i := 1; i < len(splitedStr); i++ { //delete all spaces or \t
		splitedStr[i] = strings.TrimSpace(splitedStr[i])
	}
	for i := 1; i < len(splitedStr); i++ {
		if j, err := strconv.ParseFloat(splitedStr[i], 64); err == nil && j > 0 {
			arr[i-1] = j
		} else {
			return 0, 0, 0, fmt.Errorf("there is something wrong with your side #%d. It must be dighit > 0", i)
		}
	}
	return arr[0], arr[1], arr[2], nil
}
func workCycle() {
	triangleList := []triangle{}
	var err error
	for {
		flag := ""
		fmt.Println("Would you like to add triangle? (yes/y to add new one)")
		fmt.Scan(&flag)
		if flag = strings.ToLower(flag); flag != "y" && flag != "yes" {
			sort.Slice(triangleList, func(i, j int) bool { return triangleList[i].Square < triangleList[j].Square })
			outPutTriangleList(triangleList)
			break
		}
		fmt.Println("Input name, side 1, side 2, side 3:")
		var inputStr string
		var splitedStr []string
		in := bufio.NewReader(os.Stdin)
		if inputStr, err = in.ReadString('\n'); err != nil {
			fmt.Println("Something gone wrong while reading string")
			continue
		}
		if splitedStr, err = stringValidation(inputStr); err != nil {
			fmt.Println(err)
			continue
		}
		var a, b, c float64
		var triangle triangle
		triangle.name = splitedStr[0]
		if a, b, c, err = sidesValidation(splitedStr); err != nil {
			fmt.Println(err)
			continue
		}
		if err = triangleValidation(a, b, c); err != nil {
			fmt.Println(err)
			continue
		}
		triangle.Square = getSquare(a, b, c)
		triangleList = append(triangleList, triangle)
	}
}
