// Hello World Start from here
package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {

	var iAge = 9
	if (iAge >= 1) && (iAge <= 18) {
		pl(" You have an important birthday")
	} else if (iAge == 21) && (iAge == 50) {
		pl(" You have an important birthday")
	} else if iAge >= 65 {
		pl(" You have an important birthday")
	} else {
		pl(" Not an important birthday")
	}
	pl(" !true represents ", !true)
}
