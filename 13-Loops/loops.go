package main

import (
	"fmt"
	"time"
)

func main() {	
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i", i)
		time.Sleep(time.Second)

	}

	fmt.Println(i)
	
	for j := 0; j < 10; j += 5{
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"joao", "silvia", "lucas"}

	for _, nomes := range nomes{
		fmt.Println(nomes)
	} 

	for indice, letra := range "PALAVRA"{
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome": "Leonardo",
		"sobrenome": "Silva",
	}

	for k, v := range usuario{
		fmt.Println(k, v)
	}
 }
