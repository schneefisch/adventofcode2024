package _9

import "testing"

func TestDiskFragmenter(t *testing.T) {
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
			name:    "test-input",
			args:    args{"test_input.txt"},
			want:    1928,
			want1:   2858,
			wantErr: false,
		},
		{
			name:    "input",
			args:    args{"input.txt"},
			want:    6283404590840,
			want1:   6304576012713,
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, got1, err := DiskFragmenter(test.args.filename)
			if (err != nil) != test.wantErr {
				t.Errorf("DiskFragmenter() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("DiskFragmenter() got = %v, want %v", got, test.want)
			}
			if got1 != test.want1 {
				t.Errorf("DiskFragmenter() got1 = %v, want %v", got1, test.want1)
			}
		})
	}
}
