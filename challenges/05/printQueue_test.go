package _5

import (
	"reflect"
	"testing"
)

func TestPrintQueue(t *testing.T) {
	type args struct {
		filenameOrdering string
		filenameUpdates  string
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantInvalid int
		wantErr     bool
	}{
		{
			name: "Test case 1",
			args: args{
				filenameOrdering: "test_rules.txt",
				filenameUpdates:  "test_updates.txt",
			},
			want:        143,
			wantInvalid: 123,
			wantErr:     false,
		},
		{
			name: "input",
			args: args{
				filenameOrdering: "input_rules.txt",
				filenameUpdates:  "input_updates.txt",
			},
			want:        0,
			wantInvalid: 0,
			wantErr:     false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotValid, gotInvalid, err := PrintQueue(test.args.filenameOrdering, test.args.filenameUpdates)
			if (err != nil) != test.wantErr {
				t.Errorf("PrintQueue() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if gotValid != test.want {
				t.Errorf("PrintQueue() gotValid = %v, want %v", gotValid, test.want)
			} else {
				t.Logf("PrintQueue() gotValid = %v, want %v", gotValid, test.want)
			}
			if gotInvalid != test.wantInvalid {
				t.Errorf("PrintQueue() gotInvalid = %v, want %v", gotInvalid, test.wantInvalid)
			} else {
				t.Logf("PrintQueue() gotInvalid = %v, want %v", gotInvalid, test.wantInvalid)
			}
		})
	}
}

func Test_fixPageOrdering(t *testing.T) {
	graph := &Graph{
		vertices: map[int][]int{
			29: {13},
			47: {53, 13, 61, 29},
			53: {29, 13},
			61: {13, 53, 29},
			75: {29, 53, 47, 61, 13},
			97: {13, 61, 47, 29, 53, 75},
		},
	}
	type args struct {
		update Update
		graph  *Graph
	}
	tests := []struct {
		name string
		args args
		want Update
	}{
		{
			name: "Test case 1",
			args: args{
				update: Update{pages: []int{75, 97, 47, 61, 53}},
				graph:  graph,
			},
			want: Update{pages: []int{97, 75, 47, 61, 53}},
		},
		{
			name: "Test case 2",
			args: args{
				update: Update{pages: []int{61, 13, 29}},
				graph:  graph,
			},
			want: Update{pages: []int{61, 29, 13}},
		},
		{
			name: "Test case 3",
			args: args{
				update: Update{pages: []int{97, 13, 75, 29, 47}},
				graph:  graph,
			},
			want: Update{pages: []int{97, 75, 47, 29, 13}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := fixPageOrdering(test.args.update, test.args.graph); !reflect.DeepEqual(got, test.want) {
				t.Errorf("fixPageOrdering() = %v, want %v", got, test.want)
			}
		})
	}
}
