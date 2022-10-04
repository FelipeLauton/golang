package main

import "fmt"

func main() {
	generica("String")
	generica(1)
	generica(true)
}

func generica(interf interface{}) {
	fmt.Println(interf)

	fmt.Println(1, 2, "string", false, true, float64(1234))

	mapa := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"string":     "string",
		true:         float32(12),
	}

	fmt.Println(mapa)

}
