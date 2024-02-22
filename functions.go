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

// Functions that returns erros as well

func returnError(x float64, y float64) (ans float64, err error) {
	if y == 0 {

		return 0, fmt.Errorf("Denominator can not be zero please ")

	} else {
		return x / y, nil
	}

}

// Varaidic functions are those which recieves an unknown number of parameters

func variadicFunction(x ...int) int {
	sum := 0
	for _, num := range x {
		sum = sum + num
	}

	return sum
}

func main() {

	fmt.Println(returnError(3.22, 0))
	fmt.Println(returnError(3.22, 4))
	helloWorld()
	getSum(10, 11)
	returnTwoValues(11, 4)
	fmt.Println(variadicFunction(1, 2, 3, 3, 3))
}
