// Questions:
//
//  - What type does the implicit
//  assignment give to errX and Y?
//
//  - What is the idiomamtic way
//  to handle conversion issues?
//
//  - How do array literals and
//  slices differ?
//
//  - How does append() work
//  under-the-hood?

package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
)

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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter two numbers: ")

	text, err := reader.ReadString('\n')
	if (err != nil) {
		fmt.Println("Error: ReadString failed")
		return
	}

	var nums []int
	var prev int = 0

	for i, elem := range text {
		if elem == ' ' || elem == '\n' {
			num, err := strconv.Atoi(text[prev:i])
			if (err != nil) {
				fmt.Println("Error: String conversion")
			}
			nums = append(nums, num)
			prev = i + 1
		}
	}

	var calculator Arithmetic = Calculator{}
	fmt.Println(calculator.Add(nums[0], nums[1]))
}