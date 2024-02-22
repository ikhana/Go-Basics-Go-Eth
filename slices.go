// Slices are like array but these are different in a way that these can be
// grow and the size can change
//
//	array declaration is arr := [size]type  while the slices can be declared like slice := make ([]type,size)
package main

import (
	"fmt"
)

func main() {

	slice := make([]string, 3)

	slice[0] = "I AM"
	slice[1] = "A GOLANG"
	slice[2] = "DEVELOPER"

	// Iterating over slice is as like we do over an arrray

	for i := 0; i < len(slice); i++ {

		fmt.Printf("WE Have at index %d : the string %s\n ", i, slice[i])

	}
	// Iterating Over range of slices

	for i, v := range slice {
		fmt.Printf("WE Have at index %d : the string %s\n ", i, v)
	}

	// OR IF WE DO NOT WANT THE INDEX

	for _, x := range slice {

		fmt.Println(x)
	}

	// INSIDE SLICES ARE PEICES OR SLICES OF ARRAY

	arr := []int{1, 2, 3, 4, 2}

	slice1 := arr[0:2]

	fmt.Println(slice1)

	fmt.Println(arr[:3], "\n", arr[2:])

	// IF WE CHNAGE THE ARRAY THE SLICES ARE ALSO CHANGED BUT WHAT IS MORE IMPORTANT IF WE CHNAGE THE SLICES THE ARRAYS ALSO GETS CNAGES

	fmt.Println("Before the slice change the array was : ", arr)

	slice1[0] = 8

	fmt.Println("After the slice change the array is now : ", arr)
}
