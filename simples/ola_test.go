package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("pedro")
	esperado := "Olá, pedro"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
