package _12

import "testing"

func TestGardenGroups(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want2   int
		wantErr bool
	}{
		{
			name:    "Test case 1",
			args:    args{"test.txt"},
			want:    140,
			want2:   0,
			wantErr: false,
		},
		{
			name:    "Test case 2",
			args:    args{"test2.txt"},
			want:    772,
			want2:   0,
			wantErr: false,
		},
		{
			name:    "Test case 3",
			args:    args{"test3.txt"},
			want:    1930,
			want2:   0,
			wantErr: false,
		},
		{
			name:    "input",
			args:    args{"input.txt"},
			want:    1319878,
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, got2, err := GardenGroups(test.args.filename)
			if (err != nil) != test.wantErr {
				t.Errorf("GardenGroups() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if test.want != 0 && got != test.want {
				t.Errorf("GardenGroups() got = %v, want %v", got, test.want)
			}
			if test.want2 != 0 && got2 != test.want2 {
				t.Errorf("GardenGroups() got2 = %v, want2 %v", got2, test.want2)
			}
		})
	}
}
