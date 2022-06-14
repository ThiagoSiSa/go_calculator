package aroundnumbers

import (
	"errors"
	"log"
	"strconv"
)

func Numbers(token []string, index int) (numBefore float64, numAfter float64) {
	stringBefore := token[index-1]
	stringAfter := token[index+1]
	msgErr := "Entrada inválida"

	if stringBefore == "/" || stringAfter == "/" {
		err := errors.New(msgErr)
		log.Fatal(err.Error())
	} else if stringBefore == "*" || stringAfter == "*" {
		err := errors.New(msgErr)
		log.Fatal(err.Error())
	}

	numberBefore, err := strconv.ParseFloat(token[index-1], 64)
	if err != nil {
		panic(err.Error())
	}
	numberAfter, err := strconv.ParseFloat(token[index+1], 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	return numberBefore, numberAfter
}
