// Hello World Start from here
package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main() {

	sWrod := "A word"

	replacer := strings.NewReplacer("A", "Another")
	sWord2 := replacer.Replace(sWrod)
	pl("length of the string is ", len(sWord2))

	pl("Contains Another :", strings.Contains(sWord2, "Another"))
	pl("0 index of string contains  :", strings.Index(sWord2, "o"))
	pl("0 index of string contains  :", strings.Replace(sWord2, "o", "0", -1))

	sWord3 := "\n Some Words\n"

	sWord3 = strings.TrimSpace(sWord3)

}
