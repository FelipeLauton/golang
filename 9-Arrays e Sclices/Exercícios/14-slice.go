package main

import "fmt"

func main() {

	ages := []int{35, 15, 25}
	fmt.Println("ages", ages)

	fmt.Println("ages[0:0]", ages[0:0])

	for i := 1; i < 4; i++ {
		txt := fmt.Sprintf("ages[%d:%d]", 0, i)
		fmt.Println(txt, ages[0:i])
	}

	fmt.Println("append", append(ages, 50))
}
