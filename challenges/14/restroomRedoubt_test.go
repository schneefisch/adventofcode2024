package _14

import "testing"

func Test_restroomRedoubt(t *testing.T) {
	type args struct {
		filename   string
		width      int
		height     int
		iterations int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "Example",
			args: args{
				filename:   "test.txt",
				width:      11,
				height:     7,
				iterations: 100,
			},
			want:  12,
			want1: 0,
		},
		{
			name: "input",
			args: args{
				filename:   "input.txt",
				width:      101,
				height:     103,
				iterations: 100,
			},
			want:  230172768,
			want1: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, got1 := restroomRedoubt(test.args.filename, test.args.width, test.args.height, test.args.iterations)
			if got != test.want {
				t.Errorf("restroomRedoubt() got = %v, want %v", got, test.want)
			} else {
				t.Logf("Got: %v", got)
			}
			if got1 != test.want1 {
				t.Errorf("restroomRedoubt() got1 = %v, want %v", got1, test.want1)
			} else {
				t.Logf("Got1: %v", got1)
			}
		})
	}
}
