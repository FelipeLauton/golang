package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct{
	Nome string `json:"nome"`
	Raca string `json:"raca"`
	Idade int `json:"idade"`
}
func main() {
	c := cachorro{"Rex", "Korgi", 15}
	fmt.Println(c)
	
	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroEmJSON)
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string] string{
		"nome": "Toddy",
		"raca": "Pitbull",
	}
	fmt.Println(c2)

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil{
		log.Fatal(erro)
	}
	fmt.Println(cachorro2EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}
