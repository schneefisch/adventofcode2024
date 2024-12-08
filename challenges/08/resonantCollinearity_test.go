package _8

import "testing"

func TestResonantCollinearity(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		wantErr bool
	}{
		{
			name:    "Test case 1",
			args:    args{filename: "test-input.txt"},
			want:    14,
			want1:   34,
			wantErr: false,
		},
		{
			name:    "Input",
			args:    args{filename: "input.txt"},
			want:    256,
			want1:   1005,
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, got1, err := ResonantCollinearity(test.args.filename)
			if (err != nil) != test.wantErr {
				t.Errorf("ResonantCollinearity() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("ResonantCollinearity() got = %v, want %v", got, test.want)
			}
			if got1 != test.want1 {
				t.Errorf("ResonantCollinearity() got1 = %v, want %v", got1, test.want1)
			}
		})
	}
}
