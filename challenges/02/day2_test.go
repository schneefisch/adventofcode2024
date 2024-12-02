package _2

import "testing"

func Test_redNosedReports(t *testing.T) {
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
			name: "example",
			args: args{
				filename: "test-input.csv",
			},
			want: 6,
		},
		{
			name: "input",
			args: args{
				filename: "input.csv",
			},
			want: 290,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := redNosedReports(test.args.filename)
			if (err != nil) != test.wantErr {
				t.Errorf("redNosedReports() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("redNosedReports() got = %v, want %v", got, test.want)
			} else {
				t.Logf("redNosedReports() success! got = %v, want %v", got, test.want)
			}
		})
	}
}
