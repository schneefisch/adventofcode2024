package _6

import "testing"

func TestGuardGallivant(t *testing.T) {
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
			name:    "Test 1",
			args:    args{filename: "test_input.txt"},
			want:    41,
			wantErr: false,
		},
		{
			name:    "Input",
			args:    args{filename: "input.txt"},
			want:    5030,
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GuardGallivant(test.args.filename)
			if (err != nil) != test.wantErr {
				t.Errorf("GuardGallivant() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("GuardGallivant() got = %v, want %v", got, test.want)
			}
		})
	}
}
