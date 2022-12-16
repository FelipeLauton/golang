package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Declare nil slices
//
//  1. Declare the following slices as nil slices:
//
//    1. The names of your friends (names slice)
//
//    2. The distances to locations (distances slice)
//
//    3. A data buffer (data slice)
//
//    4. Currency exchange ratios (ratios slice)
//
//    5. Up/Down status of web servers (alives slice)
//
//
//  2. Print their type, length and whether they're equal to nil value or not.
//
//
// EXPECTED OUTPUT
//  names    : []string 0 true
//  distances: []int 0 true
//  data     : []uint8 0 true
//  ratios   : []float64 0 true
//  alives   : []bool 0 true
// ---------------------------------------------------------

func main() {
	var(
		names []string
		distances []int
		data []uint8
		ratios []float64
		alives []bool
	)

	fmt.Printf("names		: %T %d %v\n", names, len(names), names == nil)
	fmt.Printf("distances	: %T %d %v\n", distances, len(distances), distances == nil)
	fmt.Printf("data		: %T %d %v\n", data, len(data), data == nil)
	fmt.Printf("ratios		: %T %d %v\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives		: %T %d %v\n", alives, len(alives), alives == nil)
}
