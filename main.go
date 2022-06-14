package main

import (
	"errors"
	"fmt"
	"go_calculator/tokanizer/tokanizer"
	"log"
	"strconv"
)

func main() {

	userInput := Input()
	token := tokanizer.StringToArray(userInput)
	result := Calculator(token)
	fmt.Printf("O resultado é, %.f \n", result)

}

func Input() string {
	var userInput string
	fmt.Scan(&userInput)
	//userInput = "4//4"
	return userInput
}

func Calculator(token []string) float64 {
	var result float64
	for {
		hasDiv, indexDiv := Includes(token, "/")
		hasMult, indexMult := Includes(token, "*")
		hasSum, indexSum := Includes(token, "+")
		hasSub, indexSub := Includes(token, "-")

		if hasDiv {
			numberBefore, numberAfter := AroundNumbers(token, indexDiv)
			result = float64(numberBefore) / float64(numberAfter)
			resultConv := strconv.FormatFloat(result, 'g', 4, 64)
			token = Splice(token, indexDiv-1, indexDiv+2, resultConv)

		} else if hasMult {
			numberBefore, numberAfter := AroundNumbers(token, indexMult)
			result = float64(numberBefore) * float64(numberAfter)
			resultConv := strconv.FormatFloat(result, 'g', 4, 64)
			token = Splice(token, indexMult-1, indexMult+2, resultConv)

		} else if hasSum {
			numberBefore, numberAfter := AroundNumbers(token, indexSum)
			result = float64(numberBefore) + float64(numberAfter)
			resultConv := strconv.FormatFloat(result, 'g', 4, 64)
			token = Splice(token, indexSum-1, indexSum+2, resultConv)

		} else if hasSub {
			numberBefore, numberAfter := AroundNumbers(token, indexSub)
			result = float64(numberBefore) - float64(numberAfter)
			resultConv := strconv.FormatFloat(result, 'g', 4, 64)
			token = Splice(token, indexSub-1, indexSub+2, resultConv)
		} else {
			break
		}
	}
	sum := Sum(token)
	return sum
}

func Includes(array []string, symbol string) (bool, int) {
	for i, str := range array {

		if str == symbol {
			return true, i
		}
	}
	return false, 0
}

func Splice(array []string, startIndex int, finalIndex int, substitute string) []string {

	var newToken []string
	firstPartArray := array[:startIndex]
	lastPartArray := array[finalIndex:]
	newToken = append(newToken, firstPartArray...)
	newToken = append(newToken, substitute)
	newToken = append(newToken, lastPartArray...)

	return newToken
}

func Sum(array []string) float64 {

	var result float64

	for _, str := range array {

		operating, _ := strconv.ParseFloat(str, 64)
		result = operating + result
	}
	return result
}

func AroundNumbers(token []string, index int) (numBefore float64, numAfter float64) {
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
