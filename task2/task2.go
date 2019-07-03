package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

//сделать метод-валидацию для прямоугольника,читать для стороны
func main() {
	for {
		flag := ""
		fmt.Println("Would you to check now? (yes/y to add new one)")
		fmt.Scan(&flag)
		if flag = strings.ToLower(flag); flag != "y" && flag != "yes" {
			break
		}
		var r1, r2 rectangle
		if err := r1.InPut(os.Stdin); err != nil {
			fmt.Println(err)
			continue
		}

		if err := r2.InPut(os.Stdin); err != nil {
			fmt.Println(err)
			continue
		}
		if r1.PutIn(r2) || r2.PutIn(r1) {
			fmt.Println("It is possible to input one into another one")
			continue
		}
		fmt.Println("It is NOT possible to input one into another one")

	}
}

type rectangle struct {
	a, b float64
}

//input and check params for sides of rectangle
func (r *rectangle) InPut(reader io.Reader) (err error) {
	fmt.Println("Input sides of rectangle (a x b)")
	fmt.Printf("Input a:")
	_, err = fmt.Fscanf(reader, "%f", &r.a)
	if err != nil {
		return fmt.Errorf("cannot scan side a: %s", err)
	}

	fmt.Printf("Input b:")
	_, err = fmt.Fscanf(reader, "%f", &r.b)
	if err != nil {
		return fmt.Errorf("cannot scan side b: %s", err)
	}
	if err = r.ValidateSides(); err != nil {
		return err
	}
	return nil
}

//sides cannot be negatuve
func (r rectangle) ValidateSides() (err error) {
	if r.a > 0 && r.b > 0 {
		return nil
	}
	return errors.New("both of sides must be > 0")
}

//try to put r1 to r2
func (r1 rectangle) PutIn(r2 rectangle) bool {
	p, q := GetMax(r1.a, r1.b)
	a, b := GetMax(r2.a, r2.b)
	if p <= a && q <= b {
		return true
	}
	formula := (2*q*a + (p*p-q*q)*math.Sqrt(p*p+q*q-a*a)) / (p*p + q*q)
	return p > a && b >= formula
}

func GetMax(a, b float64) (float64, float64) {
	if a >= b {
		return a, b
	}
	return b, a
}
