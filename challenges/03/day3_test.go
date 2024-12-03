package _3

import "testing"

func TestDay3(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
		wantErr  bool
	}{
		{
			name:     "test-inptut",
			filename: "test_input.txt",
			want:     60,
			wantErr:  false,
		},
		{
			name:     "day-input",
			filename: "input.txt",
			want:     95846796,
			wantErr:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := DaythreeMullitover(test.filename)
			if (err != nil) != test.wantErr {
				t.Errorf("Day3() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("Day3() got = %v, want %v", got, test.want)
			} else {
				t.Logf("Day3() got = %v, want %v", got, test.want)
			}
		})
	}
}
