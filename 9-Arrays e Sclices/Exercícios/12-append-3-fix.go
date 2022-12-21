package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Append #3 — Fix the problems
//
//  Fix the problems in the code below.
//
// BONUS
//
//  Simplify the code.
//
// EXPECTED OUTPUT
//
//  toppings: [black olives green peppers onions extra cheese]
//
// ---------------------------------------------------------

func main() {
	toppings := []string{"black olives", "green peppers"}

	var pizza []string
	pizza = append(pizza, toppings...)
	pizza = append(pizza, "onions")
	toppings = append(pizza, "extra", "cheese")

	fmt.Printf("toppings      : %s\n", toppings)
}
