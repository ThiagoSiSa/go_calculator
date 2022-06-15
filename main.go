package main

import (
	"fmt"
	"go_calculator/aroundnumbers"
	"go_calculator/calc"
	"go_calculator/research"
	"go_calculator/splice"
	"go_calculator/tokanizer"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

	userInput := Input()
	token := tokanizer.StringToArray(userInput)
	result := Calculator(token)
	fmt.Printf("O resultado é, %.f \n", result)

}

func Input() string {
	var userInput string
	fmt.Scan(&userInput)
	//userInput = "4/4"
	return userInput
}

func Calculator(token []string) float64 {
	var result float64
	for {
		hasDiv, indexDiv := research.Includes(token, "/")
		hasMult, indexMult := research.Includes(token, "*")
		hasSum, indexSum := research.Includes(token, "+")
		hasSub, indexSub := research.Includes(token, "-")

		if hasDiv {
			numberBefore, numberAfter := aroundnumbers.Numbers(token, indexDiv)
			result = float64(numberBefore) / float64(numberAfter)
			resultConv := strconv.FormatFloat(result, 'g', 4, 64)
			token = splice.RefactoryArray(token, indexDiv-1, indexDiv+2, resultConv)

		} else if hasMult {
			numberBefore, numberAfter := aroundnumbers.Numbers(token, indexMult)
			result = float64(numberBefore) * float64(numberAfter)
			resultConv := strconv.FormatFloat(result, 'g', 4, 64)
			token = splice.RefactoryArray(token, indexMult-1, indexMult+2, resultConv)

		} else if hasSum {
			numberBefore, numberAfter := aroundnumbers.Numbers(token, indexSum)
			result = float64(numberBefore) + float64(numberAfter)
			resultConv := strconv.FormatFloat(result, 'g', 4, 64)
			token = splice.RefactoryArray(token, indexSum-1, indexSum+2, resultConv)

		} else if hasSub {
			numberBefore, numberAfter := aroundnumbers.Numbers(token, indexSub)
			result = float64(numberBefore) - float64(numberAfter)
			resultConv := strconv.FormatFloat(result, 'g', 4, 64)
			token = splice.RefactoryArray(token, indexSub-1, indexSub+2, resultConv)
		} else {
			break
		}
	}
	sum := calc.Add(token)
	return sum
}

func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index", nil)

}
