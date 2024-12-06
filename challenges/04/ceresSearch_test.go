package _4

import "testing"

func TestCeresSearch(t *testing.T) {

	tests := []struct {
		name     string
		filename string
		want     int
		wantErr  bool
	}{
		{
			name:     "Test 1",
			filename: "test-input.txt",
			want:     9,
			wantErr:  false,
		},
		{
			name:     "Real input",
			filename: "input.txt",
			want:     1854,
			wantErr:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := CeresSearch(test.filename)
			if (err != nil) != test.wantErr {
				t.Errorf("CeresSearch() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("CeresSearch() got = %v, want %v", got, test.want)
			} else {
				t.Logf("CeresSearch() got = %v, want %v", got, test.want)
			}
		})
	}
}

func Test_findDiagonalOccurrencesInMap(t *testing.T) {
	// ignore test
	t.Skip()

	type args struct {
		matrix [][]rune
		s      string
	}
	testMatrix := [][]rune{
		[]rune("abc"),
		[]rune("def"),
		[]rune("ghi"),
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				matrix: testMatrix,
				s:      "ceg",
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				matrix: testMatrix,
				s:      "bf",
			},
			want: 1,
		},
		{
			name: "Test 3",
			args: args{
				matrix: testMatrix,
				s:      "aei",
			},
			want: 1,
		},
		{
			name: "Test 4",
			args: args{
				matrix: testMatrix,
				s:      "hf",
			},
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := findDiagonalOccurrencesInMap(test.args.matrix, test.args.s); got != test.want {
				t.Errorf("findDiagonalOccurrencesInMap() = %v, want %v", got, test.want)
			}
		})
	}
}

func Test_invertString(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Test 1",
			input: "abc",
			want:  "cba",
		},
		{
			name:  "Test 2",
			input: "de",
			want:  "ed",
		},
		{
			name:  "Test 3",
			input: "CHRISTmas",
			want:  "samTSIRHC",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := invertString(test.input); got != test.want {
				t.Errorf("invertString() = %v, want %v", got, test.want)
			}
		})
	}
}
