package main

import "fmt"

const prefOlaPT = "Olá, "

func Ola(nome string) string {
	if nome == "" {
		nome = "Mundo"
	}
	return prefOlaPT + nome
}

func main() {
	fmt.Println(Ola("Mundo"))
}
