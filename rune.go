package main

import (
	"fmt"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {

	vRune := "OneTwoThree"

	pl("Characters count in Rune", utf8.RuneCountInString(vRune))

	// If we want to do the same with For Loop then

	for i, runeValue := range vRune {
		fmt.Printf("%d : %#U : %c\n ", i, runeValue, runeValue)
	}

}
