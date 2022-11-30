package main

import (
	"fmt"
)

func main() {

	// A utilização de arrays se faz necessaria para trazer mais performance em seus códigos.
	// Em Go, utilizamos comumente slices por não se limitarem a tamanhos pré definidos.
	// Abaixo temos a declaração de um Array literal com o declarador curto de declaração, o tamanho,
	// e os dados do Arrary separados por virgula.	
	books := [4]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	fmt.Println(books)
}
