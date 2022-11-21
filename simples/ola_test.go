package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola()
	esperado := "Hello World"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
