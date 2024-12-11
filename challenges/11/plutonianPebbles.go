package _11

import (
	"adventofcode2024/challenges/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// StonePark represents the stones and their count
// the order of the stones is irrelevant, therefore we can count the occurrences of each stone in a map
type StonePark struct {
	stones map[int]int
}

func (s *StonePark) Length() int {
	l := 0
	for _, v := range s.stones {
		l += v
	}
	return l
}

func PlutionianPebbles(filename string, iterations int) (int, error) {
	input, err := util.ReadLines(filename)
	if err != nil {
		return 0, err
	}
	stones, err := parseStones(input)
	if err != nil {
		return 0, err
	}
	log.Println(stones)
	blinkTimes(stones, iterations)

	return stones.Length(), nil
}

func blinkTimes(stonePark *StonePark, times int) {
	for i := 0; i < times; i++ {
		blink(stonePark)
		log.Printf("After %d iterations got: %d stones", i+1, stonePark.Length())
	}
}

// blink takes a list of stones with engraved numbers and processes them according to the rules (in this order):
//   - If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
//   - If the stone is engraved with a number that has an even number of digits,
//     it is replaced by two stones. The left half of the digits are engraved on the new left stone,
//     and the right half of the digits are engraved on the new right stone. (
//     The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
//   - If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by
//     2024 is engraved on the new stone.
func blink(park *StonePark) {

	newStones := make(map[int]int)

	for k, v := range park.stones {
		if k == 0 {
			newStones[1] += v
		} else {
			stoneString := strconv.Itoa(k)
			if len(stoneString)%2 == 0 {
				half := len(stoneString) / 2
				left, err := strconv.Atoi(stoneString[:half])
				if err != nil {
					log.Fatal(err)
				}
				right, err := strconv.Atoi(stoneString[half:])
				if err != nil {
					log.Fatal(err)
				}
				newStones[left] += v
				newStones[right] += v
			} else {
				newStones[k*2024] += v
			}
		}
	}
	park.stones = newStones
}

func parseStones(input []string) (*StonePark, error) {
	stones := &StonePark{stones: make(map[int]int)}
	// split at whitespace
	if len(input) == 0 {
		return stones, fmt.Errorf("no stones found")
	}
	if len(input) > 1 {
		return stones, fmt.Errorf("too many lines")
	}
	stoneStrings := strings.Split(input[0], " ")
	for _, stoneString := range stoneStrings {
		stone, err := strconv.Atoi(stoneString)
		if err != nil {
			return stones, err
		}
		// check if the stone is already in the map
		if _, ok := stones.stones[stone]; ok {
			// increase the count
			stones.stones[stone]++
		} else {
			// add the stone to the map
			stones.stones[stone] = 1
		}
	}
	return stones, nil
}
