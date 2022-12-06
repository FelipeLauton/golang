package main

import "fmt"

func main() {
	type bookcase [5]int
	type boloDeCenoura [5]int

	blue := bookcase{6, 9, 3, 2, 1}
	red := boloDeCenoura{6, 9, 3, 2, 1}

	fmt.Println("Are they equal? ")

	if boloDeCenoura(blue) == red {
		fmt.Println("YESSS")
	} else {
		fmt.Println("NOOOO")
	}

	fmt.Printf("blue: %#v\n", blue)
	fmt.Printf("red: %#v\n", red)
}