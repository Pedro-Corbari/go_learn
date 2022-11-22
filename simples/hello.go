package main

import "fmt"

const prefOlaPT = "Olá, "
const prefOlaES = "Hola, "
const prefOlaFR = "Bonjour, "
const prefOlaIN = "Hello, "

const espanhol = "Espanhol"
const frances = "Frances"
const ingles = "Ingles"

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	// prefixo := prefOlaPT
	return prefSaudacao(idioma) + nome
}

func prefSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = prefOlaFR

	case espanhol:
		prefixo = prefOlaES
	case ingles:
		prefixo = prefOlaIN

	default:
		prefixo = prefOlaPT
	}

	return

}

func main() {
	fmt.Println(Ola("Mundo", "Espanhol"))
}
