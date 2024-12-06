package _1

import (
	"log"
	"testing"
)

func Test_day1(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			name:     "example",
			filename: "input_test.csv",
			want:     11,
		},
		{
			name:     "input",
			filename: "input.csv",
			want:     0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, _, err := dayOneHistorianHysteria(test.filename)
			if err != nil {
				t.Errorf("day1() error = %v", err)
			}
			if test.want != 0 && got != test.want {
				t.Errorf("day1() = %v, want %v", got, test.want)
			}
			log.Printf("\n==== Distance of %s: %d\n\n", test.name, got)
		})
	}
}

func Test_similarityScore(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			name:     "example",
			filename: "input_test.csv",
			want:     31,
		},
		{
			name:     "input",
			filename: "input.csv",
			want:     0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, got, err := dayOneHistorianHysteria(test.filename)
			if err != nil {
				t.Errorf("day1() error = %v", err)
			}
			if test.want != 0 && got != test.want {
				t.Errorf("day1() = %v, want %v", got, test.want)
			}
			log.Printf("\n==== Similarity score of %s: %d\n\n", test.name, got)
		})

	}
}
