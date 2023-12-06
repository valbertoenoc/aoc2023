package line_utils

import (
	"reflect"
	"testing"
)

func Test_SumConcatDigits(t *testing.T) {
	test_cases := []struct {
		name string
		line string
		want int
	}{
		{
			name: "one number",
			line: "five8b",
			want: 58,
		},
		{
			name: "two or more numbers",
			line: "2733vmmpknvgr",
			want: 23,
		},
		{
			name: "two or more numbers",
			line: "3oneeighttwo",
			want: 32,
		},
		{
			name: "two or more numbers",
			line: "2fourghsixptk",
			want: 26,
		},
		{
			name: "elision",
			line: "jshmdl7oneeightwocj",
			want: 72,
		},
		{
			name: "elision",
			line: "nine9twoeightonebz",
			want: 91,
		},
	}

	for _, tt := range test_cases {
		t.Run(tt.name, func(t *testing.T) {
			got := SumConcatDigits(tt.line)

			if got != tt.want {
				t.Errorf("expected %v got %v", tt.want, got)
			}
		})
	}
}

func Test_FindSpelledNumbers(t *testing.T) {
	test_cases := []struct {
		name string
		line string
		want spelledIndices
	}{
		{
			name: "one spelled number",
			line: "five8b",
			want: spelledIndices{0: 5},
		},
		{
			name: "two or more spelled numbers",
			line: "3oneeighttwo",
			want: spelledIndices{1: 1, 4: 8, 9: 2},
		},
		{
			name: "two or more spelled numbers spread",
			line: "2fourghsixptk",
			want: spelledIndices{1: 4, 7: 6},
		},
		{
			name: "elision",
			line: "jshmdl7oneeightwocj",
			want: spelledIndices{7: 1, 10: 8, 14: 2},
		},
	}

	for _, tt := range test_cases {
		t.Run(tt.name, func(t *testing.T) {
			got := FindSpelledNumbers(tt.line)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected %v got %v", tt.want, got)
			}
		})
	}
}
