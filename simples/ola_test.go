package main

import "testing"

func TestOla(t *testing.T) {

	verificaMsgCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)

		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("pedro", "Portugues")
		esperado := "Olá, pedro"
		verificaMsgCorreta(t, resultado, esperado)

	})

	t.Run("diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "Portugues")
		esperado := "Olá, Mundo"
		verificaMsgCorreta(t, resultado, esperado)

	})

	t.Run(
		"em espanhol", func(t *testing.T) {
			resultado := Ola("Emanoel", "Espanhol")
			esperado := "Hola, Emanoel"
			verificaMsgCorreta(t, resultado, esperado)
		})

	t.Run(
		"em frances", func(t *testing.T) {
			resultado := Ola("Gabriel", "Frances")
			esperado := "Bonjour, Gabriel"
			verificaMsgCorreta(t, resultado, esperado)
		})

	t.Run(
		"em ingles", func(t *testing.T) {
			resultado := Ola("Julia", "Ingles")
			esperado := "Hello, Julia"
			verificaMsgCorreta(t, resultado, esperado)
		})
}
