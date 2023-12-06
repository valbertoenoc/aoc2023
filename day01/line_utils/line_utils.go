package line_utils

import (
	"strconv"
	"unicode"
)

func ConcatDigits(line string) int {
	var digits string
	for _, char := range line {

		if !unicode.IsDigit(char) {
			continue
		}
		digits += string(char)
	}

	concatRunes := string(digits[0]) + string(digits[len(digits)-1])
	sumDigits, _ := strconv.Atoi(string(concatRunes))

	return sumDigits
}
