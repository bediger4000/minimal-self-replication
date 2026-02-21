package lexing

import "testing"

func Test_stringEval(t *testing.T) {
	tests := []struct {
		name           string
		original       string
		variableValues map[string][]byte
		want           string
	}{
		{
			name:     "single-quoted string, no variables",
			original: "'abc def ghi'",
			want:     "abc def ghi",
			variableValues: map[string][]byte{
				"abc": []byte{'W', 'R', 'O', 'N', 'G'},
			},
		},
		{
			name:     "single-quoted string, one variable",
			original: "'abc $def ghi'",
			want:     "abc $def ghi",
			variableValues: map[string][]byte{
				"abc": []byte{'W', 'R', 'O', 'N', 'G'},
				"def": []byte{'A', 'L', 'S', 'O', ' ', 'W', 'R', 'O', 'N', 'G'},
			},
		},
		{
			name:     "double-quoted string, one variable",
			original: `"abc $def ghi"`,
			want:     "abc CORRECT ghi",
			variableValues: map[string][]byte{
				"abc": []byte{'W', 'R', 'O', 'N', 'G'},
				"def": []byte{'C', 'O', 'R', 'R', 'E', 'C', 'T'},
			},
		},
		{
			name:     "unquoted string, multiple variable",
			original: `"$a$b$c"`,
			want:     "abc",
			variableValues: map[string][]byte{
				"a": []byte{'a'},
				"b": []byte{'b'},
				"c": []byte{'c'},
			},
		},
		{
			name:     "single-quoted double-quote",
			original: `'"'`,
			want:     `"`,
			variableValues: map[string][]byte{
				"abc": []byte{'W', 'R', 'O', 'N', 'G'},
				"a":   []byte{'a'},
				"b":   []byte{'b'},
				"c":   []byte{'c'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringEval(tt.original, tt.variableValues); got != tt.want {
				t.Errorf("stringEval() = %v, want %v", got, tt.want)
			}
		})
	}
}
