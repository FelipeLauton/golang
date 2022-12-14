package main

import "fmt"

func main() {
	var books [5]string

	books[0] = "dracula"
	books[1] = "1984"
	books[2] = "island"

	newBooks := [5]string{"ulysses", "fire"}
	books = newBooks

	games := []string{"sekiro", "fallen order"}
	newGames := []string{"pacman", "doom"}

	newGames = games
	games = nil

	games = []string{}

	var ok string
	for i, game := range games {
		if game != newGames[i] {
			ok = "not "
			break
		}
	}

	if len(games) != len(newGames) {
		ok = "not "
	}

	fmt.Printf("games and newGames are %sequal\n\n", ok)

	fmt.Printf("books			: %#v\n", books)
	fmt.Printf("new books		: %#v\n", newBooks)
	fmt.Printf("games			: %#v\n", games)

	fmt.Printf("games			: %T\n", games)
	fmt.Printf("games' len		: %d\n", len(games))
	fmt.Printf("nil? 			: %t\n", games == nil)
}