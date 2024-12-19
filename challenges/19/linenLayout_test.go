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
		want2   int
		wantErr bool
	}{
		{
			name:    "Test case 1",
			args:    args{"test.txt"},
			want:    6,
			want2:   16,
			wantErr: false,
		},
		{
			name:    "input",
			args:    args{"input.txt"},
			want:    358,
			want2:   600639829400603,
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, got2, err := LinenLayout(test.args.filename)
			if (err != nil) != test.wantErr {
				t.Errorf("LinenLayout() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("LinenLayout() got = %v, want %v", got, test.want)
			}
			if got2 != test.want2 {
				t.Errorf("LinenLayout() got2 = %v, want %v", got2, test.want2)
			}
		})
	}
}
