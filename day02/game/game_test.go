package game

import (
	"reflect"
	"testing"
)

func Test_GameParser(t *testing.T) {
	test_cases := []struct {
		name  string
		input string
		want  Game
	}{
		{
			name:  "normal game input",
			input: "Game 1: 2 blue, 4 green; 7 blue, 1 red, 14 green; 5 blue, 13 green, 1 red; 1 red, 7 blue, 11 green",
			want: Game{
				ID: 1, Cubes: GameSet{
					CubeSet{R: 0, G: 4, B: 2},
					CubeSet{R: 1, G: 14, B: 7},
					CubeSet{R: 1, G: 13, B: 5},
					CubeSet{R: 1, G: 11, B: 7},
				}, Valid: false,
			},
		},
		{
			name:  "another game input",
			input: "Game 65: 7 green, 1 blue; 1 red, 14 blue, 4 green; 8 blue, 6 red; 14 green, 4 red ",
			want: Game{
				ID: 65, Cubes: GameSet{
					CubeSet{R: 0, G: 7, B: 1},
					CubeSet{R: 1, G: 4, B: 14},
					CubeSet{R: 6, G: 0, B: 8},
					CubeSet{R: 4, G: 14, B: 0},
				}, Valid: false,
			},
		},
		{
			name:  "another game input",
			input: "Game 31: 10 green, 9 blue; 5 green, 9 blue, 1 red; 1 red, 8 blue ",
			want: Game{
				ID: 31, Cubes: GameSet{
					CubeSet{R: 0, G: 10, B: 9},
					CubeSet{R: 1, G: 5, B: 9},
					CubeSet{R: 1, G: 0, B: 8},
				}, Valid: false,
			},
		},
	}

	for _, tt := range test_cases {
		t.Run(tt.name, func(t *testing.T) {
			got := GameParser(tt.input)

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("expected %v got %v", tt.want, got)
			}
		})
	}
}

func Test_IsGameValid(t *testing.T) {
	test_cases := []struct {
		name  string
		input GameSet
		want  bool
	}{
		{
			name: "normal gameset",
			input: GameSet{
				CubeSet{R: 0, G: 2, B: 2},
				CubeSet{R: 1, G: 3, B: 1},
				CubeSet{R: 1, G: 1, B: 5},
				CubeSet{R: 1, G: 1, B: 1},
			},
			want: true,
		},
		{
			name: "reds greater 13",
			input: GameSet{
				CubeSet{R: 5, G: 1, B: 1},
				CubeSet{R: 5, G: 2, B: 2},
				CubeSet{R: 5, G: 3, B: 3},
				CubeSet{R: 5, G: 4, B: 4},
			},
			want: false,
		},
		{
			name: "blues greater 15",
			input: GameSet{
				CubeSet{R: 0, G: 1, B: 10},
				CubeSet{R: 0, G: 2, B: 2},
				CubeSet{R: 0, G: 3, B: 3},
				CubeSet{R: 0, G: 4, B: 4},
			},
			want: false,
		},
	}

	for _, tt := range test_cases {
		t.Run(tt.name, func(t *testing.T) {
			got := IsGameValid(tt.input)

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("expected %v got %v", tt.want, got)
			}
		})
	}
}
