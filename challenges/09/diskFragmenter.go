package _9

import (
	"adventofcode2024/challenges/util"
	"fmt"
	"log"
)

func DiskFragmenter(filename string) (int, int, error) {
	input, err := util.ReadLines(filename)
	if err != nil {
		return 0, 0, err
	}
	if len(input) != 1 {
		return 0, 0, fmt.Errorf("expected 1 line of input, got %d", len(input))
	}

	disk := NewAmphipod()
	if err = disk.ParseInput(input[0]); err != nil {
		return 0, 0, err
	}

	//disk.Print()

	disk.Fragment()
	log.Println("Fragmented disk (simple)")
	//disk.Print()

	checksum := disk.Checksum()

	// advanced mode
	disk = NewAmphipod()
	if err = disk.ParseInput(input[0]); err != nil {
		return 0, 0, err
	}

	disk.AdvancedFragment()
	log.Println("Fragmented disk (advanced)")
	//disk.Print()
	checksum2 := disk.Checksum()

	return checksum, checksum2, nil
}
