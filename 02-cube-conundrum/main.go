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
	count := 0
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
		g := game.GameParser(line)
		sumPossible += game.FindMinPossibleToValid(g.Cubes)
		fmt.Println(g)
		fmt.Println("----------------------------------------------------------------------------------------------------")

		if g.Valid {
			sumId += g.ID
		}

		count += 1
		// if count == 100 {
		// 	break
		// }

	}

	fmt.Println("part 1 - ", sumId)
	fmt.Println("part 2 - ", sumPossible)
}
