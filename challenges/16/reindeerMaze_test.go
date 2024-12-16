package _16

import "testing"

func TestReindeerMaze(t *testing.T) {
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
			args:    args{"test.txt"},
			want:    7036,
			want1:   0,
			wantErr: false,
		},
		{
			name:    "Test case 2",
			args:    args{"test2.txt"},
			want:    11048,
			want1:   0,
			wantErr: false,
		},
		{
			name:    "Input",
			args:    args{"input.txt"},
			want:    0,
			want1:   0,
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, got1, err := ReindeerMaze(test.args.filename)
			if (err != nil) != test.wantErr {
				t.Errorf("ReindeerMaze() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("ReindeerMaze() got = %v, want %v", got, test.want)
			}
			if got1 != test.want1 {
				t.Errorf("ReindeerMaze() got1 = %v, want %v", got1, test.want1)
			}
		})
	}
}

func Test_rotationScore(t *testing.T) {
	type args struct {
		dir          Direction
		newDirection Direction
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 90 degrees",
			args: args{
				dir:          North,
				newDirection: East,
			},
			want: 1000,
		},
		{
			name: "Test 180 degrees",
			args: args{
				dir:          West,
				newDirection: East,
			},
			want: 2000,
		},
		{
			name: "Test 270 degrees",
			args: args{
				dir:          South,
				newDirection: East,
			},
			want: 1000,
		},
		{
			name: "test 0 or 360 degrees",
			args: args{
				dir:          North,
				newDirection: North,
			},
			want: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := rotationScore(test.args.dir, test.args.newDirection); got != test.want {
				t.Errorf("rotationScore() = %v, want %v", got, test.want)
			}
		})
	}
}
