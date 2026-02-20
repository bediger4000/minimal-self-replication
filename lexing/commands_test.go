package lexing

import (
	"reflect"
	"testing"
)

func Test_Lexing(t *testing.T) {
	tests := []struct {
		name   string
		buffer string
		want   []string
	}{
		{
			name:   "unquoted assignment, no newline",
			buffer: `abc=defghijklmnopq`,
			want:   []string{`abc=defghijklmnopq`},
		},
		{
			name:   "unquoted assignment, with newline",
			buffer: "abc=defghijklmnopq\n",
			want:   []string{`abc=defghijklmnopq`},
		},
		{
			name:   "single-quoted assignment, with newline",
			buffer: "abc='defghijklmnopq'\n",
			want:   []string{`abc='defghijklmnopq'`},
		},
		{
			name:   "double-quoted assignment, without newline",
			buffer: `abc="defghijklmnopq"`,
			want:   []string{`abc="defghijklmnopq"`},
		},
		{
			name:   "multiple commands semicolon separated, no newline at end",
			buffer: `a=b;c=d;e="now is the time"`,
			want: []string{
				`a=b`,
				`c=d`,
				`e="now is the time"`,
			},
		},
		{
			name:   "adjacent single and double quoted string literals",
			buffer: `abc='def'"ghijklmnopq"`,
			want:   []string{`abc='def'"ghijklmnopq"`},
		},
		{
			name:   "variable interpolation in double quoted string literal",
			buffer: `abc="$defghijklmnopq"`,
			want:   []string{`abc="$defghijklmnopq"`},
		},
		{
			name:   "multiple adjacent variable interpolation",
			buffer: `abc=$defghij$klmnopq`,
			want:   []string{`abc=$defghij$klmnopq`},
		},
		{
			name:   "single-quoted double-quote",
			buffer: `sq='"'`,
			want:   []string{`sq='"'`},
		},
		{
			name:   "double-quoted single-quote",
			buffer: `dq="'"`,
			want:   []string{`dq="'"`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			lxr := NewCommandLexer([]rune(tt.buffer))
			var got []string
			for str := lxr(); str != ""; str = lxr() {
				got = append(got, str)
			}

			if len(got) != len(tt.want) {
				t.Errorf("lexing, got %d commands, want %d commands", len(got), len(tt.want))
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexing, got %v, want %v", got, tt.want)
			}

		})
	}
}
