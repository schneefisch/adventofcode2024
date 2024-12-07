package _7

import "testing"

func TestBridgeRepair(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name         string
		args         args
		want         int
		wantExtended int
		wantErr      bool
	}{
		{
			name:         "Test 1",
			args:         args{filename: "test_input.txt"},
			want:         3749,
			wantExtended: 11387,
			wantErr:      false,
		},
		{
			name:         "Input",
			args:         args{filename: "input.txt"},
			want:         1153997401072,
			wantExtended: 97902809384118,
			wantErr:      false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, gotExtended, err := BridgeRepair(test.args.filename)
			if (err != nil) != test.wantErr {
				t.Errorf("BridgeRepair() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("BridgeRepair() got = %v, want %v", got, test.want)
			}
			if gotExtended != test.wantExtended {
				t.Errorf("BridgeRepair() gotExtended = %v, want %v", gotExtended, test.wantExtended)
			}
		})
	}
}
