package main

import "fmt"

func main() {
	// ages, first and last2 have the same backing arrays
	ages := []int{35, 15, 25}
	first := ages[0:1]
	last2 := ages[1:3]

	ages[0] = 55
	ages[1] = 10
	ages[2] = 20

	// grades and ages have separate backing arrays
	grades := []int{70, 99}
	grades[0] = 50

	fmt.Println("ages", ages)
	fmt.Println("ages[0:1]", first)
	fmt.Println("ages[1:3]", last2)
	fmt.Println("grades", grades)

	// let's create a new scope
	// 'cause I'm going to use variables with the same name
	{
		// ages and agesArray have the same backing arrays
		agesArray := [3]int{35, 15, 25}
		ages := agesArray[0:3]

		ages[0] = 100
		ages[2] = 50

		fmt.Println("agesArray", agesArray[:])
		fmt.Println("agesArray's ages", ages)
	}
}
