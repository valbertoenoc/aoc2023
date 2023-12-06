package line_utils

import (
	"testing"
)

func Test_SumFirstAndLast(t *testing.T) {
	tests := []struct {
		name string
		line string
		want int
	}{
		{
			name: "one number",
			line: "five8b",
			want: 88,
		},
		{
			name: "two or more numbers",
			line: "2733vmmpknvgr",
			want: 23,
		},
		{
			name: "two or more numbers",
			line: "3oneeighttwo",
			want: 33,
		},
		{
			name: "two or more numbers",
			line: "2fourghsixptk",
			want: 22,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConcatDigits(tt.line)

			if got != tt.want {
				t.Errorf("expected %v got %v", tt.want, got)
			}
		})
	}
}
