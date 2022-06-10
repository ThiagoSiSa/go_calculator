package tokanizer

import (
	re "regexp"
)

func StringToArray(token string) []string {

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
