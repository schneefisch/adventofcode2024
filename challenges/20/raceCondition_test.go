package _20

import (
	"adventofcode2024/challenges/util"
	"testing"
)

func TestRaceCondition(t *testing.T) {
	type args struct {
		filename  string
		threshold int
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
			args:    args{filename: "test1.txt", threshold: 12},
			want:    8,
			want1:   0,
			wantErr: false,
		},
		{
			name:    "Test case 2",
			args:    args{filename: "test1.txt", threshold: 0},
			want:    44,
			want1:   0,
			wantErr: false,
		},
		{
			name:    "Test case 3",
			args:    args{filename: "test1.txt", threshold: 50},
			want:    0,
			want1:   285,
			wantErr: false,
		},
		{
			name:    "input",
			args:    args{filename: "input.txt", threshold: 100},
			want:    1346,
			want1:   2345398, // ToDo: part-two solution is not correct (too high)
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got1, got2, err := RaceCondition(test.args.filename, test.args.threshold)
			if (err != nil) != test.wantErr {
				t.Errorf("RaceCondition() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if test.want != 0 && got1 != test.want {
				t.Errorf("RaceCondition() got1 = %v, want %v", got1, test.want)
			}
			if test.want1 != 0 && got2 != test.want1 {
				t.Errorf("RaceCondition() got2 = %v, want %v", got2, test.want1)
			}
		})
	}
}

func Test_distance(t *testing.T) {
	type args struct {
		position  util.Position
		targetPos util.Position
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{position: util.Position{X: 0, Y: 0}, targetPos: util.Position{X: 1, Y: 1}},
			want: 2,
		},
		{
			name: "Test case 2",
			args: args{position: util.Position{X: 0, Y: 0}, targetPos: util.Position{X: 0, Y: 1}},
			want: 1,
		},
		{
			name: "Test case 3",
			args: args{position: util.Position{X: 15, Y: 15}, targetPos: util.Position{X: 15, Y: 15}},
			want: 0,
		},
		{
			name: "negative distance",
			args: args{position: util.Position{X: 8, Y: 8}, targetPos: util.Position{X: 2, Y: 8}},
			want: 6,
		},
		{
			name: "negative distance 2",
			args: args{position: util.Position{X: 8, Y: 8}, targetPos: util.Position{X: 2, Y: 0}},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distance(tt.args.position, tt.args.targetPos); got != tt.want {
				t.Errorf("distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
