package _18

import "testing"

func TestRamRun(t *testing.T) {
	type args struct {
		filename string
		width    int
		height   int
		bytes    int
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantBlock Position
		wantErr   bool
	}{
		{
			name: "testcase 1",
			args: args{
				"test.txt",
				7,
				7,
				12,
			},
			want:      22,
			wantBlock: Position{6, 1},
			wantErr:   false,
		},
		{
			name: "input",
			args: args{
				"input.txt",
				71,
				71,
				1024,
			},
			want:      374,
			wantBlock: Position{30, 12},
			wantErr:   false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, blockPos, err := RamRun(test.args.filename, test.args.width, test.args.height, test.args.bytes)
			if (err != nil) != test.wantErr {
				t.Errorf("RamRun() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("RamRun() got = %v, want %v", got, test.want)
			}
			if !blockPos.Equals(test.wantBlock) {
				t.Errorf("RamRun() blockPos = %v, want %v", blockPos, test.wantBlock)
			}
		})
	}
}
