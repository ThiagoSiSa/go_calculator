package main

import (
	"fmt"
	re "regexp"
	"strconv"
)

func main() {
	userInput := Input()
	token := Tokanizer(userInput)
	result := Calculator(token)
	fmt.Printf("O resultado é, %.f \n", result)
}

func Input() string {
	var userInput string
	fmt.Scan(&userInput)
	return userInput
}
func Tokanizer(token string) []string {

	regex := re.MustCompile(`[\d]{1,}|[^\d]`)
	numbers := re.MustCompile(`\d`)
	arrayCharacters := regex.FindAllString(token, -1)
	var arrayNumbersAndSignals []string
	for i := 0; i < len(arrayCharacters); i++ {
		str := arrayCharacters[i]
		if str == "+" || str == "-" && numbers.Match([]byte(arrayCharacters[i+1])) {
			strConcat := str + arrayCharacters[i+1]
			arrayNumbersAndSignals = append(arrayNumbersAndSignals, strConcat)
			arrayCharacters = arrayCharacters[i+2:]
			i = -1
		} else {
			arrayNumbersAndSignals = append(arrayNumbersAndSignals, str)
		}
	}

	return arrayNumbersAndSignals
}
func Calculator(token []string) float64 {
	var result float64
	for {
		hasDiv, indexDiv := Includes(token, "/")
		if hasDiv {
			numberBefore, _ := strconv.ParseFloat(token[indexDiv-1], 64)
			numberAfter, _ := strconv.ParseFloat(token[indexDiv+1], 64)
			result = float64(numberBefore) / float64(numberAfter)
			resultConv := strconv.FormatFloat(result, 'g', 4, 64)
			token = Splice(token, indexDiv-1, indexDiv+2, resultConv)

		}
		hasMult, indexMult := Includes(token, "*")
		if hasMult {
			numberBefore, _ := strconv.ParseFloat(token[indexMult-1], 64)
			numberAfter, _ := strconv.ParseFloat(token[indexMult+1], 64)
			result = float64(numberBefore) * float64(numberAfter)
			resultConv := strconv.FormatFloat(result, 'g', 4, 64)
			token = Splice(token, indexMult-1, indexMult+2, resultConv)

		}

		return Sum(token)
	}

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
