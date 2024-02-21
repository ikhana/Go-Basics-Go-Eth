// Hello World Start from here
package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"bufio"
	"log"
	"math/rand"
	"os"
)

var pl = fmt.Println

func main() {

	for x := 7; x >= 1; x-- {
		pl(x)
	}

	// While true loop (Infinite Loop) will be used for a guessing
	// game
	seedSecs := time.Now().Unix() // Returns seconds as int
	rand.Seed(seedSecs)

	randNum := rand.Intn(50) + 1
	for true {
		fmt.Print("Guess a number between 0 and 50 : ")
		pl("Random Number is :", randNum)
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if iGuess > randNum {
			pl("Lower")
		} else if iGuess < randNum {
			pl("Higher")
		} else {
			pl("You Guessed it")
			break
		}
	}

	arr := []int{1, 2, 3}

	for _, num := range arr {
		pl(num)

	}
}
