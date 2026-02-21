package interpolate

import "testing"

func TestVariables(t *testing.T) {
	type args struct {
		dict           map[string][]byte
		originalString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no interpolation",
			args: args{
				dict: map[string][]byte{
					"a": []byte{'a', 'b', 'c'},
				},
				originalString: "original string",
			},
			want: "original string",
		},
		{
			name: "single interpolation",
			args: args{
				dict: map[string][]byte{
					"abc": []byte("abc"),
					"def": []byte("def"),
				},
				originalString: "$abc",
			},
			want: "abc",
		},
		{
			name: "multiple, immediately adjacent interpolation",
			args: args{
				dict: map[string][]byte{
					"abc": []byte("abc"),
					"def": []byte("def"),
					"ghi": []byte("ghi"),
				},
				originalString: "$abc$def$ghi",
			},
			want: "abcdefghi",
		},
		{
			name: "interpolation before non-identifier character",
			args: args{
				dict: map[string][]byte{
					"abc": []byte("abc"),
					"def": []byte("def"),
					"ghi": []byte("ghi"),
				},
				originalString: "$abc;",
			},
			want: "abc;",
		},
		{
			name: "interpolation after non-identifier character",
			args: args{
				dict: map[string][]byte{
					"abc": []byte("abc"),
					"def": []byte("def"),
					"ghi": []byte("ghi"),
				},
				originalString: ";$ghi",
			},
			want: ";ghi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Variables(tt.args.originalString, tt.args.dict); got != tt.want {
				t.Errorf("Variables() = %v, want %v", got, tt.want)
			}
		})
	}
}
