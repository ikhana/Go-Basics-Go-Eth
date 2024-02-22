package main

import (
	"fmt"
)

// Functions are building blocks of any programming language
// The syntax for go is something like func funcName(parameters) return type {}

// Simple Hello Print function

func helloWorld() {
	fmt.Print("HELLO WORLD!\n")

}

// Function that takes parameters and returns a value

func getSum(x int, y int) int {
	fmt.Println(x + y)
	return (x + y)
}

// Functions that return multiple values

func returnTwoValues(x int, y int) (int, int) {
	fmt.Println(x+y, x*y, "\n")
	return (x + y), (x * y)
}

func main() {

	helloWorld()
	getSum(10, 11)
	returnTwoValues(11, 4)

}
