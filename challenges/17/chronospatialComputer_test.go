package _17

import "testing"

func TestChronospatialComputer(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   int
		wantErr bool
	}{
		{
			name:    "test case 1",
			args:    args{filename: "test1.txt"},
			want:    "4,6,3,5,6,3,5,2,1,0",
			want1:   0,
			wantErr: false,
		},
		{
			name:    "input",
			args:    args{filename: "input.txt"},
			want:    "2,1,3,0,5,2,3,7,1",
			want1:   0,
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, got1, err := ChronospatialComputer(test.args.filename)
			if (err != nil) != test.wantErr {
				t.Errorf("ChronospatialComputer() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("ChronospatialComputer() got = %v, want %v", got, test.want)
			} else {
				t.Logf("got = %v, want %v", got, test.want)
			}
			if got1 != test.want1 {
				t.Errorf("ChronospatialComputer() got1 = %v, want %v", got1, test.want1)
			}
		})
	}
}

func TestComputer_adv(t *testing.T) {
	type fields struct {
		A                  int
		B                  int
		C                  int
		programm           []int
		instructionPointer int
	}
	type args struct {
		comboOperator int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantA  int
	}{
		{
			name:   "adv 1",
			fields: fields{A: 8},
			args:   args{comboOperator: 2},
			wantA:  2,
		},
		{
			name:   "adv 2",
			fields: fields{A: 128, B: 5},
			args:   args{comboOperator: 5},
			wantA:  4,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &Computer{
				A:                  test.fields.A,
				B:                  test.fields.B,
				C:                  test.fields.C,
				programm:           test.fields.programm,
				instructionPointer: test.fields.instructionPointer,
			}
			c.adv(test.args.comboOperator)

			if c.A != test.wantA {
				t.Errorf("Computer.adv() = %v, want %v", c.A, test.wantA)
			}
		})
	}
}

func TestComputer_bst(t *testing.T) {
	type fields struct {
		A                  int
		B                  int
		C                  int
		programm           []int
		instructionPointer int
	}
	type args struct {
		comboOperator int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantB  int
	}{
		{
			name:   "bst 1",
			fields: fields{B: 3},
			args:   args{comboOperator: 2},
			wantB:  2,
		},
		{
			name:   "bst 2",
			fields: fields{A: 10, B: 0},
			args:   args{comboOperator: 4},
			wantB:  2,
		},
		{
			name:   "bst 3",
			fields: fields{B: 15, C: 20},
			args:   args{comboOperator: 6},
			wantB:  4,
		},
		{
			name:   "bst 4",
			fields: fields{B: 15},
			args:   args{comboOperator: 5},
			wantB:  7,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &Computer{
				A:                  test.fields.A,
				B:                  test.fields.B,
				C:                  test.fields.C,
				programm:           test.fields.programm,
				instructionPointer: test.fields.instructionPointer,
			}
			c.bst(test.args.comboOperator)

			if c.B != test.wantB {
				t.Errorf("Computer.bst() = %v, want %v", c.B, test.wantB)
			}
		})
	}
}
