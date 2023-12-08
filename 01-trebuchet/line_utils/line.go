package line_utils

import (
	"strconv"
	"unicode"
)

type (
	numbersTable map[string]int
)

var spelledNumbers numbersTable = numbersTable{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func SumConcatDigits(line string) int {
	var digits string

	for i, char := range line {

		if len(line[i:]) >= 3 {
			if val, ok := spelledNumbers[line[i:i+3]]; ok {
				digits += strconv.Itoa(val)
			}
		}

		if len(line[i:]) >= 4 {
			if val, ok := spelledNumbers[line[i:i+4]]; ok {
				digits += strconv.Itoa(val)
			}
		}

		if len(line[i:]) >= 5 {
			if val, ok := spelledNumbers[line[i:i+5]]; ok {
				digits += strconv.Itoa(val)
			}
		}

		if unicode.IsDigit(char) {
			digits += string(char)
		}
	}

	concatRunes := string(digits[0]) + string(digits[len(digits)-1])
	sumDigits, _ := strconv.Atoi(string(concatRunes))

	return sumDigits
}
