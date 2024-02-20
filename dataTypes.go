// Hello World Start from here
package main

import (
	"fmt"
	"reflect"
)

func main() {

	// int, float64, string,  rune,bool...
	// Default type 0, 0.0, "", false,

	fmt.Println(reflect.TypeOf(true)) // bool

	fmt.Println(reflect.TypeOf(2.3)) //float64

	fmt.Println(reflect.TypeOf(3)) //int

	fmt.Println(reflect.TypeOf("Hello")) // string
}
