package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(13, 17)

	if total != 30 {
		t.Errorf("Resultado da soma inválido. Resultado %d. Esperado: %d", total, 30)
	}
}
