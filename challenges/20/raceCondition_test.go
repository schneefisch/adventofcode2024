package _20

import "testing"

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
