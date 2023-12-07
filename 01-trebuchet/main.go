package main

import (
	"aoc2023/day01/line_utils"
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += line_utils.SumConcatDigits(line)
	}

	fmt.Println(sum)
}
