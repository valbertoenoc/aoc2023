package main

import (
	"fmt"
	"os"
	"strings"
)

type Coords struct {
	y int
	x int
}

type Elem struct {
	Value    string
	IsDigit  bool
	IsSymbol bool
	Coords   Coords
}

type Matrix [][]Elem

func check(err error) {
	if err != nil {
		panic(1)
	}
}

func main() {
	b, err := os.ReadFile("input.txt")
	check(err)
	fmt.Println(string(b))
	lines := strings.Split(string(b), "\n")
	fmt.Println(lines)

	m := Matrix{}
	for i, l := range lines {
		l = strings.Trim(l, " ")
		lineSlice := make([]Elem, 0)
		for j, c := range l {
			elem := Elem{
				Value:    string(c),
				IsDigit:  true,
				IsSymbol: false,
				Coords:   Coords{y: i, x: j},
			}

			lineSlice = append(lineSlice, elem)
		}
		m = append(m, lineSlice)

	}
	fmt.Println(m[2][3])
}
