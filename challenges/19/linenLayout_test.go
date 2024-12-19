package _19

import "testing"

func TestLinenLayout(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Test case 1",
			args:    args{"test.txt"},
			want:    6,
			wantErr: false,
		},
		{
			name:    "input",
			args:    args{"input.txt"},
			want:    358,
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := LinenLayout(test.args.filename)
			if (err != nil) != test.wantErr {
				t.Errorf("LinenLayout() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("LinenLayout() got = %v, want %v", got, test.want)
			}
		})
	}
}
