package _15

import "testing"

func TestWarehouse(t *testing.T) {
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
			name: "Small example",
			args: args{
				filename: "small_example.txt",
			},
			want:    2028,
			want1:   0,
			wantErr: false,
		},
		{
			name: "Example",
			args: args{
				filename: "test.txt",
			},
			want:    10092,
			want1:   0,
			wantErr: false,
		},
		{
			name: "Input",
			args: args{
				filename: "input.txt",
			},
			want:    1514333,
			want1:   0,
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, got1, err := WarehouseWoes(test.args.filename)
			if (err != nil) != test.wantErr {
				t.Errorf("Warehouse() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("Warehouse() got = %v, want %v", got, test.want)
			} else {
				t.Logf("Got: %v", got)
			}
			if got1 != test.want1 {
				t.Errorf("Warehouse() got1 = %v, want %v", got1, test.want1)
			} else {
				t.Logf("Got1: %v", got1)
			}
		})
	}
}
