// Pointers are data types that lets refrence memory in memory location , to do so we have pass by refence the value

package main

import (
	"fmt"
)

//  Focus on both values valueItslef and valueByRefrence varaibles and see how things go

func chnageValue(valueItslef int, valueByrefrence *int) (int, int) {

	valueItslef += 1
	*valueByrefrence = *valueByrefrence * 3

	return valueItslef, *valueByrefrence
}

func main() {

	// Say we declare two values one for passByValue and other one for passByRefrence

	passByValue := 4
	passByRefrence := 4

	fmt.Println("Values before of pass by value and pass by refrence respectivel :\n", passByValue, "\n", passByRefrence)

	fmt.Println(chnageValue(passByValue, &passByRefrence))

	fmt.Println("Values after of pass by value and pass by refrence respectivel :\n", passByValue, "\n", passByRefrence)
}
