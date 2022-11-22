package iteracao

import (
	"fmt"
	"testing"
)

func TestRep(t *testing.T) {

	repeticoes := Repetir("a", 6)
	esperado := "aaaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s',  resultado '%s'", repeticoes, esperado)
	}
}

func BenchmarkRepetir(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Repetir("a", 5)
	}
}

func ExampleRepetir() {
	valorResult := Repetir("a", 5)
	fmt.Println(valorResult)
	// Output: aaaaa
}
