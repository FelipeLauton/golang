package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)
func main() {
	var books [yearly]string

	books[0] = "Kafka's Revenge"
	books[1] = "Stay Golden"
	books[2] = "Everythingship"
	books[3] += books[0] + " 2nd Edition"

	fmt.Printf("books : %T\n", books)
	fmt.Println("books :", books)
	fmt.Printf("books : %q\n", books)
	fmt.Printf("books : %#v\n", books)
}