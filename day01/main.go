package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`\d+`)
		b := re.FindAllString(line, -1)

		fmt.Println(b)

		first, _ := strconv.Atoi(b[0])
		last, _ := strconv.Atoi(b[len(b)-1])

		sum += first + last
		fmt.Println(sum)
		count += 1

		if count == 3 {
			break
		}

	}
}
