// Hello World Start from here
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {

	var v1 float64 = 1.2

	var v2 = int(v1)

	pl(v2) // Will print 1

	v3 := "0000"

	v4, err := strconv.Atoi(v3)     // ASCKII to INT
	pl(v4, err, reflect.TypeOf(v4)) // Will print 0 <nil> int

	v5 := 40000

	v6 := strconv.Itoa(v5) // ineger to aski

	pl(reflect.TypeOf(v5), reflect.TypeOf(v6)) // Will print int string

	v7 := "3.1"

	if v8, err := strconv.ParseFloat(v7, 64); err == nil {
		pl(v8) // Will print 3.1

	}
}
