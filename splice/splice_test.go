package splice

import (
	"reflect"
	"testing"
)

func TestRefactoryArray(t *testing.T) {
	test := []string{"1", "+", "6", "/", "2", "+", "2"}
	firstPartArray := 2
	lastPartArray := 5
	substitute := "3"
	result := RefactoryArray(test, firstPartArray, lastPartArray, substitute)
	comp := []string{"1", "+", "3", "+", "1"}

	if !reflect.DeepEqual(comp, result) {
		t.Fatal("Valor esperado é", comp)
	}

}
