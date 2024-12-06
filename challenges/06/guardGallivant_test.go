package _6

import "testing"

func TestGuardGallivant(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name                 string
		args                 args
		want                 int
		wantObstaclePosCount int
		wantErr              bool
	}{
		{
			name:                 "Test 1",
			args:                 args{filename: "test_input.txt"},
			want:                 41,
			wantObstaclePosCount: 6,
			wantErr:              false,
		},
		{
			name:                 "Input",
			args:                 args{filename: "input.txt"},
			want:                 5030,
			wantObstaclePosCount: 1928,
			wantErr:              false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotVisited, gotObstaclePositions, err := GuardGallivant(test.args.filename)
			if (err != nil) != test.wantErr {
				t.Errorf("GuardGallivant() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if gotVisited != test.want {
				t.Errorf("GuardGallivant() gotVisited = %v, want %v", gotVisited, test.want)
			} else {
				t.Logf("GuardGallivant() visited = %v", gotVisited)
			}
			if gotObstaclePositions != test.wantObstaclePosCount {
				t.Errorf("GuardGallivant() gotObstaclePositions = %v, want %v", gotObstaclePositions, test.wantObstaclePosCount)
			} else {
				t.Logf("GuardGallivant() obstacle positions = %v", gotObstaclePositions)
			}
		})
	}
}
