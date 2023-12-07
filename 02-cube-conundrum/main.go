package main

import (
	"aoc2023/day02/game"
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(1)
	}
}

func main() {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	sumId := 0
	sumPossible := 0
	for scanner.Scan() {
		line := scanner.Text()

		g := game.GameParser(line)
		sumPossible += game.FindMinPossibleToValid(g.Cubes)

		if g.Valid {
			sumId += g.ID
		}

	}

	fmt.Println("part 1 - ", sumId)
	fmt.Println("part 2 - ", sumPossible)
}
