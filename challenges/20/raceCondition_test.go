package _20

import "testing"

func TestRaceCondition(t *testing.T) {
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
			name: "Test case 1",
			args: args{filename: "test1.txt"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := RaceCondition(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("RaceCondition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RaceCondition() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("RaceCondition() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
