package main

import "fmt"

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "warhammer40k",
		"star wars",
	}

	fmt.Println("items", items)

	top3 := items[:3]

	fmt.Println("top 3 items", top3)

	l := len(items)

	last4 := items[l-4:]

	fmt.Println("last 4 items", last4)

	fmt.Printf("slicing : %T %[1]q\n", items[2:3])
	fmt.Printf("slicing : %T %[1]q\n", items[2])
}
