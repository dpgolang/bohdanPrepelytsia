package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	for {
		flag:=""
		fmt.Println("Would you to check now? (yes/y to add new one)")
		fmt.Scan(&flag)
		if flag = strings.ToLower(flag); flag != "y" && flag != "yes" {
			break
		}
		var a, b, c, d float64
		if !validateInput(&a, &b, &c, &d){
			fmt.Println("something wrong with one of params. It must be dighit > 0!")
			continue
		}
		var r1 = rectangle{a, b}
		var r2 = rectangle{c, d}
		if r1.putIn(r2) || r2.putIn(r1) {
			fmt.Println("It is possible to input one into another one")
			continue
		}
		fmt.Println("It is NOT possible to input one into another one")

	}
}

type rectangle struct {
	a, b float64
}

func (r1 rectangle) putIn(r2 rectangle) bool {
	p, q := getMax(r1.a, r1.b)
	a, b := getMax(r2.a, r2.b)
	if p <= a && q <= b {
		return true
	}
	formula := (2*q*a + (p*p-q*q)*math.Sqrt(p*p+q*q-a*a)) / (p*p + q*q)
	if p > a && b >= formula {
		return true
	}
	return false
}
func getMax(a, b float64) (float64, float64) {
	if a >= b {
		return a, b
	}
	return b, a
}
func validateInput(a, b, c, d *float64) bool {
	var err error
	fmt.Println("Now input sides of 2 rectangles (a x b;c x d)")
	
	fmt.Printf("Input a: ")
	_, err = fmt.Scan(a)
	if err!= nil {
		return false
	}

	fmt.Printf("Input b: ")
	_, err = fmt.Scan(b)
	if err!= nil {
		return false
	}

	fmt.Printf("Input c: ")
	_, err = fmt.Scan(c)
	if err!= nil {
		return false
	}

	fmt.Printf("Input d: ")
	_, err = fmt.Scan(d)
	if err!= nil {
		return false
	}

	return validateSides(*a,*b,*c,*d)
}
func validateSides(a, b, c, d float64) bool {
	return a > 0 && b > 0 && c > 0 && d > 0
}
