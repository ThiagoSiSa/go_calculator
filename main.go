package main

import (
	"fmt"
	"go_calculator/tokanizer/tokanizer"
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
	return userInput
}

func Calculator(token []string) float64 {
	var result float64
	for {
		hasDiv, indexDiv := Includes(token, "/")
		hasMult, indexMult := Includes(token, "*")
		if hasDiv {
			numberBefore, err := strconv.ParseFloat(token[indexDiv-1], 64)
			if err != nil {
				fmt.Println("Erro", err)
				break
			}
			numberAfter, _ := strconv.ParseFloat(token[indexDiv+1], 64)

			result = float64(numberBefore) / float64(numberAfter)
			resultConv := strconv.FormatFloat(result, 'g', 4, 64)
			token = Splice(token, indexDiv-1, indexDiv+2, resultConv)

		} else if hasMult {
			numberBefore, _ := strconv.ParseFloat(token[indexMult-1], 64)
			numberAfter, _ := strconv.ParseFloat(token[indexMult+1], 64)
			result = float64(numberBefore) * float64(numberAfter)
			resultConv := strconv.FormatFloat(result, 'g', 4, 64)
			token = Splice(token, indexMult-1, indexMult+2, resultConv)

		} else {
			break
		}
	}
	return Sum(token)
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
