package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("[your name] [positive|negative]")
		return
	}

	name, mood := args[0], args[1]

	moods := [...][3]string{
		{"happy", "awesome", "good"},
		{"sad", "terrible", "horrible"},
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods))

	var mi int
	if mood != "positive" {
		mi = 1
	} 
	fmt.Printf("%s feels %s", name, moods[mi][n])
}
