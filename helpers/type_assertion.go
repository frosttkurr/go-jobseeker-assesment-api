package helpers

import "strconv"

func StringToNumber(str string) (number int) {
	var err error

	if str == "" {
		number = 0
		return number
	}

	number, err = strconv.Atoi(str)
	if err != nil {
		number = 0
		return number
	}
	return number
}
