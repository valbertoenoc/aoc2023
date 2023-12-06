package line_utils

import "testing"

func Test_InsertIntoString(t *testing.T) {
	test_cases := []struct {
		name  string
		old   string
		input string
		index int
		want  string
	}{
		{
			name:  "insert in the beginning",
			old:   "whatever",
			input: "test-",
			index: 0,
			want:  "test-whatever",
		},
		{
			name:  "insert in the end",
			old:   "whatever",
			input: "-test",
			index: 8,
			want:  "whatever-test",
		},
		{
			name:  "insert in the middle",
			old:   "whatever",
			input: "-test-",
			index: 4,
			want:  "what-test-ever",
		},
	}

	for _, tt := range test_cases {
		t.Run(tt.name, func(t *testing.T) {
			got := InsertIntoString(tt.index, tt.old, tt.input)

			if tt.want != got {
				t.Errorf("expected %v got %v", tt.want, got)
			}
		})
	}
}
