package iteracao

func Repetir(caract string, vezes int) string {
	var repeticoes string

	for i := 0; i < vezes; i++ {
		repeticoes += caract
	}

	return repeticoes
}
