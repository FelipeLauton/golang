package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Nesse código temos um pequeno desafio de receber o nome de um usuario e retornar como seu nome e como ele está se sentindo.
// Primeiro inicio meu código trazendo o uma randomização do pacote rand que muda conforme o tempo avança.
// Depois criei uma variavel do tipo os.Args para receber os argumentos do usuario passado [1:] para pegar somente o que for passado.
// Após isso criei uma validação para verificar se o usuario digitou seu nome ou não, caso não o programa encerra.

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[your name]")
		return
	}
	
	name := args[0]

	moods := [5]string{
		"sad",
		"terrible",
		"happy",
		"awesome",
		"good",
	}
	n := rand.Intn(len(moods))

	fmt.Printf("%s feels %s", name, moods[n])
}
