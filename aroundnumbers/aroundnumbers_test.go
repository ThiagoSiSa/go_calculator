package aroundnumbers

import "testing"

func TestNumbers(t *testing.T) {
	test := []string{"1", "+", "2"}
	numBefore, numAfter := Numbers(test, 1)

	if numBefore != 1 && numAfter != 2 {
		t.Fatalf("Valor esperado: Numero antes: 1 Numero depois : 2 - Valor retornado numero antes %f numero depois %f", numBefore, numAfter)

	}
}
