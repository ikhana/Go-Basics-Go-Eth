package main

import (
	"fmt"
	"time"
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
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()) // As of writng and runing this code it has orinted 2024 February 20 13 39 48

}
