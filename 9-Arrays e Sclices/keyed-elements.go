package main

import "fmt"

func main() {
	const (
		ETH = iota
		WAN
	)

	rates := [...]float64{
		ETH: 25.5,  // ethereum
		WAN: 120.5, // wanchain
	}

	fmt.Printf("1 BTC is %g EHT\n", rates[ETH])
	fmt.Printf("1 BTC is %g WAN\n", rates[WAN])

	fmt.Printf("%#v\n", rates)
}
