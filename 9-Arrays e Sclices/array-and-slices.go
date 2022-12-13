package main

import "fmt"

func main() {
	var books [5]string

	books[0] = "dracula"
	books[1] = "star wars"

	games := []string{"sekiro", "sta wars"}

	fmt.Println(books, games)
}