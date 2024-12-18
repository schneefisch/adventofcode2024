package _13

import "testing"

func TestClawContraption(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{"test.txt"},
			want: 480,
		},
		{
			name: "input",
			args: args{"input.txt"},
			want: 38839,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ClawContraption(test.args.filename)
			if got != test.want {
				t.Errorf("ClawContraption() = %v, want %v", got, test.want)
			}
		})
	}
}
