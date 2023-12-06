package line_utils

import (
	"strconv"
	"strings"
	"unicode"
)

type numbersTable map[string]int

type spelledIndices map[int]int

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
	spelledIdx := FindSpelledNumbers(line)

	for i, char := range line {
		if val, ok := spelledIdx[i]; ok {
			digits += strconv.Itoa(val)
		}

		if unicode.IsDigit(char) {
			digits += string(char)
		}
	}

	concatRunes := string(digits[0]) + string(digits[len(digits)-1])
	sumDigits, _ := strconv.Atoi(string(concatRunes))

	return sumDigits
}

func FindSpelledNumbers(line string) spelledIndices {
	spelled := spelledIndices{}
	for k, v := range spelledNumbers {
		index := strings.Index(line, k)
		if index == -1 {
			continue
		}

		spelled[index] = v
	}

	return spelled
}
