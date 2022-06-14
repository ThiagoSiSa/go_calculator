package research

func Includes(array []string, symbol string) (bool, int) {
	for i, str := range array {

		if str == symbol {
			return true, i
		}
	}
	return false, 0
}
