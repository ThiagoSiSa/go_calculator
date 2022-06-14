package calc

import "testing"

func TestAdd(t *testing.T) {
	test := []string{"1", "+", "3"}
	result := Add(test)
	comp := 4

	if result != float64(comp) {
		t.Fatalf("Valor esperado é 4")
	}
}
