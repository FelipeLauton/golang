package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	guess := 10

	for turn := 0; turn < 5; turn++{
		n := rand.Intn(guess + 1)

		if n == guess {
			fmt.Println("YOU WIN")
			return
		}
	}
	fmt.Println("You Lost... try again?")
}