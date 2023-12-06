package line_utils

import (
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
			line: "2fourghsixptk",
			want: 26,
		},
		{
			name: "elision",
			line: "jshmdl7oneeightwocj",
			want: 72,
		},
		{
			name: "spelled in the beginning",
			line: "nine9twoeightonebz",
			want: 91,
		},
		{
			name: "case",
			line: "onejzk2t4nineqstdnhlqvmrrc",
			want: 19,
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
