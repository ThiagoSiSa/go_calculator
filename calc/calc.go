package calc

import (
	"strconv"
)

func Add(array []string) float64 {

	var result float64

	for _, str := range array {

		operating, _ := strconv.ParseFloat(str, 64)
		result = operating + result
	}
	return result
}
