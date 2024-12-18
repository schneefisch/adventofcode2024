package _13

import "testing"

func TestClawContraption(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want2 int
	}{
		{
			name: "test",
			args: args{"test.txt"},
			want: 480,
		},
		{
			name:  "input",
			args:  args{"input.txt"},
			want:  38839,
			want2: 75200131617108,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, got2 := ClawContraption(test.args.filename)
			if got != test.want {
				t.Errorf("ClawContraption() = %v, want %v", got, test.want)
			}
			if test.want2 != 0 && got2 != test.want2 {
				t.Errorf("ClawContraption() = %v, want %v", got2, test.want2)
			}
		})
	}
}
