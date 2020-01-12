package main

import (
	"fmt"
	"strconv"
	"os"
)

// Questions:
//  - What type does the implicit
//  assignment give to errX and Y?
//
//  - What is the idiomamtic way
//  to handle conversion issues?

// Calculator
//   Set of method signatures
//   for a calculator

type Arithmetic interface {
	Add(x, y int) int
	Sub(x, y int) int
	Mult(x, y int) int
	Div(x, y int) int
}

type Calculator struct {}

func (calc Calculator) Add(x,y int) int {
	return x + y
}

func (calc Calculator) Sub(x,y int) int {
	return x - y
}

func (calc Calculator) Mult(x,y int) int {
	return x * y
}

func (calc Calculator) Div(x,y int) int {
	return x / y
}

func main() {
	if (len(os.Args) <= 1) {
		fmt.Println("Enter two values")
	} else {
		x, errX := strconv.Atoi(os.Args[1])
		y, errY := strconv.Atoi(os.Args[2])
		if errX == nil && errY == nil {
			var calculator Arithmetic = Calculator{}
			fmt.Println(calculator.Add(x, y))
			fmt.Println(calculator.Sub(x, y))
			fmt.Println(calculator.Mult(x, y))
			fmt.Println(calculator.Div(x, y))
		}
	}
}