package splice

func RefactoryArray(array []string, startIndex int, finalIndex int, substitute string) []string {

	var newToken []string
	firstPartArray := array[:startIndex]
	lastPartArray := array[finalIndex:]
	newToken = append(newToken, firstPartArray...)
	newToken = append(newToken, substitute)
	newToken = append(newToken, lastPartArray...)

	return newToken
}
