package tokanizer

import (
	"reflect"
	"testing"
)

func TestStringToArray(t *testing.T) {
	test := "1+2"
	result := StringToArray(test)
	comp := []string{"1", "+2"}

	if !reflect.DeepEqual(comp, result) {
		t.Fatal("Valor esperado: [1] - Valor retornado:", result)
	}
}
