package game

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	GameSet []CubeSet
)

type CubeSet struct {
	R int
	G int
	B int
}

type Game struct {
	ID    int
	Cubes GameSet
	Valid bool
}

func GameParser(input string) Game {
	inputSplit := strings.Split(input, ":")

	input1 := inputSplit[0]
	id := parseId(input1)

	input2 := inputSplit[1]
	gameSet := parseGameSet(input2)

	return Game{ID: id, Cubes: gameSet, Valid: IsGameValid(gameSet)}
}

func IsGameValid(gameSet GameSet) bool {
	for _, set := range gameSet {
		if set.R > 12 {
			return false
		}

		if set.G > 13 {
			return false
		}

		if set.B > 14 {
			return false
		}

	}

	return true
}

func FindMinPossibleToValid(gameSet GameSet) int {
	maxRed := 0
	maxGreen := 0
	maxBlue := 0
	for _, set := range gameSet {
		if set.R > maxRed {
			maxRed = set.R
		}

		if set.G > maxGreen {
			maxGreen = set.G
		}

		if set.B > maxBlue {
			maxBlue = set.B
		}
	}
	fmt.Printf("max R: %v - max G: %v - max B: %v", maxRed, maxGreen, maxBlue)

	return maxRed * maxGreen * maxBlue
}

func parseId(gameInput string) int {
	id, _ := strconv.Atoi(strings.Split(gameInput, " ")[1])
	return id
}

func parseGameSet(gameSetInput string) GameSet {
	sets := strings.Split(gameSetInput, ";")

	gameSet := parseSets(sets)

	return gameSet
}

func parseSets(cubesString []string) []CubeSet {
	var cubeSet []CubeSet
	for _, set := range cubesString {
		cSet := CubeSet{}
		for _, s := range strings.Split(set, ",") {
			s := strings.Trim(s, " ")
			cube := strings.Split(s, " ")

			if "red" == cube[1] {
				color, _ := strconv.Atoi(cube[0])
				cSet.R = color
			}

			if "green" == cube[1] {
				color, _ := strconv.Atoi(cube[0])
				cSet.G = color
			}

			if "blue" == cube[1] {
				color, _ := strconv.Atoi(cube[0])
				cSet.B = color
			}

		}

		cubeSet = append(cubeSet, cSet)
	}

	return cubeSet
}
