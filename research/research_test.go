package research

import "testing"

func TestIncludes(t *testing.T) {
	token := []string{"1", "+", "2", "/", "3"}
	symbol := "/"
	compBool := true
	compInt := 3
	hasSymbol, indexSymbol := Includes(token, symbol)
	if compBool != true || compInt != 3 {
		t.Fatalf("Valor esperado compBol: true compInt: 3 - Valor retornado hasSymbol:%t indexSymbol: %d", hasSymbol, indexSymbol)

	}

}
