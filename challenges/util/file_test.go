package util

import (
	"reflect"
	"testing"
)

var testMatrix = [][]rune{
	{'1', '2', '3'},
	{'4', '5', '6'},
	{'7', '8', '9'},
}

var expectedMatrix = [][]rune{
	{'7', '4', '1'},
	{'8', '5', '2'},
	{'9', '6', '3'},
}

func TestRotateMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]rune
		want   [][]rune
	}{
		{
			name:   "Test 3x3 Matrix",
			matrix: testMatrix,
			want:   expectedMatrix,
		},
		{
			name: "Test 2x3 matrix",
			matrix: [][]rune{
				{'1', '2', '3'},
				{'4', '5', '6'},
			},
			want: [][]rune{
				{'4', '1'},
				{'5', '2'},
				{'6', '3'},
			},
		},
		{
			name: "Test 4x4 matrix",
			matrix: [][]rune{
				{'1', '2', '3', '4'},
				{'5', '6', '7', '8'},
				{'9', '0', '1', '2'},
				{'3', '4', '5', '6'},
			},
			want: [][]rune{
				{'3', '9', '5', '1'},
				{'4', '0', '6', '2'},
				{'5', '1', '7', '3'},
				{'6', '2', '8', '4'},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := RotateMatrix(test.matrix); !reflect.DeepEqual(got, test.want) {
				t.Logf("error RotateMatrix() got:")
				PrintMap(got)
				t.Logf("error RotateMatrix() want:")
				PrintMap(test.want)
				t.Errorf("unexpected result")
			}
		})
	}
}
