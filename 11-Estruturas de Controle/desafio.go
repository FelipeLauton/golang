package main

import (
	"fmt"
	"os"
)

const (
	usage = "Usage: [username] [password]"
	errUser = "Usuário inválido para %q.\n"
	errPwd = "Senha inválida para %q.\n"
	acessOK = "Acesso garantido para %q.\n"

	user = "jack"
	pass = "1888"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]
	if u != user {
		fmt.Printf(errUser, u)
	} else if p != pass {
		fmt.Printf(errPwd, u)
	} else {
		fmt.Printf(acessOK, u)
	}
	
}