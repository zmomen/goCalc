package main

import (
	"fmt"
)

// Calculator
//   Set of method signatures
//   for a calculator
//   Each method returns a 64
// 	 -bit, integer value

type Arithmetic interface {
	Add(x, y int64) int64
	Sub(x, y int64) int64
	Mult(x, y int64) int64
	Div(x, y int64) int64
}

type Calculator struct {}

func (calc Calculator) Add(x,y int64) int64 {
	return x + y
}

func (calc Calculator) Sub(x,y int64) int64 {
	return x - y
}

func (calc Calculator) Mult(x,y int64) int64 {
	return x * y
}

func (calc Calculator) Div(x,y int64) int64 {
	return x / y
}

func main() {
	var calculator Arithmetic = Calculator{}
	var x, y int64 = 5, 6
	fmt.Println(calculator.Add(x, y))
	fmt.Println(calculator.Sub(x, y))
	fmt.Println(calculator.Mult(x, y))
	fmt.Println(calculator.Div(x, y))
}