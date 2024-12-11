package _11

import "testing"

func TestPlutionianPebbles(t *testing.T) {
	type args struct {
		filename   string
		iterations int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		wantErr bool
	}{
		{
			name: "Test case 0",
			args: args{
				filename:   "test0.txt",
				iterations: 25,
			},
			want:    55312,
			wantErr: false,
		},
		{
			name: "input",
			args: args{
				filename:   "input.txt",
				iterations: 25,
			},
			want:    186996,
			wantErr: false,
		},
		{
			name: "input with 75 iterations",
			args: args{
				filename:   "input.txt",
				iterations: 75,
			},
			want:    221683913164898,
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := PlutionianPebbles(test.args.filename, test.args.iterations)
			if (err != nil) != test.wantErr {
				t.Errorf("PlutionianPebbles() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("PlutionianPebbles() iterations: %d got = %v, want %v", test.args.iterations, got, test.want)
			}
		})
	}
}
