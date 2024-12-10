package _0

import "testing"

func TestHoofIt(t *testing.T) {
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
			args:    args{filename: "test1.txt"},
			want:    36,
			want1:   81,
			wantErr: false,
		},
		{
			name:    "Input",
			args:    args{filename: "input.txt"},
			want:    733,
			want1:   1514,
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, got1, err := HoofIt(test.args.filename)
			if (err != nil) != test.wantErr {
				t.Errorf("HoofIt() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("HoofIt() got = %v, want %v", got, test.want)
			}
			if got1 != test.want1 {
				t.Errorf("HoofIt() got1 = %v, want %v", got1, test.want1)
			}
		})
	}
}
