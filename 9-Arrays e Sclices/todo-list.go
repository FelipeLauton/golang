package main

import "fmt"

func main() {
	var todo []string

	todo = append(todo, "sing")
	todo = append(todo, "run")
	todo = append(todo, "code")
	todo = append(todo, "play")

	fmt.Printf("todo %v\n", todo)
}
