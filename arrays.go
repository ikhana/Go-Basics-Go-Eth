// Arays are coleection of same data type
// The most important is that the size of the array can not be chnaged
package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {

	var simpleArr [4]int

	simpleArr[0] = 1

	simpleArr1 := []int{1, 2, 1, 1, 1, 1, 1}

	pl("At index 0 the vaue is :", simpleArr1[0]) // will print 1

	// Running the for loop with Range

	for i, v := range simpleArr1 {

		fmt.Printf("At index %d : the value exists is %d\n ", i, v)

	}

	// Iterating over array for loops example

	for i := 0; i <= len(simpleArr1); i++ {
		fmt.Printf("At index %d : the value exists is %d\n ", i, simpleArr1[i])
	}
	// Note that if you put the first loop in the last the after while loop wil not be executed becuase when the go executes without range with index, it will not find the
	// 7th index and it will give run time error
}
